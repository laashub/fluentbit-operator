package output

import (
	"kubesphere.io/fluentbit-operator/api/v1alpha1/plugins"
)

// +kubebuilder:object:generate:=true

// The null output plugin just throws away events.
type Null struct{}

func (_ *Null) Name() string {
	return "null"
}

// implement Section() method
func (_ *Null) Params(_ plugins.SecretLoader) (*plugins.KVs, error) {
	return nil, nil
}
