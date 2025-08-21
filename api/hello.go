package api

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type Hello struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec HelloSpec `json:"spec"`
}

type HelloSpec struct {
	FileName     string `json:"fileName"`
	FileContents string `json:"fileContents"`
}
