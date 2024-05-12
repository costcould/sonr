// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package didv1

import (
	v1beta1 "cosmossdk.io/api/cosmos/auth/v1beta1"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_EthAccount              protoreflect.MessageDescriptor
	fd_EthAccount_base_account protoreflect.FieldDescriptor
	fd_EthAccount_code_hash    protoreflect.FieldDescriptor
	fd_EthAccount_controller   protoreflect.FieldDescriptor
)

func init() {
	file_did_v1_account_proto_init()
	md_EthAccount = File_did_v1_account_proto.Messages().ByName("EthAccount")
	fd_EthAccount_base_account = md_EthAccount.Fields().ByName("base_account")
	fd_EthAccount_code_hash = md_EthAccount.Fields().ByName("code_hash")
	fd_EthAccount_controller = md_EthAccount.Fields().ByName("controller")
}

var _ protoreflect.Message = (*fastReflection_EthAccount)(nil)

type fastReflection_EthAccount EthAccount

func (x *EthAccount) ProtoReflect() protoreflect.Message {
	return (*fastReflection_EthAccount)(x)
}

func (x *EthAccount) slowProtoReflect() protoreflect.Message {
	mi := &file_did_v1_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_EthAccount_messageType fastReflection_EthAccount_messageType
var _ protoreflect.MessageType = fastReflection_EthAccount_messageType{}

type fastReflection_EthAccount_messageType struct{}

func (x fastReflection_EthAccount_messageType) Zero() protoreflect.Message {
	return (*fastReflection_EthAccount)(nil)
}
func (x fastReflection_EthAccount_messageType) New() protoreflect.Message {
	return new(fastReflection_EthAccount)
}
func (x fastReflection_EthAccount_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_EthAccount
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_EthAccount) Descriptor() protoreflect.MessageDescriptor {
	return md_EthAccount
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_EthAccount) Type() protoreflect.MessageType {
	return _fastReflection_EthAccount_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_EthAccount) New() protoreflect.Message {
	return new(fastReflection_EthAccount)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_EthAccount) Interface() protoreflect.ProtoMessage {
	return (*EthAccount)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_EthAccount) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.BaseAccount != nil {
		value := protoreflect.ValueOfMessage(x.BaseAccount.ProtoReflect())
		if !f(fd_EthAccount_base_account, value) {
			return
		}
	}
	if x.CodeHash != "" {
		value := protoreflect.ValueOfString(x.CodeHash)
		if !f(fd_EthAccount_code_hash, value) {
			return
		}
	}
	if x.Controller != "" {
		value := protoreflect.ValueOfString(x.Controller)
		if !f(fd_EthAccount_controller, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_EthAccount) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "did.v1.EthAccount.base_account":
		return x.BaseAccount != nil
	case "did.v1.EthAccount.code_hash":
		return x.CodeHash != ""
	case "did.v1.EthAccount.controller":
		return x.Controller != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: did.v1.EthAccount"))
		}
		panic(fmt.Errorf("message did.v1.EthAccount does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EthAccount) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "did.v1.EthAccount.base_account":
		x.BaseAccount = nil
	case "did.v1.EthAccount.code_hash":
		x.CodeHash = ""
	case "did.v1.EthAccount.controller":
		x.Controller = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: did.v1.EthAccount"))
		}
		panic(fmt.Errorf("message did.v1.EthAccount does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_EthAccount) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "did.v1.EthAccount.base_account":
		value := x.BaseAccount
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "did.v1.EthAccount.code_hash":
		value := x.CodeHash
		return protoreflect.ValueOfString(value)
	case "did.v1.EthAccount.controller":
		value := x.Controller
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: did.v1.EthAccount"))
		}
		panic(fmt.Errorf("message did.v1.EthAccount does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EthAccount) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "did.v1.EthAccount.base_account":
		x.BaseAccount = value.Message().Interface().(*v1beta1.BaseAccount)
	case "did.v1.EthAccount.code_hash":
		x.CodeHash = value.Interface().(string)
	case "did.v1.EthAccount.controller":
		x.Controller = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: did.v1.EthAccount"))
		}
		panic(fmt.Errorf("message did.v1.EthAccount does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EthAccount) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "did.v1.EthAccount.base_account":
		if x.BaseAccount == nil {
			x.BaseAccount = new(v1beta1.BaseAccount)
		}
		return protoreflect.ValueOfMessage(x.BaseAccount.ProtoReflect())
	case "did.v1.EthAccount.code_hash":
		panic(fmt.Errorf("field code_hash of message did.v1.EthAccount is not mutable"))
	case "did.v1.EthAccount.controller":
		panic(fmt.Errorf("field controller of message did.v1.EthAccount is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: did.v1.EthAccount"))
		}
		panic(fmt.Errorf("message did.v1.EthAccount does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_EthAccount) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "did.v1.EthAccount.base_account":
		m := new(v1beta1.BaseAccount)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "did.v1.EthAccount.code_hash":
		return protoreflect.ValueOfString("")
	case "did.v1.EthAccount.controller":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: did.v1.EthAccount"))
		}
		panic(fmt.Errorf("message did.v1.EthAccount does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_EthAccount) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in did.v1.EthAccount", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_EthAccount) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EthAccount) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_EthAccount) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_EthAccount) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*EthAccount)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.BaseAccount != nil {
			l = options.Size(x.BaseAccount)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.CodeHash)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Controller)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*EthAccount)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Controller) > 0 {
			i -= len(x.Controller)
			copy(dAtA[i:], x.Controller)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Controller)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.CodeHash) > 0 {
			i -= len(x.CodeHash)
			copy(dAtA[i:], x.CodeHash)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.CodeHash)))
			i--
			dAtA[i] = 0x12
		}
		if x.BaseAccount != nil {
			encoded, err := options.Marshal(x.BaseAccount)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*EthAccount)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: EthAccount: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: EthAccount: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BaseAccount", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.BaseAccount == nil {
					x.BaseAccount = &v1beta1.BaseAccount{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.BaseAccount); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field CodeHash", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.CodeHash = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Controller", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Controller = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: did/v1/account.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// EthAccount implements the authtypes.AccountI interface and embeds an
