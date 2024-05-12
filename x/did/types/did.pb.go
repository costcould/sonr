// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: did/v1/did.proto

package types

import (
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

// DIDDocument defines a DID document
type DIDDocument struct {
	Id                   string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	VerificationMethods  []*VerificationMethod `protobuf:"bytes,2,rep,name=verification_methods,json=verificationMethods,proto3" json:"verification_methods,omitempty"`
	Authentication       []string              `protobuf:"bytes,4,rep,name=authentication,proto3" json:"authentication,omitempty"`
	AssertionMethod      []string              `protobuf:"bytes,5,rep,name=assertion_method,json=assertionMethod,proto3" json:"assertion_method,omitempty"`
	CapabilityDelegation []string              `protobuf:"bytes,7,rep,name=capability_delegation,json=capabilityDelegation,proto3" json:"capability_delegation,omitempty"`
	CapabilityInvocation []string              `protobuf:"bytes,8,rep,name=capability_invocation,json=capabilityInvocation,proto3" json:"capability_invocation,omitempty"`
}

func (m *DIDDocument) Reset()         { *m = DIDDocument{} }
func (m *DIDDocument) String() string { return proto.CompactTextString(m) }
func (*DIDDocument) ProtoMessage()    {}
func (*DIDDocument) Descriptor() ([]byte, []int) {
	return fileDescriptor_e63f2c747195152a, []int{0}
}
func (m *DIDDocument) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DIDDocument) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DIDDocument.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DIDDocument) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DIDDocument.Merge(m, src)
}
func (m *DIDDocument) XXX_Size() int {
	return m.Size()
}
func (m *DIDDocument) XXX_DiscardUnknown() {
	xxx_messageInfo_DIDDocument.DiscardUnknown(m)
}

var xxx_messageInfo_DIDDocument proto.InternalMessageInfo

func (m *DIDDocument) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DIDDocument) GetVerificationMethods() []*VerificationMethod {
	if m != nil {
		return m.VerificationMethods
	}
	return nil
}

func (m *DIDDocument) GetAuthentication() []string {
	if m != nil {
		return m.Authentication
	}
	return nil
}

func (m *DIDDocument) GetAssertionMethod() []string {
	if m != nil {
		return m.AssertionMethod
	}
	return nil
}

func (m *DIDDocument) GetCapabilityDelegation() []string {
	if m != nil {
		return m.CapabilityDelegation
	}
	return nil
}

func (m *DIDDocument) GetCapabilityInvocation() []string {
	if m != nil {
		return m.CapabilityInvocation
	}
	return nil
}

// VerificationMethod defines a verification method
type VerificationMethod struct {
	Id           string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Controller   string            `protobuf:"bytes,2,opt,name=controller,proto3" json:"controller,omitempty"`
	Type         string            `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	PublicKey    string            `protobuf:"bytes,4,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	PublicKeyJwk map[string]string `protobuf:"bytes,5,rep,name=public_key_jwk,json=publicKeyJwk,proto3" json:"public_key_jwk,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *VerificationMethod) Reset()         { *m = VerificationMethod{} }
func (m *VerificationMethod) String() string { return proto.CompactTextString(m) }
func (*VerificationMethod) ProtoMessage()    {}
func (*VerificationMethod) Descriptor() ([]byte, []int) {
	return fileDescriptor_e63f2c747195152a, []int{1}
}
func (m *VerificationMethod) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VerificationMethod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VerificationMethod.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VerificationMethod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerificationMethod.Merge(m, src)
}
func (m *VerificationMethod) XXX_Size() int {
	return m.Size()
}
func (m *VerificationMethod) XXX_DiscardUnknown() {
	xxx_messageInfo_VerificationMethod.DiscardUnknown(m)
}

var xxx_messageInfo_VerificationMethod proto.InternalMessageInfo

func (m *VerificationMethod) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *VerificationMethod) GetController() string {
	if m != nil {
		return m.Controller
	}
	return ""
}

func (m *VerificationMethod) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *VerificationMethod) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

func (m *VerificationMethod) GetPublicKeyJwk() map[string]string {
	if m != nil {
		return m.PublicKeyJwk
	}
	return nil
}

func init() {
	proto.RegisterType((*DIDDocument)(nil), "did.v1.DIDDocument")
	proto.RegisterType((*VerificationMethod)(nil), "did.v1.VerificationMethod")
	proto.RegisterMapType((map[string]string)(nil), "did.v1.VerificationMethod.PublicKeyJwkEntry")
}

func init() { proto.RegisterFile("did/v1/did.proto", fileDescriptor_e63f2c747195152a) }

