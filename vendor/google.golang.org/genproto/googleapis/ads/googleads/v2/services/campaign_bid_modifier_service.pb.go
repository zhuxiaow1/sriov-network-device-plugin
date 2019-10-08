// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/campaign_bid_modifier_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request message for [CampaignBidModifierService.GetCampaignBidModifier][google.ads.googleads.v2.services.CampaignBidModifierService.GetCampaignBidModifier].
type GetCampaignBidModifierRequest struct {
	// The resource name of the campaign bid modifier to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCampaignBidModifierRequest) Reset()         { *m = GetCampaignBidModifierRequest{} }
func (m *GetCampaignBidModifierRequest) String() string { return proto.CompactTextString(m) }
func (*GetCampaignBidModifierRequest) ProtoMessage()    {}
func (*GetCampaignBidModifierRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6133a924995ecc19, []int{0}
}

func (m *GetCampaignBidModifierRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCampaignBidModifierRequest.Unmarshal(m, b)
}
func (m *GetCampaignBidModifierRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCampaignBidModifierRequest.Marshal(b, m, deterministic)
}
func (m *GetCampaignBidModifierRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCampaignBidModifierRequest.Merge(m, src)
}
func (m *GetCampaignBidModifierRequest) XXX_Size() int {
	return xxx_messageInfo_GetCampaignBidModifierRequest.Size(m)
}
func (m *GetCampaignBidModifierRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCampaignBidModifierRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCampaignBidModifierRequest proto.InternalMessageInfo

func (m *GetCampaignBidModifierRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [CampaignBidModifierService.MutateCampaignBidModifier][].
type MutateCampaignBidModifiersRequest struct {
	// ID of the customer whose campaign bid modifiers are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual campaign bid modifiers.
	Operations []*CampaignBidModifierOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCampaignBidModifiersRequest) Reset()         { *m = MutateCampaignBidModifiersRequest{} }
func (m *MutateCampaignBidModifiersRequest) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignBidModifiersRequest) ProtoMessage()    {}
func (*MutateCampaignBidModifiersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6133a924995ecc19, []int{1}
}

func (m *MutateCampaignBidModifiersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignBidModifiersRequest.Unmarshal(m, b)
}
func (m *MutateCampaignBidModifiersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignBidModifiersRequest.Marshal(b, m, deterministic)
}
func (m *MutateCampaignBidModifiersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignBidModifiersRequest.Merge(m, src)
}
func (m *MutateCampaignBidModifiersRequest) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignBidModifiersRequest.Size(m)
}
func (m *MutateCampaignBidModifiersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignBidModifiersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignBidModifiersRequest proto.InternalMessageInfo

func (m *MutateCampaignBidModifiersRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateCampaignBidModifiersRequest) GetOperations() []*CampaignBidModifierOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateCampaignBidModifiersRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateCampaignBidModifiersRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, remove, update) on a campaign bid modifier.
type CampaignBidModifierOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*CampaignBidModifierOperation_Create
	//	*CampaignBidModifierOperation_Update
	//	*CampaignBidModifierOperation_Remove
	Operation            isCampaignBidModifierOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *CampaignBidModifierOperation) Reset()         { *m = CampaignBidModifierOperation{} }
func (m *CampaignBidModifierOperation) String() string { return proto.CompactTextString(m) }
func (*CampaignBidModifierOperation) ProtoMessage()    {}
func (*CampaignBidModifierOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_6133a924995ecc19, []int{2}
}

