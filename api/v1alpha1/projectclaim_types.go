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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ProjectClaimSpec defines the desired state of ProjectClaim
type ProjectClaimSpec struct {
	LegalEntity            LegalEntity    `json:"legalEntity"`
	GCPCredentialSecret    NamespacedName `json:"gcpCredentialSecret"`
	Region                 string         `json:"region"`
	GCPProjectID           string         `json:"gcpProjectID,omitempty"`
	ProjectReferenceCRLink NamespacedName `json:"projectReferenceCRLink,omitempty"`
	AvailabilityZones      []string       `json:"availabilityZones,omitempty"`
}

// ProjectClaimStatus defines the observed state of ProjectClaim
type ProjectClaimStatus struct {
	Conditions []Condition `json:"conditions"`
	State      ClaimStatus `json:"state"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ProjectClaim is the Schema for the projectclaims API
type ProjectClaim struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProjectClaimSpec   `json:"spec,omitempty"`
	Status ProjectClaimStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectClaimList contains a list of ProjectClaim
type ProjectClaimList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectClaim `json:"items"`
}

// ClaimStatus is a valid value from ProjectClaim.Status
type ClaimStatus string

const (
	// ClaimStatusPending pending status for a claim
	ClaimStatusPending ClaimStatus = "Pending"
	// ClaimStatusPendingProject pending project status for a claim
	ClaimStatusPendingProject ClaimStatus = "PendingProject"
	// ClaimStatusReady ready status for a claim
	ClaimStatusReady ClaimStatus = "Ready"
	// ClaimStatusError error status for a claim
	ClaimStatusError ClaimStatus = "Error"
	// ClaimStatusVerification pending verification status for a claim
	ClaimStatusVerification ClaimStatus = "Verification"
)

func init() {
	SchemeBuilder.Register(&ProjectClaim{}, &ProjectClaimList{})
}