// authtypes.BaseAccount type. It is compatible with the auth AccountKeeper.
type EthAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// base_account is an authtypes.BaseAccount
	BaseAccount *v1beta1.BaseAccount `protobuf:"bytes,1,opt,name=base_account,json=baseAccount,proto3" json:"base_account,omitempty"`
	// code_hash is the hash calculated from the code contents
	CodeHash string `protobuf:"bytes,2,opt,name=code_hash,json=codeHash,proto3" json:"code_hash,omitempty"`
	// controller is the address of the controller
	Controller string `protobuf:"bytes,3,opt,name=controller,proto3" json:"controller,omitempty"`
}

func (x *EthAccount) Reset() {
	*x = EthAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_did_v1_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthAccount) ProtoMessage() {}

// Deprecated: Use EthAccount.ProtoReflect.Descriptor instead.
func (*EthAccount) Descriptor() ([]byte, []int) {
	return file_did_v1_account_proto_rawDescGZIP(), []int{0}
}

func (x *EthAccount) GetBaseAccount() *v1beta1.BaseAccount {
	if x != nil {
		return x.BaseAccount
	}
	return nil
}

func (x *EthAccount) GetCodeHash() string {
	if x != nil {
		return x.CodeHash
	}
	return ""
}

func (x *EthAccount) GetController() string {
	if x != nil {
		return x.Controller
	}
	return ""
}

var File_did_v1_account_proto protoreflect.FileDescriptor

var file_did_v1_account_proto_rawDesc = []byte{
	0x0a, 0x14, 0x64, 0x69, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x64, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1e,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb0, 0x02, 0x0a, 0x0a, 0x45, 0x74, 0x68, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x60,
	0x0a, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x1b, 0xd0, 0xde, 0x1f, 0x01, 0xf2, 0xde, 0x1f, 0x13,
	0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x31, 0x0a, 0x09, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x14, 0xf2, 0xde, 0x1f, 0x10, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x63,
	0x6f, 0x64, 0x65, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x22, 0x52, 0x08, 0x63, 0x6f, 0x64, 0x65, 0x48,
	0x61, 0x73, 0x68, 0x12, 0x35, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0xf2, 0xde, 0x1f, 0x11, 0x79, 0x61, 0x6d,
	0x6c, 0x3a, 0x22, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x22, 0x52, 0x0a,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x3a, 0x56, 0x88, 0xa0, 0x1f, 0x00,
	0x98, 0xa0, 0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x00, 0xca, 0xb4, 0x2d, 0x46, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x78, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x42, 0x7c, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x64, 0x2e, 0x76, 0x31,
	0x42, 0x0c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x2d,
	0x64, 0x61, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x64,
	0x2f, 0x76, 0x31, 0x3b, 0x64, 0x69, 0x64, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x44, 0x58, 0x58, 0xaa,
	0x02, 0x06, 0x44, 0x69, 0x64, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x06, 0x44, 0x69, 0x64, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x12, 0x44, 0x69, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x44, 0x69, 0x64, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_did_v1_account_proto_rawDescOnce sync.Once
	file_did_v1_account_proto_rawDescData = file_did_v1_account_proto_rawDesc
)

func file_did_v1_account_proto_rawDescGZIP() []byte {
	file_did_v1_account_proto_rawDescOnce.Do(func() {
		file_did_v1_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_did_v1_account_proto_rawDescData)
	})
	return file_did_v1_account_proto_rawDescData
}

var file_did_v1_account_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_did_v1_account_proto_goTypes = []interface{}{
	(*EthAccount)(nil),          // 0: did.v1.EthAccount
	(*v1beta1.BaseAccount)(nil), // 1: cosmos.auth.v1beta1.BaseAccount
}
var file_did_v1_account_proto_depIdxs = []int32{
	1, // 0: did.v1.EthAccount.base_account:type_name -> cosmos.auth.v1beta1.BaseAccount
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_did_v1_account_proto_init() }
func file_did_v1_account_proto_init() {
	if File_did_v1_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_did_v1_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EthAccount); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_did_v1_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_did_v1_account_proto_goTypes,
		DependencyIndexes: file_did_v1_account_proto_depIdxs,
		MessageInfos:      file_did_v1_account_proto_msgTypes,
	}.Build()
	File_did_v1_account_proto = out.File
	file_did_v1_account_proto_rawDesc = nil
	file_did_v1_account_proto_goTypes = nil
	file_did_v1_account_proto_depIdxs = nil
}
