/*


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
	istioapi "istio.io/client-go/pkg/apis/networking/v1alpha3"
	appsv1 "k8s.io/api/apps/v1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ApiServiceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	DeploymentSpec appsv1.DeploymentSpec `json:"deploymentSpec"`
	ServiceSpec    corev1.ServiceSpec    `json:"serviceSpec"`
	// +optional
	PodAutoScalerSpec autoscalingv1.HorizontalPodAutoscalerSpec `json:"podAutoScalerSpec"`

	// Spec for rules. These two
	// +optional
	RoleKind string `json:"roleKind"`
	// +optional
	RoleMetadata metav1.ObjectMeta `json:"roleMetadata"`
	// +optional
	RoleRules []rbacv1.PolicyRule `json:"roleRules"`

	// Spec for role-bindings
	// +optional
	RoleBindingKind string `json:"roleBindingKind"`
	// +optional
	RoleBindingMetadata metav1.ObjectMeta `json:"roleBindingMetadata"`
	// +optional
	RoleBindingSubjects []rbacv1.Subject `json:"roleBindingSubjects"`
	// +optional
	RoleBindingRoleRef rbacv1.RoleRef `json:"roleBindingRoleRef"`

	// Spec for service account
	// +optional
	ServiceAccountMeta metav1.ObjectMeta `json:"serviceAccountMeta"`

	// Spec for istio virtual service
	// +optional
	VirtualService istioapi.VirtualService `json:"virtualService"`
}

type ApiServiceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +genclient
type ApiService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApiServiceSpec   `json:"spec,omitempty"`
	Status ApiServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

type ApiServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ApiService{}, &ApiServiceList{})
}
