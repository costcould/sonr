// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: did/v1/state.proto

package types

import (
	_ "cosmossdk.io/orm"
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Authenticator is the representation of a Credential Assertion.
type Authenticator struct {
	// The unique identifier of the assertion.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The public key associated with the assertion.
	PublicKey *PublicKey `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// The type of key associated with the assertion.
	KeyType string `protobuf:"bytes,3,opt,name=key_type,json=keyType,proto3" json:"key_type,omitempty"`
	// The transport methods associated with the assertion.
	Transport []string `protobuf:"bytes,4,rep,name=transport,proto3" json:"transport,omitempty"`
	// The authenticator associated with the assertion.
	Authenticator string `protobuf:"bytes,5,opt,name=authenticator,proto3" json:"authenticator,omitempty"`
	// The origin of the issuer associated with the assertion.
	Origin string `protobuf:"bytes,6,opt,name=origin,proto3" json:"origin,omitempty"`
	// The index associated with the assertion.
	Index uint64 `protobuf:"varint,8,opt,name=index,proto3" json:"index,omitempty"`
	// The controller of the assertion.
	Controller string `protobuf:"bytes,9,opt,name=controller,proto3" json:"controller,omitempty"`
}

func (m *Authenticator) Reset()         { *m = Authenticator{} }
func (m *Authenticator) String() string { return proto.CompactTextString(m) }
func (*Authenticator) ProtoMessage()    {}
func (*Authenticator) Descriptor() ([]byte, []int) {
	return fileDescriptor_f44bb702879c34b4, []int{0}
}
func (m *Authenticator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Authenticator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Authenticator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Authenticator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Authenticator.Merge(m, src)
}
func (m *Authenticator) XXX_Size() int {
	return m.Size()
}
func (m *Authenticator) XXX_DiscardUnknown() {
	xxx_messageInfo_Authenticator.DiscardUnknown(m)
}

var xxx_messageInfo_Authenticator proto.InternalMessageInfo

func (m *Authenticator) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Authenticator) GetPublicKey() *PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *Authenticator) GetKeyType() string {
	if m != nil {
		return m.KeyType
	}
	return ""
}

func (m *Authenticator) GetTransport() []string {
	if m != nil {
		return m.Transport
	}
	return nil
}

func (m *Authenticator) GetAuthenticator() string {
	if m != nil {
		return m.Authenticator
	}
	return ""
}

func (m *Authenticator) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Authenticator) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Authenticator) GetController() string {
	if m != nil {
		return m.Controller
	}
	return ""
}

// Keyshare is the generated MPC share references.
type Keyshare struct {
	// The unique identifier of the keyshare.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The public key associated with the keyshare.
	PublicKey *PublicKey `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// The type of key associated with the keyshare.
	KeyType string `protobuf:"bytes,3,opt,name=key_type,json=keyType,proto3" json:"key_type,omitempty"`
	// The validator address associated with the keyshare.
	ValAddress string `protobuf:"bytes,4,opt,name=val_address,json=valAddress,proto3" json:"val_address,omitempty"`
	// The index associated with the keyshare.
	Index uint64 `protobuf:"varint,5,opt,name=index,proto3" json:"index,omitempty"`
	// The controller of the keyshare.
	Controller string `protobuf:"bytes,6,opt,name=controller,proto3" json:"controller,omitempty"`
}

func (m *Keyshare) Reset()         { *m = Keyshare{} }
func (m *Keyshare) String() string { return proto.CompactTextString(m) }
func (*Keyshare) ProtoMessage()    {}
func (*Keyshare) Descriptor() ([]byte, []int) {
	return fileDescriptor_f44bb702879c34b4, []int{1}
}
func (m *Keyshare) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Keyshare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Keyshare.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Keyshare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Keyshare.Merge(m, src)
}
func (m *Keyshare) XXX_Size() int {
	return m.Size()
}
func (m *Keyshare) XXX_DiscardUnknown() {
	xxx_messageInfo_Keyshare.DiscardUnknown(m)
}

var xxx_messageInfo_Keyshare proto.InternalMessageInfo

func (m *Keyshare) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Keyshare) GetPublicKey() *PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *Keyshare) GetKeyType() string {
	if m != nil {
		return m.KeyType
	}
	return ""
}

func (m *Keyshare) GetValAddress() string {
	if m != nil {
		return m.ValAddress
	}
	return ""
}

