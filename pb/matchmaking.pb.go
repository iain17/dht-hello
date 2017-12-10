// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/matchmaking.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	pb/matchmaking.proto
	pb/platform.proto

It has these top-level messages:
	SessionInfo
	DNSessionResponse
	DNSessionRequest
	RPCUpsertSessionRequest
	RPCUpsertSessionResponse
	RPCDeleteSessionRequest
	RPCDeleteSessionResponse
	RPCSessionIdsRequest
	RPCSessionIdsResponse
	RPCGetSessionRequest
	RPCGetSessionResponse
	RPCHealthRequest
	RPCHealthReply
	RPCMessage
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SessionInfo struct {
	DnId      uint64            `protobuf:"varint,1,opt,name=dnId" json:"dnId,omitempty"`
	PId       string            `protobuf:"bytes,2,opt,name=pId" json:"pId,omitempty"`
	SessionId uint64            `protobuf:"varint,3,opt,name=sessionId" json:"sessionId,omitempty"`
	Type      uint64            `protobuf:"varint,4,opt,name=type" json:"type,omitempty"`
	Name      string            `protobuf:"bytes,5,opt,name=name" json:"name,omitempty"`
	Address   uint64            `protobuf:"varint,6,opt,name=address" json:"address,omitempty"`
	Port      uint32            `protobuf:"varint,7,opt,name=port" json:"port,omitempty"`
	Details   map[string]string `protobuf:"bytes,8,rep,name=details" json:"details,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *SessionInfo) Reset()                    { *m = SessionInfo{} }
func (m *SessionInfo) String() string            { return proto.CompactTextString(m) }
func (*SessionInfo) ProtoMessage()               {}
func (*SessionInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SessionInfo) GetDnId() uint64 {
	if m != nil {
		return m.DnId
	}
	return 0
}

func (m *SessionInfo) GetPId() string {
	if m != nil {
		return m.PId
	}
	return ""
}

func (m *SessionInfo) GetSessionId() uint64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *SessionInfo) GetType() uint64 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *SessionInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SessionInfo) GetAddress() uint64 {
	if m != nil {
		return m.Address
	}
	return 0
}

func (m *SessionInfo) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *SessionInfo) GetDetails() map[string]string {
	if m != nil {
		return m.Details
	}
	return nil
}

// DN messages
type DNSessionResponse struct {
	Results []*SessionInfo `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *DNSessionResponse) Reset()                    { *m = DNSessionResponse{} }
func (m *DNSessionResponse) String() string            { return proto.CompactTextString(m) }
func (*DNSessionResponse) ProtoMessage()               {}
func (*DNSessionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DNSessionResponse) GetResults() []*SessionInfo {
	if m != nil {
		return m.Results
	}
	return nil
}

type DNSessionRequest struct {
	Type uint64 `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
}

func (m *DNSessionRequest) Reset()                    { *m = DNSessionRequest{} }
func (m *DNSessionRequest) String() string            { return proto.CompactTextString(m) }
func (*DNSessionRequest) ProtoMessage()               {}
func (*DNSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DNSessionRequest) GetType() uint64 {
	if m != nil {
		return m.Type
	}
	return 0
}

// RPC messages
// Create/update
type RPCUpsertSessionRequest struct {
	Info *SessionInfo `protobuf:"bytes,1,opt,name=info" json:"info,omitempty"`
}

func (m *RPCUpsertSessionRequest) Reset()                    { *m = RPCUpsertSessionRequest{} }
func (m *RPCUpsertSessionRequest) String() string            { return proto.CompactTextString(m) }
func (*RPCUpsertSessionRequest) ProtoMessage()               {}
func (*RPCUpsertSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RPCUpsertSessionRequest) GetInfo() *SessionInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type RPCUpsertSessionResponse struct {
	Result    bool   `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
	SessionId uint64 `protobuf:"varint,2,opt,name=sessionId" json:"sessionId,omitempty"`
}

