/*
Copyright 2023.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OpenSimulatorSpec defines the desired state of OpenSimulator
type OpenSimulatorSpec struct {
	Name string `json:"name"`
}

// OpenSimulatorStatus defines the observed state of OpenSimulator
type OpenSimulatorStatus struct {
	Started bool `json:"started,omitempty"`
  Configured bool `json:"configured,omitempty"`
  // INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// OpenSimulator is the Schema for the opensimulators API
type OpenSimulator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OpenSimulatorSpec   `json:"spec,omitempty"`
	Status OpenSimulatorStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// OpenSimulatorList contains a list of OpenSimulator
type OpenSimulatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpenSimulator `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OpenSimulator{}, &OpenSimulatorList{})
}
