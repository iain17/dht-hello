// Code generated by protoc-gen-go. DO NOT EDIT.
// source: serve/pb/protocol.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	serve/pb/protocol.proto

It has these top-level messages:
	HealthRequest
	HealthReply
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

type HealthRequest struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *HealthRequest) Reset()                    { *m = HealthRequest{} }
func (m *HealthRequest) String() string            { return proto.CompactTextString(m) }
func (*HealthRequest) ProtoMessage()               {}
func (*HealthRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HealthRequest) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

// The response message containing the greetings
type HealthReply struct {
	Ready   bool   `protobuf:"varint,1,opt,name=ready" json:"ready,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *HealthReply) Reset()                    { *m = HealthReply{} }
func (m *HealthReply) String() string            { return proto.CompactTextString(m) }
func (*HealthReply) ProtoMessage()               {}
func (*HealthReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HealthReply) GetReady() bool {
	if m != nil {
		return m.Ready
	}
	return false
}

func (m *HealthReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type RPCMessage struct {
	Version int64 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Id      int64 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	// Types that are valid to be assigned to Msg:
	//	*RPCMessage_HealthRequest
	//	*RPCMessage_HealthReply
	Msg isRPCMessage_Msg `protobuf_oneof:"msg"`
}

func (m *RPCMessage) Reset()                    { *m = RPCMessage{} }
func (m *RPCMessage) String() string            { return proto.CompactTextString(m) }
func (*RPCMessage) ProtoMessage()               {}
func (*RPCMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isRPCMessage_Msg interface {
	isRPCMessage_Msg()
}

type RPCMessage_HealthRequest struct {
	HealthRequest *HealthRequest `protobuf:"bytes,3,opt,name=healthRequest,oneof"`
}
type RPCMessage_HealthReply struct {
	HealthReply *HealthReply `protobuf:"bytes,4,opt,name=HealthReply,oneof"`
}

func (*RPCMessage_HealthRequest) isRPCMessage_Msg() {}
func (*RPCMessage_HealthReply) isRPCMessage_Msg()   {}

func (m *RPCMessage) GetMsg() isRPCMessage_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *RPCMessage) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *RPCMessage) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RPCMessage) GetHealthRequest() *HealthRequest {
	if x, ok := m.GetMsg().(*RPCMessage_HealthRequest); ok {
		return x.HealthRequest
	}
	return nil
}

func (m *RPCMessage) GetHealthReply() *HealthReply {
	if x, ok := m.GetMsg().(*RPCMessage_HealthReply); ok {
		return x.HealthReply
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*RPCMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _RPCMessage_OneofMarshaler, _RPCMessage_OneofUnmarshaler, _RPCMessage_OneofSizer, []interface{}{
		(*RPCMessage_HealthRequest)(nil),
		(*RPCMessage_HealthReply)(nil),
	}
}

func _RPCMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*RPCMessage)
	// msg
	switch x := m.Msg.(type) {
	case *RPCMessage_HealthRequest:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.HealthRequest); err != nil {
			return err
		}
	case *RPCMessage_HealthReply:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.HealthReply); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("RPCMessage.Msg has unexpected type %T", x)
	}
	return nil
}

func _RPCMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*RPCMessage)
	switch tag {
	case 3: // msg.healthRequest
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HealthRequest)
		err := b.DecodeMessage(msg)
		m.Msg = &RPCMessage_HealthRequest{msg}
		return true, err
	case 4: // msg.HealthReply
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HealthReply)
		err := b.DecodeMessage(msg)
		m.Msg = &RPCMessage_HealthReply{msg}
		return true, err
	default:
		return false, nil
	}
}

func _RPCMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*RPCMessage)
	// msg
	switch x := m.Msg.(type) {
	case *RPCMessage_HealthRequest:
		s := proto.Size(x.HealthRequest)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *RPCMessage_HealthReply:
		s := proto.Size(x.HealthReply)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*HealthRequest)(nil), "pb.HealthRequest")
	proto.RegisterType((*HealthReply)(nil), "pb.HealthReply")
	proto.RegisterType((*RPCMessage)(nil), "pb.RPCMessage")
}

func init() { proto.RegisterFile("serve/pb/protocol.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x4e, 0x2d, 0x2a,
	0x4b, 0xd5, 0x2f, 0x48, 0xd2, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x03, 0x33,
	0x84, 0x98, 0x0a, 0x92, 0x94, 0x14, 0xb9, 0x78, 0x3d, 0x52, 0x13, 0x73, 0x4a, 0x32, 0x82, 0x52,
	0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x04, 0xb8, 0x98, 0x73, 0x8b, 0xd3, 0x25, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x83, 0x40, 0x4c, 0x25, 0x5b, 0x2e, 0x6e, 0x98, 0x92, 0x82, 0x9c, 0x4a, 0x21, 0x11,
	0x2e, 0xd6, 0xa2, 0xd4, 0xc4, 0x94, 0x4a, 0xb0, 0x12, 0x8e, 0x20, 0x08, 0x47, 0x48, 0x82, 0x8b,
	0x3d, 0x37, 0xb5, 0xb8, 0x38, 0x31, 0x3d, 0x55, 0x82, 0x09, 0xac, 0x15, 0xc6, 0x55, 0x5a, 0xcb,
	0xc8, 0xc5, 0x15, 0x14, 0xe0, 0xec, 0x0b, 0xe1, 0x82, 0x14, 0x96, 0xa5, 0x16, 0x15, 0x67, 0xe6,
	0xe7, 0x81, 0x0d, 0x60, 0x0e, 0x82, 0x71, 0x85, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0xc0, 0xba, 0x99,
	0x83, 0x98, 0x32, 0x53, 0x84, 0x2c, 0xb9, 0x78, 0x33, 0x90, 0x9d, 0x26, 0xc1, 0xac, 0xc0, 0xa8,
	0xc1, 0x6d, 0x24, 0xa8, 0x57, 0x90, 0xa4, 0x87, 0xe2, 0x66, 0x0f, 0x86, 0x20, 0x54, 0x95, 0x42,
	0xc6, 0x28, 0x4e, 0x96, 0x60, 0x01, 0x6b, 0xe4, 0x47, 0xd6, 0x58, 0x90, 0x53, 0xe9, 0xc1, 0x10,
	0x84, 0xac, 0xca, 0x89, 0x15, 0xec, 0x73, 0x27, 0x26, 0x0f, 0xe6, 0x24, 0x36, 0x70, 0x00, 0x19,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x7b, 0x0c, 0xd7, 0xe1, 0x3b, 0x01, 0x00, 0x00,
}