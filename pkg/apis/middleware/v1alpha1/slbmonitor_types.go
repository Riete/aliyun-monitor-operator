package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	SlbExporterImage     string = "riet/aliyun-slb-exporter"
	SlbExporterNamespace string = "ops-monitor"
	SlbExporterPort      int32  = 10002
)

type SlbAliyunAccount struct {
	Name            string   `json:"name"`
	AccessKeyId     string   `json:"accessKeyId"`
	AccessKeySecret string   `json:"accessKeySecret"`
	RegionId        string   `json:"regionId"`
	InstanceId      []string `json:"instanceId,omitempty"`
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SlbMonitorSpec defines the desired state of SlbMonitor
type SlbMonitorSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	SlbAliyunAccount []SlbAliyunAccount `json:"slbAliyunAccount"`
}

// SlbMonitorStatus defines the observed state of SlbMonitor
type SlbMonitorStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SlbMonitor is the Schema for the slbmonitors API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=slbmonitors,scope=Namespaced
type SlbMonitor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SlbMonitorSpec   `json:"spec,omitempty"`
	Status SlbMonitorStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SlbMonitorList contains a list of SlbMonitor
type SlbMonitorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SlbMonitor `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SlbMonitor{}, &SlbMonitorList{})
}
