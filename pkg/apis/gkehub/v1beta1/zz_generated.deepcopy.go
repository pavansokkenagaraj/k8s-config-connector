// +build !ignore_autogenerated

// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package v1beta1

import (
	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureMulticlusteringress) DeepCopyInto(out *FeatureMulticlusteringress) {
	*out = *in
	out.ConfigMembershipRef = in.ConfigMembershipRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureMulticlusteringress.
func (in *FeatureMulticlusteringress) DeepCopy() *FeatureMulticlusteringress {
	if in == nil {
		return nil
	}
	out := new(FeatureMulticlusteringress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureResourceStateStatus) DeepCopyInto(out *FeatureResourceStateStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureResourceStateStatus.
func (in *FeatureResourceStateStatus) DeepCopy() *FeatureResourceStateStatus {
	if in == nil {
		return nil
	}
	out := new(FeatureResourceStateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureSpec) DeepCopyInto(out *FeatureSpec) {
	*out = *in
	if in.Multiclusteringress != nil {
		in, out := &in.Multiclusteringress, &out.Multiclusteringress
		*out = new(FeatureMulticlusteringress)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureSpec.
func (in *FeatureSpec) DeepCopy() *FeatureSpec {
	if in == nil {
		return nil
	}
	out := new(FeatureSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureStateStatus) DeepCopyInto(out *FeatureStateStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureStateStatus.
func (in *FeatureStateStatus) DeepCopy() *FeatureStateStatus {
	if in == nil {
		return nil
	}
	out := new(FeatureStateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GKEHubFeature) DeepCopyInto(out *GKEHubFeature) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GKEHubFeature.
func (in *GKEHubFeature) DeepCopy() *GKEHubFeature {
	if in == nil {
		return nil
	}
	out := new(GKEHubFeature)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GKEHubFeature) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GKEHubFeatureList) DeepCopyInto(out *GKEHubFeatureList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GKEHubFeature, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GKEHubFeatureList.
func (in *GKEHubFeatureList) DeepCopy() *GKEHubFeatureList {
	if in == nil {
		return nil
	}
	out := new(GKEHubFeatureList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GKEHubFeatureList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GKEHubFeatureSpec) DeepCopyInto(out *GKEHubFeatureSpec) {
	*out = *in
	out.ProjectRef = in.ProjectRef
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(FeatureSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GKEHubFeatureSpec.
func (in *GKEHubFeatureSpec) DeepCopy() *GKEHubFeatureSpec {
	if in == nil {
		return nil
	}
	out := new(GKEHubFeatureSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GKEHubFeatureStatus) DeepCopyInto(out *GKEHubFeatureStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	out.ResourceState = in.ResourceState
	out.State = in.State
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GKEHubFeatureStatus.
func (in *GKEHubFeatureStatus) DeepCopy() *GKEHubFeatureStatus {
	if in == nil {
		return nil
	}
	out := new(GKEHubFeatureStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GKEHubMembership) DeepCopyInto(out *GKEHubMembership) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GKEHubMembership.
func (in *GKEHubMembership) DeepCopy() *GKEHubMembership {
	if in == nil {
		return nil
	}
	out := new(GKEHubMembership)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GKEHubMembership) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GKEHubMembershipList) DeepCopyInto(out *GKEHubMembershipList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GKEHubMembership, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GKEHubMembershipList.
func (in *GKEHubMembershipList) DeepCopy() *GKEHubMembershipList {
	if in == nil {
		return nil
	}
	out := new(GKEHubMembershipList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GKEHubMembershipList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GKEHubMembershipSpec) DeepCopyInto(out *GKEHubMembershipSpec) {
	*out = *in
	if in.Authority != nil {
		in, out := &in.Authority, &out.Authority
		*out = new(MembershipAuthority)
		(*in).DeepCopyInto(*out)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(MembershipEndpoint)
		(*in).DeepCopyInto(*out)
	}
	if in.ExternalId != nil {
		in, out := &in.ExternalId, &out.ExternalId
		*out = new(string)
		**out = **in
	}
	if in.InfrastructureType != nil {
		in, out := &in.InfrastructureType, &out.InfrastructureType
		*out = new(string)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GKEHubMembershipSpec.
func (in *GKEHubMembershipSpec) DeepCopy() *GKEHubMembershipSpec {
	if in == nil {
		return nil
	}
	out := new(GKEHubMembershipSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GKEHubMembershipStatus) DeepCopyInto(out *GKEHubMembershipStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	out.Authority = in.Authority
	in.Endpoint.DeepCopyInto(&out.Endpoint)
	out.State = in.State
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GKEHubMembershipStatus.
func (in *GKEHubMembershipStatus) DeepCopy() *GKEHubMembershipStatus {
	if in == nil {
		return nil
	}
	out := new(GKEHubMembershipStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MembershipAuthority) DeepCopyInto(out *MembershipAuthority) {
	*out = *in
	if in.Issuer != nil {
		in, out := &in.Issuer, &out.Issuer
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MembershipAuthority.
func (in *MembershipAuthority) DeepCopy() *MembershipAuthority {
	if in == nil {
		return nil
	}
	out := new(MembershipAuthority)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MembershipAuthorityStatus) DeepCopyInto(out *MembershipAuthorityStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MembershipAuthorityStatus.
func (in *MembershipAuthorityStatus) DeepCopy() *MembershipAuthorityStatus {
	if in == nil {
		return nil
	}
	out := new(MembershipAuthorityStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MembershipConnectResourcesStatus) DeepCopyInto(out *MembershipConnectResourcesStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MembershipConnectResourcesStatus.
func (in *MembershipConnectResourcesStatus) DeepCopy() *MembershipConnectResourcesStatus {
	if in == nil {
		return nil
	}
	out := new(MembershipConnectResourcesStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MembershipEndpoint) DeepCopyInto(out *MembershipEndpoint) {
	*out = *in
	if in.GkeCluster != nil {
		in, out := &in.GkeCluster, &out.GkeCluster
		*out = new(MembershipGkeCluster)
		(*in).DeepCopyInto(*out)
	}
	if in.KubernetesResource != nil {
		in, out := &in.KubernetesResource, &out.KubernetesResource
		*out = new(MembershipKubernetesResource)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MembershipEndpoint.
func (in *MembershipEndpoint) DeepCopy() *MembershipEndpoint {
	if in == nil {
		return nil
	}
	out := new(MembershipEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MembershipEndpointStatus) DeepCopyInto(out *MembershipEndpointStatus) {
	*out = *in
	out.KubernetesMetadata = in.KubernetesMetadata
	in.KubernetesResource.DeepCopyInto(&out.KubernetesResource)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MembershipEndpointStatus.
func (in *MembershipEndpointStatus) DeepCopy() *MembershipEndpointStatus {
	if in == nil {
		return nil
	}
	out := new(MembershipEndpointStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MembershipGkeCluster) DeepCopyInto(out *MembershipGkeCluster) {
	*out = *in
	if in.ResourceRef != nil {
		in, out := &in.ResourceRef, &out.ResourceRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MembershipGkeCluster.
func (in *MembershipGkeCluster) DeepCopy() *MembershipGkeCluster {
	if in == nil {
		return nil
	}
	out := new(MembershipGkeCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MembershipKubernetesMetadataStatus) DeepCopyInto(out *MembershipKubernetesMetadataStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MembershipKubernetesMetadataStatus.
func (in *MembershipKubernetesMetadataStatus) DeepCopy() *MembershipKubernetesMetadataStatus {
	if in == nil {
		return nil
	}
	out := new(MembershipKubernetesMetadataStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MembershipKubernetesResource) DeepCopyInto(out *MembershipKubernetesResource) {
	*out = *in
	if in.MembershipCrManifest != nil {
		in, out := &in.MembershipCrManifest, &out.MembershipCrManifest
		*out = new(string)
		**out = **in
	}
	if in.ResourceOptions != nil {
		in, out := &in.ResourceOptions, &out.ResourceOptions
		*out = new(MembershipResourceOptions)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MembershipKubernetesResource.
func (in *MembershipKubernetesResource) DeepCopy() *MembershipKubernetesResource {
	if in == nil {
		return nil
	}
	out := new(MembershipKubernetesResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MembershipKubernetesResourceStatus) DeepCopyInto(out *MembershipKubernetesResourceStatus) {
	*out = *in
	if in.ConnectResources != nil {
		in, out := &in.ConnectResources, &out.ConnectResources
		*out = make([]MembershipConnectResourcesStatus, len(*in))
		copy(*out, *in)
	}
	if in.MembershipResources != nil {
		in, out := &in.MembershipResources, &out.MembershipResources
		*out = make([]MembershipMembershipResourcesStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MembershipKubernetesResourceStatus.
func (in *MembershipKubernetesResourceStatus) DeepCopy() *MembershipKubernetesResourceStatus {
	if in == nil {
		return nil
	}
	out := new(MembershipKubernetesResourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MembershipMembershipResourcesStatus) DeepCopyInto(out *MembershipMembershipResourcesStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MembershipMembershipResourcesStatus.
func (in *MembershipMembershipResourcesStatus) DeepCopy() *MembershipMembershipResourcesStatus {
	if in == nil {
		return nil
	}
	out := new(MembershipMembershipResourcesStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MembershipResourceOptions) DeepCopyInto(out *MembershipResourceOptions) {
	*out = *in
	if in.ConnectVersion != nil {
		in, out := &in.ConnectVersion, &out.ConnectVersion
		*out = new(string)
		**out = **in
	}
	if in.V1beta1Crd != nil {
		in, out := &in.V1beta1Crd, &out.V1beta1Crd
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MembershipResourceOptions.
func (in *MembershipResourceOptions) DeepCopy() *MembershipResourceOptions {
	if in == nil {
		return nil
	}
	out := new(MembershipResourceOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MembershipStateStatus) DeepCopyInto(out *MembershipStateStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MembershipStateStatus.
func (in *MembershipStateStatus) DeepCopy() *MembershipStateStatus {
	if in == nil {
		return nil
	}
	out := new(MembershipStateStatus)
	in.DeepCopyInto(out)
	return out
}
