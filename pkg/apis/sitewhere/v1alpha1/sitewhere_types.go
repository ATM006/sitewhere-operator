package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SiteWhereSpec defines the desired state of SiteWhere
// +k8s:openapi-gen=true
type SiteWhereSpec struct {
	// Size is the size of the memcached deployment
	Size int32 `json:"size"`
}

// SiteWhereStatus defines the observed state of SiteWhere
// +k8s:openapi-gen=true
type SiteWhereStatus struct {
	// Nodes are the names of the memcached pods
	Nodes []string `json:"nodes"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SiteWhere is the Schema for the sitewheres API
// +k8s:openapi-gen=true
type SiteWhere struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SiteWhereSpec   `json:"spec,omitempty"`
	Status SiteWhereStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SiteWhereList contains a list of SiteWhere
type SiteWhereList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SiteWhere `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SiteWhere{}, &SiteWhereList{})
}
