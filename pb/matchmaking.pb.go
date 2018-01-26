// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/matchmaking.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Session struct {
	PId       string            `protobuf:"bytes,1,opt,name=pId" json:"pId,omitempty"`
	DnId      uint64            `protobuf:"varint,2,opt,name=dnId" json:"dnId,omitempty"`
	SessionId uint64            `protobuf:"varint,3,opt,name=sessionId" json:"sessionId,omitempty"`
	Type      uint64            `protobuf:"varint,4,opt,name=type" json:"type,omitempty"`
	Name      string            `protobuf:"bytes,5,opt,name=name" json:"name,omitempty"`
	Address   uint32            `protobuf:"varint,6,opt,name=address" json:"address,omitempty"`
	Port      uint32            `protobuf:"varint,7,opt,name=port" json:"port,omitempty"`
	Details   map[string]string `protobuf:"bytes,8,rep,name=details" json:"details,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Published uint64            `protobuf:"varint,9,opt,name=published" json:"published,omitempty"`
}

func (m *Session) Reset()                    { *m = Session{} }
func (m *Session) String() string            { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()               {}
func (*Session) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Session) GetPId() string {
	if m != nil {
		return m.PId
	}
	return ""
}

func (m *Session) GetDnId() uint64 {
	if m != nil {
		return m.DnId
	}
	return 0
}

func (m *Session) GetSessionId() uint64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *Session) GetType() uint64 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Session) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Session) GetAddress() uint32 {
	if m != nil {
		return m.Address
	}
	return 0
}

func (m *Session) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Session) GetDetails() map[string]string {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *Session) GetPublished() uint64 {
	if m != nil {
		return m.Published
	}
	return 0
}

// DN messages
type DNSessionStore struct {
	Sessions []*Session `protobuf:"bytes,1,rep,name=sessions" json:"sessions,omitempty"`
}

func (m *DNSessionStore) Reset()                    { *m = DNSessionStore{} }
func (m *DNSessionStore) String() string            { return proto.CompactTextString(m) }
func (*DNSessionStore) ProtoMessage()               {}
func (*DNSessionStore) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *DNSessionStore) GetSessions() []*Session {
	if m != nil {
		return m.Sessions
	}
	return nil
}

type DNSessionsRecord struct {
	Published uint64     `protobuf:"varint,1,opt,name=published" json:"published,omitempty"`
	Results   []*Session `protobuf:"bytes,2,rep,name=results" json:"results,omitempty"`
}

func (m *DNSessionsRecord) Reset()                    { *m = DNSessionsRecord{} }
func (m *DNSessionsRecord) String() string            { return proto.CompactTextString(m) }
func (*DNSessionsRecord) ProtoMessage()               {}
func (*DNSessionsRecord) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *DNSessionsRecord) GetPublished() uint64 {
	if m != nil {
		return m.Published
	}
	return 0
}

func (m *DNSessionsRecord) GetResults() []*Session {
	if m != nil {
		return m.Results
	}
	return nil
}

type DNSessionsResponse struct {
	Results []*Session `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *DNSessionsResponse) Reset()                    { *m = DNSessionsResponse{} }
func (m *DNSessionsResponse) String() string            { return proto.CompactTextString(m) }
func (*DNSessionsResponse) ProtoMessage()               {}
func (*DNSessionsResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *DNSessionsResponse) GetResults() []*Session {
	if m != nil {
		return m.Results
	}
	return nil
}

type DNSessionsRequest struct {
	Type       uint64   `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	SessionIds []uint64 `protobuf:"varint,2,rep,packed,name=sessionIds" json:"sessionIds,omitempty"`
}

func (m *DNSessionsRequest) Reset()                    { *m = DNSessionsRequest{} }
func (m *DNSessionsRequest) String() string            { return proto.CompactTextString(m) }
func (*DNSessionsRequest) ProtoMessage()               {}
func (*DNSessionsRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *DNSessionsRequest) GetType() uint64 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *DNSessionsRequest) GetSessionIds() []uint64 {
	if m != nil {
		return m.SessionIds
	}
	return nil
}

type DNSessionidsResponse struct {
	SessionIds []uint64 `protobuf:"varint,1,rep,packed,name=sessionIds" json:"sessionIds,omitempty"`
}

func (m *DNSessionidsResponse) Reset()                    { *m = DNSessionidsResponse{} }
func (m *DNSessionidsResponse) String() string            { return proto.CompactTextString(m) }
func (*DNSessionidsResponse) ProtoMessage()               {}
func (*DNSessionidsResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *DNSessionidsResponse) GetSessionIds() []uint64 {
	if m != nil {
		return m.SessionIds
	}
	return nil
}

type DNSessionidsRequest struct {
	Type uint64 `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
}

