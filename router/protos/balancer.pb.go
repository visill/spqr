// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/balancer.proto

package proto

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

type ReloadRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReloadRequest) Reset()         { *m = ReloadRequest{} }
func (m *ReloadRequest) String() string { return proto.CompactTextString(m) }
func (*ReloadRequest) ProtoMessage()    {}
func (*ReloadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b80f900138edf81a, []int{0}
}

func (m *ReloadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReloadRequest.Unmarshal(m, b)
}
func (m *ReloadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReloadRequest.Marshal(b, m, deterministic)
}
func (m *ReloadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReloadRequest.Merge(m, src)
}
func (m *ReloadRequest) XXX_Size() int {
	return xxx_messageInfo_ReloadRequest.Size(m)
}
func (m *ReloadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReloadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReloadRequest proto.InternalMessageInfo

type ReloadReply struct {
	ReloadRequired       bool     `protobuf:"varint,1,opt,name=reloadRequired,proto3" json:"reloadRequired,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReloadReply) Reset()         { *m = ReloadReply{} }
func (m *ReloadReply) String() string { return proto.CompactTextString(m) }
func (*ReloadReply) ProtoMessage()    {}
func (*ReloadReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_b80f900138edf81a, []int{1}
}

func (m *ReloadReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReloadReply.Unmarshal(m, b)
}
func (m *ReloadReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReloadReply.Marshal(b, m, deterministic)
}
func (m *ReloadReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReloadReply.Merge(m, src)
}
func (m *ReloadReply) XXX_Size() int {
	return xxx_messageInfo_ReloadReply.Size(m)
}
func (m *ReloadReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ReloadReply.DiscardUnknown(m)
}

var xxx_messageInfo_ReloadReply proto.InternalMessageInfo

func (m *ReloadReply) GetReloadRequired() bool {
	if m != nil {
		return m.ReloadRequired
	}
	return false
}

func init() {
	proto.RegisterType((*ReloadRequest)(nil), "yandex.spqr.ReloadRequest")
	proto.RegisterType((*ReloadReply)(nil), "yandex.spqr.ReloadReply")
}

func init() { proto.RegisterFile("protos/balancer.proto", fileDescriptor_b80f900138edf81a) }

var fileDescriptor_b80f900138edf81a = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x4f, 0x4a, 0xcc, 0x49, 0xcc, 0x4b, 0x4e, 0x2d, 0xd2, 0x03, 0xf3, 0x85, 0xb8,
	0x2b, 0x13, 0xf3, 0x52, 0x52, 0x2b, 0xf4, 0x8a, 0x0b, 0x0a, 0x8b, 0x94, 0xf8, 0xb9, 0x78, 0x83,
	0x52, 0x73, 0xf2, 0x13, 0x53, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x94, 0x4c, 0xb9, 0xb8,
	0x61, 0x02, 0x05, 0x39, 0x95, 0x42, 0x6a, 0x5c, 0x7c, 0x45, 0x70, 0xf9, 0xcc, 0xa2, 0xd4, 0x14,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0x8e, 0x20, 0x34, 0x51, 0xa3, 0x68, 0x2e, 0x7e, 0x27, 0xa8, 0x35,
	0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0x1e, 0x5c, 0x7c, 0x41, 0x28, 0x8a, 0x84, 0xa4,
	0xf4, 0x90, 0xac, 0xd6, 0x43, 0xb1, 0x57, 0x4a, 0x02, 0xab, 0x5c, 0x41, 0x4e, 0xa5, 0x12, 0x83,
	0x93, 0x70, 0x94, 0x20, 0x44, 0x52, 0x1f, 0x24, 0xa9, 0x0f, 0xf6, 0x46, 0x12, 0x1b, 0x98, 0x32,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xac, 0xbc, 0xc3, 0x38, 0xe6, 0x00, 0x00, 0x00,
}
