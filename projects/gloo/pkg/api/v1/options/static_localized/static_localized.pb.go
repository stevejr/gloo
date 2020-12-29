// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/static_localized/static_localized.proto

package static_localized

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	options "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Static localized upstreams are identical to static upstreams except that hosts can be configured
//  with a locality that is used for load balancing requests
//
// Similar to static upstreams, static localized upstreams:
//  - are used to route request to services listening at fixed IP/Host & Port pairs.
//  - can be used to proxy any kind of service, and therefore contain a ServiceSpec for additional service-specific configuration.
//  - must be created manually by users
type UpstreamSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of addresses and ports, separated by locality
	// at least one must be specified
	LocalizedHosts []*LocalizedHosts `protobuf:"bytes,1,rep,name=localized_hosts,json=localizedHosts,proto3" json:"localized_hosts,omitempty"`
	// Attempt to use outbound TLS
	// Gloo will automatically set this to true for port 443
	UseTls bool `protobuf:"varint,3,opt,name=use_tls,json=useTls,proto3" json:"use_tls,omitempty"`
	// An optional Service Spec describing the service listening at this address
	ServiceSpec *options.ServiceSpec `protobuf:"bytes,5,opt,name=service_spec,json=serviceSpec,proto3" json:"service_spec,omitempty"`
	// When set, automatically set the sni address to use to the addr field.
	// If both this and host.sni_addr are set, host.sni_addr has priority.
	// defaults to "true".
	AutoSniRewrite *wrappers.BoolValue `protobuf:"bytes,6,opt,name=auto_sni_rewrite,json=autoSniRewrite,proto3" json:"auto_sni_rewrite,omitempty"`
}

func (x *UpstreamSpec) Reset() {
	*x = UpstreamSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpstreamSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpstreamSpec) ProtoMessage() {}

func (x *UpstreamSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpstreamSpec.ProtoReflect.Descriptor instead.
func (*UpstreamSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_rawDescGZIP(), []int{0}
}

func (x *UpstreamSpec) GetLocalizedHosts() []*LocalizedHosts {
	if x != nil {
		return x.LocalizedHosts
	}
	return nil
}

func (x *UpstreamSpec) GetUseTls() bool {
	if x != nil {
		return x.UseTls
	}
	return false
}

func (x *UpstreamSpec) GetServiceSpec() *options.ServiceSpec {
	if x != nil {
		return x.ServiceSpec
	}
	return nil
}

func (x *UpstreamSpec) GetAutoSniRewrite() *wrappers.BoolValue {
	if x != nil {
		return x.AutoSniRewrite
	}
	return nil
}

// Represents a set of instances of an upstream, running in a particular locality
type LocalizedHosts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifies where the hosts run.
	Locality *Locality `protobuf:"bytes,1,opt,name=locality,proto3" json:"locality,omitempty"`
	// A list of addresses and ports
	// at least one must be specified
	Hosts []*Host `protobuf:"bytes,2,rep,name=hosts,proto3" json:"hosts,omitempty"`
}

func (x *LocalizedHosts) Reset() {
	*x = LocalizedHosts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalizedHosts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalizedHosts) ProtoMessage() {}

func (x *LocalizedHosts) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalizedHosts.ProtoReflect.Descriptor instead.
func (*LocalizedHosts) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_rawDescGZIP(), []int{1}
}

func (x *LocalizedHosts) GetLocality() *Locality {
	if x != nil {
		return x.Locality
	}
	return nil
}

func (x *LocalizedHosts) GetHosts() []*Host {
	if x != nil {
		return x.Hosts
	}
	return nil
}

// Represents a single instance of an upstream
type Host struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Address (hostname or IP)
	Addr string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	// Port the instance is listening on
	Port uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	// Address to use for SNI if using ssl.
	SniAddr string `protobuf:"bytes,4,opt,name=sni_addr,json=sniAddr,proto3" json:"sni_addr,omitempty"`
	// (Enterprise Only): Host specific health checking configuration.
	HealthCheckConfig *Host_HealthCheckConfig `protobuf:"bytes,3,opt,name=health_check_config,json=healthCheckConfig,proto3" json:"health_check_config,omitempty"`
}

func (x *Host) Reset() {
	*x = Host{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Host) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Host) ProtoMessage() {}

func (x *Host) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Host.ProtoReflect.Descriptor instead.
func (*Host) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_rawDescGZIP(), []int{2}
}

func (x *Host) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Host) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Host) GetSniAddr() string {
	if x != nil {
		return x.SniAddr
	}
	return ""
}

func (x *Host) GetHealthCheckConfig() *Host_HealthCheckConfig {
	if x != nil {
		return x.HealthCheckConfig
	}
	return nil
}

