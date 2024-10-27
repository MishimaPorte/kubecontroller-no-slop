// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	context "context"
	supermodulev1alpha1 "kek/apis/supermodule/v1alpha1"
	scheme "kek/generated/clientset/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// SuperModulesGetter has a method to return a SuperModuleInterface.
// A group's client should implement this interface.
type SuperModulesGetter interface {
	SuperModules(namespace string) SuperModuleInterface
}

// SuperModuleInterface has methods to work with SuperModule resources.
type SuperModuleInterface interface {
	Create(ctx context.Context, superModule *supermodulev1alpha1.SuperModule, opts v1.CreateOptions) (*supermodulev1alpha1.SuperModule, error)
	Update(ctx context.Context, superModule *supermodulev1alpha1.SuperModule, opts v1.UpdateOptions) (*supermodulev1alpha1.SuperModule, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, superModule *supermodulev1alpha1.SuperModule, opts v1.UpdateOptions) (*supermodulev1alpha1.SuperModule, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*supermodulev1alpha1.SuperModule, error)
	List(ctx context.Context, opts v1.ListOptions) (*supermodulev1alpha1.SuperModuleList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *supermodulev1alpha1.SuperModule, err error)
	SuperModuleExpansion
}

// superModules implements SuperModuleInterface
type superModules struct {
	*gentype.ClientWithList[*supermodulev1alpha1.SuperModule, *supermodulev1alpha1.SuperModuleList]
}

// newSuperModules returns a SuperModules
func newSuperModules(c *SpV1alpha1Client, namespace string) *superModules {
	return &superModules{
		gentype.NewClientWithList[*supermodulev1alpha1.SuperModule, *supermodulev1alpha1.SuperModuleList](
			"supermodules",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *supermodulev1alpha1.SuperModule { return &supermodulev1alpha1.SuperModule{} },
			func() *supermodulev1alpha1.SuperModuleList { return &supermodulev1alpha1.SuperModuleList{} },
		),
	}
}