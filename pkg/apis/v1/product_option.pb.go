// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: metalstack/io/accounting/api/v1/product_option.proto

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

type ProductOption int32

const (
	ProductOption_PRODUCT_OPTION_UNSPECIFIED       ProductOption = 0
	ProductOption_PRODUCT_OPTION_AUDIT_EXTENSION   ProductOption = 1
	ProductOption_PRODUCT_OPTION_HA_CONTROL_PLANE  ProductOption = 2
	ProductOption_PRODUCT_OPTION_MONITORING_STACK  ProductOption = 3
	ProductOption_PRODUCT_OPTION_LOGGING_STACK     ProductOption = 4
	ProductOption_PRODUCT_OPTION_CLUSTER_ISOLATION ProductOption = 5
)

// Enum value maps for ProductOption.
var (
	ProductOption_name = map[int32]string{
		0: "PRODUCT_OPTION_UNSPECIFIED",
		1: "PRODUCT_OPTION_AUDIT_EXTENSION",
		2: "PRODUCT_OPTION_HA_CONTROL_PLANE",
		3: "PRODUCT_OPTION_MONITORING_STACK",
		4: "PRODUCT_OPTION_LOGGING_STACK",
		5: "PRODUCT_OPTION_CLUSTER_ISOLATION",
	}
	ProductOption_value = map[string]int32{
		"PRODUCT_OPTION_UNSPECIFIED":       0,
		"PRODUCT_OPTION_AUDIT_EXTENSION":   1,
		"PRODUCT_OPTION_HA_CONTROL_PLANE":  2,
		"PRODUCT_OPTION_MONITORING_STACK":  3,
		"PRODUCT_OPTION_LOGGING_STACK":     4,
		"PRODUCT_OPTION_CLUSTER_ISOLATION": 5,
	}
)

func (x ProductOption) Enum() *ProductOption {
	p := new(ProductOption)
	*p = x
	return p
}

func (x ProductOption) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProductOption) Descriptor() protoreflect.EnumDescriptor {
	return file_metalstack_io_accounting_api_v1_product_option_proto_enumTypes[0].Descriptor()
}

func (ProductOption) Type() protoreflect.EnumType {
	return &file_metalstack_io_accounting_api_v1_product_option_proto_enumTypes[0]
}

func (x ProductOption) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProductOption.Descriptor instead.
func (ProductOption) EnumDescriptor() ([]byte, []int) {
	return file_metalstack_io_accounting_api_v1_product_option_proto_rawDescGZIP(), []int{0}
}

type ProductOptionReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Report      *Report       `protobuf:"bytes,1,opt,name=report,proto3" json:"report,omitempty"`
	Id          ProductOption `protobuf:"varint,2,opt,name=id,proto3,enum=metalstack.io.accounting.api.v1.ProductOption" json:"id,omitempty"`
	Annotations []string      `protobuf:"bytes,3,rep,name=annotations,proto3" json:"annotations,omitempty"`
}

func (x *ProductOptionReport) Reset() {
	*x = ProductOptionReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_accounting_api_v1_product_option_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductOptionReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductOptionReport) ProtoMessage() {}

func (x *ProductOptionReport) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_accounting_api_v1_product_option_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductOptionReport.ProtoReflect.Descriptor instead.
func (*ProductOptionReport) Descriptor() ([]byte, []int) {
	return file_metalstack_io_accounting_api_v1_product_option_proto_rawDescGZIP(), []int{0}
}

func (x *ProductOptionReport) GetReport() *Report {
	if x != nil {
		return x.Report
	}
	return nil
}

func (x *ProductOptionReport) GetId() ProductOption {
	if x != nil {
		return x.Id
	}
	return ProductOption_PRODUCT_OPTION_UNSPECIFIED
}

func (x *ProductOptionReport) GetAnnotations() []string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

type ProductOptionUsageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query *UsageQuery   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Id    ProductOption `protobuf:"varint,2,opt,name=id,proto3,enum=metalstack.io.accounting.api.v1.ProductOption" json:"id,omitempty"`
}

