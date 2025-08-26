//go:generate ../../hack/generate-crd.sh

// +kubebuilder:object:generate=true
// +groupName=volume.pdok.nl
package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=pdok
type AzureVolumePopulator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// AzureVolumePopulatorSpec specification of an Azure storage account resource
	Spec AzureVolumePopulatorSpec `json:"spec"`
}

type AzureVolumePopulatorSpec struct {
	// Copy all blobs with this prefix
	// +kubebuilder:validation:Required
	BlobPrefix string `json:"blobPrefix"`

	// Destination path on the volume
	// +kubebuilder:validation:Required
	VolumePath string `json:"volumePath"`
}

// AzureVolumePopulatorList contains a list of AzureVolumePopulator
// +kubebuilder:object:root=true
type AzureVolumePopulatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureVolumePopulator `json:"items"`
}
