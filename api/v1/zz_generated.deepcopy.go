//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInfo) DeepCopyInto(out *NetworkInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInfo.
func (in *NetworkInfo) DeepCopy() *NetworkInfo {
	if in == nil {
		return nil
	}
	out := new(NetworkInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenSimulator) DeepCopyInto(out *OpenSimulator) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenSimulator.
func (in *OpenSimulator) DeepCopy() *OpenSimulator {
	if in == nil {
		return nil
	}
	out := new(OpenSimulator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenSimulator) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenSimulatorList) DeepCopyInto(out *OpenSimulatorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenSimulator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenSimulatorList.
func (in *OpenSimulatorList) DeepCopy() *OpenSimulatorList {
	if in == nil {
		return nil
	}
	out := new(OpenSimulatorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenSimulatorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenSimulatorSpec) DeepCopyInto(out *OpenSimulatorSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenSimulatorSpec.
func (in *OpenSimulatorSpec) DeepCopy() *OpenSimulatorSpec {
	if in == nil {
		return nil
	}
	out := new(OpenSimulatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenSimulatorStatus) DeepCopyInto(out *OpenSimulatorStatus) {
	*out = *in
	out.NetworkInfo = in.NetworkInfo
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenSimulatorStatus.
func (in *OpenSimulatorStatus) DeepCopy() *OpenSimulatorStatus {
	if in == nil {
		return nil
	}
	out := new(OpenSimulatorStatus)
	in.DeepCopyInto(out)
	return out
}