var fileDescriptor_e63f2c747195152a = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x41, 0x8b, 0xda, 0x40,
	0x14, 0xc7, 0x4d, 0xa2, 0xb6, 0x3e, 0x8b, 0xb5, 0x53, 0x0b, 0x41, 0x68, 0x10, 0x0f, 0xc5, 0x42,
	0x9b, 0xa0, 0x5e, 0x4a, 0x29, 0x14, 0x4a, 0x7a, 0xb0, 0x45, 0x28, 0x39, 0xf4, 0xb0, 0x97, 0x90,
	0x64, 0x66, 0x75, 0xd6, 0x98, 0x09, 0xc9, 0x24, 0x6e, 0xbe, 0xc3, 0x1e, 0xf6, 0x6b, 0xec, 0x37,
	0xd9, 0xa3, 0xc7, 0x3d, 0x2e, 0xfa, 0x45, 0x96, 0x4c, 0x5c, 0x0d, 0x86, 0xbd, 0xbd, 0xfc, 0x7f,
	0xff, 0xf7, 0x32, 0xff, 0xc7, 0x83, 0x2e, 0xa6, 0xd8, 0x48, 0xc7, 0x06, 0xa6, 0x58, 0x0f, 0x23,
	0xc6, 0x19, 0x6a, 0xe6, 0x65, 0x3a, 0x1e, 0xde, 0xc9, 0xd0, 0x36, 0x67, 0xa6, 0xc9, 0xbc, 0x64,
	0x4d, 0x02, 0x8e, 0x3a, 0x20, 0x53, 0xac, 0x4a, 0x03, 0x69, 0xd4, 0xb2, 0x64, 0x8a, 0xd1, 0x1c,
	0x7a, 0x29, 0x89, 0xe8, 0x25, 0xf5, 0x1c, 0x4e, 0x59, 0x60, 0xaf, 0x09, 0x5f, 0x32, 0x1c, 0xab,
	0xf2, 0x40, 0x19, 0xb5, 0x27, 0x7d, 0xbd, 0x18, 0xa3, 0xff, 0x2f, 0x79, 0xe6, 0xc2, 0x62, 0xbd,
	0x4f, 0x2b, 0x5a, 0x8c, 0x3e, 0x41, 0xc7, 0x49, 0xf8, 0x92, 0x04, 0xfc, 0x00, 0xd4, 0xfa, 0x40,
	0x19, 0xb5, 0xac, 0x33, 0x15, 0x7d, 0x86, 0xae, 0x13, 0xc7, 0x24, 0x2a, 0xfd, 0x53, 0x6d, 0x08,
	0xe7, 0xdb, 0xa3, 0x5e, 0xcc, 0x44, 0x53, 0xf8, 0xe0, 0x39, 0xa1, 0xe3, 0x52, 0x9f, 0xf2, 0xcc,
	0xc6, 0xc4, 0x27, 0x8b, 0x62, 0xf2, 0x2b, 0xe1, 0xef, 0x9d, 0xa0, 0x79, 0x64, 0x67, 0x4d, 0x34,
	0x48, 0xd9, 0xe1, 0x39, 0xaf, 0xcf, 0x9b, 0x66, 0x47, 0x36, 0xbc, 0x91, 0x01, 0x55, 0x83, 0x56,
	0x56, 0xa6, 0x01, 0x78, 0x2c, 0xe0, 0x11, 0xf3, 0x7d, 0x12, 0xa9, 0xb2, 0xd0, 0x4b, 0x0a, 0x42,
	0x50, 0xe7, 0x59, 0x48, 0x54, 0x45, 0x10, 0x51, 0xa3, 0x8f, 0x00, 0x61, 0xe2, 0xfa, 0xd4, 0xb3,
	0x57, 0x24, 0x53, 0xeb, 0x82, 0xb4, 0x0a, 0xe5, 0x2f, 0xc9, 0x90, 0x05, 0x9d, 0x13, 0xb6, 0xaf,
	0x36, 0x2b, 0xb1, 0x8c, 0xf6, 0xe4, 0xcb, 0xcb, 0xfb, 0xd7, 0xff, 0x3d, 0x77, 0xff, 0xd9, 0xac,
	0x7e, 0x07, 0x3c, 0xca, 0xac, 0x37, 0x61, 0x49, 0xea, 0xff, 0x84, 0x77, 0x15, 0x0b, 0xea, 0x82,
	0x92, 0x3f, 0xa0, 0x08, 0x93, 0x97, 0xa8, 0x07, 0x8d, 0xd4, 0xf1, 0x13, 0x72, 0x08, 0x52, 0x7c,
	0x7c, 0x97, 0xbf, 0x49, 0xbf, 0x7e, 0xdc, 0xef, 0x34, 0x69, 0xbb, 0xd3, 0xa4, 0xc7, 0x9d, 0x26,
	0xdd, 0xee, 0xb5, 0xda, 0x76, 0xaf, 0xd5, 0x1e, 0xf6, 0x5a, 0xed, 0x62, 0xb8, 0xa0, 0x7c, 0x99,
	0xb8, 0xba, 0xc7, 0xd6, 0x06, 0xa6, 0x5f, 0xb1, 0xc3, 0x0c, 0x8f, 0x45, 0xc4, 0xb8, 0xce, 0x0f,
	0xd0, 0xc8, 0x03, 0xc7, 0x6e, 0x53, 0xdc, 0xe1, 0xf4, 0x29, 0x00, 0x00, 0xff, 0xff, 0xc8, 0xad,
	0x78, 0x31, 0x9b, 0x02, 0x00, 0x00,
}

