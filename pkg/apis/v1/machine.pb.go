// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: metalstack/io/accounting/api/v1/machine.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MachineReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Report      *Report                `protobuf:"bytes,1,opt,name=report,proto3" json:"report,omitempty"`
	Start       *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=start,proto3" json:"start,omitempty"`
	End         *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=end,proto3" json:"end,omitempty"`
	Machineid   string                 `protobuf:"bytes,4,opt,name=machineid,proto3" json:"machineid,omitempty"`
	Machinename string                 `protobuf:"bytes,5,opt,name=machinename,proto3" json:"machinename,omitempty"`
	Sizeid      string                 `protobuf:"bytes,6,opt,name=sizeid,proto3" json:"sizeid,omitempty"`
	Imageid     string                 `protobuf:"bytes,7,opt,name=imageid,proto3" json:"imageid,omitempty"`
}

func (x *MachineReport) Reset() {
	*x = MachineReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_accounting_api_v1_machine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachineReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineReport) ProtoMessage() {}

func (x *MachineReport) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_accounting_api_v1_machine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineReport.ProtoReflect.Descriptor instead.
func (*MachineReport) Descriptor() ([]byte, []int) {
	return file_metalstack_io_accounting_api_v1_machine_proto_rawDescGZIP(), []int{0}
}

func (x *MachineReport) GetReport() *Report {
	if x != nil {
		return x.Report
	}
	return nil
}

func (x *MachineReport) GetStart() *timestamppb.Timestamp {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *MachineReport) GetEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.End
	}
	return nil
}

func (x *MachineReport) GetMachineid() string {
	if x != nil {
		return x.Machineid
	}
	return ""
}

func (x *MachineReport) GetMachinename() string {
	if x != nil {
		return x.Machinename
	}
	return ""
}

func (x *MachineReport) GetSizeid() string {
	if x != nil {
		return x.Sizeid
	}
	return ""
}

func (x *MachineReport) GetImageid() string {
	if x != nil {
		return x.Imageid
	}
	return ""
}

type MachineUsageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query     *UsageQuery `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Sizeid    *string     `protobuf:"bytes,2,opt,name=sizeid,proto3,oneof" json:"sizeid,omitempty"`
	Machineid *string     `protobuf:"bytes,3,opt,name=machineid,proto3,oneof" json:"machineid,omitempty"`
	Partition *string     `protobuf:"bytes,4,opt,name=partition,proto3,oneof" json:"partition,omitempty"`
}

func (x *MachineUsageRequest) Reset() {
	*x = MachineUsageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_accounting_api_v1_machine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachineUsageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineUsageRequest) ProtoMessage() {}

func (x *MachineUsageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_accounting_api_v1_machine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineUsageRequest.ProtoReflect.Descriptor instead.
func (*MachineUsageRequest) Descriptor() ([]byte, []int) {
	return file_metalstack_io_accounting_api_v1_machine_proto_rawDescGZIP(), []int{1}
}

func (x *MachineUsageRequest) GetQuery() *UsageQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *MachineUsageRequest) GetSizeid() string {
	if x != nil && x.Sizeid != nil {
		return *x.Sizeid
	}
	return ""
}

func (x *MachineUsageRequest) GetMachineid() string {
	if x != nil && x.Machineid != nil {
		return *x.Machineid
	}
	return ""
}

func (x *MachineUsageRequest) GetPartition() string {
	if x != nil && x.Partition != nil {
		return *x.Partition
	}
	return ""
}

type MachineUsageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From             *timestamppb.Timestamp    `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To               *timestamppb.Timestamp    `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Usage            []*MachineUsage           `protobuf:"bytes,3,rep,name=usage,proto3" json:"usage,omitempty"`
	AccumulatedUsage *MachineUsageAccumuluated `protobuf:"bytes,4,opt,name=accumulated_usage,json=accumulatedUsage,proto3" json:"accumulated_usage,omitempty"`
}

func (x *MachineUsageResponse) Reset() {
	*x = MachineUsageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_accounting_api_v1_machine_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachineUsageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineUsageResponse) ProtoMessage() {}

func (x *MachineUsageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_accounting_api_v1_machine_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineUsageResponse.ProtoReflect.Descriptor instead.
func (*MachineUsageResponse) Descriptor() ([]byte, []int) {
	return file_metalstack_io_accounting_api_v1_machine_proto_rawDescGZIP(), []int{2}
}

func (x *MachineUsageResponse) GetFrom() *timestamppb.Timestamp {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *MachineUsageResponse) GetTo() *timestamppb.Timestamp {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *MachineUsageResponse) GetUsage() []*MachineUsage {
	if x != nil {
		return x.Usage
	}
	return nil
}

func (x *MachineUsageResponse) GetAccumulatedUsage() *MachineUsageAccumuluated {
	if x != nil {
		return x.AccumulatedUsage
	}
	return nil
}

type MachineUsage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Partition      string                 `protobuf:"bytes,1,opt,name=partition,proto3" json:"partition,omitempty"`
	Tenant         string                 `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Projectid      string                 `protobuf:"bytes,3,opt,name=projectid,proto3" json:"projectid,omitempty"`
	Machineid      string                 `protobuf:"bytes,4,opt,name=machineid,proto3" json:"machineid,omitempty"`
	Machinename    string                 `protobuf:"bytes,5,opt,name=machinename,proto3" json:"machinename,omitempty"`
	Sizeid         string                 `protobuf:"bytes,6,opt,name=sizeid,proto3" json:"sizeid,omitempty"`
	Imageid        string                 `protobuf:"bytes,7,opt,name=imageid,proto3" json:"imageid,omitempty"`
	Clusterid      string                 `protobuf:"bytes,8,opt,name=clusterid,proto3" json:"clusterid,omitempty"`
	Clustername    string                 `protobuf:"bytes,9,opt,name=clustername,proto3" json:"clustername,omitempty"`
	Start          *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=start,proto3" json:"start,omitempty"`
	End            *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=end,proto3" json:"end,omitempty"`
	Lifetime       *durationpb.Duration   `protobuf:"bytes,12,opt,name=lifetime,proto3" json:"lifetime,omitempty"`
	Projectname    string                 `protobuf:"bytes,13,opt,name=projectname,proto3" json:"projectname,omitempty"`
	Tenantname     string                 `protobuf:"bytes,14,opt,name=tenantname,proto3" json:"tenantname,omitempty"`
	Contractnumber string                 `protobuf:"bytes,15,opt,name=contractnumber,proto3" json:"contractnumber,omitempty"`
	Debtorid       string                 `protobuf:"bytes,16,opt,name=debtorid,proto3" json:"debtorid,omitempty"`
}

