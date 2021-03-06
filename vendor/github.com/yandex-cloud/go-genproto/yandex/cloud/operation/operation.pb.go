// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/operation/operation.proto

package operation

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	status "google.golang.org/genproto/googleapis/rpc/status"
	math "math"
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

// An Operation resource. For more information, see [Operation](/docs/api-design-guide/concepts/operation).
type Operation struct {
	// ID of the operation.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Description of the operation. 0-256 characters long.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Creation timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// ID of the user or service account who initiated the operation.
	CreatedBy string `protobuf:"bytes,4,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	// The time when the Operation resource was last modified.
	ModifiedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=modified_at,json=modifiedAt,proto3" json:"modified_at,omitempty"`
	// If the value is `false`, it means the operation is still in progress.
	// If `true`, the operation is completed, and either `error` or `response` is available.
	Done bool `protobuf:"varint,6,opt,name=done,proto3" json:"done,omitempty"`
	// Service-specific metadata associated with the operation.
	// It typically contains the ID of the target resource that the operation is performed on.
	// Any method that returns a long-running operation should document the metadata type, if any.
	Metadata *any.Any `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// The operation result.
	// If `done == false` and there was no failure detected, neither `error` nor `response` is set.
	// If `done == false` and there was a failure detected, `error` is set.
	// If `done == true`, exactly one of `error` or `response` is set.
	//
	// Types that are valid to be assigned to Result:
	//	*Operation_Error
	//	*Operation_Response
	Result               isOperation_Result `protobuf_oneof:"result"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Operation) Reset()         { *m = Operation{} }
func (m *Operation) String() string { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()    {}
func (*Operation) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a0501b73987a137, []int{0}
}

func (m *Operation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Operation.Unmarshal(m, b)
}
func (m *Operation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Operation.Marshal(b, m, deterministic)
}
func (m *Operation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Operation.Merge(m, src)
}
func (m *Operation) XXX_Size() int {
	return xxx_messageInfo_Operation.Size(m)
}
func (m *Operation) XXX_DiscardUnknown() {
	xxx_messageInfo_Operation.DiscardUnknown(m)
}

var xxx_messageInfo_Operation proto.InternalMessageInfo

func (m *Operation) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Operation) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Operation) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Operation) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *Operation) GetModifiedAt() *timestamp.Timestamp {
	if m != nil {
		return m.ModifiedAt
	}
	return nil
}

func (m *Operation) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

func (m *Operation) GetMetadata() *any.Any {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type isOperation_Result interface {
	isOperation_Result()
}

type Operation_Error struct {
	Error *status.Status `protobuf:"bytes,8,opt,name=error,proto3,oneof"`
}

type Operation_Response struct {
	Response *any.Any `protobuf:"bytes,9,opt,name=response,proto3,oneof"`
}

func (*Operation_Error) isOperation_Result() {}

func (*Operation_Response) isOperation_Result() {}

func (m *Operation) GetResult() isOperation_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *Operation) GetError() *status.Status {
	if x, ok := m.GetResult().(*Operation_Error); ok {
		return x.Error
	}
	return nil
}

func (m *Operation) GetResponse() *any.Any {
	if x, ok := m.GetResult().(*Operation_Response); ok {
		return x.Response
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Operation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Operation_Error)(nil),
		(*Operation_Response)(nil),
	}
}

func init() {
	proto.RegisterType((*Operation)(nil), "yandex.cloud.operation.Operation")
}

func init() {
	proto.RegisterFile("yandex/cloud/operation/operation.proto", fileDescriptor_8a0501b73987a137)
}

