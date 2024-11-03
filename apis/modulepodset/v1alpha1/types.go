package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Supermodule is a module with batteries (pod-controller, i mean)
// included.
type ModulePodSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   ModulePodSetSpec   `json:"spec"`
	Status ModulePodSetStatus `json:"status"`
}

type ReplicaAmount int32

type ModulePodSetStatus struct {
	ReadyReplicas    ReplicaAmount `json:"ready"`
	FaultyReplicas   ReplicaAmount `json:"error"`
	StartingReplicas ReplicaAmount `json:"inProgress"`
}

type ModulePodSetSpec struct {
	// 2^31 replicas ought to be enough for everybody (C) steve gates probably
	Replicas ReplicaAmount `json:"replicas"`
	Template corev1.PodSpec
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ModulePodSetList is a list of supermodules
type ModulePodSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ModulePodSet `json:"items"`
}