func (x *MachineUsage) Reset() {
	*x = MachineUsage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_accounting_api_v1_machine_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachineUsage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineUsage) ProtoMessage() {}

func (x *MachineUsage) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_accounting_api_v1_machine_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineUsage.ProtoReflect.Descriptor instead.
func (*MachineUsage) Descriptor() ([]byte, []int) {
	return file_metalstack_io_accounting_api_v1_machine_proto_rawDescGZIP(), []int{3}
}

func (x *MachineUsage) GetPartition() string {
	if x != nil {
		return x.Partition
	}
	return ""
}

func (x *MachineUsage) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *MachineUsage) GetProjectid() string {
	if x != nil {
		return x.Projectid
	}
	return ""
}

func (x *MachineUsage) GetMachineid() string {
	if x != nil {
		return x.Machineid
	}
	return ""
}

func (x *MachineUsage) GetMachinename() string {
	if x != nil {
		return x.Machinename
	}
	return ""
}

func (x *MachineUsage) GetSizeid() string {
	if x != nil {
		return x.Sizeid
	}
	return ""
}

func (x *MachineUsage) GetImageid() string {
	if x != nil {
		return x.Imageid
	}
	return ""
}

func (x *MachineUsage) GetClusterid() string {
	if x != nil {
		return x.Clusterid
	}
	return ""
}

func (x *MachineUsage) GetClustername() string {
	if x != nil {
		return x.Clustername
	}
	return ""
}

func (x *MachineUsage) GetStart() *timestamppb.Timestamp {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *MachineUsage) GetEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.End
	}
	return nil
}

func (x *MachineUsage) GetLifetime() *durationpb.Duration {
	if x != nil {
		return x.Lifetime
	}
	return nil
}

func (x *MachineUsage) GetProjectname() string {
	if x != nil {
		return x.Projectname
	}
	return ""
}

func (x *MachineUsage) GetTenantname() string {
	if x != nil {
		return x.Tenantname
	}
	return ""
}

func (x *MachineUsage) GetContractnumber() string {
	if x != nil {
		return x.Contractnumber
	}
	return ""
}

func (x *MachineUsage) GetDebtorid() string {
	if x != nil {
		return x.Debtorid
	}
	return ""
}

type MachineUsageAccumuluated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lifetime *durationpb.Duration `protobuf:"bytes,1,opt,name=lifetime,proto3" json:"lifetime,omitempty"`
}

func (x *MachineUsageAccumuluated) Reset() {
	*x = MachineUsageAccumuluated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_accounting_api_v1_machine_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachineUsageAccumuluated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineUsageAccumuluated) ProtoMessage() {}

