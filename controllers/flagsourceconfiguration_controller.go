/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	"strings"
	"time"

	"k8s.io/api/apps/v1beta1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	"github.com/go-logr/logr"
	corev1alpha1 "github.com/open-feature/open-feature-operator/apis/core/v1alpha1"
)

const (
	OpenFeatureAnnotationPath         = "metadata.annotations.openfeature.dev/openfeature.dev"
	FlagSourceConfigurationAnnotation = "flagsourceconfiguration"
)

// FlagSourceConfigurationReconciler reconciles a FlagSourceConfiguration object
type FlagSourceConfigurationReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	// ReqLogger contains the Logger of this controller
	Log logr.Logger
}

//+kubebuilder:rbac:groups=core.openfeature.dev,resources=flagsourceconfigurations,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core.openfeature.dev,resources=flagsourceconfigurations/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=core.openfeature.dev,resources=flagsourceconfigurations/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the FlagSourceConfiguration object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.0/pkg/reconcile
func (r *FlagSourceConfigurationReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	r.Log = log.FromContext(ctx)

	fsConfig := &corev1alpha1.FlagSourceConfiguration{}
	if err := r.Client.Get(ctx, req.NamespacedName, fsConfig); err != nil {
		if errors.IsNotFound(err) {
			// taking down all associated K8s resources is handled by K8s
			r.Log.Info(crdName + " resource not found. Ignoring since object must be deleted")
			return r.finishReconcile(nil, false)
		}
		r.Log.Error(err, "Failed to get the "+crdName)
		return r.finishReconcile(err, false)
	}

	// Object has been updated, so, we can restart any deployments that are using this annotation
	deployList := &v1beta1.DeploymentList{}
	if err := r.Client.List(ctx, deployList, client.MatchingFields{
		fmt.Sprintf("%s/%s", OpenFeatureAnnotationPath, FlagSourceConfigurationAnnotation): "true",
	}); err != nil {
		r.Log.Error(err, fmt.Sprintf("Failed to get the pods with annotation %s/%s", OpenFeatureAnnotationPath, FlagSourceConfigurationAnnotation))
		return r.finishReconcile(err, false)
	}

	for _, deployment := range deployList.Items {
		annotations := deployment.GetAnnotations()
		annotation, ok := annotations[FlagSourceConfigurationAnnotation]
		if !ok {
			continue
		}
		if isUsingConfiguration(fsConfig.Namespace, fsConfig.Name, deployment.Namespace, annotation) {
			deployment.Spec.Template.ObjectMeta.Annotations["kubectl.kubernetes.io/restartedAt"] = time.Now().Format(time.RFC3339)
			if err := r.Client.Update(ctx, &deployment); err != nil {
				r.Log.V(1).Error(err, fmt.Sprintf("Failed to update Deployment: %s", err.Error()))
				continue
			}
		}
	}

	return r.finishReconcile(nil, false)
}

func isUsingConfiguration(namespace string, name string, deploymentNamespace string, annotation string) bool {
	s := strings.Split(annotation, ",")
	for _, target := range s {
		ss := strings.Split(strings.TrimSpace(target), "/")
		if len(ss) != 2 {
			target = fmt.Sprintf("%s/%s", deploymentNamespace, target)
		}
		if target == fmt.Sprintf("%s/%s", namespace, name) {
			return true
		}
	}
	return false
}

func (r *FlagSourceConfigurationReconciler) finishReconcile(err error, requeueImmediate bool) (ctrl.Result, error) {
	if err != nil {
		interval := reconcileErrorInterval
		if requeueImmediate {
			interval = 0
		}
		r.Log.Error(err, "Finished Reconciling "+crdName+" with error: %w")
		return ctrl.Result{Requeue: true, RequeueAfter: interval}, err
	}
	interval := reconcileSuccessInterval
	if requeueImmediate {
		interval = 0
	}
	r.Log.Info("Finished Reconciling " + crdName)
	return ctrl.Result{Requeue: true, RequeueAfter: interval}, nil
}

func FlagSourceConfigurationAnnotationFilter(o client.Object) []string {
	pod := o.(*v1.Pod)
	if pod.ObjectMeta.Annotations == nil {
		return []string{
			"false",
		}
	}
	if _, ok := pod.ObjectMeta.Annotations[fmt.Sprintf("openfeature.dev/%s", FlagSourceConfigurationAnnotation)]; ok {
		return []string{
			"true",
		}
	}
	return []string{
		"false",
	}
}

// SetupWithManager sets up the controller with the Manager.
func (r *FlagSourceConfigurationReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&corev1alpha1.FlagSourceConfiguration{}).
		// we are only interested in update events for this reconciliation loop
		WithEventFilter(predicate.GenerationChangedPredicate{}).
		Complete(r)
}
