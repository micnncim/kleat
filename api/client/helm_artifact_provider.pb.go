// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.20.1
// 	protoc        v3.11.2
// source: helm_artifact_provider.proto

package client

import (
	proto "github.com/golang/protobuf/proto"
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

// Configuration for the Helm artifact provider.
type HelmArtifactProvider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether the Helm artifact provider is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The list of configured Helm accounts.
	Accounts []*HelmArtifactAccount `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (x *HelmArtifactProvider) Reset() {
	*x = HelmArtifactProvider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helm_artifact_provider_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelmArtifactProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelmArtifactProvider) ProtoMessage() {}

func (x *HelmArtifactProvider) ProtoReflect() protoreflect.Message {
	mi := &file_helm_artifact_provider_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelmArtifactProvider.ProtoReflect.Descriptor instead.
func (*HelmArtifactProvider) Descriptor() ([]byte, []int) {
	return file_helm_artifact_provider_proto_rawDescGZIP(), []int{0}
}

func (x *HelmArtifactProvider) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *HelmArtifactProvider) GetAccounts() []*HelmArtifactAccount {
	if x != nil {
		return x.Accounts
	}
	return nil
}

// Configuration for a Helm artifact account. Either `username` and `password`
// or `usernamePasswordFile` should be specified as means of authentication.
type HelmArtifactAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the account.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The Helm chart repository URL.
	Repository string `protobuf:"bytes,2,opt,name=repository,proto3" json:"repository,omitempty"`
	// A username for Helm chart repository basic auth.
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	// A password for Helm chart repository basic auth.
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	// The path to a file containing the username and password for Helm chart
	// repository basic auth. Must be in the format `${username}:${password}`.
	UsernamePasswordFile string `protobuf:"bytes,5,opt,name=usernamePasswordFile,proto3" json:"usernamePasswordFile,omitempty"`
}

func (x *HelmArtifactAccount) Reset() {
	*x = HelmArtifactAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helm_artifact_provider_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelmArtifactAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelmArtifactAccount) ProtoMessage() {}

func (x *HelmArtifactAccount) ProtoReflect() protoreflect.Message {
	mi := &file_helm_artifact_provider_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelmArtifactAccount.ProtoReflect.Descriptor instead.
func (*HelmArtifactAccount) Descriptor() ([]byte, []int) {
	return file_helm_artifact_provider_proto_rawDescGZIP(), []int{1}
}

func (x *HelmArtifactAccount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HelmArtifactAccount) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *HelmArtifactAccount) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *HelmArtifactAccount) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *HelmArtifactAccount) GetUsernamePasswordFile() string {
	if x != nil {
		return x.UsernamePasswordFile
	}
	return ""
}

var File_helm_artifact_provider_proto protoreflect.FileDescriptor

var file_helm_artifact_provider_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x68, 0x65, 0x6c, 0x6d, 0x5f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x5f,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x14, 0x48, 0x65, 0x6c, 0x6d, 0x41, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x36, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6d, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22,
	0xb5, 0x01, 0x0a, 0x13, 0x48, 0x65, 0x6c, 0x6d, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x32, 0x0a, 0x14, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x14, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f,
	0x6b, 0x6c, 0x65, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_helm_artifact_provider_proto_rawDescOnce sync.Once
	file_helm_artifact_provider_proto_rawDescData = file_helm_artifact_provider_proto_rawDesc
)

func file_helm_artifact_provider_proto_rawDescGZIP() []byte {
	file_helm_artifact_provider_proto_rawDescOnce.Do(func() {
		file_helm_artifact_provider_proto_rawDescData = protoimpl.X.CompressGZIP(file_helm_artifact_provider_proto_rawDescData)
	})
	return file_helm_artifact_provider_proto_rawDescData
}

var file_helm_artifact_provider_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_helm_artifact_provider_proto_goTypes = []interface{}{
	(*HelmArtifactProvider)(nil), // 0: proto.HelmArtifactProvider
	(*HelmArtifactAccount)(nil),  // 1: proto.HelmArtifactAccount
}
var file_helm_artifact_provider_proto_depIdxs = []int32{
	1, // 0: proto.HelmArtifactProvider.accounts:type_name -> proto.HelmArtifactAccount
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_helm_artifact_provider_proto_init() }
func file_helm_artifact_provider_proto_init() {
	if File_helm_artifact_provider_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_helm_artifact_provider_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelmArtifactProvider); i {
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
		file_helm_artifact_provider_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelmArtifactAccount); i {
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
			RawDescriptor: file_helm_artifact_provider_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_helm_artifact_provider_proto_goTypes,
		DependencyIndexes: file_helm_artifact_provider_proto_depIdxs,
		MessageInfos:      file_helm_artifact_provider_proto_msgTypes,
	}.Build()
	File_helm_artifact_provider_proto = out.File
	file_helm_artifact_provider_proto_rawDesc = nil
	file_helm_artifact_provider_proto_goTypes = nil
	file_helm_artifact_provider_proto_depIdxs = nil
}
