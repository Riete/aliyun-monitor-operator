package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	RDS_EXPORTER_IMAGE     string = "riet/aliyun-rds-exporter"
	RDS_EXPORTER_NAMESPACE string = "ops-monitor"
	RDS_EXPORTER_PORT      int32  = 10001
)

type RdsAliyunAccount struct {
	Name            string `json:"name"`
	AccessKeyId     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	RegionId        string `json:"regionId"`
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RdsMonitorSpec defines the desired state of RdsMonitor
type RdsMonitorSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	RdsAliyunAccount []RdsAliyunAccount `json:"rdsAliyunAccount"`
}

// RdsMonitorStatus defines the observed state of RdsMonitor
type RdsMonitorStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RdsMonitor is the Schema for the rdsmonitors API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=rdsmonitors,scope=Namespaced
type RdsMonitor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RdsMonitorSpec   `json:"spec,omitempty"`
	Status RdsMonitorStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RdsMonitorList contains a list of RdsMonitor
type RdsMonitorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RdsMonitor `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RdsMonitor{}, &RdsMonitorList{})
}
