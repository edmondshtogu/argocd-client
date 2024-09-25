package v1alpha1

import (
	"github.com/edmondshtogu/argocd-client/pkg/apis"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion                   = schema.GroupVersion{Group: apis.Group, Version: "v1alpha1"}
	ApplicationSchemaGroupVersionKind    = schema.GroupVersionKind{Group: apis.Group, Version: "v1alpha1", Kind: apis.ApplicationKind}
	AppProjectSchemaGroupVersionKind     = schema.GroupVersionKind{Group: apis.Group, Version: "v1alpha1", Kind: apis.AppProjectKind}
	ApplicationSetSchemaGroupVersionKind = schema.GroupVersionKind{Group: apis.Group, Version: "v1alpha1", Kind: apis.ApplicationSetKind}
)

// Resource takes an unqualified resource and returns a Group-qualified GroupResource.
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

// addKnownTypes adds the set of types defined in this package to the supplied scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&Application{},
		&ApplicationList{},
		&AppProject{},
		&AppProjectList{},
		&ApplicationSet{},
		&ApplicationSetList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
