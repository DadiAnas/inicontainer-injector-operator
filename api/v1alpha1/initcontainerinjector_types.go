/*
Copyright 2024.

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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DeploymentAnnotations defines the supported annotations
type DeploymentAnnotations struct {
	// InitContainerInjectorArgs contains the arguments for the init container
	// +required
	InitContainerInjectorArgs string `json:"initcontainer_injector_args"`

	// InitContainerInjectorImage specifies the container image
	// +optional
	InitContainerInjectorImage string `json:"initcontainer_injector_image,omitempty"`

	// InitContainerInjectorRegistry specifies the container registry
	// +optional
	InitContainerInjectorRegistry string `json:"initcontainer_injector_registry,omitempty"`

	// InitContainerInjectorCommand specifies the command to run
	// +optional
	InitContainerInjectorCommand string `json:"initcontainer_injector_command,omitempty"`
}

// InitContainerInjectorSpec defines the desired state of InitContainerInjector
type InitContainerInjectorSpec struct {
	// Template defines the init container to be injected
	Template corev1.Container `json:"template"`

	// Annotations holds the deployment annotations configuration
	Annotations DeploymentAnnotations `json:"annotations"`
}

// InitContainerInjectorStatus defines the observed state of InitContainerInjector
type InitContainerInjectorStatus struct {
	// InjectedDeployments is the number of deployments that have been injected
	InjectedDeployments int `json:"injectedDeployments"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// InitContainerInjector is the Schema for the initcontainerinjectors API
type InitContainerInjector struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InitContainerInjectorSpec   `json:"spec,omitempty"`
	Status InitContainerInjectorStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// InitContainerInjectorList contains a list of InitContainerInjector
type InitContainerInjectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InitContainerInjector `json:"items"`
}

func init() {
	SchemeBuilder.Register(&InitContainerInjector{}, &InitContainerInjectorList{})
}
