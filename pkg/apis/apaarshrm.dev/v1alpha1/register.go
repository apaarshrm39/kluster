package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	scheme "k8s.io/apimachinery/pkg/runtime/schema"
)

var SchemeGroupVersion = scheme.GroupVersion{Group: "apaarshrm.dev", Version: "v1apha1"}

var (
	SchemeBuilder runtime.SchemeBuilder
)

func init() {
	SchemeBuilder.Register(addKnownTypes)
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion, &Kluster{}, &KlusterList{})
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
