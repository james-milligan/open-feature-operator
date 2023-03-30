package flagsourceconfiguration

import (
	"context"
	"fmt"
	"os"

	"github.com/go-logr/logr"
	"github.com/open-feature/open-feature-operator/pkg/utils"
	appsV1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	ManagedByAnnotationValue = "open-feature-operator"
)

type KubeFlagdProxyHandler struct {
	client.Client
	config *KubeProxyConfiguration
	Log    logr.Logger
}

type KubeProxyConfiguration struct {
	Port         int
	MetricsPort  int
	DebugLogging bool
	Image        string
	Tag          string
	Namespace    string
}

func NewKubeProxyConfiguration() (*KubeProxyConfiguration, error) {
	config := &KubeProxyConfiguration{
		Port:         defaultKubeProxyPort,
		MetricsPort:  defaultKubeProxyMetricsPort,
		DebugLogging: defaultKubeProxyDebugLogging,
		Image:        defaultKubeProxyImage,
		Tag:          defaultKubeProxyTag,
		Namespace:    defaultKubeProxyNamespace,
	}
	ns, ok := os.LookupEnv(envVarPodNamespace)
	if ok {
		config.Namespace = ns
	}
	kpi, ok := os.LookupEnv(envVarProxyImage)
	if ok {
		config.Image = kpi
	}
	kpt, ok := os.LookupEnv(envVarProxyTag)
	if ok {
		config.Tag = kpt
	}
	port, err := utils.GetIntEnvVar(envVarProxyPort)
	if err != nil {
		return config, err
	}
	config.Port = port

	metricsPort, err := utils.GetIntEnvVar(envVarProxyMetricsPort)
	if err != nil {
		return config, err
	}
	config.MetricsPort = metricsPort

	kpDebugLogging, err := utils.GetBoolEnvVar(envVarProxyDebugLogging)
	if err != nil {
		return config, err
	}
	config.DebugLogging = kpDebugLogging

	return config, nil
}

func NewKubeFlagdProxyHandler(config *KubeProxyConfiguration, client client.Client, logger logr.Logger) *KubeFlagdProxyHandler {
	return &KubeFlagdProxyHandler{
		config: config,
		Client: client,
		Log:    logger,
	}
}

func (k *KubeFlagdProxyHandler) Config() *KubeProxyConfiguration {
	return k.config
}

func (k *KubeFlagdProxyHandler) handleKubeProxy(ctx context.Context) error {
	exists, err := k.doesKubeProxyExist(ctx)
	if err != nil {
		return err
	}
	if !exists {
		return k.deployKubeProxy(ctx)
	}
	return nil
}

func (k *KubeFlagdProxyHandler) deployKubeProxy(ctx context.Context) error {
	k.Log.Info("deploying the kube-flagd-proxy")
	if err := k.Client.Create(ctx, k.newFlagdKubeProxyManifest()); err != nil && !errors.IsAlreadyExists(err) {
		return err
	}
	k.Log.Info("deploying the kube-flagd-proxy service")
	if err := k.Client.Create(ctx, k.newFlagdKubeProxyServiceManifest()); err != nil && !errors.IsAlreadyExists(err) {
		return err
	}
	return nil
}

func (k *KubeFlagdProxyHandler) newFlagdKubeProxyServiceManifest() *corev1.Service {
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      KubeProxyServiceName,
			Namespace: k.config.Namespace,
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				"app.kubernetes.io/name":       KubeProxyDeploymentName,
				"app.kubernetes.io/managed-by": ManagedByAnnotationValue,
			},
			Ports: []corev1.ServicePort{
				{
					Name:       "kube-flagd-proxy",
					Port:       int32(k.config.Port),
					TargetPort: intstr.FromInt(k.config.Port),
				},
			},
		},
	}
}

func (k *KubeFlagdProxyHandler) newFlagdKubeProxyManifest() *appsV1.Deployment {
	replicas := int32(1)
	args := []string{
		"start",
		"--metrics-port",
		fmt.Sprintf("%d", k.config.MetricsPort),
	}
	if k.config.DebugLogging {
		args = append(args, "--debug")
	}
	return &appsV1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      KubeProxyDeploymentName,
			Namespace: k.config.Namespace,
			Labels: map[string]string{
				"app":                          KubeProxyDeploymentName,
				"app.kubernetes.io/managed-by": ManagedByAnnotationValue,
				"app.kubernetes.io/version":    k.config.Tag,
			},
		},
		Spec: appsV1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": KubeProxyDeploymentName,
				},
			},

			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":                          KubeProxyDeploymentName,
						"app.kubernetes.io/name":       KubeProxyDeploymentName,
						"app.kubernetes.io/managed-by": ManagedByAnnotationValue,
						"app.kubernetes.io/version":    k.config.Tag,
					},
				},
				Spec: corev1.PodSpec{
					ServiceAccountName: KubeProxyServiceAccountName,
					Containers: []corev1.Container{
						{
							Image: fmt.Sprintf("%s:%s", k.config.Image, k.config.Tag),
							Name:  KubeProxyDeploymentName,
							Ports: []corev1.ContainerPort{
								{
									Name:          "port",
									ContainerPort: int32(k.config.Port),
								},
								{
									Name:          "metrics-port",
									ContainerPort: int32(k.config.MetricsPort),
								},
							},
							Args: args,
						},
					},
				},
			},
		},
	}
}

func (r *KubeFlagdProxyHandler) doesKubeProxyExist(ctx context.Context) (bool, error) {
	r.Client.Scheme()
	d := appsV1.Deployment{}
	err := r.Client.Get(ctx, client.ObjectKey{Name: KubeProxyDeploymentName, Namespace: r.config.Namespace}, &d)
	if err != nil {
		if errors.IsNotFound(err) {
			// does not exist, is not ready, no error
			return false, nil
		}
		// does not exist, is not ready, is in error
		return false, err
	}
	// exists, at least one replica ready, no error
	return true, nil
}
