package v1alpha1

import (
	stable "github.com/shohagrana64/goThings/crd/pkg/apis/stable.example.com"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var SchemaGroupVersion = schema.GroupVersion{Group: stable.GroupName, Version: "v1"}

var (
	SchemaBuilder      runtime.SchemeBuilder
	localSchemaBuilder = &SchemaBuilder
	AddToSchema        = localSchemaBuilder.AddToScheme
)

func init() {
	localSchemaBuilder.Register(addKnownTypes)
}

func Resource(resource string) schema.GroupResource {
	return SchemaGroupVersion.WithResource(resource).GroupResource()
}

func addKnownTypes(schema *runtime.Scheme) error {
	schema.AddKnownTypes(SchemaGroupVersion,
		&Disappointment{},
		&DisappointmentList{},
	)
	schema.AddKnownTypes(SchemaGroupVersion,
		&metav1.Status{},
	)

	metav1.AddToGroupVersion(schema, SchemaGroupVersion)
	return nil
}
