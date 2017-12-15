// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/addressbook.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	pb/addressbook.proto
	pb/matchmaking.proto
	pb/platform.proto

It has these top-level messages:
	Peer
	DNPeerResponse
	DNPeerRequest
	RPCUpsertPeerRequest
	RPCUpsertPeerResponse
	RPCGetPeerIdsRequest
	RPCGetPeerIdsResponse
	RPCGetPeerRequest
	RPCGetPeerResponse
	Session
	DNSessionResponse
	DNSessionRequest
	RPCUpsertSessionRequest
	RPCUpsertSessionResponse
	RPCDeleteSessionRequest
	RPCDeleteSessionResponse
	RPCGetSessionIdsRequest
	RPCGetSessionIdsResponse
	RPCGetSessionRequest
	RPCGetSessionResponse
	RPCHealthRequest
	RPCHealthReply
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

type Peer struct {
	// If pId == "self". We'll resolve it to our peerId.
	PId     string            `protobuf:"bytes,1,opt,name=pId" json:"pId,omitempty"`
	DnId    uint64            `protobuf:"varint,2,opt,name=dnId" json:"dnId,omitempty"`
	Details map[string]string `protobuf:"bytes,3,rep,name=details" json:"details,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Peer) Reset()                    { *m = Peer{} }
func (m *Peer) String() string            { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()               {}
func (*Peer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Peer) GetPId() string {
	if m != nil {
		return m.PId
	}
	return ""
}

func (m *Peer) GetDnId() uint64 {
	if m != nil {
		return m.DnId
	}
	return 0
}

func (m *Peer) GetDetails() map[string]string {
	if m != nil {
		return m.Details
	}
	return nil
}

type DNPeerResponse struct {
	Peer *Peer `protobuf:"bytes,1,opt,name=peer" json:"peer,omitempty"`
}

func (m *DNPeerResponse) Reset()                    { *m = DNPeerResponse{} }
func (m *DNPeerResponse) String() string            { return proto.CompactTextString(m) }
func (*DNPeerResponse) ProtoMessage()               {}
func (*DNPeerResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DNPeerResponse) GetPeer() *Peer {
	if m != nil {
		return m.Peer
	}
	return nil
}

type DNPeerRequest struct {
}

func (m *DNPeerRequest) Reset()                    { *m = DNPeerRequest{} }
func (m *DNPeerRequest) String() string            { return proto.CompactTextString(m) }
func (*DNPeerRequest) ProtoMessage()               {}
func (*DNPeerRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// upsert
type RPCUpsertPeerRequest struct {
	Peer *Peer `protobuf:"bytes,1,opt,name=peer" json:"peer,omitempty"`
}

func (m *RPCUpsertPeerRequest) Reset()                    { *m = RPCUpsertPeerRequest{} }
func (m *RPCUpsertPeerRequest) String() string            { return proto.CompactTextString(m) }
func (*RPCUpsertPeerRequest) ProtoMessage()               {}
func (*RPCUpsertPeerRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RPCUpsertPeerRequest) GetPeer() *Peer {
	if m != nil {
		return m.Peer
	}
	return nil
}

type RPCUpsertPeerResponse struct {
	Result bool `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
}

