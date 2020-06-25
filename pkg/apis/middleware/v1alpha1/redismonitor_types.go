package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	RedisExporterImage     string = "riet/aliyun-redis-exporter"
	RedisExporterNamespace string = "ops-monitor"
	RedisExporterPort      int32  = 10003
)

type RedisAliyunAccount struct {
	Name            string   `json:"name"`
	AccessKeyId     string   `json:"accessKeyId"`
	AccessKeySecret string   `json:"accessKeySecret"`
	RegionId        string   `json:"regionId"`
	ExtraMetric     []string `json:"extraMetric,omitempty"`
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RedisMonitorSpec defines the desired state of RedisMonitor
type RedisMonitorSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	RedisAliyunAccount []RedisAliyunAccount `json:"redisAliyunAccount"`
}

// RedisMonitorStatus defines the observed state of RedisMonitor
type RedisMonitorStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RedisMonitor is the Schema for the redismonitors API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=redismonitors,scope=Namespaced
type RedisMonitor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RedisMonitorSpec   `json:"spec,omitempty"`
	Status RedisMonitorStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RedisMonitorList contains a list of RedisMonitor
type RedisMonitorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisMonitor `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RedisMonitor{}, &RedisMonitorList{})
}
