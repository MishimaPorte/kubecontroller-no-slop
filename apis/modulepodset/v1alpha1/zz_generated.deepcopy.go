//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModulePodSet) DeepCopyInto(out *ModulePodSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModulePodSet.
func (in *ModulePodSet) DeepCopy() *ModulePodSet {
	if in == nil {
		return nil
	}
	out := new(ModulePodSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ModulePodSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModulePodSetList) DeepCopyInto(out *ModulePodSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ModulePodSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModulePodSetList.
func (in *ModulePodSetList) DeepCopy() *ModulePodSetList {
	if in == nil {
		return nil
	}
	out := new(ModulePodSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ModulePodSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModulePodSetSpec) DeepCopyInto(out *ModulePodSetSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModulePodSetSpec.
func (in *ModulePodSetSpec) DeepCopy() *ModulePodSetSpec {
	if in == nil {
		return nil
	}
	out := new(ModulePodSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModulePodSetStatus) DeepCopyInto(out *ModulePodSetStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModulePodSetStatus.
func (in *ModulePodSetStatus) DeepCopy() *ModulePodSetStatus {
	if in == nil {
		return nil
	}
	out := new(ModulePodSetStatus)
	in.DeepCopyInto(out)
	return out
}
