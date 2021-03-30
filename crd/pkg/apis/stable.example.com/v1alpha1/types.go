package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=kubeapis,singular:kubeapi,shortName=kapi,categories={api,all}
// +kubebuilder:subresources:status
// +kubebuilder:printcolumen:name="Version",type="string",JSONPath=".spec.version"
// +kubebuilder:printcolumn:name="Status",type="string",JSONPATH=".spec.phase"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPATH=".metadata.creationTimestamp"

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Disappointment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DisappointmentSpec   `json:"spec"`
	Status DisappointmentStatus `json:"status"`
}

type DisappointmentSpec struct {
	Version     string        `json:"version"`
	Replicas    *int32        `json:"replicas,omitempty"`
	HostUrl     string        `json:"hostUrl"`
	ServiceType string        `json:"serviceType"`
	Container   ContainerSpec `json:"container"`
}

type ContainerSpec struct {
	Image         string `json:"image"`
	ContainerPort int32  `json:"containerPort"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type DisappointmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Disappointment `json:"items"`
}

type DisappointmentStatus struct {
	// +optional
	Phase string `json:"phase,omitempty"`
}