func (x *ProductOptionUsageRequest) Reset() {
	*x = ProductOptionUsageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_accounting_api_v1_product_option_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductOptionUsageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductOptionUsageRequest) ProtoMessage() {}

func (x *ProductOptionUsageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_accounting_api_v1_product_option_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductOptionUsageRequest.ProtoReflect.Descriptor instead.
func (*ProductOptionUsageRequest) Descriptor() ([]byte, []int) {
	return file_metalstack_io_accounting_api_v1_product_option_proto_rawDescGZIP(), []int{1}
}

func (x *ProductOptionUsageRequest) GetQuery() *UsageQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *ProductOptionUsageRequest) GetId() ProductOption {
	if x != nil {
		return x.Id
	}
	return ProductOption_PRODUCT_OPTION_UNSPECIFIED
}

type ProductOptionUsageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From             *timestamppb.Timestamp          `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To               *timestamppb.Timestamp          `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Usage            []*ProductOptionUsage           `protobuf:"bytes,3,rep,name=usage,proto3" json:"usage,omitempty"`
	AccumulatedUsage *ProductOptionUsageAccumuluated `protobuf:"bytes,4,opt,name=accumulated_usage,json=accumulatedUsage,proto3" json:"accumulated_usage,omitempty"`
}

func (x *ProductOptionUsageResponse) Reset() {
	*x = ProductOptionUsageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_accounting_api_v1_product_option_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductOptionUsageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductOptionUsageResponse) ProtoMessage() {}