func (m *DNSessionidsRequest) Reset()                    { *m = DNSessionidsRequest{} }
func (m *DNSessionidsRequest) String() string            { return proto.CompactTextString(m) }
func (*DNSessionidsRequest) ProtoMessage()               {}
func (*DNSessionidsRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *DNSessionidsRequest) GetType() uint64 {
	if m != nil {
		return m.Type
	}
	return 0
}

// Upsert
type RPCUpsertSessionRequest struct {
	Session *Session `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
}

func (m *RPCUpsertSessionRequest) Reset()                    { *m = RPCUpsertSessionRequest{} }
func (m *RPCUpsertSessionRequest) String() string            { return proto.CompactTextString(m) }
func (*RPCUpsertSessionRequest) ProtoMessage()               {}
func (*RPCUpsertSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *RPCUpsertSessionRequest) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

type RPCUpsertSessionResponse struct {
	SessionId uint64 `protobuf:"varint,1,opt,name=sessionId" json:"sessionId,omitempty"`
}

func (m *RPCUpsertSessionResponse) Reset()                    { *m = RPCUpsertSessionResponse{} }
func (m *RPCUpsertSessionResponse) String() string            { return proto.CompactTextString(m) }
func (*RPCUpsertSessionResponse) ProtoMessage()               {}
func (*RPCUpsertSessionResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

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
func (*RPCDeleteSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

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
func (*RPCDeleteSessionResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func (m *RPCDeleteSessionResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

// Get session ids
type RPCGetSessionIdsByDetailsRequest struct {
	Type  uint64 `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	Key   string `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
}

func (m *RPCGetSessionIdsByDetailsRequest) Reset()         { *m = RPCGetSessionIdsByDetailsRequest{} }
func (m *RPCGetSessionIdsByDetailsRequest) String() string { return proto.CompactTextString(m) }
func (*RPCGetSessionIdsByDetailsRequest) ProtoMessage()    {}
func (*RPCGetSessionIdsByDetailsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{11}
}

func (m *RPCGetSessionIdsByDetailsRequest) GetType() uint64 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *RPCGetSessionIdsByDetailsRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *RPCGetSessionIdsByDetailsRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type RPCGetSessionIdsByPeerIdsRequest struct {
	PeerIds []string `protobuf:"bytes,1,rep,name=peerIds" json:"peerIds,omitempty"`
}

func (m *RPCGetSessionIdsByPeerIdsRequest) Reset()         { *m = RPCGetSessionIdsByPeerIdsRequest{} }
func (m *RPCGetSessionIdsByPeerIdsRequest) String() string { return proto.CompactTextString(m) }
func (*RPCGetSessionIdsByPeerIdsRequest) ProtoMessage()    {}
func (*RPCGetSessionIdsByPeerIdsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{12}
}

func (m *RPCGetSessionIdsByPeerIdsRequest) GetPeerIds() []string {
	if m != nil {
		return m.PeerIds
	}
	return nil
}

type RPCGetSessionIdsResponse struct {
	SessionIds []uint64 `protobuf:"varint,1,rep,packed,name=sessionIds" json:"sessionIds,omitempty"`
}

func (m *RPCGetSessionIdsResponse) Reset()                    { *m = RPCGetSessionIdsResponse{} }
func (m *RPCGetSessionIdsResponse) String() string            { return proto.CompactTextString(m) }
func (*RPCGetSessionIdsResponse) ProtoMessage()               {}
func (*RPCGetSessionIdsResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{13} }

func (m *RPCGetSessionIdsResponse) GetSessionIds() []uint64 {
	if m != nil {
		return m.SessionIds
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
func (*RPCGetSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{14} }

func (m *RPCGetSessionRequest) GetSessionId() uint64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

type RPCGetSessionResponse struct {
	Session *Session `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
}

func (m *RPCGetSessionResponse) Reset()                    { *m = RPCGetSessionResponse{} }
func (m *RPCGetSessionResponse) String() string            { return proto.CompactTextString(m) }
func (*RPCGetSessionResponse) ProtoMessage()               {}
func (*RPCGetSessionResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{15} }

func (m *RPCGetSessionResponse) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

