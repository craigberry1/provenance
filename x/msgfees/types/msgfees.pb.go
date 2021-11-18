// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: provenance/msgfees/v1/msgfees.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/x/auth/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Params defines the set of params for the msgfees module.
type Params struct {
	// indicates if governance based controls of msgFees is allowed.
	EnableGovernance bool `protobuf:"varint,1,opt,name=enable_governance,json=enableGovernance,proto3" json:"enable_governance,omitempty"`
	// constant used to calculate fees when gas fees shares denom with msg fee
	MinGasPrice uint32 `protobuf:"varint,2,opt,name=min_gas_price,json=minGasPrice,proto3" json:"min_gas_price,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c6265859d114362, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetEnableGovernance() bool {
	if m != nil {
		return m.EnableGovernance
	}
	return false
}

func (m *Params) GetMinGasPrice() uint32 {
	if m != nil {
		return m.MinGasPrice
	}
	return 0
}

// MsgBasedFee is the core of what gets stored on the blockchain
// it consists of two parts
// 1. the msg type url, i.e. /cosmos.bank.v1beta1.MsgSend
// 2. minimum additional fees(can be of any denom)
type MsgBasedFee struct {
	MsgTypeUrl string `protobuf:"bytes,1,opt,name=msg_type_url,json=msgTypeUrl,proto3" json:"msg_type_url,omitempty"`
	// can pay in any Coin( basically a Denom and Amount, Amount can be zero)
	AdditionalFee types.Coin `protobuf:"bytes,2,opt,name=additional_fee,json=additionalFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"additional_fee" yaml:"additional_fee"`
}

func (m *MsgBasedFee) Reset()         { *m = MsgBasedFee{} }
func (m *MsgBasedFee) String() string { return proto.CompactTextString(m) }
func (*MsgBasedFee) ProtoMessage()    {}
func (*MsgBasedFee) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c6265859d114362, []int{1}
}
func (m *MsgBasedFee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBasedFee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBasedFee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBasedFee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBasedFee.Merge(m, src)
}
func (m *MsgBasedFee) XXX_Size() int {
	return m.Size()
}
func (m *MsgBasedFee) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBasedFee.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBasedFee proto.InternalMessageInfo

func (m *MsgBasedFee) GetMsgTypeUrl() string {
	if m != nil {
		return m.MsgTypeUrl
	}
	return ""
}

func (m *MsgBasedFee) GetAdditionalFee() types.Coin {
	if m != nil {
		return m.AdditionalFee
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*Params)(nil), "provenance.msgfees.v1.Params")
	proto.RegisterType((*MsgBasedFee)(nil), "provenance.msgfees.v1.MsgBasedFee")
}

func init() {
	proto.RegisterFile("provenance/msgfees/v1/msgfees.proto", fileDescriptor_0c6265859d114362)
}

