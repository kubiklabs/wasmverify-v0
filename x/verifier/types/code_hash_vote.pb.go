// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: verifier/verifier/code_hash_vote.proto

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

type CodeHashVote struct {
	ApplicationId uint64 `protobuf:"varint,1,opt,name=applicationId,proto3" json:"applicationId,omitempty"`
	CodeHash      string `protobuf:"bytes,2,opt,name=codeHash,proto3" json:"codeHash,omitempty"`
	Voter         string `protobuf:"bytes,3,opt,name=voter,proto3" json:"voter,omitempty"`
}

func (m *CodeHashVote) Reset()         { *m = CodeHashVote{} }
func (m *CodeHashVote) String() string { return proto.CompactTextString(m) }
func (*CodeHashVote) ProtoMessage()    {}
func (*CodeHashVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f5042bd9ed61f9c, []int{0}
}
func (m *CodeHashVote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CodeHashVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CodeHashVote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CodeHashVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeHashVote.Merge(m, src)
}
func (m *CodeHashVote) XXX_Size() int {
	return m.Size()
}
func (m *CodeHashVote) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeHashVote.DiscardUnknown(m)
}

var xxx_messageInfo_CodeHashVote proto.InternalMessageInfo

func (m *CodeHashVote) GetApplicationId() uint64 {
	if m != nil {
		return m.ApplicationId
	}
	return 0
}

func (m *CodeHashVote) GetCodeHash() string {
	if m != nil {
		return m.CodeHash
	}
	return ""
}

func (m *CodeHashVote) GetVoter() string {
	if m != nil {
		return m.Voter
	}
	return ""
}

func init() {
	proto.RegisterType((*CodeHashVote)(nil), "verifier.verifier.CodeHashVote")
}

func init() {
	proto.RegisterFile("verifier/verifier/code_hash_vote.proto", fileDescriptor_5f5042bd9ed61f9c)
}

var fileDescriptor_5f5042bd9ed61f9c = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2b, 0x4b, 0x2d, 0xca,
	0x4c, 0xcb, 0x4c, 0x2d, 0xd2, 0x87, 0x33, 0x92, 0xf3, 0x53, 0x52, 0xe3, 0x33, 0x12, 0x8b, 0x33,
	0xe2, 0xcb, 0xf2, 0x4b, 0x52, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x04, 0x61, 0xd2, 0x7a,
	0x30, 0x86, 0x52, 0x1a, 0x17, 0x8f, 0x73, 0x7e, 0x4a, 0xaa, 0x47, 0x62, 0x71, 0x46, 0x58, 0x7e,
	0x49, 0xaa, 0x90, 0x0a, 0x17, 0x6f, 0x62, 0x41, 0x41, 0x4e, 0x66, 0x72, 0x62, 0x49, 0x66, 0x7e,
	0x9e, 0x67, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x4b, 0x10, 0xaa, 0xa0, 0x90, 0x14, 0x17, 0x47,
	0x32, 0x54, 0x97, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x9c, 0x2f, 0x24, 0xc2, 0xc5, 0x0a,
	0xb2, 0xb2, 0x48, 0x82, 0x19, 0x2c, 0x01, 0xe1, 0x38, 0x19, 0x9f, 0x78, 0x24, 0xc7, 0x78, 0xe1,
	0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70,
	0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x24, 0xdc, 0xcd, 0x15, 0x08, 0xe7, 0x97, 0x54, 0x16, 0xa4, 0x16,
	0x27, 0xb1, 0x81, 0x9d, 0x6d, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x54, 0xda, 0x79, 0x43, 0xe0,
	0x00, 0x00, 0x00,
}

func (m *CodeHashVote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CodeHashVote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CodeHashVote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Voter) > 0 {
		i -= len(m.Voter)
		copy(dAtA[i:], m.Voter)
		i = encodeVarintCodeHashVote(dAtA, i, uint64(len(m.Voter)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.CodeHash) > 0 {
		i -= len(m.CodeHash)
		copy(dAtA[i:], m.CodeHash)
		i = encodeVarintCodeHashVote(dAtA, i, uint64(len(m.CodeHash)))
		i--
		dAtA[i] = 0x12
	}
	if m.ApplicationId != 0 {
		i = encodeVarintCodeHashVote(dAtA, i, uint64(m.ApplicationId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCodeHashVote(dAtA []byte, offset int, v uint64) int {
	offset -= sovCodeHashVote(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CodeHashVote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ApplicationId != 0 {
		n += 1 + sovCodeHashVote(uint64(m.ApplicationId))
	}
	l = len(m.CodeHash)
	if l > 0 {
		n += 1 + l + sovCodeHashVote(uint64(l))
	}
	l = len(m.Voter)
	if l > 0 {
		n += 1 + l + sovCodeHashVote(uint64(l))
	}
	return n
}

func sovCodeHashVote(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCodeHashVote(x uint64) (n int) {
	return sovCodeHashVote(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CodeHashVote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodeHashVote
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
			return fmt.Errorf("proto: CodeHashVote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CodeHashVote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApplicationId", wireType)
			}
			m.ApplicationId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodeHashVote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ApplicationId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodeHashVote
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
				return ErrInvalidLengthCodeHashVote
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCodeHashVote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CodeHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Voter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodeHashVote
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
				return ErrInvalidLengthCodeHashVote
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCodeHashVote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Voter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodeHashVote(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCodeHashVote
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
func skipCodeHashVote(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCodeHashVote
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
					return 0, ErrIntOverflowCodeHashVote
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
					return 0, ErrIntOverflowCodeHashVote
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
				return 0, ErrInvalidLengthCodeHashVote
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCodeHashVote
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCodeHashVote
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCodeHashVote        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCodeHashVote          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCodeHashVote = fmt.Errorf("proto: unexpected end of group")
)