func init() {
	proto.RegisterType((*Session)(nil), "pb.Session")
	proto.RegisterType((*DNSessionStore)(nil), "pb.DNSessionStore")
	proto.RegisterType((*DNSessionsRecord)(nil), "pb.DNSessionsRecord")
	proto.RegisterType((*DNSessionsResponse)(nil), "pb.DNSessionsResponse")
	proto.RegisterType((*DNSessionsRequest)(nil), "pb.DNSessionsRequest")
	proto.RegisterType((*DNSessionidsResponse)(nil), "pb.DNSessionidsResponse")
	proto.RegisterType((*DNSessionidsRequest)(nil), "pb.DNSessionidsRequest")
	proto.RegisterType((*RPCUpsertSessionRequest)(nil), "pb.RPCUpsertSessionRequest")
	proto.RegisterType((*RPCUpsertSessionResponse)(nil), "pb.RPCUpsertSessionResponse")
	proto.RegisterType((*RPCDeleteSessionRequest)(nil), "pb.RPCDeleteSessionRequest")
	proto.RegisterType((*RPCDeleteSessionResponse)(nil), "pb.RPCDeleteSessionResponse")
	proto.RegisterType((*RPCGetSessionIdsByDetailsRequest)(nil), "pb.RPCGetSessionIdsByDetailsRequest")
	proto.RegisterType((*RPCGetSessionIdsByPeerIdsRequest)(nil), "pb.RPCGetSessionIdsByPeerIdsRequest")
	proto.RegisterType((*RPCGetSessionIdsResponse)(nil), "pb.RPCGetSessionIdsResponse")
	proto.RegisterType((*RPCGetSessionRequest)(nil), "pb.RPCGetSessionRequest")
	proto.RegisterType((*RPCGetSessionResponse)(nil), "pb.RPCGetSessionResponse")
}

