// +build !ignore_autogenerated

/*
Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package core

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupBucket) DeepCopyInto(out *BackupBucket) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupBucket.
func (in *BackupBucket) DeepCopy() *BackupBucket {
	if in == nil {
		return nil
	}
	out := new(BackupBucket)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupBucket) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupBucketList) DeepCopyInto(out *BackupBucketList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BackupBucket, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupBucketList.
func (in *BackupBucketList) DeepCopy() *BackupBucketList {
	if in == nil {
		return nil
	}
	out := new(BackupBucketList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupBucketList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupBucketProvider) DeepCopyInto(out *BackupBucketProvider) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupBucketProvider.
func (in *BackupBucketProvider) DeepCopy() *BackupBucketProvider {
	if in == nil {
		return nil
	}
	out := new(BackupBucketProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupBucketSpec) DeepCopyInto(out *BackupBucketSpec) {
	*out = *in
	out.Provider = in.Provider
	out.SecretRef = in.SecretRef
	if in.SeedName != nil {
		in, out := &in.SeedName, &out.SeedName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupBucketSpec.
func (in *BackupBucketSpec) DeepCopy() *BackupBucketSpec {
	if in == nil {
		return nil
	}
	out := new(BackupBucketSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupBucketStatus) DeepCopyInto(out *BackupBucketStatus) {
	*out = *in
	if in.LastOperation != nil {
		in, out := &in.LastOperation, &out.LastOperation
		*out = new(LastOperation)
		(*in).DeepCopyInto(*out)
	}
	if in.LastError != nil {
		in, out := &in.LastError, &out.LastError
		*out = new(LastError)
		(*in).DeepCopyInto(*out)
	}
	if in.GeneratedSecretRef != nil {
		in, out := &in.GeneratedSecretRef, &out.GeneratedSecretRef
		*out = new(v1.SecretReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupBucketStatus.
func (in *BackupBucketStatus) DeepCopy() *BackupBucketStatus {
	if in == nil {
		return nil
	}
	out := new(BackupBucketStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupEntry) DeepCopyInto(out *BackupEntry) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupEntry.
func (in *BackupEntry) DeepCopy() *BackupEntry {
	if in == nil {
		return nil
	}
	out := new(BackupEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupEntry) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupEntryList) DeepCopyInto(out *BackupEntryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BackupEntry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupEntryList.
func (in *BackupEntryList) DeepCopy() *BackupEntryList {
	if in == nil {
		return nil
	}
	out := new(BackupEntryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupEntryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupEntrySpec) DeepCopyInto(out *BackupEntrySpec) {
	*out = *in
	if in.SeedName != nil {
		in, out := &in.SeedName, &out.SeedName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupEntrySpec.
func (in *BackupEntrySpec) DeepCopy() *BackupEntrySpec {
	if in == nil {
		return nil
	}
	out := new(BackupEntrySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupEntryStatus) DeepCopyInto(out *BackupEntryStatus) {
	*out = *in
	if in.LastOperation != nil {
		in, out := &in.LastOperation, &out.LastOperation
		*out = new(LastOperation)
		(*in).DeepCopyInto(*out)
	}
	if in.LastError != nil {
		in, out := &in.LastError, &out.LastError
		*out = new(LastError)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupEntryStatus.
func (in *BackupEntryStatus) DeepCopy() *BackupEntryStatus {
	if in == nil {
		return nil
	}
	out := new(BackupEntryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudInfo) DeepCopyInto(out *CloudInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudInfo.
func (in *CloudInfo) DeepCopy() *CloudInfo {
	if in == nil {
		return nil
	}
	out := new(CloudInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterInfo) DeepCopyInto(out *ClusterInfo) {
	*out = *in
	out.Cloud = in.Cloud
	out.Kubernetes = in.Kubernetes
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterInfo.
func (in *ClusterInfo) DeepCopy() *ClusterInfo {
	if in == nil {
		return nil
	}
	out := new(ClusterInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerDeployment) DeepCopyInto(out *ControllerDeployment) {
	*out = *in
	if in.ProviderConfig != nil {
		in, out := &in.ProviderConfig, &out.ProviderConfig
		*out = new(ProviderConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerDeployment.
func (in *ControllerDeployment) DeepCopy() *ControllerDeployment {
	if in == nil {
		return nil
	}
	out := new(ControllerDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerInstallation) DeepCopyInto(out *ControllerInstallation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerInstallation.
func (in *ControllerInstallation) DeepCopy() *ControllerInstallation {
	if in == nil {
		return nil
	}
	out := new(ControllerInstallation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControllerInstallation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerInstallationList) DeepCopyInto(out *ControllerInstallationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ControllerInstallation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerInstallationList.
func (in *ControllerInstallationList) DeepCopy() *ControllerInstallationList {
	if in == nil {
		return nil
	}
	out := new(ControllerInstallationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControllerInstallationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerInstallationSpec) DeepCopyInto(out *ControllerInstallationSpec) {
	*out = *in
	out.RegistrationRef = in.RegistrationRef
	out.SeedRef = in.SeedRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerInstallationSpec.
func (in *ControllerInstallationSpec) DeepCopy() *ControllerInstallationSpec {
	if in == nil {
		return nil
	}
	out := new(ControllerInstallationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerInstallationStatus) DeepCopyInto(out *ControllerInstallationStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ProviderStatus != nil {
		in, out := &in.ProviderStatus, &out.ProviderStatus
		*out = new(ProviderConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerInstallationStatus.
func (in *ControllerInstallationStatus) DeepCopy() *ControllerInstallationStatus {
	if in == nil {
		return nil
	}
	out := new(ControllerInstallationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerRegistration) DeepCopyInto(out *ControllerRegistration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerRegistration.
func (in *ControllerRegistration) DeepCopy() *ControllerRegistration {
	if in == nil {
		return nil
	}
	out := new(ControllerRegistration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControllerRegistration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerRegistrationList) DeepCopyInto(out *ControllerRegistrationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ControllerRegistration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerRegistrationList.
func (in *ControllerRegistrationList) DeepCopy() *ControllerRegistrationList {
	if in == nil {
		return nil
	}
	out := new(ControllerRegistrationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControllerRegistrationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerRegistrationSpec) DeepCopyInto(out *ControllerRegistrationSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ControllerResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Deployment != nil {
		in, out := &in.Deployment, &out.Deployment
		*out = new(ControllerDeployment)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerRegistrationSpec.
func (in *ControllerRegistrationSpec) DeepCopy() *ControllerRegistrationSpec {
	if in == nil {
		return nil
	}
	out := new(ControllerRegistrationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerResource) DeepCopyInto(out *ControllerResource) {
	*out = *in
	if in.GloballyEnabled != nil {
		in, out := &in.GloballyEnabled, &out.GloballyEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ReconcileTimeout != nil {
		in, out := &in.ReconcileTimeout, &out.ReconcileTimeout
		*out = new(metav1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerResource.
func (in *ControllerResource) DeepCopy() *ControllerResource {
	if in == nil {
		return nil
	}
	out := new(ControllerResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Endpoint) DeepCopyInto(out *Endpoint) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Endpoint.
func (in *Endpoint) DeepCopy() *Endpoint {
	if in == nil {
		return nil
	}
	out := new(Endpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtensionResourceState) DeepCopyInto(out *ExtensionResourceState) {
	*out = *in
	if in.Purpose != nil {
		in, out := &in.Purpose, &out.Purpose
		*out = new(string)
		**out = **in
	}
	in.State.DeepCopyInto(&out.State)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtensionResourceState.
func (in *ExtensionResourceState) DeepCopy() *ExtensionResourceState {
	if in == nil {
		return nil
	}
	out := new(ExtensionResourceState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GardenerResourceData) DeepCopyInto(out *GardenerResourceData) {
	*out = *in
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GardenerResourceData.
func (in *GardenerResourceData) DeepCopy() *GardenerResourceData {
	if in == nil {
		return nil
	}
	out := new(GardenerResourceData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesInfo) DeepCopyInto(out *KubernetesInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesInfo.
func (in *KubernetesInfo) DeepCopy() *KubernetesInfo {
	if in == nil {
		return nil
	}
	out := new(KubernetesInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LastError) DeepCopyInto(out *LastError) {
	*out = *in
	if in.TaskID != nil {
		in, out := &in.TaskID, &out.TaskID
		*out = new(string)
		**out = **in
	}
	if in.Codes != nil {
		in, out := &in.Codes, &out.Codes
		*out = make([]ErrorCode, len(*in))
		copy(*out, *in)
	}
	if in.LastUpdateTime != nil {
		in, out := &in.LastUpdateTime, &out.LastUpdateTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LastError.
func (in *LastError) DeepCopy() *LastError {
	if in == nil {
		return nil
	}
	out := new(LastError)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LastOperation) DeepCopyInto(out *LastOperation) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LastOperation.
func (in *LastOperation) DeepCopy() *LastOperation {
	if in == nil {
		return nil
	}
	out := new(LastOperation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Plant) DeepCopyInto(out *Plant) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Plant.
func (in *Plant) DeepCopy() *Plant {
	if in == nil {
		return nil
	}
	out := new(Plant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Plant) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlantList) DeepCopyInto(out *PlantList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Plant, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlantList.
func (in *PlantList) DeepCopy() *PlantList {
	if in == nil {
		return nil
	}
	out := new(PlantList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PlantList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlantSpec) DeepCopyInto(out *PlantSpec) {
	*out = *in
	out.SecretRef = in.SecretRef
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]Endpoint, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlantSpec.
func (in *PlantSpec) DeepCopy() *PlantSpec {
	if in == nil {
		return nil
	}
	out := new(PlantSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlantStatus) DeepCopyInto(out *PlantStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.ClusterInfo != nil {
		in, out := &in.ClusterInfo, &out.ClusterInfo
		*out = new(ClusterInfo)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlantStatus.
func (in *PlantStatus) DeepCopy() *PlantStatus {
	if in == nil {
		return nil
	}
	out := new(PlantStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderConfig) DeepCopyInto(out *ProviderConfig) {
	*out = *in
	in.RawExtension.DeepCopyInto(&out.RawExtension)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderConfig.
func (in *ProviderConfig) DeepCopy() *ProviderConfig {
	if in == nil {
		return nil
	}
	out := new(ProviderConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootState) DeepCopyInto(out *ShootState) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootState.
func (in *ShootState) DeepCopy() *ShootState {
	if in == nil {
		return nil
	}
	out := new(ShootState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ShootState) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootStateList) DeepCopyInto(out *ShootStateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ShootState, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootStateList.
func (in *ShootStateList) DeepCopy() *ShootStateList {
	if in == nil {
		return nil
	}
	out := new(ShootStateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ShootStateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootStateSpec) DeepCopyInto(out *ShootStateSpec) {
	*out = *in
	if in.Gardener != nil {
		in, out := &in.Gardener, &out.Gardener
		*out = make([]GardenerResourceData, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Extensions != nil {
		in, out := &in.Extensions, &out.Extensions
		*out = make([]ExtensionResourceState, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootStateSpec.
func (in *ShootStateSpec) DeepCopy() *ShootStateSpec {
	if in == nil {
		return nil
	}
	out := new(ShootStateSpec)
	in.DeepCopyInto(out)
	return out
}
