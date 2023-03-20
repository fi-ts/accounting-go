// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        (unknown)
// source: metalstack/io/accounting/api/v1/volume.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Volume struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Class       string   `protobuf:"bytes,1,opt,name=class,proto3" json:"class,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type        string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Size        string   `protobuf:"bytes,4,opt,name=size,proto3" json:"size,omitempty"`
	Annotations []string `protobuf:"bytes,5,rep,name=annotations,proto3" json:"annotations,omitempty"`
	Uuid        string   `protobuf:"bytes,6,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *Volume) Reset() {
	*x = Volume{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_accounting_api_v1_volume_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Volume) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Volume) ProtoMessage() {}

func (x *Volume) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_accounting_api_v1_volume_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Volume.ProtoReflect.Descriptor instead.
func (*Volume) Descriptor() ([]byte, []int) {
	return file_metalstack_io_accounting_api_v1_volume_proto_rawDescGZIP(), []int{0}
}

func (x *Volume) GetClass() string {
	if x != nil {
		return x.Class
	}
	return ""
}

func (x *Volume) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Volume) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Volume) GetSize() string {
	if x != nil {
		return x.Size
	}
	return ""
}

func (x *Volume) GetAnnotations() []string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *Volume) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type VolumeReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KubernetesReport *KubernetesResourceReport `protobuf:"bytes,1,opt,name=kubernetes_report,json=kubernetesReport,proto3" json:"kubernetes_report,omitempty"`
	Volume           *Volume                   `protobuf:"bytes,2,opt,name=volume,proto3" json:"volume,omitempty"`
}

func (x *VolumeReport) Reset() {
	*x = VolumeReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_accounting_api_v1_volume_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumeReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumeReport) ProtoMessage() {}

func (x *VolumeReport) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_accounting_api_v1_volume_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumeReport.ProtoReflect.Descriptor instead.
func (*VolumeReport) Descriptor() ([]byte, []int) {
	return file_metalstack_io_accounting_api_v1_volume_proto_rawDescGZIP(), []int{1}
}

func (x *VolumeReport) GetKubernetesReport() *KubernetesResourceReport {
	if x != nil {
		return x.KubernetesReport
	}
	return nil
}

func (x *VolumeReport) GetVolume() *Volume {
	if x != nil {
		return x.Volume
	}
	return nil
}

type VolumeUsageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query        *UsageQuery             `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Namespace    *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Storageclass *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=storageclass,proto3" json:"storageclass,omitempty"`
	Type         *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	Annotations  []string                `protobuf:"bytes,3,rep,name=annotations,proto3" json:"annotations,omitempty"`
}

func (x *VolumeUsageRequest) Reset() {
	*x = VolumeUsageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_accounting_api_v1_volume_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumeUsageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumeUsageRequest) ProtoMessage() {}

func (x *VolumeUsageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_accounting_api_v1_volume_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumeUsageRequest.ProtoReflect.Descriptor instead.
func (*VolumeUsageRequest) Descriptor() ([]byte, []int) {
	return file_metalstack_io_accounting_api_v1_volume_proto_rawDescGZIP(), []int{2}
}

func (x *VolumeUsageRequest) GetQuery() *UsageQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *VolumeUsageRequest) GetNamespace() *wrapperspb.StringValue {
	if x != nil {
		return x.Namespace
	}
	return nil
}

func (x *VolumeUsageRequest) GetStorageclass() *wrapperspb.StringValue {
	if x != nil {
		return x.Storageclass
	}
	return nil
}

func (x *VolumeUsageRequest) GetType() *wrapperspb.StringValue {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *VolumeUsageRequest) GetAnnotations() []string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

type VolumeUsageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From             *timestamppb.Timestamp   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To               *timestamppb.Timestamp   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Usage            []*VolumeUsage           `protobuf:"bytes,3,rep,name=usage,proto3" json:"usage,omitempty"`
	AccumulatedUsage *VolumeUsageAccumuluated `protobuf:"bytes,4,opt,name=accumulated_usage,json=accumulatedUsage,proto3" json:"accumulated_usage,omitempty"`
}