func (m *DIDDocument) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DIDDocument) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DIDDocument) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CapabilityInvocation) > 0 {
		for iNdEx := len(m.CapabilityInvocation) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.CapabilityInvocation[iNdEx])
			copy(dAtA[i:], m.CapabilityInvocation[iNdEx])
			i = encodeVarintDid(dAtA, i, uint64(len(m.CapabilityInvocation[iNdEx])))
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.CapabilityDelegation) > 0 {
		for iNdEx := len(m.CapabilityDelegation) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.CapabilityDelegation[iNdEx])
			copy(dAtA[i:], m.CapabilityDelegation[iNdEx])
			i = encodeVarintDid(dAtA, i, uint64(len(m.CapabilityDelegation[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.AssertionMethod) > 0 {
		for iNdEx := len(m.AssertionMethod) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AssertionMethod[iNdEx])
			copy(dAtA[i:], m.AssertionMethod[iNdEx])
			i = encodeVarintDid(dAtA, i, uint64(len(m.AssertionMethod[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Authentication) > 0 {
		for iNdEx := len(m.Authentication) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Authentication[iNdEx])
			copy(dAtA[i:], m.Authentication[iNdEx])
			i = encodeVarintDid(dAtA, i, uint64(len(m.Authentication[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.VerificationMethods) > 0 {
		for iNdEx := len(m.VerificationMethods) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VerificationMethods[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDid(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintDid(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *VerificationMethod) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VerificationMethod) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VerificationMethod) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PublicKeyJwk) > 0 {
		for k := range m.PublicKeyJwk {
			v := m.PublicKeyJwk[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintDid(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintDid(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintDid(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.PublicKey) > 0 {
		i -= len(m.PublicKey)
		copy(dAtA[i:], m.PublicKey)
		i = encodeVarintDid(dAtA, i, uint64(len(m.PublicKey)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintDid(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Controller) > 0 {
		i -= len(m.Controller)
		copy(dAtA[i:], m.Controller)
		i = encodeVarintDid(dAtA, i, uint64(len(m.Controller)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintDid(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDid(dAtA []byte, offset int, v uint64) int {
	offset -= sovDid(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DIDDocument) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	if len(m.VerificationMethods) > 0 {
		for _, e := range m.VerificationMethods {
			l = e.Size()
			n += 1 + l + sovDid(uint64(l))
		}
	}
	if len(m.Authentication) > 0 {
		for _, s := range m.Authentication {
			l = len(s)
			n += 1 + l + sovDid(uint64(l))
		}
	}
	if len(m.AssertionMethod) > 0 {
		for _, s := range m.AssertionMethod {
			l = len(s)
			n += 1 + l + sovDid(uint64(l))
		}
	}
	if len(m.CapabilityDelegation) > 0 {
		for _, s := range m.CapabilityDelegation {
			l = len(s)
			n += 1 + l + sovDid(uint64(l))
		}
	}
	if len(m.CapabilityInvocation) > 0 {
		for _, s := range m.CapabilityInvocation {
			l = len(s)
			n += 1 + l + sovDid(uint64(l))
		}
	}
	return n
}

func (m *VerificationMethod) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	l = len(m.Controller)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	l = len(m.PublicKey)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	if len(m.PublicKeyJwk) > 0 {
		for k, v := range m.PublicKeyJwk {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovDid(uint64(len(k))) + 1 + len(v) + sovDid(uint64(len(v)))
			n += mapEntrySize + 1 + sovDid(uint64(mapEntrySize))
		}
	}
	return n
}

func sovDid(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDid(x uint64) (n int) {
	return sovDid(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DIDDocument) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDid
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
			return fmt.Errorf("proto: DIDDocument: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DIDDocument: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerificationMethods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VerificationMethods = append(m.VerificationMethods, &VerificationMethod{})
			if err := m.VerificationMethods[len(m.VerificationMethods)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authentication", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authentication = append(m.Authentication, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssertionMethod", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssertionMethod = append(m.AssertionMethod, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CapabilityDelegation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CapabilityDelegation = append(m.CapabilityDelegation, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CapabilityInvocation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CapabilityInvocation = append(m.CapabilityInvocation, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDid
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
func (m *VerificationMethod) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDid
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
			return fmt.Errorf("proto: VerificationMethod: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VerificationMethod: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Controller", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Controller = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKeyJwk", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PublicKeyJwk == nil {
				m.PublicKeyJwk = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowDid
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowDid
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthDid
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthDid
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowDid
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthDid
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthDid
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipDid(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthDid
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.PublicKeyJwk[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDid
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
func skipDid(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDid
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
					return 0, ErrIntOverflowDid
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
					return 0, ErrIntOverflowDid
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
				return 0, ErrInvalidLengthDid
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDid
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDid
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDid        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDid          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDid = fmt.Errorf("proto: unexpected end of group")
)