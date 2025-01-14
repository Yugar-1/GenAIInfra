/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha3

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ServiceoperatorSpec defines the desired state of Serviceoperator.
type ServiceoperatorSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Image        Image                        `json:"image,omitempty"`
	Model        string                       `json:"model,omitempty"`
	ReplicaCount int32                        `json:"replicaCount,omitempty"`
	AccelDevice  string                       `json:"accelDevice,omitempty"`
	Resource     *corev1.ResourceRequirements `json:"resources,omitempty"`
}

type Image struct {
	Repository string `json:"repository,omitempty"`
	Tag        string `json:"tag,omitempty"`
	PullPolicy string `json:"pullPolicy,omitempty"`
}

// ServiceoperatorStatus defines the observed state of Serviceoperator.
type ServiceoperatorStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Serviceoperator is the Schema for the serviceoperators API.
type Serviceoperator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceoperatorSpec   `json:"spec,omitempty"`
	Status ServiceoperatorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceoperatorList contains a list of Serviceoperator.
type ServiceoperatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Serviceoperator `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Serviceoperator{}, &ServiceoperatorList{})
}
