// Code generated by protoc-gen-go. DO NOT EDIT.
// source: groupcache.proto

package groupcachepb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type GetRequest struct {
	Group                *string  `protobuf:"bytes,1,req,name=group" json:"group,omitempty"`
	Key                  *string  `protobuf:"bytes,2,req,name=key" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aba3e029bfa16198, []int{0}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetGroup() string {
	if m != nil && m.Group != nil {
		return *m.Group
	}
	return ""
}

func (m *GetRequest) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

type GetResponse struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	MinuteQps            *float64 `protobuf:"fixed64,2,opt,name=minute_qps,json=minuteQps" json:"minute_qps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aba3e029bfa16198, []int{1}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *GetResponse) GetMinuteQps() float64 {
	if m != nil && m.MinuteQps != nil {
		return *m.MinuteQps
	}
	return 0
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "groupcachepb.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "groupcachepb.GetResponse")
}

func init() { proto.RegisterFile("groupcache.proto", fileDescriptor_aba3e029bfa16198) }

var fileDescriptor_aba3e029bfa16198 = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x2f, 0xca, 0x2f,
	0x2d, 0x48, 0x4e, 0x4c, 0xce, 0x48, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x41, 0x88,
	0x14, 0x24, 0x29, 0x99, 0x70, 0x71, 0xb9, 0xa7, 0x96, 0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97,
	0x08, 0x89, 0x70, 0xb1, 0x82, 0x65, 0x25, 0x18, 0x15, 0x98, 0x34, 0x38, 0x83, 0x20, 0x1c, 0x21,
	0x01, 0x2e, 0xe6, 0xec, 0xd4, 0x4a, 0x09, 0x26, 0xb0, 0x18, 0x88, 0xa9, 0xe4, 0xc4, 0xc5, 0x0d,
	0xd6, 0x55, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x0a, 0xd2, 0x56, 0x96, 0x98, 0x53, 0x9a, 0x2a, 0xc1,
	0xa8, 0xc0, 0xa8, 0xc1, 0x13, 0x04, 0xe1, 0x08, 0xc9, 0x72, 0x71, 0xe5, 0x66, 0xe6, 0x95, 0x96,
	0xa4, 0xc6, 0x17, 0x16, 0x14, 0x4b, 0x30, 0x29, 0x30, 0x6a, 0x30, 0x06, 0x71, 0x42, 0x44, 0x02,
	0x0b, 0x8a, 0x8d, 0xbc, 0xb8, 0xb8, 0xdc, 0x41, 0xc6, 0x3b, 0x83, 0x5c, 0x22, 0x64, 0xc3, 0xc5,
	0xec, 0x9e, 0x5a, 0x22, 0x24, 0xa1, 0x87, 0xec, 0x3a, 0x3d, 0x84, 0xd3, 0xa4, 0x24, 0xb1, 0xc8,
	0x40, 0xac, 0x57, 0x62, 0x00, 0x04, 0x00, 0x00, 0xff, 0xff, 0x85, 0xb0, 0x33, 0x9f, 0xe6, 0x00,
	0x00, 0x00,
}