func (m *CampaignBidModifierOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignBidModifierOperation.Unmarshal(m, b)
}
func (m *CampaignBidModifierOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignBidModifierOperation.Marshal(b, m, deterministic)
}
func (m *CampaignBidModifierOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignBidModifierOperation.Merge(m, src)
}
func (m *CampaignBidModifierOperation) XXX_Size() int {
	return xxx_messageInfo_CampaignBidModifierOperation.Size(m)
}
func (m *CampaignBidModifierOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignBidModifierOperation.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignBidModifierOperation proto.InternalMessageInfo

func (m *CampaignBidModifierOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isCampaignBidModifierOperation_Operation interface {
	isCampaignBidModifierOperation_Operation()
}

type CampaignBidModifierOperation_Create struct {
	Create *resources.CampaignBidModifier `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type CampaignBidModifierOperation_Update struct {
	Update *resources.CampaignBidModifier `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type CampaignBidModifierOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*CampaignBidModifierOperation_Create) isCampaignBidModifierOperation_Operation() {}

func (*CampaignBidModifierOperation_Update) isCampaignBidModifierOperation_Operation() {}

func (*CampaignBidModifierOperation_Remove) isCampaignBidModifierOperation_Operation() {}

func (m *CampaignBidModifierOperation) GetOperation() isCampaignBidModifierOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *CampaignBidModifierOperation) GetCreate() *resources.CampaignBidModifier {
	if x, ok := m.GetOperation().(*CampaignBidModifierOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *CampaignBidModifierOperation) GetUpdate() *resources.CampaignBidModifier {
	if x, ok := m.GetOperation().(*CampaignBidModifierOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *CampaignBidModifierOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*CampaignBidModifierOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CampaignBidModifierOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CampaignBidModifierOperation_Create)(nil),
		(*CampaignBidModifierOperation_Update)(nil),
		(*CampaignBidModifierOperation_Remove)(nil),
	}
}

// Response message for campaign bid modifiers mutate.
type MutateCampaignBidModifiersResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateCampaignBidModifierResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *MutateCampaignBidModifiersResponse) Reset()         { *m = MutateCampaignBidModifiersResponse{} }
func (m *MutateCampaignBidModifiersResponse) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignBidModifiersResponse) ProtoMessage()    {}
func (*MutateCampaignBidModifiersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6133a924995ecc19, []int{3}
}

func (m *MutateCampaignBidModifiersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignBidModifiersResponse.Unmarshal(m, b)
}
func (m *MutateCampaignBidModifiersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignBidModifiersResponse.Marshal(b, m, deterministic)
}
func (m *MutateCampaignBidModifiersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignBidModifiersResponse.Merge(m, src)
}
func (m *MutateCampaignBidModifiersResponse) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignBidModifiersResponse.Size(m)
}
func (m *MutateCampaignBidModifiersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignBidModifiersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignBidModifiersResponse proto.InternalMessageInfo

func (m *MutateCampaignBidModifiersResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateCampaignBidModifiersResponse) GetResults() []*MutateCampaignBidModifierResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the criterion mutate.
type MutateCampaignBidModifierResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCampaignBidModifierResult) Reset()         { *m = MutateCampaignBidModifierResult{} }
func (m *MutateCampaignBidModifierResult) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignBidModifierResult) ProtoMessage()    {}
func (*MutateCampaignBidModifierResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_6133a924995ecc19, []int{4}
}

func (m *MutateCampaignBidModifierResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignBidModifierResult.Unmarshal(m, b)
}
func (m *MutateCampaignBidModifierResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignBidModifierResult.Marshal(b, m, deterministic)
}
func (m *MutateCampaignBidModifierResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignBidModifierResult.Merge(m, src)
}
func (m *MutateCampaignBidModifierResult) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignBidModifierResult.Size(m)
}
func (m *MutateCampaignBidModifierResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignBidModifierResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignBidModifierResult proto.InternalMessageInfo

func (m *MutateCampaignBidModifierResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCampaignBidModifierRequest)(nil), "google.ads.googleads.v2.services.GetCampaignBidModifierRequest")
	proto.RegisterType((*MutateCampaignBidModifiersRequest)(nil), "google.ads.googleads.v2.services.MutateCampaignBidModifiersRequest")
	proto.RegisterType((*CampaignBidModifierOperation)(nil), "google.ads.googleads.v2.services.CampaignBidModifierOperation")
	proto.RegisterType((*MutateCampaignBidModifiersResponse)(nil), "google.ads.googleads.v2.services.MutateCampaignBidModifiersResponse")
	proto.RegisterType((*MutateCampaignBidModifierResult)(nil), "google.ads.googleads.v2.services.MutateCampaignBidModifierResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/campaign_bid_modifier_service.proto", fileDescriptor_6133a924995ecc19)
}

