//go:generate ../../hack/generate-crd.sh

// +kubebuilder:object:generate=true
// +groupName=volume.pdok.nl
package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// AzureVolumePopulator is the Schema for the AzureVolumePopulatorSpec
// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=pdok
type AzureVolumePopulator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// AzureVolumePopulatorSpec specification of an Azure storage account resource
	Spec AzureVolumePopulatorSpec `json:"spec"`
}

// AzureVolumePopulatorSpec is the custom resource spec for provision a volume with data from Azure Blob Storage
type AzureVolumePopulatorSpec struct {
	// Copy all blobs with this prefix
	// +kubebuilder:validation:Required
	BlobPrefix string `json:"blobPrefix"`

	// Advanced settings to tune download behavior from Azure Blob Storage
	// +optional
	BlobDownloadOptions BlobDownloadOptions `json:"blobDownloadOptions"`

	// Destination path on the volume
	// +kubebuilder:validation:Required
	VolumePath string `json:"volumePath"`
}

// BlobDownloadOptions advanced settings to tune download behavior from Azure Blob Storage
// +kubebuilder:object:generate=true
type BlobDownloadOptions struct {
	// The maximum chunk size used for downloading a blob. Defaults to 4MiB
	// +optional
	// +kubebuilder:default=4194304
	BlockSize int `json:"blockSize"`

	// The maximum number of subtransfers that can be used in parallel. Defaults to 5.
	// +optional
	// +kubebuilder:default=5
	Concurrency int `json:"concurrency"`
}

// AzureVolumePopulatorList contains a list of AzureVolumePopulator
// +kubebuilder:object:root=true
type AzureVolumePopulatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureVolumePopulator `json:"items"`
}
