//go:build !ignore_autogenerated

/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Alchemist) DeepCopyInto(out *Alchemist) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Alchemist.
func (in *Alchemist) DeepCopy() *Alchemist {
	if in == nil {
		return nil
	}
	out := new(Alchemist)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Alchemist) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlchemistList) DeepCopyInto(out *AlchemistList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Alchemist, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlchemistList.
func (in *AlchemistList) DeepCopy() *AlchemistList {
	if in == nil {
		return nil
	}
	out := new(AlchemistList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlchemistList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlchemistSpec) DeepCopyInto(out *AlchemistSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.ContainerPort != nil {
		in, out := &in.ContainerPort, &out.ContainerPort
		*out = new(int32)
		**out = **in
	}
	if in.ServicePort != nil {
		in, out := &in.ServicePort, &out.ServicePort
		*out = new(int32)
		**out = **in
	}
	if in.TargetPort != nil {
		in, out := &in.TargetPort, &out.TargetPort
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlchemistSpec.
func (in *AlchemistSpec) DeepCopy() *AlchemistSpec {
	if in == nil {
		return nil
	}
	out := new(AlchemistSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlchemistStatus) DeepCopyInto(out *AlchemistStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlchemistStatus.
func (in *AlchemistStatus) DeepCopy() *AlchemistStatus {
	if in == nil {
		return nil
	}
	out := new(AlchemistStatus)
	in.DeepCopyInto(out)
	return out
}
