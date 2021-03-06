// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.1
// source: cloudprovider/dcos.proto

package cloudprovider

import (
	proto "github.com/golang/protobuf/proto"
	client "github.com/spinnaker/kleat/api/client"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Configuration for the DC/OS (Distributed Cloud Operating System) provider.
type Dcos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether the provider is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The list of configured accounts.
	Accounts []*DcosAccount `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
	// The name of the primary account.
	PrimaryAccount string `protobuf:"bytes,3,opt,name=primaryAccount,proto3" json:"primaryAccount,omitempty"`
	// The list of configured clusters.
	Clusters []*DcosCluster `protobuf:"bytes,4,rep,name=clusters,proto3" json:"clusters,omitempty"`
}

func (x *Dcos) Reset() {
	*x = Dcos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_dcos_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dcos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dcos) ProtoMessage() {}

func (x *Dcos) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_dcos_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dcos.ProtoReflect.Descriptor instead.
func (*Dcos) Descriptor() ([]byte, []int) {
	return file_cloudprovider_dcos_proto_rawDescGZIP(), []int{0}
}

func (x *Dcos) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Dcos) GetAccounts() []*DcosAccount {
	if x != nil {
		return x.Accounts
	}
	return nil
}

func (x *Dcos) GetPrimaryAccount() string {
	if x != nil {
		return x.PrimaryAccount
	}
	return ""
}

func (x *Dcos) GetClusters() []*DcosCluster {
	if x != nil {
		return x.Clusters
	}
	return nil
}

// Credentials to authenticate against one or more DC/OS clusters.
type DcosAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) The name of the account.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// (Required) The clusters against which this account will authenticate.
	Clusters []*DcosAccountCluster `protobuf:"bytes,2,rep,name=clusters,proto3" json:"clusters,omitempty"`
	// The environment name for the account. Many accounts can share the
	// same environment (e.g., dev, test, prod).
	Environment string `protobuf:"bytes,3,opt,name=environment,proto3" json:"environment,omitempty"`
	// (Required) The list of Docker registries to use with this DC/OS account.
	DockerRegistries []*DcosAccountDockerRegistry `protobuf:"bytes,4,rep,name=dockerRegistries,proto3" json:"dockerRegistries,omitempty"`
	// Fiat permissions configuration.
	Permissions *client.Permissions `protobuf:"bytes,5,opt,name=permissions,proto3" json:"permissions,omitempty"`
	// (Deprecated) List of required Fiat permission groups. Configure
	// `permissions` instead.
	RequiredGroupMemberships []string `protobuf:"bytes,6,rep,name=requiredGroupMemberships,proto3" json:"requiredGroupMemberships,omitempty"`
}

func (x *DcosAccount) Reset() {
	*x = DcosAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_dcos_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DcosAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DcosAccount) ProtoMessage() {}

func (x *DcosAccount) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_dcos_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DcosAccount.ProtoReflect.Descriptor instead.
func (*DcosAccount) Descriptor() ([]byte, []int) {
	return file_cloudprovider_dcos_proto_rawDescGZIP(), []int{1}
}

func (x *DcosAccount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DcosAccount) GetClusters() []*DcosAccountCluster {
	if x != nil {
		return x.Clusters
	}
	return nil
}

func (x *DcosAccount) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *DcosAccount) GetDockerRegistries() []*DcosAccountDockerRegistry {
	if x != nil {
		return x.DockerRegistries
	}
	return nil
}

func (x *DcosAccount) GetPermissions() *client.Permissions {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *DcosAccount) GetRequiredGroupMemberships() []string {
	if x != nil {
		return x.RequiredGroupMemberships
	}
	return nil
}

// Configuration for a DC/OS cluster associated with a `DcosAccount`.
type DcosAccountCluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) The name of the cluster. Must match the name of a
	// `DcosCluster` defined for this provider.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// (Required) User or service account identifier.
	Uid string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	// Path to a file containing the secret key for service account
	// authentication. If set, `password` should not be set.
	ServiceKeyFile string `protobuf:"bytes,3,opt,name=serviceKeyFile,proto3" json:"serviceKeyFile,omitempty"`
	// Password for a user account. If set, `serviceKeyFile` should not be set.
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *DcosAccountCluster) Reset() {
	*x = DcosAccountCluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_dcos_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DcosAccountCluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DcosAccountCluster) ProtoMessage() {}

func (x *DcosAccountCluster) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_dcos_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DcosAccountCluster.ProtoReflect.Descriptor instead.
func (*DcosAccountCluster) Descriptor() ([]byte, []int) {
	return file_cloudprovider_dcos_proto_rawDescGZIP(), []int{2}
}

func (x *DcosAccountCluster) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DcosAccountCluster) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *DcosAccountCluster) GetServiceKeyFile() string {
	if x != nil {
		return x.ServiceKeyFile
	}
	return ""
}

func (x *DcosAccountCluster) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// Configuration for a DC/OS cluster.
type DcosCluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) The name of the cluster.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Root certificate file to trust for connections to the cluster.
	CaCertFile string `protobuf:"bytes,2,opt,name=caCertFile,proto3" json:"caCertFile,omitempty"`
	// (Required) URL of the endpoint for the DC/OS cluster's admin router.
	DcosUrl string `protobuf:"bytes,3,opt,name=dcosUrl,proto3" json:"dcosUrl,omitempty"`
	// Configuration for a DC/OS load balancer.
	LoadBalancer *DcosClusterLoadBalancer `protobuf:"bytes,4,opt,name=loadBalancer,proto3" json:"loadBalancer,omitempty"`
	// If `true`, disables verification of certificates from the cluster
	// (insecure).
	InsecureSkipTlsVerify bool `protobuf:"varint,5,opt,name=insecureSkipTlsVerify,proto3" json:"insecureSkipTlsVerify,omitempty"`
}

func (x *DcosCluster) Reset() {
	*x = DcosCluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_dcos_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DcosCluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DcosCluster) ProtoMessage() {}

func (x *DcosCluster) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_dcos_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DcosCluster.ProtoReflect.Descriptor instead.
func (*DcosCluster) Descriptor() ([]byte, []int) {
	return file_cloudprovider_dcos_proto_rawDescGZIP(), []int{3}
}

func (x *DcosCluster) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DcosCluster) GetCaCertFile() string {
	if x != nil {
		return x.CaCertFile
	}
	return ""
}

func (x *DcosCluster) GetDcosUrl() string {
	if x != nil {
		return x.DcosUrl
	}
	return ""
}

func (x *DcosCluster) GetLoadBalancer() *DcosClusterLoadBalancer {
	if x != nil {
		return x.LoadBalancer
	}
	return nil
}

func (x *DcosCluster) GetInsecureSkipTlsVerify() bool {
	if x != nil {
		return x.InsecureSkipTlsVerify
	}
	return false
}

// Configuration for a DC/OS load balancer.
type DcosClusterLoadBalancer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Marathon-lb image to use when creating a load balancer with Spinnaker.
	Image string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	// Name of the secret to use for allowing marathon-lb to authenticate with
	// the cluster. Only necessary for clusters with strict or permissive
	// security.
	ServiceAccountSecret string `protobuf:"bytes,2,opt,name=serviceAccountSecret,proto3" json:"serviceAccountSecret,omitempty"`
}

func (x *DcosClusterLoadBalancer) Reset() {
	*x = DcosClusterLoadBalancer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_dcos_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DcosClusterLoadBalancer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DcosClusterLoadBalancer) ProtoMessage() {}

func (x *DcosClusterLoadBalancer) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_dcos_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DcosClusterLoadBalancer.ProtoReflect.Descriptor instead.
func (*DcosClusterLoadBalancer) Descriptor() ([]byte, []int) {
	return file_cloudprovider_dcos_proto_rawDescGZIP(), []int{4}
}

func (x *DcosClusterLoadBalancer) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *DcosClusterLoadBalancer) GetServiceAccountSecret() string {
	if x != nil {
		return x.ServiceAccountSecret
	}
	return ""
}

// Configuration for a Docker registry associated with a `DcosAccount`.
type DcosAccountDockerRegistry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the Docker registry. Must be the name of an account
	// configured with the Docker registry provider.
	AccountName string `protobuf:"bytes,1,opt,name=accountName,proto3" json:"accountName,omitempty"`
}

func (x *DcosAccountDockerRegistry) Reset() {
	*x = DcosAccountDockerRegistry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_dcos_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DcosAccountDockerRegistry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DcosAccountDockerRegistry) ProtoMessage() {}

func (x *DcosAccountDockerRegistry) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_dcos_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DcosAccountDockerRegistry.ProtoReflect.Descriptor instead.
func (*DcosAccountDockerRegistry) Descriptor() ([]byte, []int) {
	return file_cloudprovider_dcos_proto_rawDescGZIP(), []int{5}
}

func (x *DcosAccountDockerRegistry) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

var File_cloudprovider_dcos_proto protoreflect.FileDescriptor

var file_cloudprovider_dcos_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f,
	0x64, 0x63, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x1a,
	0x11, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xc4, 0x01, 0x0a, 0x04, 0x44, 0x63, 0x6f, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x3c, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x44, 0x63,
	0x6f, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x08, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x2e, 0x44, 0x63, 0x6f, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x22, 0xd6, 0x02, 0x0a, 0x0b, 0x44, 0x63,
	0x6f, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x43, 0x0a,
	0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x44, 0x63, 0x6f, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x5a, 0x0a, 0x10, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x44, 0x63, 0x6f, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x10,
	0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x12, 0x34, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3a, 0x0a, 0x18, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69,
	0x70, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x18, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69,
	0x70, 0x73, 0x22, 0x7e, 0x0a, 0x12, 0x44, 0x63, 0x6f, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x26,
	0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x46, 0x69, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4b,
	0x65, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0xe3, 0x01, 0x0a, 0x0b, 0x44, 0x63, 0x6f, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x43, 0x65, 0x72, 0x74,
	0x46, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x43, 0x65,
	0x72, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x63, 0x6f, 0x73, 0x55, 0x72,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x63, 0x6f, 0x73, 0x55, 0x72, 0x6c,
	0x12, 0x50, 0x0a, 0x0c, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x44, 0x63, 0x6f,
	0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x72, 0x52, 0x0c, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x72, 0x12, 0x34, 0x0a, 0x15, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x53, 0x6b,
	0x69, 0x70, 0x54, 0x6c, 0x73, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x15, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x53, 0x6b, 0x69, 0x70, 0x54,
	0x6c, 0x73, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x22, 0x63, 0x0a, 0x17, 0x44, 0x63, 0x6f, 0x73,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x3d, 0x0a,
	0x19, 0x44, 0x63, 0x6f, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x6f, 0x63, 0x6b,
	0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x35, 0x5a, 0x33,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e,
	0x61, 0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloudprovider_dcos_proto_rawDescOnce sync.Once
	file_cloudprovider_dcos_proto_rawDescData = file_cloudprovider_dcos_proto_rawDesc
)

func file_cloudprovider_dcos_proto_rawDescGZIP() []byte {
	file_cloudprovider_dcos_proto_rawDescOnce.Do(func() {
		file_cloudprovider_dcos_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloudprovider_dcos_proto_rawDescData)
	})
	return file_cloudprovider_dcos_proto_rawDescData
}

var file_cloudprovider_dcos_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cloudprovider_dcos_proto_goTypes = []interface{}{
	(*Dcos)(nil),                      // 0: proto.cloudprovider.Dcos
	(*DcosAccount)(nil),               // 1: proto.cloudprovider.DcosAccount
	(*DcosAccountCluster)(nil),        // 2: proto.cloudprovider.DcosAccountCluster
	(*DcosCluster)(nil),               // 3: proto.cloudprovider.DcosCluster
	(*DcosClusterLoadBalancer)(nil),   // 4: proto.cloudprovider.DcosClusterLoadBalancer
	(*DcosAccountDockerRegistry)(nil), // 5: proto.cloudprovider.DcosAccountDockerRegistry
	(*client.Permissions)(nil),        // 6: proto.Permissions
}
var file_cloudprovider_dcos_proto_depIdxs = []int32{
	1, // 0: proto.cloudprovider.Dcos.accounts:type_name -> proto.cloudprovider.DcosAccount
	3, // 1: proto.cloudprovider.Dcos.clusters:type_name -> proto.cloudprovider.DcosCluster
	2, // 2: proto.cloudprovider.DcosAccount.clusters:type_name -> proto.cloudprovider.DcosAccountCluster
	5, // 3: proto.cloudprovider.DcosAccount.dockerRegistries:type_name -> proto.cloudprovider.DcosAccountDockerRegistry
	6, // 4: proto.cloudprovider.DcosAccount.permissions:type_name -> proto.Permissions
	4, // 5: proto.cloudprovider.DcosCluster.loadBalancer:type_name -> proto.cloudprovider.DcosClusterLoadBalancer
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_cloudprovider_dcos_proto_init() }
func file_cloudprovider_dcos_proto_init() {
	if File_cloudprovider_dcos_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloudprovider_dcos_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dcos); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cloudprovider_dcos_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DcosAccount); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cloudprovider_dcos_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DcosAccountCluster); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cloudprovider_dcos_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DcosCluster); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cloudprovider_dcos_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DcosClusterLoadBalancer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cloudprovider_dcos_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DcosAccountDockerRegistry); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloudprovider_dcos_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloudprovider_dcos_proto_goTypes,
		DependencyIndexes: file_cloudprovider_dcos_proto_depIdxs,
		MessageInfos:      file_cloudprovider_dcos_proto_msgTypes,
	}.Build()
	File_cloudprovider_dcos_proto = out.File
	file_cloudprovider_dcos_proto_rawDesc = nil
	file_cloudprovider_dcos_proto_goTypes = nil
	file_cloudprovider_dcos_proto_depIdxs = nil
}