var fileDescriptor_8a0501b73987a137 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xd7, 0xfd, 0xb2, 0xcd, 0xc0, 0x43, 0x10, 0x8d, 0x05, 0xb1, 0x78, 0x90, 0x21, 0x2c,
	0x95, 0x79, 0x92, 0x9d, 0x36, 0x3c, 0xec, 0x26, 0x54, 0x4f, 0x5e, 0x24, 0x6d, 0xb2, 0x1a, 0x58,
	0x9b, 0x90, 0xbe, 0x82, 0xbd, 0xfb, 0x87, 0x8b, 0xc9, 0xda, 0xf9, 0x13, 0x6f, 0x8f, 0x7e, 0x3f,
	0xef, 0xf3, 0x85, 0xe6, 0xa1, 0xcb, 0x86, 0x95, 0x5c, 0xbc, 0xc6, 0xd9, 0x56, 0xd5, 0x3c, 0x56,
	0x5a, 0x18, 0x06, 0x52, 0x95, 0xfb, 0x89, 0x6a, 0xa3, 0x40, 0xe1, 0x63, 0xc7, 0x51, 0xcb, 0xd1,
	0x2e, 0x0d, 0x4f, 0x73, 0xa5, 0xf2, 0xad, 0x88, 0x2d, 0x95, 0xd6, 0x9b, 0x98, 0x95, 0x8d, 0x5b,
	0x09, 0x4f, 0x76, 0x91, 0xd1, 0x59, 0x5c, 0x01, 0x83, 0xba, 0xda, 0x05, 0xe7, 0xdf, 0x77, 0x40,
	0x16, 0xa2, 0x02, 0x56, 0x68, 0x07, 0x5c, 0xbc, 0x0d, 0x50, 0x70, 0xdf, 0x56, 0xe0, 0x43, 0xd4,
	0x97, 0x9c, 0x78, 0x91, 0x37, 0x0d, 0x92, 0xbe, 0xe4, 0x38, 0x42, 0x13, 0x2e, 0xaa, 0xcc, 0x48,
	0xfd, 0x11, 0x93, 0xbe, 0x0d, 0x3e, 0x7f, 0xc2, 0xb7, 0x08, 0x65, 0x46, 0x30, 0x10, 0xfc, 0x99,
	0x01, 0x19, 0x44, 0xde, 0x74, 0x32, 0x0f, 0xa9, 0x6b, 0xa5, 0x6d, 0x2b, 0x7d, 0x6c, 0x5b, 0x93,
	0x60, 0x47, 0x2f, 0x01, 0x9f, 0xed, 0x57, 0xd3, 0x86, 0x0c, 0xad, 0xbb, 0x8d, 0x57, 0x0d, 0x5e,
	0xa0, 0x49, 0xa1, 0xb8, 0xdc, 0x48, 0xa7, 0x1e, 0xfd, 0xab, 0x46, 0x2d, 0xbe, 0x04, 0x8c, 0xd1,
	0x90, 0xab, 0x52, 0x90, 0x71, 0xe4, 0x4d, 0xfd, 0xc4, 0xce, 0xf8, 0x1a, 0xf9, 0x85, 0x00, 0xc6,
	0x19, 0x30, 0x72, 0x60, 0x6d, 0x47, 0x3f, 0x6c, 0xcb, 0xb2, 0x49, 0x3a, 0x0a, 0x5f, 0xa1, 0x91,
	0x30, 0x46, 0x19, 0xe2, 0x5b, 0x1c, 0xb7, 0xb8, 0xd1, 0x19, 0x7d, 0xb0, 0xbf, 0x79, 0xdd, 0x4b,
	0x1c, 0x82, 0xe7, 0xc8, 0x37, 0xa2, 0xd2, 0xaa, 0xac, 0x04, 0x09, 0xfe, 0xb6, 0xaf, 0x7b, 0x49,
	0xc7, 0xad, 0x7c, 0x34, 0x36, 0xa2, 0xaa, 0xb7, 0xb0, 0x4a, 0x51, 0xf8, 0xe5, 0xd5, 0x99, 0x96,
	0xfb, 0x97, 0x7f, 0xba, 0xcb, 0x25, 0xbc, 0xd4, 0x29, 0xcd, 0x54, 0x11, 0x3b, 0x6c, 0xe6, 0x8e,
	0x28, 0x57, 0xb3, 0x5c, 0x94, 0xd6, 0x1f, 0xff, 0x7e, 0x5d, 0x8b, 0x6e, 0x4a, 0xc7, 0x96, 0xbb,
	0x79, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x10, 0xa6, 0x51, 0xfd, 0x88, 0x02, 0x00, 0x00,
}