func (m *Keyshare) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Keyshare) GetController() string {
	if m != nil {
		return m.Controller
	}
	return ""
}

// Proof is the zero-knowledge proof for properties of the keyshare.
type Proof struct {
	// The unique identifier of the proof.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The public key associated with the proof.
	PublicKey *PublicKey `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// The type of key associated with the proof.
	KeyType string `protobuf:"bytes,3,opt,name=key_type,json=keyType,proto3" json:"key_type,omitempty"`
	// The witness associated with the proof.
	Witness string `protobuf:"bytes,4,opt,name=witness,proto3" json:"witness,omitempty"`
	// The controller of the proof.
	Controller string `protobuf:"bytes,5,opt,name=controller,proto3" json:"controller,omitempty"`
}

func (m *Proof) Reset()         { *m = Proof{} }
func (m *Proof) String() string { return proto.CompactTextString(m) }
func (*Proof) ProtoMessage()    {}
func (*Proof) Descriptor() ([]byte, []int) {
	return fileDescriptor_f44bb702879c34b4, []int{2}
}
func (m *Proof) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Proof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Proof.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Proof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proof.Merge(m, src)
}
func (m *Proof) XXX_Size() int {
	return m.Size()
}
func (m *Proof) XXX_DiscardUnknown() {
	xxx_messageInfo_Proof.DiscardUnknown(m)
}

var xxx_messageInfo_Proof proto.InternalMessageInfo

func (m *Proof) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Proof) GetPublicKey() *PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *Proof) GetKeyType() string {
	if m != nil {
		return m.KeyType
	}
	return ""
}

func (m *Proof) GetWitness() string {
	if m != nil {
		return m.Witness
	}
	return ""
}

func (m *Proof) GetController() string {
	if m != nil {
		return m.Controller
	}
	return ""
}

func init() {
	proto.RegisterType((*Authenticator)(nil), "did.v1.Authenticator")
	proto.RegisterType((*Keyshare)(nil), "did.v1.Keyshare")
	proto.RegisterType((*Proof)(nil), "did.v1.Proof")
}

func init() { proto.RegisterFile("did/v1/state.proto", fileDescriptor_f44bb702879c34b4) }

var fileDescriptor_f44bb702879c34b4 = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0x41, 0x6e, 0xd3, 0x40,
	0x14, 0x86, 0x33, 0x8e, 0xed, 0xc6, 0xaf, 0x6a, 0x25, 0x46, 0x15, 0x0c, 0x15, 0x32, 0x51, 0xc4,
	0xc2, 0x1b, 0x6c, 0x0a, 0x3b, 0xc4, 0xa6, 0x6c, 0xbb, 0xa9, 0x2c, 0x56, 0x6c, 0xa2, 0x89, 0x67,
	0x48, 0x46, 0xb1, 0xe7, 0x59, 0x33, 0x13, 0x53, 0x5f, 0xa2, 0xe2, 0x04, 0x1c, 0x81, 0x73, 0xb0,
	0x42, 0x95, 0xd8, 0xb0, 0x44, 0xc9, 0x0d, 0x38, 0x01, 0x8a, 0xdd, 0x2a, 0x05, 0xa9, 0xea, 0xaa,
	0xbb, 0x79, 0xff, 0x7b, 0xff, 0xcc, 0xff, 0x8d, 0x1e, 0x50, 0xa1, 0x44, 0xd6, 0x9c, 0x64, 0xd6,
	0x71, 0x27, 0xd3, 0xda, 0xa0, 0x43, 0x1a, 0x0a, 0x25, 0xd2, 0xe6, 0xe4, 0xf8, 0x49, 0x81, 0xb6,
	0x42, 0x9b, 0xa1, 0xa9, 0xb6, 0x23, 0x68, 0xaa, 0x7e, 0xe0, 0xf8, 0xe8, 0xda, 0x34, 0x97, 0x5a,
	0x5a, 0x65, 0x7b, 0x75, 0x72, 0xe9, 0xc1, 0xc1, 0xe9, 0xca, 0x2d, 0xa4, 0x76, 0xaa, 0xe0, 0x0e,
	0x0d, 0x3d, 0x04, 0x4f, 0x09, 0x46, 0xc6, 0x24, 0x89, 0x72, 0x4f, 0x09, 0xfa, 0x0a, 0xa0, 0x5e,
	0xcd, 0x4a, 0x55, 0x4c, 0x97, 0xb2, 0x65, 0xde, 0x98, 0x24, 0xfb, 0xaf, 0x1f, 0xa5, 0xfd, 0x6b,
	0xe9, 0x79, 0xd7, 0x39, 0x93, 0x6d, 0x1e, 0xd5, 0x37, 0x47, 0xfa, 0x14, 0x46, 0x4b, 0xd9, 0x4e,
	0x5d, 0x5b, 0x4b, 0x36, 0xec, 0xee, 0xd9, 0x5b, 0xca, 0xf6, 0x43, 0x5b, 0x4b, 0xfa, 0x0c, 0x22,
	0x67, 0xb8, 0xb6, 0x35, 0x1a, 0xc7, 0xfc, 0xf1, 0x30, 0x89, 0xf2, 0x9d, 0x40, 0x5f, 0xc0, 0x01,
	0xbf, 0x9d, 0x85, 0x05, 0x9d, 0xfb, 0x5f, 0x91, 0x3e, 0x86, 0x10, 0x8d, 0x9a, 0x2b, 0xcd, 0xc2,
	0xae, 0x7d, 0x5d, 0xd1, 0x23, 0x08, 0x94, 0x16, 0xf2, 0x82, 0x8d, 0xc6, 0x24, 0xf1, 0xf3, 0xbe,
	0xa0, 0x31, 0x40, 0x81, 0xda, 0x19, 0x2c, 0x4b, 0x69, 0x58, 0xd4, 0x39, 0x6e, 0x29, 0x6f, 0x0f,
	0xff, 0x7c, 0xfd, 0x79, 0x39, 0x1c, 0x81, 0xdf, 0x63, 0x4f, 0x7e, 0x10, 0x18, 0x9d, 0xc9, 0xd6,
	0x2e, 0xb8, 0x91, 0x0f, 0xfb, 0x17, 0xcf, 0x61, 0xbf, 0xe1, 0xe5, 0x94, 0x0b, 0x61, 0xa4, 0xb5,
	0xcc, 0xef, 0xa3, 0x35, 0xbc, 0x3c, 0xed, 0x95, 0x1d, 0x50, 0x70, 0x37, 0x50, 0x78, 0x0f, 0x90,
	0x37, 0xf9, 0x46, 0x20, 0x38, 0x37, 0x88, 0x9f, 0x1e, 0x96, 0x86, 0xc1, 0xde, 0x67, 0xe5, 0xf4,
	0x8e, 0xe4, 0xa6, 0xfc, 0x2f, 0x70, 0x70, 0x4f, 0x60, 0xff, 0xfd, 0xbb, 0xef, 0xeb, 0x98, 0x5c,
	0xad, 0x63, 0xf2, 0x7b, 0x1d, 0x93, 0x2f, 0x9b, 0x78, 0x70, 0xb5, 0x89, 0x07, 0xbf, 0x36, 0xf1,
	0xe0, 0xe3, 0x64, 0xae, 0xdc, 0x62, 0x35, 0x4b, 0x0b, 0xac, 0x32, 0xa1, 0x5e, 0x0a, 0x8e, 0x99,
	0x45, 0x6d, 0xb2, 0x8b, 0x6c, 0xbb, 0xdb, 0xdb, 0x54, 0x76, 0x16, 0x76, 0x7b, 0xfd, 0xe6, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x50, 0xd8, 0x9c, 0xad, 0x24, 0x03, 0x00, 0x00,
}

func (m *Authenticator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Authenticator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Authenticator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Controller) > 0 {
		i -= len(m.Controller)
		copy(dAtA[i:], m.Controller)
		i = encodeVarintState(dAtA, i, uint64(len(m.Controller)))
		i--
		dAtA[i] = 0x4a
	}
	if m.Index != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x40
	}
	if len(m.Origin) > 0 {
		i -= len(m.Origin)
		copy(dAtA[i:], m.Origin)
		i = encodeVarintState(dAtA, i, uint64(len(m.Origin)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Authenticator) > 0 {
		i -= len(m.Authenticator)
		copy(dAtA[i:], m.Authenticator)
		i = encodeVarintState(dAtA, i, uint64(len(m.Authenticator)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Transport) > 0 {
		for iNdEx := len(m.Transport) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Transport[iNdEx])
			copy(dAtA[i:], m.Transport[iNdEx])
			i = encodeVarintState(dAtA, i, uint64(len(m.Transport[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.KeyType) > 0 {
		i -= len(m.KeyType)
		copy(dAtA[i:], m.KeyType)
		i = encodeVarintState(dAtA, i, uint64(len(m.KeyType)))
		i--
		dAtA[i] = 0x1a
	}
	if m.PublicKey != nil {
		{
			size, err := m.PublicKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintState(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintState(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Keyshare) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Keyshare) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Keyshare) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Controller) > 0 {
		i -= len(m.Controller)
		copy(dAtA[i:], m.Controller)
		i = encodeVarintState(dAtA, i, uint64(len(m.Controller)))
		i--
		dAtA[i] = 0x32
	}
	if m.Index != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x28
	}
	if len(m.ValAddress) > 0 {
		i -= len(m.ValAddress)
		copy(dAtA[i:], m.ValAddress)
		i = encodeVarintState(dAtA, i, uint64(len(m.ValAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.KeyType) > 0 {
		i -= len(m.KeyType)
		copy(dAtA[i:], m.KeyType)
		i = encodeVarintState(dAtA, i, uint64(len(m.KeyType)))
		i--
		dAtA[i] = 0x1a
	}
	if m.PublicKey != nil {
		{
			size, err := m.PublicKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintState(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintState(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Proof) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Proof) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Proof) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Controller) > 0 {
		i -= len(m.Controller)
		copy(dAtA[i:], m.Controller)
		i = encodeVarintState(dAtA, i, uint64(len(m.Controller)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Witness) > 0 {
		i -= len(m.Witness)
		copy(dAtA[i:], m.Witness)
		i = encodeVarintState(dAtA, i, uint64(len(m.Witness)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.KeyType) > 0 {
		i -= len(m.KeyType)
		copy(dAtA[i:], m.KeyType)
		i = encodeVarintState(dAtA, i, uint64(len(m.KeyType)))
		i--
		dAtA[i] = 0x1a
	}
	if m.PublicKey != nil {
		{
			size, err := m.PublicKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintState(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintState(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintState(dAtA []byte, offset int, v uint64) int {
	offset -= sovState(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Authenticator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	if m.PublicKey != nil {
		l = m.PublicKey.Size()
		n += 1 + l + sovState(uint64(l))
	}
	l = len(m.KeyType)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	if len(m.Transport) > 0 {
		for _, s := range m.Transport {
			l = len(s)
			n += 1 + l + sovState(uint64(l))
		}
	}
	l = len(m.Authenticator)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	l = len(m.Origin)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	if m.Index != 0 {
		n += 1 + sovState(uint64(m.Index))
	}
	l = len(m.Controller)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	return n
}

func (m *Keyshare) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	if m.PublicKey != nil {
		l = m.PublicKey.Size()
		n += 1 + l + sovState(uint64(l))
	}
	l = len(m.KeyType)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	l = len(m.ValAddress)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	if m.Index != 0 {
		n += 1 + sovState(uint64(m.Index))
	}
	l = len(m.Controller)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	return n
}

func (m *Proof) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	if m.PublicKey != nil {
		l = m.PublicKey.Size()
		n += 1 + l + sovState(uint64(l))
	}
	l = len(m.KeyType)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	l = len(m.Witness)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	l = len(m.Controller)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	return n
}

func sovState(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozState(x uint64) (n int) {
	return sovState(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Authenticator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Authenticator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Authenticator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PublicKey == nil {
				m.PublicKey = &PublicKey{}
			}
			if err := m.PublicKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transport", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Transport = append(m.Transport, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authenticator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authenticator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Origin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Origin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Controller", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Controller = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthState
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Keyshare) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Keyshare: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Keyshare: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PublicKey == nil {
				m.PublicKey = &PublicKey{}
			}
			if err := m.PublicKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Controller", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Controller = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthState
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Proof) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Proof: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Proof: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PublicKey == nil {
				m.PublicKey = &PublicKey{}
			}
			if err := m.PublicKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Witness", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Witness = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Controller", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Controller = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthState
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipState(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowState
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowState
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowState
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthState
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupState
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthState
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthState        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowState          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupState = fmt.Errorf("proto: unexpected end of group")
)