func (x *MachineUsageAccumuluated) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_accounting_api_v1_machine_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineUsageAccumuluated.ProtoReflect.Descriptor instead.
func (*MachineUsageAccumuluated) Descriptor() ([]byte, []int) {
	return file_metalstack_io_accounting_api_v1_machine_proto_rawDescGZIP(), []int{4}
}

func (x *MachineUsageAccumuluated) GetLifetime() *durationpb.Duration {
	if x != nil {
		return x.Lifetime
	}
	return nil
}

var File_metalstack_io_accounting_api_v1_machine_proto protoreflect.FileDescriptor

var file_metalstack_io_accounting_api_v1_machine_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x69, 0x6f, 0x2f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1f, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2c, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x69, 0x6f,
	0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa2, 0x02, 0x0a, 0x0d, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x3f, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69,
	0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x12, 0x2c, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x65,
	0x6e, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x7a, 0x65, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x69, 0x7a, 0x65, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x69, 0x64, 0x22, 0xe2, 0x01, 0x0a, 0x13, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x05,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6d, 0x65,
	0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12,
	0x1b, 0x0a, 0x06, 0x73, 0x69, 0x7a, 0x65, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x06, 0x73, 0x69, 0x7a, 0x65, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09,
	0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x09, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x21, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x02, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x88,
	0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x69, 0x64, 0x42, 0x0c, 0x0a,
	0x0a, 0x5f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x5f,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x9f, 0x02, 0x0a, 0x14, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x43,
	0x0a, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x75, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x66, 0x0a, 0x11, 0x61, 0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x63,
	0x75, 0x6d, 0x75, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x64, 0x52, 0x10, 0x61, 0x63, 0x63, 0x75, 0x6d,
	0x75, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x55, 0x73, 0x61, 0x67, 0x65, 0x22, 0xb1, 0x04, 0x0a, 0x0c,
	0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x7a, 0x65, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x69, 0x7a, 0x65, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x69, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x69, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x12, 0x2c, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x65,
	0x6e, 0x64, 0x12, 0x35, 0x0a, 0x08, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x08, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x62, 0x74, 0x6f, 0x72, 0x69, 0x64, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x62, 0x74, 0x6f, 0x72, 0x69, 0x64, 0x22,
	0x51, 0x0a, 0x18, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x41,
	0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x64, 0x12, 0x35, 0x0a, 0x08, 0x6c,
	0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69,
	0x6d, 0x65, 0x32, 0xae, 0x03, 0x0a, 0x0e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x65, 0x64, 0x12, 0x2e,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x26,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x62, 0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x12, 0x2e, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e,
	0x69, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x1a, 0x26, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e,
	0x69, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x61, 0x0a, 0x07, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x2e, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x26, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x74, 0x0a,
	0x05, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x34, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_metalstack_io_accounting_api_v1_machine_proto_rawDescOnce sync.Once
	file_metalstack_io_accounting_api_v1_machine_proto_rawDescData = file_metalstack_io_accounting_api_v1_machine_proto_rawDesc
)

func file_metalstack_io_accounting_api_v1_machine_proto_rawDescGZIP() []byte {
	file_metalstack_io_accounting_api_v1_machine_proto_rawDescOnce.Do(func() {
		file_metalstack_io_accounting_api_v1_machine_proto_rawDescData = protoimpl.X.CompressGZIP(file_metalstack_io_accounting_api_v1_machine_proto_rawDescData)
	})
	return file_metalstack_io_accounting_api_v1_machine_proto_rawDescData
}