func (m *RPCUpsertPeerResponse) Reset()                    { *m = RPCUpsertPeerResponse{} }
func (m *RPCUpsertPeerResponse) String() string            { return proto.CompactTextString(m) }
func (*RPCUpsertPeerResponse) ProtoMessage()               {}
func (*RPCUpsertPeerResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RPCUpsertPeerResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

// Get peer ids
type RPCGetPeerIdsRequest struct {
	// Key in the details map
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	// Value in the details map.
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *RPCGetPeerIdsRequest) Reset()                    { *m = RPCGetPeerIdsRequest{} }
func (m *RPCGetPeerIdsRequest) String() string            { return proto.CompactTextString(m) }
func (*RPCGetPeerIdsRequest) ProtoMessage()               {}
func (*RPCGetPeerIdsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RPCGetPeerIdsRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *RPCGetPeerIdsRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type RPCGetPeerIdsResponse struct {
	PeerIds []string `protobuf:"bytes,1,rep,name=peerIds" json:"peerIds,omitempty"`
}

func (m *RPCGetPeerIdsResponse) Reset()                    { *m = RPCGetPeerIdsResponse{} }
func (m *RPCGetPeerIdsResponse) String() string            { return proto.CompactTextString(m) }
func (*RPCGetPeerIdsResponse) ProtoMessage()               {}
func (*RPCGetPeerIdsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RPCGetPeerIdsResponse) GetPeerIds() []string {
	if m != nil {
		return m.PeerIds
	}
	return nil
}

// Get individual peer
type RPCGetPeerRequest struct {
	// If pId == "self". We'll resolve it to our peerId.
	// Either one must be filled in. If both are specified. preference goes to pId
	PId  string `protobuf:"bytes,1,opt,name=pId" json:"pId,omitempty"`
	DnId uint64 `protobuf:"varint,2,opt,name=dnId" json:"dnId,omitempty"`
}

func (m *RPCGetPeerRequest) Reset()                    { *m = RPCGetPeerRequest{} }
func (m *RPCGetPeerRequest) String() string            { return proto.CompactTextString(m) }
func (*RPCGetPeerRequest) ProtoMessage()               {}
func (*RPCGetPeerRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *RPCGetPeerRequest) GetPId() string {
	if m != nil {
		return m.PId
	}
	return ""
}

func (m *RPCGetPeerRequest) GetDnId() uint64 {
	if m != nil {
		return m.DnId
	}
	return 0
}

type RPCGetPeerResponse struct {
	Peer *Peer `protobuf:"bytes,1,opt,name=peer" json:"peer,omitempty"`
}

func (m *RPCGetPeerResponse) Reset()                    { *m = RPCGetPeerResponse{} }
func (m *RPCGetPeerResponse) String() string            { return proto.CompactTextString(m) }
func (*RPCGetPeerResponse) ProtoMessage()               {}
func (*RPCGetPeerResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *RPCGetPeerResponse) GetPeer() *Peer {
	if m != nil {
		return m.Peer
	}
	return nil
}

func init() {
	proto.RegisterType((*Peer)(nil), "pb.Peer")
	proto.RegisterType((*DNPeerResponse)(nil), "pb.DNPeerResponse")
	proto.RegisterType((*DNPeerRequest)(nil), "pb.DNPeerRequest")
	proto.RegisterType((*RPCUpsertPeerRequest)(nil), "pb.RPCUpsertPeerRequest")
	proto.RegisterType((*RPCUpsertPeerResponse)(nil), "pb.RPCUpsertPeerResponse")
	proto.RegisterType((*RPCGetPeerIdsRequest)(nil), "pb.RPCGetPeerIdsRequest")
	proto.RegisterType((*RPCGetPeerIdsResponse)(nil), "pb.RPCGetPeerIdsResponse")
	proto.RegisterType((*RPCGetPeerRequest)(nil), "pb.RPCGetPeerRequest")
	proto.RegisterType((*RPCGetPeerResponse)(nil), "pb.RPCGetPeerResponse")
}

func init() { proto.RegisterFile("pb/addressbook.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcd, 0x4e, 0x84, 0x30,
	0x10, 0xc7, 0xd3, 0x05, 0xf7, 0x63, 0xd6, 0xcf, 0x66, 0xd7, 0x10, 0xe3, 0x81, 0xf4, 0xc4, 0x89,
	0x8d, 0xe8, 0x41, 0xf7, 0xe0, 0xc5, 0x35, 0x86, 0x8b, 0xd9, 0x34, 0xf1, 0x01, 0x20, 0x9d, 0x83,
	0x81, 0x40, 0x6d, 0x8b, 0x09, 0x8f, 0xe2, 0xdb, 0x1a, 0x0a, 0x8d, 0xc4, 0xcb, 0xee, 0x6d, 0xe6,
	0xcf, 0xfc, 0xe6, 0x17, 0xda, 0xc2, 0x4a, 0xe6, 0x9b, 0x4c, 0x08, 0x85, 0x5a, 0xe7, 0x75, 0x5d,
	0xc4, 0x52, 0xd5, 0xa6, 0xa6, 0x13, 0x99, 0xb3, 0x1f, 0x02, 0xfe, 0x1e, 0x51, 0xd1, 0x4b, 0xf0,
	0x64, 0x2a, 0x02, 0x12, 0x92, 0x68, 0xc1, 0xbb, 0x92, 0x52, 0xf0, 0x45, 0x95, 0x8a, 0x60, 0x12,
	0x92, 0xc8, 0xe7, 0xb6, 0xa6, 0x1b, 0x98, 0x09, 0x34, 0xd9, 0x67, 0xa9, 0x03, 0x2f, 0xf4, 0xa2,
	0x65, 0xb2, 0x8e, 0x65, 0x1e, 0x77, 0x0b, 0xe2, 0x5d, 0x9f, 0xbf, 0x56, 0x46, 0xb5, 0xdc, 0x4d,
	0xdd, 0x6c, 0xe1, 0x74, 0xfc, 0xa1, 0xd3, 0x14, 0xd8, 0x3a, 0x4d, 0x81, 0x2d, 0x5d, 0xc1, 0xc9,
	0x77, 0x56, 0x36, 0x68, 0x3d, 0x0b, 0xde, 0x37, 0xdb, 0xc9, 0x23, 0x61, 0x31, 0x9c, 0xef, 0xde,
	0xbb, 0xdd, 0x1c, 0xb5, 0xac, 0x2b, 0x8d, 0xf4, 0x16, 0x7c, 0x89, 0xa8, 0x2c, 0xbe, 0x4c, 0xe6,
	0xce, 0xcd, 0x6d, 0xca, 0x2e, 0xe0, 0xcc, 0xcd, 0x7f, 0x35, 0xa8, 0x0d, 0x7b, 0x80, 0x15, 0xdf,
	0xbf, 0x7c, 0x48, 0x8d, 0xca, 0x8c, 0xf2, 0x03, 0x6b, 0x36, 0xb0, 0xfe, 0x47, 0x0d, 0xf6, 0x6b,
	0x98, 0x2a, 0xd4, 0x4d, 0x69, 0x2c, 0x38, 0xe7, 0x43, 0xc7, 0x9e, 0xad, 0xe6, 0x0d, 0xed, 0x74,
	0x2a, 0xb4, 0xd3, 0x1c, 0xf9, 0xaf, 0xec, 0xce, 0x0a, 0xc7, 0xfc, 0x20, 0x0c, 0x60, 0x26, 0xfb,
	0x28, 0x20, 0xa1, 0x17, 0x2d, 0xb8, 0x6b, 0xd9, 0x13, 0x5c, 0xfd, 0x21, 0x23, 0xdf, 0xe1, 0x2b,
	0x64, 0x09, 0xd0, 0x31, 0x7a, 0xcc, 0xc9, 0xe6, 0x53, 0xfb, 0x60, 0xee, 0x7f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x20, 0x1d, 0x97, 0xec, 0x48, 0x02, 0x00, 0x00,
}