func (x *VolumeUsageResponse) Reset() {
	*x = VolumeUsageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_accounting_api_v1_volume_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumeUsageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumeUsageResponse) ProtoMessage() {}

func (x *VolumeUsageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_accounting_api_v1_volume_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumeUsageResponse.ProtoReflect.Descriptor instead.
func (*VolumeUsageResponse) Descriptor() ([]byte, []int) {
	return file_metalstack_io_accounting_api_v1_volume_proto_rawDescGZIP(), []int{3}
}

func (x *VolumeUsageResponse) GetFrom() *timestamppb.Timestamp {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *VolumeUsageResponse) GetTo() *timestamppb.Timestamp {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *VolumeUsageResponse) GetUsage() []*VolumeUsage {
	if x != nil {
		return x.Usage
	}
	return nil
}

func (x *VolumeUsageResponse) GetAccumulatedUsage() *VolumeUsageAccumuluated {
	if x != nil {
		return x.AccumulatedUsage
	}
	return nil
}

type VolumeUsage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Partition       string                 `protobuf:"bytes,1,opt,name=partition,proto3" json:"partition,omitempty"`
	Tenant          string                 `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Projectid       string                 `protobuf:"bytes,3,opt,name=projectid,proto3" json:"projectid,omitempty"`
	Clusterid       string                 `protobuf:"bytes,4,opt,name=clusterid,proto3" json:"clusterid,omitempty"`
	Clustername     string                 `protobuf:"bytes,5,opt,name=clustername,proto3" json:"clustername,omitempty"`
	Class           string                 `protobuf:"bytes,6,opt,name=class,proto3" json:"class,omitempty"`
	Type            string                 `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	Name            string                 `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	Start           *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=start,proto3" json:"start,omitempty"`
	End             *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=end,proto3" json:"end,omitempty"`
	Lifetime        *durationpb.Duration   `protobuf:"bytes,11,opt,name=lifetime,proto3" json:"lifetime,omitempty"`
	CapacitySeconds string                 `protobuf:"bytes,12,opt,name=capacity_seconds,json=capacitySeconds,proto3" json:"capacity_seconds,omitempty"`
	Warnings        []string               `protobuf:"bytes,13,rep,name=warnings,proto3" json:"warnings,omitempty"`
	Annotations     []string               `protobuf:"bytes,14,rep,name=annotations,proto3" json:"annotations,omitempty"`
	Uuid            string                 `protobuf:"bytes,15,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Projectname     string                 `protobuf:"bytes,16,opt,name=projectname,proto3" json:"projectname,omitempty"`
	Tenantname      string                 `protobuf:"bytes,17,opt,name=tenantname,proto3" json:"tenantname,omitempty"`
	Contractnumber  string                 `protobuf:"bytes,18,opt,name=contractnumber,proto3" json:"contractnumber,omitempty"`
	Debtorid        string                 `protobuf:"bytes,19,opt,name=debtorid,proto3" json:"debtorid,omitempty"`
}

func (x *VolumeUsage) Reset() {
	*x = VolumeUsage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_accounting_api_v1_volume_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumeUsage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumeUsage) ProtoMessage() {}

func (x *VolumeUsage) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_accounting_api_v1_volume_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumeUsage.ProtoReflect.Descriptor instead.
func (*VolumeUsage) Descriptor() ([]byte, []int) {
	return file_metalstack_io_accounting_api_v1_volume_proto_rawDescGZIP(), []int{4}
}

func (x *VolumeUsage) GetPartition() string {
	if x != nil {
		return x.Partition
	}
	return ""
}

func (x *VolumeUsage) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *VolumeUsage) GetProjectid() string {
	if x != nil {
		return x.Projectid
	}
	return ""
}

func (x *VolumeUsage) GetClusterid() string {
	if x != nil {
		return x.Clusterid
	}
	return ""
}

func (x *VolumeUsage) GetClustername() string {
	if x != nil {
		return x.Clustername
	}
	return ""
}

func (x *VolumeUsage) GetClass() string {
	if x != nil {
		return x.Class
	}
	return ""
}

func (x *VolumeUsage) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *VolumeUsage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *VolumeUsage) GetStart() *timestamppb.Timestamp {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *VolumeUsage) GetEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.End
	}
	return nil
}

func (x *VolumeUsage) GetLifetime() *durationpb.Duration {
	if x != nil {
		return x.Lifetime
	}
	return nil
}

func (x *VolumeUsage) GetCapacitySeconds() string {
	if x != nil {
		return x.CapacitySeconds
	}
	return ""
}

func (x *VolumeUsage) GetWarnings() []string {
	if x != nil {
		return x.Warnings
	}
	return nil
}

func (x *VolumeUsage) GetAnnotations() []string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *VolumeUsage) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *VolumeUsage) GetProjectname() string {
	if x != nil {
		return x.Projectname
	}
	return ""
}

func (x *VolumeUsage) GetTenantname() string {
	if x != nil {
		return x.Tenantname
	}
	return ""
}

func (x *VolumeUsage) GetContractnumber() string {
	if x != nil {
		return x.Contractnumber
	}
	return ""
}

func (x *VolumeUsage) GetDebtorid() string {
	if x != nil {
		return x.Debtorid
	}
	return ""
}

type VolumeUsageAccumuluated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lifetime        *durationpb.Duration `protobuf:"bytes,1,opt,name=lifetime,proto3" json:"lifetime,omitempty"`
	CapacitySeconds string               `protobuf:"bytes,2,opt,name=capacity_seconds,json=capacitySeconds,proto3" json:"capacity_seconds,omitempty"`
}

func (x *VolumeUsageAccumuluated) Reset() {
	*x = VolumeUsageAccumuluated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_accounting_api_v1_volume_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumeUsageAccumuluated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumeUsageAccumuluated) ProtoMessage() {}

func (x *VolumeUsageAccumuluated) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_accounting_api_v1_volume_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumeUsageAccumuluated.ProtoReflect.Descriptor instead.
func (*VolumeUsageAccumuluated) Descriptor() ([]byte, []int) {
	return file_metalstack_io_accounting_api_v1_volume_proto_rawDescGZIP(), []int{5}
}

func (x *VolumeUsageAccumuluated) GetLifetime() *durationpb.Duration {
	if x != nil {
		return x.Lifetime
	}
	return nil
}

func (x *VolumeUsageAccumuluated) GetCapacitySeconds() string {
	if x != nil {
		return x.CapacitySeconds
	}
	return ""
}

var File_metalstack_io_accounting_api_v1_volume_proto protoreflect.FileDescriptor

var file_metalstack_io_accounting_api_v1_volume_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x69, 0x6f, 0x2f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f,
	0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2c, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x69, 0x6f, 0x2f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90,
	0x01, 0x0a, 0x06, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x22, 0xb7, 0x01, 0x0a, 0x0c, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x66, 0x0a, 0x11, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x10, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x3f, 0x0a, 0x06, 0x76, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x22, 0xa9, 0x02, 0x0a, 0x12,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x41, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2b, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69,
	0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x3a, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x40, 0x0a, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x9c, 0x02, 0x0a, 0x13, 0x56, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2e, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12,
	0x2a, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x42, 0x0a, 0x05, 0x75,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x65, 0x0a, 0x11, 0x61, 0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x75,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x75,
	0x61, 0x74, 0x65, 0x64, 0x52, 0x10, 0x61, 0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65,
	0x64, 0x55, 0x73, 0x61, 0x67, 0x65, 0x22, 0xf9, 0x04, 0x0a, 0x0b, 0x56, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x2c, 0x0a, 0x03, 0x65, 0x6e,
	0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x35, 0x0a, 0x08, 0x6c, 0x69, 0x66, 0x65,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x29, 0x0a, 0x10, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x61, 0x70, 0x61, 0x63,
	0x69, 0x74, 0x79, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x61,
	0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x77, 0x61,
	0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26,
	0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x62, 0x74, 0x6f, 0x72,
	0x69, 0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x62, 0x74, 0x6f, 0x72,
	0x69, 0x64, 0x22, 0x7b, 0x0a, 0x17, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x41, 0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x64, 0x12, 0x35, 0x0a,
	0x08, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x69, 0x66, 0x65,
	0x74, 0x69, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79,
	0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x32,
	0xa8, 0x03, 0x0a, 0x0d, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x5e, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x65, 0x64, 0x12, 0x2d, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x26, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x61, 0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x2d, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x26, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x60, 0x0a, 0x07, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12,
	0x2d, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x26,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x72, 0x0a, 0x05, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x33, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x70, 0x6b,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_metalstack_io_accounting_api_v1_volume_proto_rawDescOnce sync.Once
	file_metalstack_io_accounting_api_v1_volume_proto_rawDescData = file_metalstack_io_accounting_api_v1_volume_proto_rawDesc
)

func file_metalstack_io_accounting_api_v1_volume_proto_rawDescGZIP() []byte {
	file_metalstack_io_accounting_api_v1_volume_proto_rawDescOnce.Do(func() {
		file_metalstack_io_accounting_api_v1_volume_proto_rawDescData = protoimpl.X.CompressGZIP(file_metalstack_io_accounting_api_v1_volume_proto_rawDescData)
	})
	return file_metalstack_io_accounting_api_v1_volume_proto_rawDescData
}

var file_metalstack_io_accounting_api_v1_volume_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_metalstack_io_accounting_api_v1_volume_proto_goTypes = []interface{}{
	(*Volume)(nil),                   // 0: metalstack.io.accounting.api.v1.Volume
	(*VolumeReport)(nil),             // 1: metalstack.io.accounting.api.v1.VolumeReport
	(*VolumeUsageRequest)(nil),       // 2: metalstack.io.accounting.api.v1.VolumeUsageRequest
	(*VolumeUsageResponse)(nil),      // 3: metalstack.io.accounting.api.v1.VolumeUsageResponse
	(*VolumeUsage)(nil),              // 4: metalstack.io.accounting.api.v1.VolumeUsage
	(*VolumeUsageAccumuluated)(nil),  // 5: metalstack.io.accounting.api.v1.VolumeUsageAccumuluated
	(*KubernetesResourceReport)(nil), // 6: metalstack.io.accounting.api.v1.KubernetesResourceReport
	(*UsageQuery)(nil),               // 7: metalstack.io.accounting.api.v1.UsageQuery
	(*wrapperspb.StringValue)(nil),   // 8: google.protobuf.StringValue
	(*timestamppb.Timestamp)(nil),    // 9: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),      // 10: google.protobuf.Duration
	(*Empty)(nil),                    // 11: metalstack.io.accounting.api.v1.Empty
}
var file_metalstack_io_accounting_api_v1_volume_proto_depIdxs = []int32{
	6,  // 0: metalstack.io.accounting.api.v1.VolumeReport.kubernetes_report:type_name -> metalstack.io.accounting.api.v1.KubernetesResourceReport
	0,  // 1: metalstack.io.accounting.api.v1.VolumeReport.volume:type_name -> metalstack.io.accounting.api.v1.Volume
	7,  // 2: metalstack.io.accounting.api.v1.VolumeUsageRequest.query:type_name -> metalstack.io.accounting.api.v1.UsageQuery
	8,  // 3: metalstack.io.accounting.api.v1.VolumeUsageRequest.namespace:type_name -> google.protobuf.StringValue
	8,  // 4: metalstack.io.accounting.api.v1.VolumeUsageRequest.storageclass:type_name -> google.protobuf.StringValue
	8,  // 5: metalstack.io.accounting.api.v1.VolumeUsageRequest.type:type_name -> google.protobuf.StringValue
	9,  // 6: metalstack.io.accounting.api.v1.VolumeUsageResponse.from:type_name -> google.protobuf.Timestamp
	9,  // 7: metalstack.io.accounting.api.v1.VolumeUsageResponse.to:type_name -> google.protobuf.Timestamp
	4,  // 8: metalstack.io.accounting.api.v1.VolumeUsageResponse.usage:type_name -> metalstack.io.accounting.api.v1.VolumeUsage
	5,  // 9: metalstack.io.accounting.api.v1.VolumeUsageResponse.accumulated_usage:type_name -> metalstack.io.accounting.api.v1.VolumeUsageAccumuluated
	9,  // 10: metalstack.io.accounting.api.v1.VolumeUsage.start:type_name -> google.protobuf.Timestamp
	9,  // 11: metalstack.io.accounting.api.v1.VolumeUsage.end:type_name -> google.protobuf.Timestamp
	10, // 12: metalstack.io.accounting.api.v1.VolumeUsage.lifetime:type_name -> google.protobuf.Duration
	10, // 13: metalstack.io.accounting.api.v1.VolumeUsageAccumuluated.lifetime:type_name -> google.protobuf.Duration
	1,  // 14: metalstack.io.accounting.api.v1.VolumeService.Added:input_type -> metalstack.io.accounting.api.v1.VolumeReport
	1,  // 15: metalstack.io.accounting.api.v1.VolumeService.Modified:input_type -> metalstack.io.accounting.api.v1.VolumeReport
	1,  // 16: metalstack.io.accounting.api.v1.VolumeService.Deleted:input_type -> metalstack.io.accounting.api.v1.VolumeReport
	2,  // 17: metalstack.io.accounting.api.v1.VolumeService.Usage:input_type -> metalstack.io.accounting.api.v1.VolumeUsageRequest
	11, // 18: metalstack.io.accounting.api.v1.VolumeService.Added:output_type -> metalstack.io.accounting.api.v1.Empty
	11, // 19: metalstack.io.accounting.api.v1.VolumeService.Modified:output_type -> metalstack.io.accounting.api.v1.Empty
	11, // 20: metalstack.io.accounting.api.v1.VolumeService.Deleted:output_type -> metalstack.io.accounting.api.v1.Empty
	3,  // 21: metalstack.io.accounting.api.v1.VolumeService.Usage:output_type -> metalstack.io.accounting.api.v1.VolumeUsageResponse
	18, // [18:22] is the sub-list for method output_type
	14, // [14:18] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_metalstack_io_accounting_api_v1_volume_proto_init() }
func file_metalstack_io_accounting_api_v1_volume_proto_init() {
	if File_metalstack_io_accounting_api_v1_volume_proto != nil {
		return
	}
	file_metalstack_io_accounting_api_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_metalstack_io_accounting_api_v1_volume_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Volume); i {
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
		file_metalstack_io_accounting_api_v1_volume_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumeReport); i {
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
		file_metalstack_io_accounting_api_v1_volume_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumeUsageRequest); i {
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
		file_metalstack_io_accounting_api_v1_volume_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumeUsageResponse); i {
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
		file_metalstack_io_accounting_api_v1_volume_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumeUsage); i {
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
		file_metalstack_io_accounting_api_v1_volume_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumeUsageAccumuluated); i {
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
			RawDescriptor: file_metalstack_io_accounting_api_v1_volume_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_metalstack_io_accounting_api_v1_volume_proto_goTypes,
		DependencyIndexes: file_metalstack_io_accounting_api_v1_volume_proto_depIdxs,
		MessageInfos:      file_metalstack_io_accounting_api_v1_volume_proto_msgTypes,
	}.Build()
	File_metalstack_io_accounting_api_v1_volume_proto = out.File
	file_metalstack_io_accounting_api_v1_volume_proto_rawDesc = nil
	file_metalstack_io_accounting_api_v1_volume_proto_goTypes = nil
	file_metalstack_io_accounting_api_v1_volume_proto_depIdxs = nil
}