func init() { proto.RegisterFile("pb/matchmaking.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 510 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xd1, 0x8b, 0xd3, 0x4e,
	0x10, 0xc7, 0xd9, 0xb4, 0xd7, 0xb4, 0x73, 0xbf, 0x9f, 0x9c, 0x6b, 0xd5, 0x45, 0x44, 0xca, 0xc2,
	0x61, 0x7d, 0xa9, 0x50, 0x45, 0xcf, 0x2a, 0x22, 0x5e, 0xe5, 0xe8, 0x8b, 0x94, 0x3d, 0xc4, 0x37,
	0x21, 0xbd, 0x1d, 0xbc, 0xd0, 0x34, 0x59, 0xb3, 0x5b, 0x21, 0x7f, 0x96, 0xff, 0xa1, 0x64, 0xb3,
	0x49, 0x93, 0xf4, 0x2c, 0x7d, 0xdb, 0x9d, 0xd9, 0xcf, 0xcc, 0x77, 0xf2, 0x1d, 0x02, 0x43, 0xb5,
	0x7a, 0xb9, 0x09, 0xcc, 0xcd, 0xed, 0x26, 0x58, 0x87, 0xf1, 0xcf, 0x89, 0x4a, 0x13, 0x93, 0x50,
	0x4f, 0xad, 0xf8, 0x1f, 0x0f, 0xfc, 0x6b, 0xd4, 0x3a, 0x4c, 0x62, 0x7a, 0x06, 0x1d, 0xb5, 0x90,
	0x8c, 0x8c, 0xc8, 0x78, 0x20, 0xf2, 0x23, 0xa5, 0xd0, 0x95, 0xf1, 0x42, 0x32, 0x6f, 0x44, 0xc6,
	0x5d, 0x61, 0xcf, 0xf4, 0x29, 0x0c, 0x74, 0x01, 0x2c, 0x24, 0xeb, 0xd8, 0xc4, 0x2e, 0x90, 0x13,
	0x26, 0x53, 0xc8, 0xba, 0x05, 0x91, 0x9f, 0xf3, 0x58, 0x1c, 0x6c, 0x90, 0x9d, 0xd8, 0xc2, 0xf6,
	0x4c, 0x19, 0xf8, 0x81, 0x94, 0x29, 0x6a, 0xcd, 0x7a, 0x23, 0x32, 0xfe, 0x5f, 0x94, 0xd7, 0xfc,
	0xb5, 0x4a, 0x52, 0xc3, 0x7c, 0x1b, 0xb6, 0x67, 0x3a, 0x05, 0x5f, 0xa2, 0x09, 0xc2, 0x48, 0xb3,
	0xfe, 0xa8, 0x33, 0x3e, 0x9d, 0xb2, 0x89, 0x5a, 0x4d, 0x9c, 0xee, 0xc9, 0xbc, 0x48, 0x7d, 0x89,
	0x4d, 0x9a, 0x89, 0xf2, 0x61, 0xae, 0x53, 0x6d, 0x57, 0x51, 0xa8, 0x6f, 0x51, 0xb2, 0x41, 0xa1,
	0xb3, 0x0a, 0x3c, 0x99, 0xc1, 0x7f, 0x75, 0x2c, 0x9f, 0x7d, 0x8d, 0x59, 0x39, 0xfb, 0x1a, 0x33,
	0x3a, 0x84, 0x93, 0xdf, 0x41, 0xb4, 0x45, 0x3b, 0xfc, 0x40, 0x14, 0x97, 0x99, 0x77, 0x41, 0xf8,
	0x3b, 0xb8, 0x37, 0xff, 0xea, 0x9a, 0x5f, 0x9b, 0x24, 0x45, 0xfa, 0x1c, 0xfa, 0xee, 0x13, 0x68,
	0x46, 0xac, 0xc0, 0xd3, 0x9a, 0x40, 0x51, 0x25, 0xf9, 0x77, 0x38, 0xab, 0x50, 0x2d, 0xf0, 0x26,
	0x49, 0x65, 0x53, 0x28, 0x69, 0x09, 0xa5, 0xe7, 0xe0, 0xa7, 0xa8, 0xb7, 0x91, 0xd1, 0xcc, 0xdb,
	0xaf, 0x5c, 0xe6, 0xf8, 0x7b, 0xa0, 0xf5, 0xc2, 0x5a, 0x25, 0xb1, 0xc6, 0x3a, 0x4c, 0x0e, 0xc0,
	0x57, 0x70, 0xbf, 0x0e, 0xff, 0xda, 0xa2, 0x36, 0x95, 0x93, 0xa4, 0xe6, 0xe4, 0x33, 0x80, 0xca,
	0xea, 0x42, 0x4f, 0x57, 0xd4, 0x22, 0xfc, 0x0d, 0x0c, 0xab, 0x42, 0xa1, 0xdc, 0xe9, 0x68, 0x72,
	0x64, 0x8f, 0x7b, 0x01, 0x0f, 0x9a, 0xdc, 0x3f, 0x25, 0xf0, 0x4f, 0xf0, 0x58, 0x2c, 0x2f, 0xbf,
	0x29, 0x8d, 0xa9, 0x29, 0x07, 0x71, 0xcf, 0xcf, 0xc1, 0x77, 0x35, 0x2d, 0xd1, 0x9e, 0xd6, 0xe5,
	0xf8, 0x05, 0xb0, 0xfd, 0x0a, 0x4e, 0x68, 0x63, 0xb9, 0x49, 0x6b, 0xb9, 0xf9, 0x5b, 0xdb, 0x7b,
	0x8e, 0x11, 0x1a, 0x6c, 0xf5, 0x3e, 0x0c, 0x4e, 0x6d, 0xcb, 0x16, 0xe8, 0x5a, 0x3e, 0x82, 0x5e,
	0xe1, 0x83, 0xc5, 0xfa, 0xc2, 0xdd, 0xf8, 0x0f, 0x18, 0x89, 0xe5, 0xe5, 0x15, 0x96, 0x1a, 0x17,
	0x52, 0x7f, 0xce, 0xdc, 0xce, 0x1e, 0xf2, 0xc8, 0x6d, 0xb2, 0x77, 0xc7, 0x26, 0x77, 0x6a, 0x9b,
	0xcc, 0x3f, 0xdc, 0x55, 0x7f, 0x89, 0x98, 0x2e, 0x76, 0x06, 0x30, 0xf0, 0x55, 0x11, 0xb1, 0xa6,
	0x0d, 0x44, 0x79, 0xe5, 0x33, 0x3b, 0x51, 0x83, 0x3e, 0xda, 0xed, 0xd7, 0x30, 0x6c, 0xb0, 0xc7,
	0x7d, 0xc3, 0x8f, 0xf0, 0xb0, 0x45, 0xed, 0x96, 0xfc, 0x08, 0xdb, 0x57, 0x3d, 0xfb, 0xd3, 0x7b,
	0xf5, 0x37, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x66, 0x16, 0xf6, 0x0c, 0x05, 0x00, 0x00,
}