var fileDescriptor_6133a924995ecc19 = []byte{
	// 737 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0x4f, 0x6b, 0xd4, 0x4c,
	0x1c, 0xc7, 0x9f, 0x64, 0x1f, 0xfa, 0x3c, 0x9d, 0xad, 0x0a, 0x23, 0xea, 0x12, 0x2b, 0x5d, 0x63,
	0xc1, 0xb2, 0x87, 0x04, 0x22, 0x14, 0x4d, 0x69, 0x65, 0xb7, 0xeb, 0xb6, 0x1e, 0x6a, 0x4b, 0x0a,
	0x3d, 0xe8, 0x62, 0x98, 0x26, 0xb3, 0xcb, 0xd0, 0x24, 0x13, 0x67, 0x26, 0x0b, 0xa5, 0xf4, 0xe2,
	0xc1, 0x37, 0xe0, 0x3b, 0xf0, 0xe0, 0xc1, 0x77, 0x62, 0x6f, 0xe2, 0x5b, 0xf0, 0x24, 0x78, 0xf1,
	0x2e, 0x48, 0x32, 0x99, 0xed, 0x1f, 0x36, 0x5d, 0xb1, 0xb7, 0xdf, 0xce, 0x7c, 0xf3, 0xf9, 0xfd,
	0x9d, 0xdf, 0x82, 0xee, 0x90, 0xd2, 0x61, 0x84, 0x6d, 0x14, 0x72, 0x5b, 0x9a, 0xb9, 0x35, 0x72,
	0x6c, 0x8e, 0xd9, 0x88, 0x04, 0x98, 0xdb, 0x01, 0x8a, 0x53, 0x44, 0x86, 0x89, 0xbf, 0x4f, 0x42,
	0x3f, 0xa6, 0x21, 0x19, 0x10, 0xcc, 0xfc, 0xf2, 0xda, 0x4a, 0x19, 0x15, 0x14, 0x36, 0xe5, 0xa7,
	0x16, 0x0a, 0xb9, 0x35, 0xa6, 0x58, 0x23, 0xc7, 0x52, 0x14, 0x63, 0xb5, 0xca, 0x0f, 0xc3, 0x9c,
	0x66, 0xac, 0xd2, 0x91, 0x74, 0x60, 0xcc, 0xab, 0xcf, 0x53, 0x62, 0xa3, 0x24, 0xa1, 0x02, 0x09,
	0x42, 0x13, 0x5e, 0xde, 0x96, 0xee, 0xed, 0xe2, 0xd7, 0x7e, 0x36, 0xb0, 0x07, 0x04, 0x47, 0xa1,
	0x1f, 0x23, 0x7e, 0x50, 0x2a, 0xee, 0x94, 0x0a, 0x96, 0x06, 0x36, 0x17, 0x48, 0x64, 0xfc, 0xc2,
	0x45, 0x0e, 0x0e, 0x22, 0x82, 0x13, 0x21, 0x2f, 0xcc, 0x2e, 0xb8, 0xb7, 0x81, 0xc5, 0x7a, 0x19,
	0x53, 0x87, 0x84, 0x5b, 0x65, 0x44, 0x1e, 0x7e, 0x93, 0x61, 0x2e, 0xe0, 0x03, 0x70, 0x4d, 0xc5,
	0xee, 0x27, 0x28, 0xc6, 0x0d, 0xad, 0xa9, 0x2d, 0xcd, 0x7a, 0x73, 0xea, 0xf0, 0x05, 0x8a, 0xb1,
	0xf9, 0x53, 0x03, 0xf7, 0xb7, 0x32, 0x81, 0x04, 0x9e, 0x40, 0xe2, 0x0a, 0xb5, 0x00, 0xea, 0x41,
	0xc6, 0x05, 0x8d, 0x31, 0xf3, 0x49, 0x58, 0x82, 0x80, 0x3a, 0x7a, 0x1e, 0xc2, 0xd7, 0x00, 0xd0,
	0x14, 0x33, 0x99, 0x74, 0x43, 0x6f, 0xd6, 0x96, 0xea, 0xce, 0x9a, 0x35, 0xad, 0xe8, 0xd6, 0x04,
	0x9f, 0xdb, 0x0a, 0xe3, 0x9d, 0x21, 0xc2, 0x87, 0xe0, 0x46, 0x8a, 0x98, 0x20, 0x28, 0xf2, 0x07,
	0x88, 0x44, 0x19, 0xc3, 0x8d, 0x5a, 0x53, 0x5b, 0xfa, 0xdf, 0xbb, 0x5e, 0x1e, 0xf7, 0xe4, 0x69,
	0x9e, 0xf4, 0x08, 0x45, 0x24, 0x44, 0x02, 0xfb, 0x34, 0x89, 0x0e, 0x1b, 0xff, 0x16, 0xb2, 0x39,
	0x75, 0xb8, 0x9d, 0x44, 0x87, 0xe6, 0x47, 0x1d, 0xcc, 0x5f, 0xe6, 0x1a, 0xae, 0x80, 0x7a, 0x96,
	0x16, 0x8c, 0xbc, 0x45, 0x05, 0xa3, 0xee, 0x18, 0x2a, 0x1f, 0xd5, 0x45, 0xab, 0x97, 0x77, 0x71,
	0x0b, 0xf1, 0x03, 0x0f, 0x48, 0x79, 0x6e, 0xc3, 0x1d, 0x30, 0x13, 0x30, 0x8c, 0x84, 0x2c, 0x78,
	0xdd, 0x59, 0xae, 0xac, 0xc3, 0x78, 0xb4, 0x26, 0x15, 0x62, 0xf3, 0x1f, 0xaf, 0xe4, 0xe4, 0x44,
	0xc9, 0x6f, 0xe8, 0x57, 0x25, 0x4a, 0x0e, 0x6c, 0x80, 0x19, 0x86, 0x63, 0x3a, 0x92, 0x65, 0x9c,
	0xcd, 0x6f, 0xe4, 0xef, 0x4e, 0x1d, 0xcc, 0x8e, 0xeb, 0x6e, 0x7e, 0xd6, 0x80, 0x79, 0xd9, 0x74,
	0xf0, 0x94, 0x26, 0x1c, 0xc3, 0x1e, 0xb8, 0x75, 0xa1, 0x3b, 0x3e, 0x66, 0x8c, 0xb2, 0x02, 0x5e,
	0x77, 0xa0, 0x0a, 0x97, 0xa5, 0x81, 0xb5, 0x5b, 0x0c, 0xb7, 0x77, 0xf3, 0x7c, 0xdf, 0x9e, 0xe5,
	0x72, 0xf8, 0x0a, 0xfc, 0xc7, 0x30, 0xcf, 0x22, 0xa1, 0x46, 0xa8, 0x3d, 0x7d, 0x84, 0x2a, 0xc3,
	0xf3, 0x0a, 0x92, 0xa7, 0x88, 0x66, 0x0f, 0x2c, 0x4c, 0xd1, 0xfe, 0xd1, 0x8b, 0x71, 0x7e, 0xd5,
	0x80, 0x31, 0x01, 0xb1, 0x2b, 0x03, 0x82, 0x5f, 0x34, 0x70, 0x7b, 0xf2, 0xbb, 0x84, 0x4f, 0xa7,
	0x67, 0x73, 0xe9, 0x8b, 0x36, 0xfe, 0xb2, 0xef, 0xe6, 0xda, 0xdb, 0xaf, 0xdf, 0xde, 0xeb, 0x8f,
	0xe1, 0x72, 0xbe, 0xcf, 0x8e, 0xce, 0xa5, 0xb8, 0xaa, 0x1e, 0x31, 0xb7, 0x5b, 0xe3, 0x05, 0x77,
	0xb6, 0xc9, 0x76, 0xeb, 0x18, 0xfe, 0xd0, 0x80, 0x51, 0x3d, 0x06, 0x70, 0xfd, 0x0a, 0x5d, 0x52,
	0x2b, 0xc6, 0xe8, 0x5e, 0x0d, 0x22, 0x27, 0xd1, 0xec, 0x16, 0x99, 0xae, 0x99, 0x4f, 0xf2, 0x4c,
	0x4f, 0x53, 0x3b, 0x3a, 0xb3, 0xbd, 0x56, 0x5b, 0xc7, 0x13, 0x13, 0x75, 0xe3, 0x02, 0xef, 0x6a,
	0x2d, 0xe3, 0xee, 0x49, 0xbb, 0x71, 0x1a, 0x42, 0x69, 0xa5, 0x84, 0x5b, 0x01, 0x8d, 0x3b, 0xef,
	0x74, 0xb0, 0x18, 0xd0, 0x78, 0x6a, 0xb8, 0x9d, 0x85, 0xea, 0x29, 0xd9, 0xc9, 0x37, 0xc8, 0x8e,
	0xf6, 0x72, 0xb3, 0x84, 0x0c, 0x69, 0x84, 0x92, 0xa1, 0x45, 0xd9, 0xd0, 0x1e, 0xe2, 0xa4, 0xd8,
	0x2f, 0xf6, 0xa9, 0xdb, 0xea, 0xff, 0xbe, 0x15, 0x65, 0x7c, 0xd0, 0x6b, 0x1b, 0xed, 0xf6, 0x27,
	0xbd, 0xb9, 0x21, 0x81, 0xed, 0x90, 0x5b, 0xd2, 0xcc, 0xad, 0x3d, 0xc7, 0x2a, 0x1d, 0xf3, 0x13,
	0x25, 0xe9, 0xb7, 0x43, 0xde, 0x1f, 0x4b, 0xfa, 0x7b, 0x4e, 0x5f, 0x49, 0xbe, 0xeb, 0x8b, 0xf2,
	0xdc, 0x75, 0xdb, 0x21, 0x77, 0xdd, 0xb1, 0xc8, 0x75, 0xf7, 0x1c, 0xd7, 0x55, 0xb2, 0xfd, 0x99,
	0x22, 0xce, 0x47, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xb1, 0x2f, 0xdc, 0xa2, 0x07, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CampaignBidModifierServiceClient is the client API for CampaignBidModifierService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CampaignBidModifierServiceClient interface {
	// Returns the requested campaign bid modifier in full detail.
	GetCampaignBidModifier(ctx context.Context, in *GetCampaignBidModifierRequest, opts ...grpc.CallOption) (*resources.CampaignBidModifier, error)
	// Creates, updates, or removes campaign bid modifiers.
	// Operation statuses are returned.
	MutateCampaignBidModifiers(ctx context.Context, in *MutateCampaignBidModifiersRequest, opts ...grpc.CallOption) (*MutateCampaignBidModifiersResponse, error)
}

type campaignBidModifierServiceClient struct {
	cc *grpc.ClientConn
}

func NewCampaignBidModifierServiceClient(cc *grpc.ClientConn) CampaignBidModifierServiceClient {
	return &campaignBidModifierServiceClient{cc}
}

func (c *campaignBidModifierServiceClient) GetCampaignBidModifier(ctx context.Context, in *GetCampaignBidModifierRequest, opts ...grpc.CallOption) (*resources.CampaignBidModifier, error) {
	out := new(resources.CampaignBidModifier)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.CampaignBidModifierService/GetCampaignBidModifier", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campaignBidModifierServiceClient) MutateCampaignBidModifiers(ctx context.Context, in *MutateCampaignBidModifiersRequest, opts ...grpc.CallOption) (*MutateCampaignBidModifiersResponse, error) {
	out := new(MutateCampaignBidModifiersResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.CampaignBidModifierService/MutateCampaignBidModifiers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignBidModifierServiceServer is the server API for CampaignBidModifierService service.
type CampaignBidModifierServiceServer interface {
	// Returns the requested campaign bid modifier in full detail.
	GetCampaignBidModifier(context.Context, *GetCampaignBidModifierRequest) (*resources.CampaignBidModifier, error)
	// Creates, updates, or removes campaign bid modifiers.
	// Operation statuses are returned.
	MutateCampaignBidModifiers(context.Context, *MutateCampaignBidModifiersRequest) (*MutateCampaignBidModifiersResponse, error)
}

// UnimplementedCampaignBidModifierServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCampaignBidModifierServiceServer struct {
}

func (*UnimplementedCampaignBidModifierServiceServer) GetCampaignBidModifier(ctx context.Context, req *GetCampaignBidModifierRequest) (*resources.CampaignBidModifier, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetCampaignBidModifier not implemented")
}
func (*UnimplementedCampaignBidModifierServiceServer) MutateCampaignBidModifiers(ctx context.Context, req *MutateCampaignBidModifiersRequest) (*MutateCampaignBidModifiersResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateCampaignBidModifiers not implemented")
}

func RegisterCampaignBidModifierServiceServer(s *grpc.Server, srv CampaignBidModifierServiceServer) {
	s.RegisterService(&_CampaignBidModifierService_serviceDesc, srv)
}

func _CampaignBidModifierService_GetCampaignBidModifier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCampaignBidModifierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignBidModifierServiceServer).GetCampaignBidModifier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.CampaignBidModifierService/GetCampaignBidModifier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignBidModifierServiceServer).GetCampaignBidModifier(ctx, req.(*GetCampaignBidModifierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampaignBidModifierService_MutateCampaignBidModifiers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCampaignBidModifiersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignBidModifierServiceServer).MutateCampaignBidModifiers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.CampaignBidModifierService/MutateCampaignBidModifiers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignBidModifierServiceServer).MutateCampaignBidModifiers(ctx, req.(*MutateCampaignBidModifiersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CampaignBidModifierService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.CampaignBidModifierService",
	HandlerType: (*CampaignBidModifierServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCampaignBidModifier",
			Handler:    _CampaignBidModifierService_GetCampaignBidModifier_Handler,
		},
		{
			MethodName: "MutateCampaignBidModifiers",
			Handler:    _CampaignBidModifierService_MutateCampaignBidModifiers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/campaign_bid_modifier_service.proto",
}