func (m *RPCUpsertSessionResponse) Reset()                    { *m = RPCUpsertSessionResponse{} }
func (m *RPCUpsertSessionResponse) String() string            { return proto.CompactTextString(m) }
func (*RPCUpsertSessionResponse) ProtoMessage()               {}
func (*RPCUpsertSessionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RPCUpsertSessionResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func (m *RPCUpsertSessionResponse) GetSessionId() uint64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

// Delete
type RPCDeleteSessionRequest struct {
	SessionId uint64 `protobuf:"varint,1,opt,name=sessionId" json:"sessionId,omitempty"`
}

func (m *RPCDeleteSessionRequest) Reset()                    { *m = RPCDeleteSessionRequest{} }
func (m *RPCDeleteSessionRequest) String() string            { return proto.CompactTextString(m) }
func (*RPCDeleteSessionRequest) ProtoMessage()               {}
func (*RPCDeleteSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RPCDeleteSessionRequest) GetSessionId() uint64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

type RPCDeleteSessionResponse struct {
	Result bool `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
}

func (m *RPCDeleteSessionResponse) Reset()                    { *m = RPCDeleteSessionResponse{} }
func (m *RPCDeleteSessionResponse) String() string            { return proto.CompactTextString(m) }
func (*RPCDeleteSessionResponse) ProtoMessage()               {}
func (*RPCDeleteSessionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RPCDeleteSessionResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

// Fetch session ids
type RPCSessionIdsRequest struct {
	Type    uint64            `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	Details map[string]string `protobuf:"bytes,2,rep,name=details" json:"details,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *RPCSessionIdsRequest) Reset()                    { *m = RPCSessionIdsRequest{} }
func (m *RPCSessionIdsRequest) String() string            { return proto.CompactTextString(m) }
func (*RPCSessionIdsRequest) ProtoMessage()               {}
func (*RPCSessionIdsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *RPCSessionIdsRequest) GetType() uint64 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *RPCSessionIdsRequest) GetDetails() map[string]string {
	if m != nil {
		return m.Details
	}
	return nil
}

type RPCSessionIdsResponse struct {
	SessionId []uint64 `protobuf:"varint,1,rep,packed,name=sessionId" json:"sessionId,omitempty"`
}

func (m *RPCSessionIdsResponse) Reset()                    { *m = RPCSessionIdsResponse{} }
func (m *RPCSessionIdsResponse) String() string            { return proto.CompactTextString(m) }
func (*RPCSessionIdsResponse) ProtoMessage()               {}
func (*RPCSessionIdsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *RPCSessionIdsResponse) GetSessionId() []uint64 {
	if m != nil {
		return m.SessionId
	}
	return nil
}

// Fetch an individual session
type RPCGetSessionRequest struct {
	SessionId uint64 `protobuf:"varint,1,opt,name=sessionId" json:"sessionId,omitempty"`
}

func (m *RPCGetSessionRequest) Reset()                    { *m = RPCGetSessionRequest{} }
func (m *RPCGetSessionRequest) String() string            { return proto.CompactTextString(m) }
func (*RPCGetSessionRequest) ProtoMessage()               {}
func (*RPCGetSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *RPCGetSessionRequest) GetSessionId() uint64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

type RPCGetSessionResponse struct {
	Found  bool         `protobuf:"varint,1,opt,name=found" json:"found,omitempty"`
	Result *SessionInfo `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
}

func (m *RPCGetSessionResponse) Reset()                    { *m = RPCGetSessionResponse{} }
func (m *RPCGetSessionResponse) String() string            { return proto.CompactTextString(m) }
func (*RPCGetSessionResponse) ProtoMessage()               {}
func (*RPCGetSessionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *RPCGetSessionResponse) GetFound() bool {
	if m != nil {
		return m.Found
	}
	return false
}

func (m *RPCGetSessionResponse) GetResult() *SessionInfo {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*SessionInfo)(nil), "pb.SessionInfo")
	proto.RegisterType((*DNSessionResponse)(nil), "pb.DNSessionResponse")
	proto.RegisterType((*DNSessionRequest)(nil), "pb.DNSessionRequest")
	proto.RegisterType((*RPCUpsertSessionRequest)(nil), "pb.RPCUpsertSessionRequest")
	proto.RegisterType((*RPCUpsertSessionResponse)(nil), "pb.RPCUpsertSessionResponse")
	proto.RegisterType((*RPCDeleteSessionRequest)(nil), "pb.RPCDeleteSessionRequest")
	proto.RegisterType((*RPCDeleteSessionResponse)(nil), "pb.RPCDeleteSessionResponse")
	proto.RegisterType((*RPCSessionIdsRequest)(nil), "pb.RPCSessionIdsRequest")
	proto.RegisterType((*RPCSessionIdsResponse)(nil), "pb.RPCSessionIdsResponse")
	proto.RegisterType((*RPCGetSessionRequest)(nil), "pb.RPCGetSessionRequest")
	proto.RegisterType((*RPCGetSessionResponse)(nil), "pb.RPCGetSessionResponse")
}

