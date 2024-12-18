// Code generated by gen. DO NOT EDIT.

package v1alpha1

import (
	modulepodsetv1alpha1 "kek/apis/modulepodset/v1alpha1"

	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// ModulePodSetLister helps list ModulePodSets.
// All objects returned here must be treated as read-only.
type ModulePodSetLister interface {
	// List lists all ModulePodSets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*modulepodsetv1alpha1.ModulePodSet, err error)
	// ModulePodSets returns an object that can list and get ModulePodSets.
	ModulePodSets(namespace string) ModulePodSetNamespaceLister
	ModulePodSetListerExpansion
}

// modulePodSetLister implements the ModulePodSetLister interface.
type modulePodSetLister struct {
	listers.ResourceIndexer[*modulepodsetv1alpha1.ModulePodSet]
}

// NewModulePodSetLister returns a new ModulePodSetLister.
func NewModulePodSetLister(indexer cache.Indexer) ModulePodSetLister {
	return &modulePodSetLister{listers.New[*modulepodsetv1alpha1.ModulePodSet](indexer, modulepodsetv1alpha1.Resource("modulepodset"))}
}

// ModulePodSets returns an object that can list and get ModulePodSets.
func (s *modulePodSetLister) ModulePodSets(namespace string) ModulePodSetNamespaceLister {
	return modulePodSetNamespaceLister{listers.NewNamespaced[*modulepodsetv1alpha1.ModulePodSet](s.ResourceIndexer, namespace)}
}

// ModulePodSetNamespaceLister helps list and get ModulePodSets.
// All objects returned here must be treated as read-only.
type ModulePodSetNamespaceLister interface {
	// List lists all ModulePodSets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*modulepodsetv1alpha1.ModulePodSet, err error)
	// Get retrieves the ModulePodSet from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*modulepodsetv1alpha1.ModulePodSet, error)
	ModulePodSetNamespaceListerExpansion
}

// modulePodSetNamespaceLister implements the ModulePodSetNamespaceLister
// interface.
type modulePodSetNamespaceLister struct {
	listers.ResourceIndexer[*modulepodsetv1alpha1.ModulePodSet]
}