// Identifies location of where either Envoy runs or where upstream hosts run.
type Locality struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Region this zone belongs to.
	Region string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	// Defines the local service zone where Envoy is running. The meaning of zone
	// is context dependent, e.g. `Availability Zone (AZ)
	// <https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html>`_
	// on AWS, `Zone <https://cloud.google.com/compute/docs/regions-zones/>`_ on
	// GCP, etc.
	Zone string `protobuf:"bytes,2,opt,name=zone,proto3" json:"zone,omitempty"`
	// When used for locality of upstream hosts, this field further splits zone
	// into smaller chunks of sub-zones so they can be load balanced
	// independently.
	SubZone string `protobuf:"bytes,3,opt,name=sub_zone,json=subZone,proto3" json:"sub_zone,omitempty"`
}

func (x *Locality) Reset() {
	*x = Locality{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Locality) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Locality) ProtoMessage() {}

func (x *Locality) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Locality.ProtoReflect.Descriptor instead.
func (*Locality) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_rawDescGZIP(), []int{3}
}

func (x *Locality) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Locality) GetZone() string {
	if x != nil {
		return x.Zone
	}
	return ""
}

func (x *Locality) GetSubZone() string {
	if x != nil {
		return x.SubZone
	}
	return ""
}

type Host_HealthCheckConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Enterprise Only): Path to use when health checking this specific host.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *Host_HealthCheckConfig) Reset() {
	*x = Host_HealthCheckConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Host_HealthCheckConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Host_HealthCheckConfig) ProtoMessage() {}

func (x *Host_HealthCheckConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Host_HealthCheckConfig.ProtoReflect.Descriptor instead.
func (*Host_HealthCheckConfig) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Host_HealthCheckConfig) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_rawDesc = []byte{
	0x0a, 0x5c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64,
	0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x93, 0x02, 0x0a, 0x0c, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53,
	0x70, 0x65, 0x63, 0x12, 0x5e, 0x0a, 0x0f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64,
	0x5f, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x2e,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x48, 0x6f,
	0x73, 0x74, 0x73, 0x52, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x48, 0x6f,
	0x73, 0x74, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x5f, 0x74, 0x6c, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x75, 0x73, 0x65, 0x54, 0x6c, 0x73, 0x12, 0x44, 0x0a, 0x0c,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x70,
	0x65, 0x63, 0x12, 0x44, 0x0a, 0x10, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x73, 0x6e, 0x69, 0x5f, 0x72,
	0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42,
	0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x61, 0x75, 0x74, 0x6f, 0x53, 0x6e,
	0x69, 0x52, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x22, 0xa0, 0x01, 0x0a, 0x0e, 0x4c, 0x6f, 0x63,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x4b, 0x0a, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64,
	0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x41, 0x0a, 0x05, 0x68, 0x6f, 0x73, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63,
	0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x48, 0x6f, 0x73, 0x74, 0x52, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x22, 0xe1, 0x01, 0x0a, 0x04,
	0x48, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x73, 0x6e, 0x69, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x6e, 0x69, 0x41, 0x64, 0x64, 0x72, 0x12, 0x6d, 0x0a, 0x13, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x48, 0x6f, 0x73,
	0x74, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x11, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x27, 0x0a, 0x11, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22,
	0x51, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x5f, 0x7a,
	0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x5a, 0x6f,
	0x6e, 0x65, 0x42, 0x53, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0xc0,
	0xf5, 0x04, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_goTypes = []interface{}{
	(*UpstreamSpec)(nil),           // 0: static_localized.options.gloo.solo.io.UpstreamSpec
	(*LocalizedHosts)(nil),         // 1: static_localized.options.gloo.solo.io.LocalizedHosts
	(*Host)(nil),                   // 2: static_localized.options.gloo.solo.io.Host
	(*Locality)(nil),               // 3: static_localized.options.gloo.solo.io.Locality
	(*Host_HealthCheckConfig)(nil), // 4: static_localized.options.gloo.solo.io.Host.HealthCheckConfig
	(*options.ServiceSpec)(nil),    // 5: options.gloo.solo.io.ServiceSpec
	(*wrappers.BoolValue)(nil),     // 6: google.protobuf.BoolValue
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_depIdxs = []int32{
	1, // 0: static_localized.options.gloo.solo.io.UpstreamSpec.localized_hosts:type_name -> static_localized.options.gloo.solo.io.LocalizedHosts
	5, // 1: static_localized.options.gloo.solo.io.UpstreamSpec.service_spec:type_name -> options.gloo.solo.io.ServiceSpec
	6, // 2: static_localized.options.gloo.solo.io.UpstreamSpec.auto_sni_rewrite:type_name -> google.protobuf.BoolValue
	3, // 3: static_localized.options.gloo.solo.io.LocalizedHosts.locality:type_name -> static_localized.options.gloo.solo.io.Locality
	2, // 4: static_localized.options.gloo.solo.io.LocalizedHosts.hosts:type_name -> static_localized.options.gloo.solo.io.Host
	4, // 5: static_localized.options.gloo.solo.io.Host.health_check_config:type_name -> static_localized.options.gloo.solo.io.Host.HealthCheckConfig
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_init()
}
func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpstreamSpec); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalizedHosts); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Host); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Locality); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Host_HealthCheckConfig); i {
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
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_localized_static_localized_proto_depIdxs = nil
}
