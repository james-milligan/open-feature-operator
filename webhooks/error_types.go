package webhooks

import (
	goErr "errors"
	"net/http"
)

type flagdProxyDeferError struct{}

func (d *flagdProxyDeferError) Error() string {
	return "flagd-proxy is not ready, deferring pod admission"
}

func (m *PodMutator) IsReady(_ *http.Request) error {
	if m.ready {
		return nil
	}
	return goErr.New("pod mutator is not ready")
}
