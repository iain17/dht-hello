// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/publisher.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// DN
type PublisherDefinition struct {
	Created int64             `protobuf:"varint,1,opt,name=created" json:"created,omitempty"`
	Files   map[string][]byte `protobuf:"bytes,2,rep,name=files" json:"files,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Links   map[string]string `protobuf:"bytes,3,rep,name=links" json:"links,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Details map[string]string `protobuf:"bytes,4,rep,name=details" json:"details,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Status  bool              `protobuf:"varint,5,opt,name=status" json:"status,omitempty"`
}

func (m *PublisherDefinition) Reset()                    { *m = PublisherDefinition{} }
func (m *PublisherDefinition) String() string            { return proto.CompactTextString(m) }
func (*PublisherDefinition) ProtoMessage()               {}
func (*PublisherDefinition) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *PublisherDefinition) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *PublisherDefinition) GetFiles() map[string][]byte {
	if m != nil {
		return m.Files
	}
	return nil
}

func (m *PublisherDefinition) GetLinks() map[string]string {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *PublisherDefinition) GetDetails() map[string]string {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *PublisherDefinition) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

type PublisherUpdate struct {
	Signature  []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	Definition []byte `protobuf:"bytes,3,opt,name=definition,proto3" json:"definition,omitempty"`
}

func (m *PublisherUpdate) Reset()                    { *m = PublisherUpdate{} }
func (m *PublisherUpdate) String() string            { return proto.CompactTextString(m) }
func (*PublisherUpdate) ProtoMessage()               {}
func (*PublisherUpdate) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *PublisherUpdate) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *PublisherUpdate) GetDefinition() []byte {
	if m != nil {
		return m.Definition
	}
	return nil
}

type DNPublisherUpdateRequest struct {
}