func (x *ProductOptionUsageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_accounting_api_v1_product_option_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductOptionUsageResponse.ProtoReflect.Descriptor instead.
func (*ProductOptionUsageResponse) Descriptor() ([]byte, []int) {
	return file_metalstack_io_accounting_api_v1_product_option_proto_rawDescGZIP(), []int{2}
}

func (x *ProductOptionUsageResponse) GetFrom() *timestamppb.Timestamp {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *ProductOptionUsageResponse) GetTo() *timestamppb.Timestamp {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *ProductOptionUsageResponse) GetUsage() []*ProductOptionUsage {
	if x != nil {
		return x.Usage
	}
	return nil
}

func (x *ProductOptionUsageResponse) GetAccumulatedUsage() *ProductOptionUsageAccumuluated {
	if x != nil {
		return x.AccumulatedUsage
	}
	return nil
}

type ProductOptionUsage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Partition      string               `protobuf:"bytes,1,opt,name=partition,proto3" json:"partition,omitempty"`
	Tenant         string               `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Projectid      string               `protobuf:"bytes,3,opt,name=projectid,proto3" json:"projectid,omitempty"`
	Clusterid      string               `protobuf:"bytes,4,opt,name=clusterid,proto3" json:"clusterid,omitempty"`
	Clustername    string               `protobuf:"bytes,5,opt,name=clustername,proto3" json:"clustername,omitempty"`
	Lifetime       *durationpb.Duration `protobuf:"bytes,6,opt,name=lifetime,proto3" json:"lifetime,omitempty"`
	Id             ProductOption        `protobuf:"varint,7,opt,name=id,proto3,enum=metalstack.io.accounting.api.v1.ProductOption" json:"id,omitempty"`
	Projectname    string               `protobuf:"bytes,8,opt,name=projectname,proto3" json:"projectname,omitempty"`
	Annotations    []string             `protobuf:"bytes,9,rep,name=annotations,proto3" json:"annotations,omitempty"`
	Tenantname     string               `protobuf:"bytes,10,opt,name=tenantname,proto3" json:"tenantname,omitempty"`
	Contractnumber string               `protobuf:"bytes,11,opt,name=contractnumber,proto3" json:"contractnumber,omitempty"`
	Debtorid       string               `protobuf:"bytes,12,opt,name=debtorid,proto3" json:"debtorid,omitempty"`
}

func (x *ProductOptionUsage) Reset() {
	*x = ProductOptionUsage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_accounting_api_v1_product_option_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductOptionUsage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductOptionUsage) ProtoMessage() {}

func (x *ProductOptionUsage) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_accounting_api_v1_product_option_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductOptionUsage.ProtoReflect.Descriptor instead.
func (*ProductOptionUsage) Descriptor() ([]byte, []int) {
	return file_metalstack_io_accounting_api_v1_product_option_proto_rawDescGZIP(), []int{3}
}

func (x *ProductOptionUsage) GetPartition() string {
	if x != nil {
		return x.Partition
	}
	return ""
}

func (x *ProductOptionUsage) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *ProductOptionUsage) GetProjectid() string {
	if x != nil {
		return x.Projectid
	}
	return ""
}

func (x *ProductOptionUsage) GetClusterid() string {
	if x != nil {
		return x.Clusterid
	}
	return ""
}

func (x *ProductOptionUsage) GetClustername() string {
	if x != nil {
		return x.Clustername
	}
	return ""
}

func (x *ProductOptionUsage) GetLifetime() *durationpb.Duration {
	if x != nil {
		return x.Lifetime
	}
	return nil
}

func (x *ProductOptionUsage) GetId() ProductOption {
	if x != nil {
		return x.Id
	}
	return ProductOption_PRODUCT_OPTION_UNSPECIFIED
}

func (x *ProductOptionUsage) GetProjectname() string {
	if x != nil {
		return x.Projectname
	}
	return ""
}

func (x *ProductOptionUsage) GetAnnotations() []string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *ProductOptionUsage) GetTenantname() string {
	if x != nil {
		return x.Tenantname
	}
	return ""
}

func (x *ProductOptionUsage) GetContractnumber() string {
	if x != nil {
		return x.Contractnumber
	}
	return ""
}

func (x *ProductOptionUsage) GetDebtorid() string {
	if x != nil {
		return x.Debtorid
	}
	return ""
}

type ProductOptionUsageAccumuluated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lifetime *durationpb.Duration `protobuf:"bytes,1,opt,name=lifetime,proto3" json:"lifetime,omitempty"`
}

func (x *ProductOptionUsageAccumuluated) Reset() {
	*x = ProductOptionUsageAccumuluated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_accounting_api_v1_product_option_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductOptionUsageAccumuluated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductOptionUsageAccumuluated) ProtoMessage() {}

func (x *ProductOptionUsageAccumuluated) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_accounting_api_v1_product_option_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductOptionUsageAccumuluated.ProtoReflect.Descriptor instead.
func (*ProductOptionUsageAccumuluated) Descriptor() ([]byte, []int) {
	return file_metalstack_io_accounting_api_v1_product_option_proto_rawDescGZIP(), []int{4}
}

func (x *ProductOptionUsageAccumuluated) GetLifetime() *durationpb.Duration {
	if x != nil {
		return x.Lifetime
	}
	return nil
}

var File_metalstack_io_accounting_api_v1_product_option_proto protoreflect.FileDescriptor

var file_metalstack_io_accounting_api_v1_product_option_proto_rawDesc = []byte{
	0x0a, 0x34, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x69, 0x6f, 0x2f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x2f, 0x69, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69,
	0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x01, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x3f,
	0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x3e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x6d, 0x65,
	0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0x9e, 0x01, 0x0a, 0x19, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x41, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x12, 0x3e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x02,
	0x69, 0x64, 0x22, 0xb1, 0x02, 0x0a, 0x1a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2e, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x49, 0x0a,
	0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x12, 0x6c, 0x0a, 0x11, 0x61, 0x63, 0x63, 0x75,
	0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x75,
	0x61, 0x74, 0x65, 0x64, 0x52, 0x10, 0x61, 0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65,
	0x64, 0x55, 0x73, 0x61, 0x67, 0x65, 0x22, 0xc7, 0x03, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x69, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x35, 0x0a, 0x08, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1e, 0x0a, 0x0a,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x62, 0x74, 0x6f, 0x72, 0x69, 0x64,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x62, 0x74, 0x6f, 0x72, 0x69, 0x64,
	0x22, 0x57, 0x0a, 0x1e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x75, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x35, 0x0a, 0x08, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x08, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x2a, 0xe5, 0x01, 0x0a, 0x0d, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x1a, 0x50,
	0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x22, 0x0a, 0x1e, 0x50,
	0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x55,
	0x44, 0x49, 0x54, 0x5f, 0x45, 0x58, 0x54, 0x45, 0x4e, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12,
	0x23, 0x0a, 0x1f, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x48, 0x41, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x5f, 0x50, 0x4c, 0x41,
	0x4e, 0x45, 0x10, 0x02, 0x12, 0x23, 0x0a, 0x1f, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f,
	0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x4f, 0x4e, 0x49, 0x54, 0x4f, 0x52, 0x49, 0x4e,
	0x47, 0x5f, 0x53, 0x54, 0x41, 0x43, 0x4b, 0x10, 0x03, 0x12, 0x20, 0x0a, 0x1c, 0x50, 0x52, 0x4f,
	0x44, 0x55, 0x43, 0x54, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x4f, 0x47, 0x47,
	0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x43, 0x4b, 0x10, 0x04, 0x12, 0x24, 0x0a, 0x20, 0x50,
	0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4c,
	0x55, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x49, 0x53, 0x4f, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10,
	0x05, 0x32, 0x83, 0x02, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x68, 0x0a, 0x08, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x34, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x26, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x80, 0x01, 0x0a, 0x05, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3a,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x70, 0x6b, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_metalstack_io_accounting_api_v1_product_option_proto_rawDescOnce sync.Once
	file_metalstack_io_accounting_api_v1_product_option_proto_rawDescData = file_metalstack_io_accounting_api_v1_product_option_proto_rawDesc
)

func file_metalstack_io_accounting_api_v1_product_option_proto_rawDescGZIP() []byte {
	file_metalstack_io_accounting_api_v1_product_option_proto_rawDescOnce.Do(func() {
		file_metalstack_io_accounting_api_v1_product_option_proto_rawDescData = protoimpl.X.CompressGZIP(file_metalstack_io_accounting_api_v1_product_option_proto_rawDescData)
	})
	return file_metalstack_io_accounting_api_v1_product_option_proto_rawDescData
}

var file_metalstack_io_accounting_api_v1_product_option_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_metalstack_io_accounting_api_v1_product_option_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_metalstack_io_accounting_api_v1_product_option_proto_goTypes = []interface{}{
	(ProductOption)(0),                     // 0: metalstack.io.accounting.api.v1.ProductOption
	(*ProductOptionReport)(nil),            // 1: metalstack.io.accounting.api.v1.ProductOptionReport
	(*ProductOptionUsageRequest)(nil),      // 2: metalstack.io.accounting.api.v1.ProductOptionUsageRequest
	(*ProductOptionUsageResponse)(nil),     // 3: metalstack.io.accounting.api.v1.ProductOptionUsageResponse
	(*ProductOptionUsage)(nil),             // 4: metalstack.io.accounting.api.v1.ProductOptionUsage
	(*ProductOptionUsageAccumuluated)(nil), // 5: metalstack.io.accounting.api.v1.ProductOptionUsageAccumuluated
	(*Report)(nil),                         // 6: metalstack.io.accounting.api.v1.Report
	(*UsageQuery)(nil),                     // 7: metalstack.io.accounting.api.v1.UsageQuery
	(*timestamppb.Timestamp)(nil),          // 8: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),            // 9: google.protobuf.Duration
	(*Empty)(nil),                          // 10: metalstack.io.accounting.api.v1.Empty
}
var file_metalstack_io_accounting_api_v1_product_option_proto_depIdxs = []int32{
	6,  // 0: metalstack.io.accounting.api.v1.ProductOptionReport.report:type_name -> metalstack.io.accounting.api.v1.Report
	0,  // 1: metalstack.io.accounting.api.v1.ProductOptionReport.id:type_name -> metalstack.io.accounting.api.v1.ProductOption
	7,  // 2: metalstack.io.accounting.api.v1.ProductOptionUsageRequest.query:type_name -> metalstack.io.accounting.api.v1.UsageQuery
	0,  // 3: metalstack.io.accounting.api.v1.ProductOptionUsageRequest.id:type_name -> metalstack.io.accounting.api.v1.ProductOption
	8,  // 4: metalstack.io.accounting.api.v1.ProductOptionUsageResponse.from:type_name -> google.protobuf.Timestamp
	8,  // 5: metalstack.io.accounting.api.v1.ProductOptionUsageResponse.to:type_name -> google.protobuf.Timestamp
	4,  // 6: metalstack.io.accounting.api.v1.ProductOptionUsageResponse.usage:type_name -> metalstack.io.accounting.api.v1.ProductOptionUsage
	5,  // 7: metalstack.io.accounting.api.v1.ProductOptionUsageResponse.accumulated_usage:type_name -> metalstack.io.accounting.api.v1.ProductOptionUsageAccumuluated
	9,  // 8: metalstack.io.accounting.api.v1.ProductOptionUsage.lifetime:type_name -> google.protobuf.Duration
	0,  // 9: metalstack.io.accounting.api.v1.ProductOptionUsage.id:type_name -> metalstack.io.accounting.api.v1.ProductOption
	9,  // 10: metalstack.io.accounting.api.v1.ProductOptionUsageAccumuluated.lifetime:type_name -> google.protobuf.Duration
	1,  // 11: metalstack.io.accounting.api.v1.ProductOptionService.Modified:input_type -> metalstack.io.accounting.api.v1.ProductOptionReport
	2,  // 12: metalstack.io.accounting.api.v1.ProductOptionService.Usage:input_type -> metalstack.io.accounting.api.v1.ProductOptionUsageRequest
	10, // 13: metalstack.io.accounting.api.v1.ProductOptionService.Modified:output_type -> metalstack.io.accounting.api.v1.Empty
	3,  // 14: metalstack.io.accounting.api.v1.ProductOptionService.Usage:output_type -> metalstack.io.accounting.api.v1.ProductOptionUsageResponse
	13, // [13:15] is the sub-list for method output_type
	11, // [11:13] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_metalstack_io_accounting_api_v1_product_option_proto_init() }
func file_metalstack_io_accounting_api_v1_product_option_proto_init() {
	if File_metalstack_io_accounting_api_v1_product_option_proto != nil {
		return
	}
	file_metalstack_io_accounting_api_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_metalstack_io_accounting_api_v1_product_option_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductOptionReport); i {
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
		file_metalstack_io_accounting_api_v1_product_option_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductOptionUsageRequest); i {
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
		file_metalstack_io_accounting_api_v1_product_option_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductOptionUsageResponse); i {
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
		file_metalstack_io_accounting_api_v1_product_option_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductOptionUsage); i {
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
		file_metalstack_io_accounting_api_v1_product_option_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductOptionUsageAccumuluated); i {
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
			RawDescriptor: file_metalstack_io_accounting_api_v1_product_option_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_metalstack_io_accounting_api_v1_product_option_proto_goTypes,
		DependencyIndexes: file_metalstack_io_accounting_api_v1_product_option_proto_depIdxs,
		EnumInfos:         file_metalstack_io_accounting_api_v1_product_option_proto_enumTypes,
		MessageInfos:      file_metalstack_io_accounting_api_v1_product_option_proto_msgTypes,
	}.Build()
	File_metalstack_io_accounting_api_v1_product_option_proto = out.File
	file_metalstack_io_accounting_api_v1_product_option_proto_rawDesc = nil
	file_metalstack_io_accounting_api_v1_product_option_proto_goTypes = nil
	file_metalstack_io_accounting_api_v1_product_option_proto_depIdxs = nil
}