var file_metalstack_io_accounting_api_v1_machine_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_metalstack_io_accounting_api_v1_machine_proto_goTypes = []interface{}{
	(*MachineReport)(nil),            // 0: metalstack.io.accounting.api.v1.MachineReport
	(*MachineUsageRequest)(nil),      // 1: metalstack.io.accounting.api.v1.MachineUsageRequest
	(*MachineUsageResponse)(nil),     // 2: metalstack.io.accounting.api.v1.MachineUsageResponse
	(*MachineUsage)(nil),             // 3: metalstack.io.accounting.api.v1.MachineUsage
	(*MachineUsageAccumuluated)(nil), // 4: metalstack.io.accounting.api.v1.MachineUsageAccumuluated
	(*Report)(nil),                   // 5: metalstack.io.accounting.api.v1.Report
	(*timestamppb.Timestamp)(nil),    // 6: google.protobuf.Timestamp
	(*UsageQuery)(nil),               // 7: metalstack.io.accounting.api.v1.UsageQuery
	(*durationpb.Duration)(nil),      // 8: google.protobuf.Duration
	(*Empty)(nil),                    // 9: metalstack.io.accounting.api.v1.Empty
}
var file_metalstack_io_accounting_api_v1_machine_proto_depIdxs = []int32{
	5,  // 0: metalstack.io.accounting.api.v1.MachineReport.report:type_name -> metalstack.io.accounting.api.v1.Report
	6,  // 1: metalstack.io.accounting.api.v1.MachineReport.start:type_name -> google.protobuf.Timestamp
	6,  // 2: metalstack.io.accounting.api.v1.MachineReport.end:type_name -> google.protobuf.Timestamp
	7,  // 3: metalstack.io.accounting.api.v1.MachineUsageRequest.query:type_name -> metalstack.io.accounting.api.v1.UsageQuery
	6,  // 4: metalstack.io.accounting.api.v1.MachineUsageResponse.from:type_name -> google.protobuf.Timestamp
	6,  // 5: metalstack.io.accounting.api.v1.MachineUsageResponse.to:type_name -> google.protobuf.Timestamp
	3,  // 6: metalstack.io.accounting.api.v1.MachineUsageResponse.usage:type_name -> metalstack.io.accounting.api.v1.MachineUsage
	4,  // 7: metalstack.io.accounting.api.v1.MachineUsageResponse.accumulated_usage:type_name -> metalstack.io.accounting.api.v1.MachineUsageAccumuluated
	6,  // 8: metalstack.io.accounting.api.v1.MachineUsage.start:type_name -> google.protobuf.Timestamp
	6,  // 9: metalstack.io.accounting.api.v1.MachineUsage.end:type_name -> google.protobuf.Timestamp
	8,  // 10: metalstack.io.accounting.api.v1.MachineUsage.lifetime:type_name -> google.protobuf.Duration
	8,  // 11: metalstack.io.accounting.api.v1.MachineUsageAccumuluated.lifetime:type_name -> google.protobuf.Duration
	0,  // 12: metalstack.io.accounting.api.v1.MachineService.Added:input_type -> metalstack.io.accounting.api.v1.MachineReport
	0,  // 13: metalstack.io.accounting.api.v1.MachineService.Modified:input_type -> metalstack.io.accounting.api.v1.MachineReport
	0,  // 14: metalstack.io.accounting.api.v1.MachineService.Deleted:input_type -> metalstack.io.accounting.api.v1.MachineReport
	1,  // 15: metalstack.io.accounting.api.v1.MachineService.Usage:input_type -> metalstack.io.accounting.api.v1.MachineUsageRequest
	9,  // 16: metalstack.io.accounting.api.v1.MachineService.Added:output_type -> metalstack.io.accounting.api.v1.Empty
	9,  // 17: metalstack.io.accounting.api.v1.MachineService.Modified:output_type -> metalstack.io.accounting.api.v1.Empty
	9,  // 18: metalstack.io.accounting.api.v1.MachineService.Deleted:output_type -> metalstack.io.accounting.api.v1.Empty
	2,  // 19: metalstack.io.accounting.api.v1.MachineService.Usage:output_type -> metalstack.io.accounting.api.v1.MachineUsageResponse
	16, // [16:20] is the sub-list for method output_type
	12, // [12:16] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_metalstack_io_accounting_api_v1_machine_proto_init() }
func file_metalstack_io_accounting_api_v1_machine_proto_init() {
	if File_metalstack_io_accounting_api_v1_machine_proto != nil {
		return
	}
	file_metalstack_io_accounting_api_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_metalstack_io_accounting_api_v1_machine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MachineReport); i {
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
		file_metalstack_io_accounting_api_v1_machine_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MachineUsageRequest); i {
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
		file_metalstack_io_accounting_api_v1_machine_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MachineUsageResponse); i {
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
		file_metalstack_io_accounting_api_v1_machine_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MachineUsage); i {
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
		file_metalstack_io_accounting_api_v1_machine_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MachineUsageAccumuluated); i {
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
	file_metalstack_io_accounting_api_v1_machine_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_metalstack_io_accounting_api_v1_machine_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_metalstack_io_accounting_api_v1_machine_proto_goTypes,
		DependencyIndexes: file_metalstack_io_accounting_api_v1_machine_proto_depIdxs,
		MessageInfos:      file_metalstack_io_accounting_api_v1_machine_proto_msgTypes,
	}.Build()
	File_metalstack_io_accounting_api_v1_machine_proto = out.File
	file_metalstack_io_accounting_api_v1_machine_proto_rawDesc = nil
	file_metalstack_io_accounting_api_v1_machine_proto_goTypes = nil
	file_metalstack_io_accounting_api_v1_machine_proto_depIdxs = nil
}