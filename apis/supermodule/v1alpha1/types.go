package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Supermodule is a module with batteries (pod-controller, i mean)
// included.
type SuperModule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   SuperModuleSpec   `json:"spec"`
	Status SuperModuleStatus `json:"status"`
}

type ReplicaAmount int32

type SuperModuleStatus struct {
	ReadyReplicas    ReplicaAmount `json:"readyReplicas"`
	FaultyReplicas   ReplicaAmount `json:"faultyReplicas"`
	StartingReplicas ReplicaAmount `json:"startingReplicas"`
}

type SubThing struct {
	StringThing string `json:"stringThing"`
}
type SuperModuleSpec struct {
	// 2^31 replicas ought to be enough for everybody (C) steve gates probably
	Replicas ReplicaAmount `json:"replicas"`
	SubThing SubThing      `json:"subThing"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SuperModuleList is a list of supermodules
type SuperModuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []SuperModule `json:"items"`
}