func init() { proto.RegisterFile("pb/matchmaking.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x4d, 0x8b, 0x13, 0x41,
	0x10, 0xa5, 0x67, 0xf2, 0xb1, 0xa9, 0x55, 0x5c, 0x9b, 0xa8, 0x8d, 0xec, 0x21, 0x8c, 0xa8, 0xf1,
	0x12, 0x61, 0xfd, 0x64, 0x0f, 0x2b, 0xb8, 0x11, 0xcd, 0x45, 0x42, 0x8b, 0xde, 0x27, 0x76, 0x45,
	0xc3, 0x26, 0xdd, 0xed, 0x74, 0x47, 0xc8, 0x6f, 0xf1, 0x07, 0xf8, 0x37, 0xa5, 0x6b, 0x3a, 0x9b,
	0xa4, 0x0d, 0x01, 0x0f, 0xde, 0xaa, 0xa6, 0xea, 0x55, 0xbd, 0xf7, 0xba, 0x18, 0xe8, 0xda, 0xc9,
	0xd3, 0x45, 0xe9, 0xbf, 0x7e, 0x5f, 0x94, 0x57, 0x33, 0xfd, 0x6d, 0x60, 0x2b, 0xe3, 0x0d, 0xcf,
	0xec, 0xa4, 0xf8, 0x95, 0xc1, 0xf1, 0x27, 0x74, 0x6e, 0x66, 0xf4, 0x48, 0x4f, 0x0d, 0xe7, 0xd0,
	0x50, 0x7a, 0xa4, 0x04, 0xeb, 0xb1, 0x7e, 0x43, 0x52, 0xcc, 0x4f, 0x20, 0xb7, 0x23, 0x25, 0xb2,
	0x1e, 0xeb, 0x77, 0x64, 0x08, 0xf9, 0x29, 0x74, 0x5c, 0x04, 0x29, 0x91, 0x53, 0xeb, 0xe6, 0x43,
	0x98, 0xe1, 0x57, 0x16, 0x45, 0xa3, 0x9e, 0x11, 0xe2, 0xf0, 0x4d, 0x97, 0x0b, 0x14, 0x4d, 0x1a,
	0x42, 0x31, 0x17, 0xd0, 0x2e, 0x95, 0xaa, 0xd0, 0x39, 0xd1, 0xa2, 0xd6, 0x75, 0x1a, 0xba, 0xad,
	0xa9, 0xbc, 0x68, 0xf7, 0x58, 0xff, 0xa6, 0xa4, 0x98, 0xbf, 0x84, 0xb6, 0x42, 0x5f, 0xce, 0xe6,
	0x4e, 0x1c, 0xf5, 0xf2, 0xfe, 0xf1, 0xd9, 0xe9, 0xc0, 0x4e, 0x06, 0x5b, 0xdc, 0x07, 0xc3, 0xba,
	0xfc, 0x4e, 0xfb, 0x6a, 0x25, 0xd7, 0xcd, 0xf7, 0xcf, 0xe1, 0xc6, 0x76, 0x21, 0xa8, 0xb9, 0xc2,
	0x15, 0x09, 0xec, 0xc8, 0x10, 0xf2, 0x2e, 0x34, 0x7f, 0x96, 0xf3, 0x25, 0x46, 0x85, 0x75, 0x72,
	0x9e, 0xbd, 0x66, 0xc5, 0x05, 0xdc, 0x1e, 0x7e, 0x8c, 0x2b, 0x24, 0x3a, 0x6b, 0xb4, 0x43, 0xfe,
	0x04, 0xda, 0x15, 0xba, 0xe5, 0xdc, 0x3b, 0xc1, 0x88, 0xc8, 0xad, 0x84, 0x88, 0x5c, 0xd7, 0x8b,
	0x47, 0x70, 0xb2, 0x85, 0xff, 0xb1, 0x44, 0xe7, 0xaf, 0xdd, 0x61, 0x1b, 0x77, 0x8a, 0x0b, 0xb8,
	0x27, 0xc7, 0x97, 0x9f, 0xad, 0xc3, 0xca, 0x27, 0xed, 0x0f, 0xa0, 0x31, 0xd3, 0x53, 0x43, 0xed,
	0x7b, 0x56, 0x51, 0xb1, 0x18, 0x83, 0xf8, 0x1b, 0x1f, 0xe9, 0xde, 0x85, 0x56, 0x4d, 0x87, 0x46,
	0x1c, 0xc9, 0x98, 0xed, 0xbe, 0x61, 0x96, 0xbc, 0x61, 0xf1, 0x8a, 0x18, 0x0d, 0x71, 0x8e, 0x1e,
	0x13, 0x46, 0x3b, 0x40, 0x96, 0x02, 0xcf, 0x88, 0x4a, 0x02, 0x3c, 0x4c, 0xa5, 0xf8, 0xcd, 0xa0,
	0x2b, 0xc7, 0x97, 0x6b, 0x5d, 0xca, 0x1d, 0xf0, 0x8a, 0xbf, 0xd9, 0xdc, 0x41, 0x46, 0xf6, 0x3f,
	0x0c, 0x9e, 0xec, 0x83, 0xff, 0x87, 0x83, 0x78, 0x01, 0x77, 0x92, 0x4d, 0x51, 0x5a, 0x62, 0x4a,
	0xbe, 0x6b, 0xca, 0x73, 0xd2, 0xf7, 0x1e, 0xfd, 0x3f, 0x59, 0xf9, 0x85, 0x96, 0x6d, 0xa3, 0xe2,
	0xb2, 0x2e, 0x34, 0xa7, 0x66, 0xa9, 0x55, 0xb4, 0xb1, 0x4e, 0xf8, 0xe3, 0x6b, 0x77, 0xb3, 0xfd,
	0xb7, 0x12, 0xcb, 0x6f, 0xb3, 0x0f, 0xf9, 0xa4, 0x45, 0xbf, 0x80, 0x67, 0x7f, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x32, 0xf3, 0x61, 0x3f, 0x1a, 0x04, 0x00, 0x00,
}
