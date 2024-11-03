// Code generated by gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "kek/generated/clientset/typed/modulepodsets/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeSpV1alpha1 struct {
	*testing.Fake
}

func (c *FakeSpV1alpha1) ModulePodSets(namespace string) v1alpha1.ModulePodSetInterface {
	return &FakeModulePodSets{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeSpV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}