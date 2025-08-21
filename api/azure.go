package api

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type AzureVolumePopulator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec AzureVolumePopulatorSpec `json:"spec"`
}

type AzureVolumePopulatorSpec struct {
	BlobPrefix string `json:"blobPrefix"`
	VolumePath string `json:"volumePath"`
}