var fileDescriptor_0c6265859d114362 = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x3f, 0x6f, 0xd4, 0x30,
	0x18, 0xc6, 0x63, 0x86, 0xaa, 0xf8, 0x38, 0x04, 0x11, 0x95, 0xae, 0x15, 0x4a, 0x4e, 0x61, 0x39,
	0x09, 0x35, 0xd6, 0xd1, 0xad, 0xe3, 0x21, 0xb5, 0x62, 0x40, 0x3a, 0x9d, 0x60, 0x61, 0x89, 0xde,
	0x24, 0xef, 0xb9, 0x16, 0xb1, 0x1d, 0xc5, 0x49, 0x44, 0x3e, 0x03, 0x0b, 0x23, 0x63, 0x67, 0xbe,
	0x03, 0x7b, 0xc7, 0x8e, 0x4c, 0x05, 0xdd, 0x2d, 0xcc, 0x7c, 0x02, 0x14, 0x27, 0xf7, 0x07, 0x3a,
	0xc5, 0xaf, 0x7f, 0xcf, 0x9b, 0xc7, 0x8f, 0xfd, 0xd2, 0x17, 0x79, 0xa1, 0x6b, 0x54, 0xa0, 0x12,
	0x64, 0xd2, 0xf0, 0x25, 0xa2, 0x61, 0xf5, 0x74, 0xb3, 0x0c, 0xf3, 0x42, 0x97, 0xda, 0x3d, 0xda,
	0x89, 0xc2, 0x0d, 0xa9, 0xa7, 0x27, 0xcf, 0xb8, 0xe6, 0xda, 0x2a, 0x58, 0xbb, 0xea, 0xc4, 0x27,
	0x5e, 0xa2, 0x8d, 0xd4, 0x86, 0x41, 0x55, 0x5e, 0xb1, 0x7a, 0x1a, 0x63, 0x09, 0x53, 0x5b, 0xf4,
	0xfc, 0xb8, 0xe3, 0x51, 0xd7, 0xd8, 0x15, 0x1b, 0xc4, 0xb5, 0xe6, 0x19, 0x32, 0x5b, 0xc5, 0xd5,
	0x92, 0x81, 0x6a, 0x7a, 0xf4, 0xbc, 0x47, 0x90, 0x0b, 0x06, 0x4a, 0xe9, 0x12, 0x4a, 0xa1, 0x95,
	0xf9, 0xcf, 0x33, 0x06, 0x83, 0x5b, 0xcf, 0x44, 0x0b, 0xd5, 0xf1, 0x20, 0xa1, 0x07, 0x73, 0x28,
	0x40, 0x1a, 0xf7, 0x25, 0x7d, 0x8a, 0x0a, 0xe2, 0x0c, 0x23, 0xae, 0x6b, 0x2c, 0x6c, 0xa6, 0x11,
	0x19, 0x93, 0xc9, 0xe1, 0xe2, 0x49, 0x07, 0x2e, 0xb7, 0xfb, 0x6e, 0x40, 0x87, 0x52, 0xa8, 0x88,
	0x43, 0x7b, 0x5a, 0x91, 0xe0, 0xe8, 0xc1, 0x98, 0x4c, 0x86, 0x8b, 0x81, 0x14, 0xea, 0x12, 0xcc,
	0xbc, 0xdd, 0x3a, 0x3f, 0xfc, 0x7a, 0xed, 0x3b, 0xbf, 0xaf, 0x7d, 0x27, 0xf8, 0x4e, 0xe8, 0xe0,
	0xad, 0xe1, 0x33, 0x30, 0x98, 0x5e, 0x20, 0xba, 0x63, 0xfa, 0x48, 0x1a, 0x1e, 0x95, 0x4d, 0x8e,
	0x51, 0x55, 0x64, 0xd6, 0xe5, 0xe1, 0x82, 0x4a, 0xc3, 0xdf, 0x35, 0x39, 0xbe, 0x2f, 0x32, 0xf7,
	0x33, 0xa1, 0x8f, 0x21, 0x4d, 0x45, 0x1b, 0x05, 0xb2, 0x68, 0x89, 0x9d, 0xc3, 0xe0, 0xd5, 0x71,
	0xd8, 0xdf, 0x4b, 0x1b, 0x28, 0xec, 0x03, 0x85, 0xaf, 0xb5, 0x50, 0xb3, 0x37, 0x37, 0x77, 0xbe,
	0xf3, 0xe7, 0xce, 0x3f, 0x6a, 0x40, 0x66, 0xe7, 0xc1, 0xbf, 0xed, 0xc1, 0xb7, 0x9f, 0xfe, 0x84,
	0x8b, 0xf2, 0xaa, 0x8a, 0xc3, 0x44, 0xcb, 0xfe, 0x76, 0xfb, 0xcf, 0xa9, 0x49, 0x3f, 0xb2, 0xf6,
	0x34, 0xc6, 0xfe, 0xc9, 0x2c, 0x86, 0xbb, 0xe6, 0x0b, 0xc4, 0x99, 0xb8, 0x59, 0x79, 0xe4, 0x76,
	0xe5, 0x91, 0x5f, 0x2b, 0x8f, 0x7c, 0x59, 0x7b, 0xce, 0xed, 0xda, 0x73, 0x7e, 0xac, 0x3d, 0x87,
	0x8e, 0x84, 0x7d, 0xdc, 0xfb, 0x23, 0x30, 0x27, 0x1f, 0xce, 0xf6, 0xec, 0x76, 0x9a, 0x53, 0xa1,
	0xf7, 0x2a, 0xf6, 0x69, 0x3b, 0x5b, 0xd6, 0x3f, 0x3e, 0xb0, 0xcf, 0x72, 0xf6, 0x37, 0x00, 0x00,
	0xff, 0xff, 0xa4, 0x2b, 0xce, 0x57, 0x7e, 0x02, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MinGasPrice != 0 {
		i = encodeVarintMsgfees(dAtA, i, uint64(m.MinGasPrice))
		i--
		dAtA[i] = 0x10
	}
	if m.EnableGovernance {
		i--
		if m.EnableGovernance {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgBasedFee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBasedFee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBasedFee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.AdditionalFee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsgfees(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.MsgTypeUrl) > 0 {
		i -= len(m.MsgTypeUrl)
		copy(dAtA[i:], m.MsgTypeUrl)
		i = encodeVarintMsgfees(dAtA, i, uint64(len(m.MsgTypeUrl)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMsgfees(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgfees(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EnableGovernance {
		n += 2
	}
	if m.MinGasPrice != 0 {
		n += 1 + sovMsgfees(uint64(m.MinGasPrice))
	}
	return n
}

func (m *MsgBasedFee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MsgTypeUrl)
	if l > 0 {
		n += 1 + l + sovMsgfees(uint64(l))
	}
	l = m.AdditionalFee.Size()
	n += 1 + l + sovMsgfees(uint64(l))
	return n
}

func sovMsgfees(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgfees(x uint64) (n int) {
	return sovMsgfees(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgfees
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableGovernance", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgfees
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EnableGovernance = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinGasPrice", wireType)
			}
			m.MinGasPrice = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgfees
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinGasPrice |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMsgfees(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgfees
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
func (m *MsgBasedFee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgfees
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
			return fmt.Errorf("proto: MsgBasedFee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBasedFee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgTypeUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgfees
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
				return ErrInvalidLengthMsgfees
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgfees
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgTypeUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdditionalFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgfees
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
				return ErrInvalidLengthMsgfees
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgfees
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AdditionalFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgfees(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgfees
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
func skipMsgfees(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgfees
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
					return 0, ErrIntOverflowMsgfees
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
					return 0, ErrIntOverflowMsgfees
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
				return 0, ErrInvalidLengthMsgfees
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgfees
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgfees
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgfees        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgfees          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgfees = fmt.Errorf("proto: unexpected end of group")
)
