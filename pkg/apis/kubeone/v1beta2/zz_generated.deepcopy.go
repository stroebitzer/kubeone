//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KubeOne Authors.

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

package v1beta2

import (
	json "encoding/json"

	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIEndpoint) DeepCopyInto(out *APIEndpoint) {
	*out = *in
	if in.AlternativeNames != nil {
		in, out := &in.AlternativeNames, &out.AlternativeNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIEndpoint.
func (in *APIEndpoint) DeepCopy() *APIEndpoint {
	if in == nil {
		return nil
	}
	out := new(APIEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSSpec) DeepCopyInto(out *AWSSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSSpec.
func (in *AWSSpec) DeepCopy() *AWSSpec {
	if in == nil {
		return nil
	}
	out := new(AWSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Addon) DeepCopyInto(out *Addon) {
	*out = *in
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Addon.
func (in *Addon) DeepCopy() *Addon {
	if in == nil {
		return nil
	}
	out := new(Addon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Addons) DeepCopyInto(out *Addons) {
	*out = *in
	if in.GlobalParams != nil {
		in, out := &in.GlobalParams, &out.GlobalParams
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Addons != nil {
		in, out := &in.Addons, &out.Addons
		*out = make([]Addon, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Addons.
func (in *Addons) DeepCopy() *Addons {
	if in == nil {
		return nil
	}
	out := new(Addons)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureSpec) DeepCopyInto(out *AzureSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureSpec.
func (in *AzureSpec) DeepCopy() *AzureSpec {
	if in == nil {
		return nil
	}
	out := new(AzureSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BinaryAsset) DeepCopyInto(out *BinaryAsset) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BinaryAsset.
func (in *BinaryAsset) DeepCopy() *BinaryAsset {
	if in == nil {
		return nil
	}
	out := new(BinaryAsset)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CNI) DeepCopyInto(out *CNI) {
	*out = *in
	if in.Canal != nil {
		in, out := &in.Canal, &out.Canal
		*out = new(CanalSpec)
		**out = **in
	}
	if in.Cilium != nil {
		in, out := &in.Cilium, &out.Cilium
		*out = new(CiliumSpec)
		**out = **in
	}
	if in.WeaveNet != nil {
		in, out := &in.WeaveNet, &out.WeaveNet
		*out = new(WeaveNetSpec)
		**out = **in
	}
	if in.External != nil {
		in, out := &in.External, &out.External
		*out = new(ExternalCNISpec)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CNI.
func (in *CNI) DeepCopy() *CNI {
	if in == nil {
		return nil
	}
	out := new(CNI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanalSpec) DeepCopyInto(out *CanalSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanalSpec.
func (in *CanalSpec) DeepCopy() *CanalSpec {
	if in == nil {
		return nil
	}
	out := new(CanalSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CiliumSpec) DeepCopyInto(out *CiliumSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CiliumSpec.
func (in *CiliumSpec) DeepCopy() *CiliumSpec {
	if in == nil {
		return nil
	}
	out := new(CiliumSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudProviderSpec) DeepCopyInto(out *CloudProviderSpec) {
	*out = *in
	if in.AWS != nil {
		in, out := &in.AWS, &out.AWS
		*out = new(AWSSpec)
		**out = **in
	}
	if in.Azure != nil {
		in, out := &in.Azure, &out.Azure
		*out = new(AzureSpec)
		**out = **in
	}
	if in.DigitalOcean != nil {
		in, out := &in.DigitalOcean, &out.DigitalOcean
		*out = new(DigitalOceanSpec)
		**out = **in
	}
	if in.GCE != nil {
		in, out := &in.GCE, &out.GCE
		*out = new(GCESpec)
		**out = **in
	}
	if in.Hetzner != nil {
		in, out := &in.Hetzner, &out.Hetzner
		*out = new(HetznerSpec)
		**out = **in
	}
	if in.Openstack != nil {
		in, out := &in.Openstack, &out.Openstack
		*out = new(OpenstackSpec)
		**out = **in
	}
	if in.Packet != nil {
		in, out := &in.Packet, &out.Packet
		*out = new(PacketSpec)
		**out = **in
	}
	if in.Vsphere != nil {
		in, out := &in.Vsphere, &out.Vsphere
		*out = new(VsphereSpec)
		**out = **in
	}
	if in.None != nil {
		in, out := &in.None, &out.None
		*out = new(NoneSpec)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudProviderSpec.
func (in *CloudProviderSpec) DeepCopy() *CloudProviderSpec {
	if in == nil {
		return nil
	}
	out := new(CloudProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterNetworkConfig) DeepCopyInto(out *ClusterNetworkConfig) {
	*out = *in
	if in.CNI != nil {
		in, out := &in.CNI, &out.CNI
		*out = new(CNI)
		(*in).DeepCopyInto(*out)
	}
	if in.KubeProxy != nil {
		in, out := &in.KubeProxy, &out.KubeProxy
		*out = new(KubeProxyConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterNetworkConfig.
func (in *ClusterNetworkConfig) DeepCopy() *ClusterNetworkConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterNetworkConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerRuntimeConfig) DeepCopyInto(out *ContainerRuntimeConfig) {
	*out = *in
	if in.Docker != nil {
		in, out := &in.Docker, &out.Docker
		*out = new(ContainerRuntimeDocker)
		(*in).DeepCopyInto(*out)
	}
	if in.Containerd != nil {
		in, out := &in.Containerd, &out.Containerd
		*out = new(ContainerRuntimeContainerd)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerRuntimeConfig.
func (in *ContainerRuntimeConfig) DeepCopy() *ContainerRuntimeConfig {
	if in == nil {
		return nil
	}
	out := new(ContainerRuntimeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerRuntimeContainerd) DeepCopyInto(out *ContainerRuntimeContainerd) {
	*out = *in
	if in.Registries != nil {
		in, out := &in.Registries, &out.Registries
		*out = make(map[string]ContainerdRegistry, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerRuntimeContainerd.
func (in *ContainerRuntimeContainerd) DeepCopy() *ContainerRuntimeContainerd {
	if in == nil {
		return nil
	}
	out := new(ContainerRuntimeContainerd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerRuntimeDocker) DeepCopyInto(out *ContainerRuntimeDocker) {
	*out = *in
	if in.RegistryMirrors != nil {
		in, out := &in.RegistryMirrors, &out.RegistryMirrors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerRuntimeDocker.
func (in *ContainerRuntimeDocker) DeepCopy() *ContainerRuntimeDocker {
	if in == nil {
		return nil
	}
	out := new(ContainerRuntimeDocker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerdRegistry) DeepCopyInto(out *ContainerdRegistry) {
	*out = *in
	if in.Mirrors != nil {
		in, out := &in.Mirrors, &out.Mirrors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TLSConfig != nil {
		in, out := &in.TLSConfig, &out.TLSConfig
		*out = new(ContainerdTLSConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerdRegistry.
func (in *ContainerdRegistry) DeepCopy() *ContainerdRegistry {
	if in == nil {
		return nil
	}
	out := new(ContainerdRegistry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerdTLSConfig) DeepCopyInto(out *ContainerdTLSConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerdTLSConfig.
func (in *ContainerdTLSConfig) DeepCopy() *ContainerdTLSConfig {
	if in == nil {
		return nil
	}
	out := new(ContainerdTLSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlaneConfig) DeepCopyInto(out *ControlPlaneConfig) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]HostConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlaneConfig.
func (in *ControlPlaneConfig) DeepCopy() *ControlPlaneConfig {
	if in == nil {
		return nil
	}
	out := new(ControlPlaneConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSConfig) DeepCopyInto(out *DNSConfig) {
	*out = *in
	if in.Servers != nil {
		in, out := &in.Servers, &out.Servers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSConfig.
func (in *DNSConfig) DeepCopy() *DNSConfig {
	if in == nil {
		return nil
	}
	out := new(DNSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DigitalOceanSpec) DeepCopyInto(out *DigitalOceanSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DigitalOceanSpec.
func (in *DigitalOceanSpec) DeepCopy() *DigitalOceanSpec {
	if in == nil {
		return nil
	}
	out := new(DigitalOceanSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamicAuditLog) DeepCopyInto(out *DynamicAuditLog) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamicAuditLog.
func (in *DynamicAuditLog) DeepCopy() *DynamicAuditLog {
	if in == nil {
		return nil
	}
	out := new(DynamicAuditLog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamicWorkerConfig) DeepCopyInto(out *DynamicWorkerConfig) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int)
		**out = **in
	}
	in.Config.DeepCopyInto(&out.Config)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamicWorkerConfig.
func (in *DynamicWorkerConfig) DeepCopy() *DynamicWorkerConfig {
	if in == nil {
		return nil
	}
	out := new(DynamicWorkerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionProviders) DeepCopyInto(out *EncryptionProviders) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionProviders.
func (in *EncryptionProviders) DeepCopy() *EncryptionProviders {
	if in == nil {
		return nil
	}
	out := new(EncryptionProviders)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalCNISpec) DeepCopyInto(out *ExternalCNISpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalCNISpec.
func (in *ExternalCNISpec) DeepCopy() *ExternalCNISpec {
	if in == nil {
		return nil
	}
	out := new(ExternalCNISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Features) DeepCopyInto(out *Features) {
	*out = *in
	if in.PodNodeSelector != nil {
		in, out := &in.PodNodeSelector, &out.PodNodeSelector
		*out = new(PodNodeSelector)
		**out = **in
	}
	if in.PodSecurityPolicy != nil {
		in, out := &in.PodSecurityPolicy, &out.PodSecurityPolicy
		*out = new(PodSecurityPolicy)
		**out = **in
	}
	if in.StaticAuditLog != nil {
		in, out := &in.StaticAuditLog, &out.StaticAuditLog
		*out = new(StaticAuditLog)
		**out = **in
	}
	if in.DynamicAuditLog != nil {
		in, out := &in.DynamicAuditLog, &out.DynamicAuditLog
		*out = new(DynamicAuditLog)
		**out = **in
	}
	if in.MetricsServer != nil {
		in, out := &in.MetricsServer, &out.MetricsServer
		*out = new(MetricsServer)
		**out = **in
	}
	if in.OpenIDConnect != nil {
		in, out := &in.OpenIDConnect, &out.OpenIDConnect
		*out = new(OpenIDConnect)
		**out = **in
	}
	if in.EncryptionProviders != nil {
		in, out := &in.EncryptionProviders, &out.EncryptionProviders
		*out = new(EncryptionProviders)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Features.
func (in *Features) DeepCopy() *Features {
	if in == nil {
		return nil
	}
	out := new(Features)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCESpec) DeepCopyInto(out *GCESpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCESpec.
func (in *GCESpec) DeepCopy() *GCESpec {
	if in == nil {
		return nil
	}
	out := new(GCESpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HetznerSpec) DeepCopyInto(out *HetznerSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HetznerSpec.
func (in *HetznerSpec) DeepCopy() *HetznerSpec {
	if in == nil {
		return nil
	}
	out := new(HetznerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostConfig) DeepCopyInto(out *HostConfig) {
	*out = *in
	if in.Taints != nil {
		in, out := &in.Taints, &out.Taints
		*out = make([]v1.Taint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Kubelet.DeepCopyInto(&out.Kubelet)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostConfig.
func (in *HostConfig) DeepCopy() *HostConfig {
	if in == nil {
		return nil
	}
	out := new(HostConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPTables) DeepCopyInto(out *IPTables) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPTables.
func (in *IPTables) DeepCopy() *IPTables {
	if in == nil {
		return nil
	}
	out := new(IPTables)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPVSConfig) DeepCopyInto(out *IPVSConfig) {
	*out = *in
	if in.ExcludeCIDRs != nil {
		in, out := &in.ExcludeCIDRs, &out.ExcludeCIDRs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.TCPTimeout = in.TCPTimeout
	out.TCPFinTimeout = in.TCPFinTimeout
	out.UDPTimeout = in.UDPTimeout
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPVSConfig.
func (in *IPVSConfig) DeepCopy() *IPVSConfig {
	if in == nil {
		return nil
	}
	out := new(IPVSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageAsset) DeepCopyInto(out *ImageAsset) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageAsset.
func (in *ImageAsset) DeepCopy() *ImageAsset {
	if in == nil {
		return nil
	}
	out := new(ImageAsset)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeOneCluster) DeepCopyInto(out *KubeOneCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ControlPlane.DeepCopyInto(&out.ControlPlane)
	in.APIEndpoint.DeepCopyInto(&out.APIEndpoint)
	in.CloudProvider.DeepCopyInto(&out.CloudProvider)
	out.Versions = in.Versions
	in.ContainerRuntime.DeepCopyInto(&out.ContainerRuntime)
	in.ClusterNetwork.DeepCopyInto(&out.ClusterNetwork)
	out.Proxy = in.Proxy
	in.StaticWorkers.DeepCopyInto(&out.StaticWorkers)
	if in.DynamicWorkers != nil {
		in, out := &in.DynamicWorkers, &out.DynamicWorkers
		*out = make([]DynamicWorkerConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MachineController != nil {
		in, out := &in.MachineController, &out.MachineController
		*out = new(MachineControllerConfig)
		**out = **in
	}
	in.Features.DeepCopyInto(&out.Features)
	if in.Addons != nil {
		in, out := &in.Addons, &out.Addons
		*out = new(Addons)
		(*in).DeepCopyInto(*out)
	}
	if in.SystemPackages != nil {
		in, out := &in.SystemPackages, &out.SystemPackages
		*out = new(SystemPackages)
		**out = **in
	}
	if in.RegistryConfiguration != nil {
		in, out := &in.RegistryConfiguration, &out.RegistryConfiguration
		*out = new(RegistryConfiguration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeOneCluster.
func (in *KubeOneCluster) DeepCopy() *KubeOneCluster {
	if in == nil {
		return nil
	}
	out := new(KubeOneCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeOneCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeProxyConfig) DeepCopyInto(out *KubeProxyConfig) {
	*out = *in
	if in.IPVS != nil {
		in, out := &in.IPVS, &out.IPVS
		*out = new(IPVSConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.IPTables != nil {
		in, out := &in.IPTables, &out.IPTables
		*out = new(IPTables)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeProxyConfig.
func (in *KubeProxyConfig) DeepCopy() *KubeProxyConfig {
	if in == nil {
		return nil
	}
	out := new(KubeProxyConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletConfig) DeepCopyInto(out *KubeletConfig) {
	*out = *in
	if in.SystemReserved != nil {
		in, out := &in.SystemReserved, &out.SystemReserved
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.KubeReserved != nil {
		in, out := &in.KubeReserved, &out.KubeReserved
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EvictionHard != nil {
		in, out := &in.EvictionHard, &out.EvictionHard
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletConfig.
func (in *KubeletConfig) DeepCopy() *KubeletConfig {
	if in == nil {
		return nil
	}
	out := new(KubeletConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineControllerConfig) DeepCopyInto(out *MachineControllerConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineControllerConfig.
func (in *MachineControllerConfig) DeepCopy() *MachineControllerConfig {
	if in == nil {
		return nil
	}
	out := new(MachineControllerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsServer) DeepCopyInto(out *MetricsServer) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsServer.
func (in *MetricsServer) DeepCopy() *MetricsServer {
	if in == nil {
		return nil
	}
	out := new(MetricsServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NoneSpec) DeepCopyInto(out *NoneSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NoneSpec.
func (in *NoneSpec) DeepCopy() *NoneSpec {
	if in == nil {
		return nil
	}
	out := new(NoneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenIDConnect) DeepCopyInto(out *OpenIDConnect) {
	*out = *in
	out.Config = in.Config
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenIDConnect.
func (in *OpenIDConnect) DeepCopy() *OpenIDConnect {
	if in == nil {
		return nil
	}
	out := new(OpenIDConnect)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenIDConnectConfig) DeepCopyInto(out *OpenIDConnectConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenIDConnectConfig.
func (in *OpenIDConnectConfig) DeepCopy() *OpenIDConnectConfig {
	if in == nil {
		return nil
	}
	out := new(OpenIDConnectConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenstackSpec) DeepCopyInto(out *OpenstackSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenstackSpec.
func (in *OpenstackSpec) DeepCopy() *OpenstackSpec {
	if in == nil {
		return nil
	}
	out := new(OpenstackSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PacketSpec) DeepCopyInto(out *PacketSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PacketSpec.
func (in *PacketSpec) DeepCopy() *PacketSpec {
	if in == nil {
		return nil
	}
	out := new(PacketSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodNodeSelector) DeepCopyInto(out *PodNodeSelector) {
	*out = *in
	out.Config = in.Config
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodNodeSelector.
func (in *PodNodeSelector) DeepCopy() *PodNodeSelector {
	if in == nil {
		return nil
	}
	out := new(PodNodeSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodNodeSelectorConfig) DeepCopyInto(out *PodNodeSelectorConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodNodeSelectorConfig.
func (in *PodNodeSelectorConfig) DeepCopy() *PodNodeSelectorConfig {
	if in == nil {
		return nil
	}
	out := new(PodNodeSelectorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSecurityPolicy) DeepCopyInto(out *PodSecurityPolicy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSecurityPolicy.
func (in *PodSecurityPolicy) DeepCopy() *PodSecurityPolicy {
	if in == nil {
		return nil
	}
	out := new(PodSecurityPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderSpec) DeepCopyInto(out *ProviderSpec) {
	*out = *in
	if in.CloudProviderSpec != nil {
		in, out := &in.CloudProviderSpec, &out.CloudProviderSpec
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.MachineAnnotations != nil {
		in, out := &in.MachineAnnotations, &out.MachineAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Taints != nil {
		in, out := &in.Taints, &out.Taints
		*out = make([]v1.Taint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SSHPublicKeys != nil {
		in, out := &in.SSHPublicKeys, &out.SSHPublicKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.OperatingSystemSpec != nil {
		in, out := &in.OperatingSystemSpec, &out.OperatingSystemSpec
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(ProviderStaticNetworkConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.OverwriteCloudConfig != nil {
		in, out := &in.OverwriteCloudConfig, &out.OverwriteCloudConfig
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderSpec.
func (in *ProviderSpec) DeepCopy() *ProviderSpec {
	if in == nil {
		return nil
	}
	out := new(ProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderStaticNetworkConfig) DeepCopyInto(out *ProviderStaticNetworkConfig) {
	*out = *in
	in.DNS.DeepCopyInto(&out.DNS)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderStaticNetworkConfig.
func (in *ProviderStaticNetworkConfig) DeepCopy() *ProviderStaticNetworkConfig {
	if in == nil {
		return nil
	}
	out := new(ProviderStaticNetworkConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxyConfig) DeepCopyInto(out *ProxyConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxyConfig.
func (in *ProxyConfig) DeepCopy() *ProxyConfig {
	if in == nil {
		return nil
	}
	out := new(ProxyConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryConfiguration) DeepCopyInto(out *RegistryConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryConfiguration.
func (in *RegistryConfiguration) DeepCopy() *RegistryConfiguration {
	if in == nil {
		return nil
	}
	out := new(RegistryConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StaticAuditLog) DeepCopyInto(out *StaticAuditLog) {
	*out = *in
	out.Config = in.Config
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StaticAuditLog.
func (in *StaticAuditLog) DeepCopy() *StaticAuditLog {
	if in == nil {
		return nil
	}
	out := new(StaticAuditLog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StaticAuditLogConfig) DeepCopyInto(out *StaticAuditLogConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StaticAuditLogConfig.
func (in *StaticAuditLogConfig) DeepCopy() *StaticAuditLogConfig {
	if in == nil {
		return nil
	}
	out := new(StaticAuditLogConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StaticWorkersConfig) DeepCopyInto(out *StaticWorkersConfig) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]HostConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StaticWorkersConfig.
func (in *StaticWorkersConfig) DeepCopy() *StaticWorkersConfig {
	if in == nil {
		return nil
	}
	out := new(StaticWorkersConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SystemPackages) DeepCopyInto(out *SystemPackages) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SystemPackages.
func (in *SystemPackages) DeepCopy() *SystemPackages {
	if in == nil {
		return nil
	}
	out := new(SystemPackages)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersionConfig) DeepCopyInto(out *VersionConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersionConfig.
func (in *VersionConfig) DeepCopy() *VersionConfig {
	if in == nil {
		return nil
	}
	out := new(VersionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VsphereSpec) DeepCopyInto(out *VsphereSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VsphereSpec.
func (in *VsphereSpec) DeepCopy() *VsphereSpec {
	if in == nil {
		return nil
	}
	out := new(VsphereSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WeaveNetSpec) DeepCopyInto(out *WeaveNetSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WeaveNetSpec.
func (in *WeaveNetSpec) DeepCopy() *WeaveNetSpec {
	if in == nil {
		return nil
	}
	out := new(WeaveNetSpec)
	in.DeepCopyInto(out)
	return out
}