func (m *DNPublisherUpdateRequest) Reset()                    { *m = DNPublisherUpdateRequest{} }
func (m *DNPublisherUpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*DNPublisherUpdateRequest) ProtoMessage()               {}
func (*DNPublisherUpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

type DNPublisherUpdateResponse struct {
	Update *PublisherUpdate `protobuf:"bytes,1,opt,name=update" json:"update,omitempty"`
}

func (m *DNPublisherUpdateResponse) Reset()                    { *m = DNPublisherUpdateResponse{} }
func (m *DNPublisherUpdateResponse) String() string            { return proto.CompactTextString(m) }
func (*DNPublisherUpdateResponse) ProtoMessage()               {}
func (*DNPublisherUpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *DNPublisherUpdateResponse) GetUpdate() *PublisherUpdate {
	if m != nil {
		return m.Update
	}
	return nil
}

// RPC
type RPCPublishPublisherUpdateRequest struct {
	Definition *PublisherDefinition `protobuf:"bytes,1,opt,name=definition" json:"definition,omitempty"`
}

func (m *RPCPublishPublisherUpdateRequest) Reset()         { *m = RPCPublishPublisherUpdateRequest{} }
func (m *RPCPublishPublisherUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*RPCPublishPublisherUpdateRequest) ProtoMessage()    {}
func (*RPCPublishPublisherUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor4, []int{4}
}

func (m *RPCPublishPublisherUpdateRequest) GetDefinition() *PublisherDefinition {
	if m != nil {
		return m.Definition
	}
	return nil
}

func init() {
	proto.RegisterType((*PublisherDefinition)(nil), "pb.PublisherDefinition")
	proto.RegisterType((*PublisherUpdate)(nil), "pb.PublisherUpdate")
	proto.RegisterType((*DNPublisherUpdateRequest)(nil), "pb.DNPublisherUpdateRequest")
	proto.RegisterType((*DNPublisherUpdateResponse)(nil), "pb.DNPublisherUpdateResponse")
	proto.RegisterType((*RPCPublishPublisherUpdateRequest)(nil), "pb.RPCPublishPublisherUpdateRequest")
}

func init() { proto.RegisterFile("pb/publisher.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x49, 0x62, 0x5b, 0x3b, 0x2d, 0x28, 0x5b, 0xd1, 0xb5, 0x88, 0x84, 0xe0, 0x21, 0x20,
	0x44, 0xa8, 0x07, 0x4b, 0x0f, 0x5e, 0xac, 0xe2, 0x41, 0xb4, 0x2c, 0x78, 0xf2, 0x94, 0x98, 0xa9,
	0x2e, 0x0d, 0xc9, 0x9a, 0xdd, 0x08, 0xfd, 0x4d, 0xfe, 0x49, 0xc9, 0x26, 0x69, 0x53, 0x49, 0x15,
	0x6f, 0x99, 0x37, 0xef, 0x9b, 0xbc, 0xc9, 0x04, 0x88, 0x08, 0x2e, 0x44, 0x16, 0x44, 0x5c, 0xbe,
	0x63, 0xea, 0x89, 0x34, 0x51, 0x09, 0x31, 0x45, 0xe0, 0x7c, 0x59, 0x30, 0x98, 0x55, 0xfa, 0x14,
	0xe7, 0x3c, 0xe6, 0x8a, 0x27, 0x31, 0xa1, 0xd0, 0x79, 0x4d, 0xd1, 0x57, 0x18, 0x52, 0xc3, 0x36,
	0x5c, 0x8b, 0x55, 0x25, 0x19, 0x43, 0x6b, 0xce, 0x23, 0x94, 0xd4, 0xb4, 0x2d, 0xb7, 0x37, 0x72,
	0x3c, 0x11, 0x78, 0x0d, 0x13, 0xbc, 0xbb, 0xdc, 0x74, 0x1b, 0xab, 0x74, 0xc9, 0x0a, 0x20, 0x27,
	0x23, 0x1e, 0x2f, 0x24, 0xb5, 0x7e, 0x27, 0x1f, 0x72, 0x53, 0x49, 0x6a, 0x80, 0x5c, 0x43, 0x27,
	0x44, 0xe5, 0xf3, 0x48, 0xd2, 0x1d, 0xcd, 0x9e, 0x6d, 0x63, 0xa7, 0x85, 0xad, 0xa0, 0x2b, 0x88,
	0x1c, 0x42, 0x5b, 0x2a, 0x5f, 0x65, 0x92, 0xb6, 0x6c, 0xc3, 0xdd, 0x65, 0x65, 0x35, 0x1c, 0x03,
	0xac, 0x63, 0x92, 0x7d, 0xb0, 0x16, 0xb8, 0xd4, 0xfb, 0x76, 0x59, 0xfe, 0x48, 0x0e, 0xa0, 0xf5,
	0xe9, 0x47, 0x19, 0x52, 0xd3, 0x36, 0xdc, 0x3e, 0x2b, 0x8a, 0x89, 0x39, 0x36, 0x72, 0x72, 0x1d,
	0xf3, 0x2f, 0xb2, 0x5b, 0x27, 0x27, 0xd0, 0xaf, 0x87, 0xfc, 0x0f, 0xeb, 0x3c, 0xc1, 0xde, 0x6a,
	0xe9, 0x67, 0x11, 0xfa, 0x0a, 0xc9, 0x09, 0x74, 0x25, 0x7f, 0x8b, 0x7d, 0x95, 0xa5, 0x55, 0xcc,
	0xb5, 0x40, 0x4e, 0x01, 0xc2, 0xd5, 0xc7, 0xa1, 0x96, 0x6e, 0xd7, 0x14, 0x67, 0x08, 0x74, 0xfa,
	0xf8, 0x63, 0x24, 0xc3, 0x8f, 0x0c, 0xa5, 0x72, 0xee, 0xe1, 0xb8, 0xa1, 0x27, 0x45, 0x12, 0x4b,
	0x24, 0xe7, 0xd0, 0xce, 0xb4, 0xa2, 0x83, 0xf7, 0x46, 0x83, 0x8d, 0x83, 0x94, 0xe6, 0xd2, 0xe2,
	0xbc, 0x80, 0xcd, 0x66, 0x37, 0x65, 0xb7, 0xf9, 0x6d, 0xe4, 0x6a, 0x23, 0x69, 0x31, 0xf4, 0x68,
	0xcb, 0x95, 0xeb, 0x2b, 0x04, 0x6d, 0xfd, 0x33, 0x5f, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0xfe,
	0xad, 0xa9, 0xc3, 0xe2, 0x02, 0x00, 0x00,
}
