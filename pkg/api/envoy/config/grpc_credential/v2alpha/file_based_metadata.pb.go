// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/grpc_credential/v2alpha/file_based_metadata.proto

package envoy_config_grpc_credential_v2alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/datawire/ambassador/pkg/api/envoy/api/v2/core"
	proto "github.com/gogo/protobuf/proto"
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

type FileBasedMetadataConfig struct {
	// Location or inline data of secret to use for authentication of the Google gRPC connection
	// this secret will be attached to a header of the gRPC connection
	SecretData *core.DataSource `protobuf:"bytes,1,opt,name=secret_data,json=secretData,proto3" json:"secret_data,omitempty"`
	// Metadata header key to use for sending the secret data
	// if no header key is set, "authorization" header will be used
	HeaderKey string `protobuf:"bytes,2,opt,name=header_key,json=headerKey,proto3" json:"header_key,omitempty"`
	// Prefix to prepend to the secret in the metadata header
	// if no prefix is set, the default is to use no prefix
	HeaderPrefix         string   `protobuf:"bytes,3,opt,name=header_prefix,json=headerPrefix,proto3" json:"header_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileBasedMetadataConfig) Reset()         { *m = FileBasedMetadataConfig{} }
func (m *FileBasedMetadataConfig) String() string { return proto.CompactTextString(m) }
func (*FileBasedMetadataConfig) ProtoMessage()    {}
func (*FileBasedMetadataConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f2b21de9d357383, []int{0}
}
func (m *FileBasedMetadataConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FileBasedMetadataConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FileBasedMetadataConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FileBasedMetadataConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileBasedMetadataConfig.Merge(m, src)
}
func (m *FileBasedMetadataConfig) XXX_Size() int {
	return m.Size()
}
func (m *FileBasedMetadataConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FileBasedMetadataConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FileBasedMetadataConfig proto.InternalMessageInfo

func (m *FileBasedMetadataConfig) GetSecretData() *core.DataSource {
	if m != nil {
		return m.SecretData
	}
	return nil
}

func (m *FileBasedMetadataConfig) GetHeaderKey() string {
	if m != nil {
		return m.HeaderKey
	}
	return ""
}

func (m *FileBasedMetadataConfig) GetHeaderPrefix() string {
	if m != nil {
		return m.HeaderPrefix
	}
	return ""
}

func init() {
	proto.RegisterType((*FileBasedMetadataConfig)(nil), "envoy.config.grpc_credential.v2alpha.FileBasedMetadataConfig")
}

func init() {
	proto.RegisterFile("envoy/config/grpc_credential/v2alpha/file_based_metadata.proto", fileDescriptor_0f2b21de9d357383)
}

var fileDescriptor_0f2b21de9d357383 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0xe5, 0xfe, 0x52, 0xa5, 0xba, 0x3f, 0x4b, 0x06, 0xa8, 0x2a, 0x5a, 0x55, 0xc0, 0xd0,
	0xc9, 0x96, 0xc2, 0xce, 0x10, 0x50, 0x17, 0x84, 0x14, 0x95, 0x8d, 0x25, 0xba, 0x75, 0x6e, 0x5a,
	0x8b, 0x60, 0x5b, 0x8e, 0x1b, 0x35, 0xcf, 0xc0, 0x6b, 0xc0, 0x33, 0xc0, 0x23, 0x30, 0xf2, 0x08,
	0x28, 0x4f, 0x82, 0x1c, 0x67, 0x82, 0x85, 0xf5, 0xdc, 0xfb, 0xdd, 0x73, 0xee, 0xa1, 0x57, 0xa8,
	0x6a, 0xdd, 0x70, 0xa1, 0x55, 0x21, 0xb7, 0x7c, 0x6b, 0x8d, 0xc8, 0x84, 0xc5, 0x1c, 0x95, 0x93,
	0x50, 0xf2, 0x3a, 0x86, 0xd2, 0xec, 0x80, 0x17, 0xb2, 0xc4, 0x6c, 0x03, 0x15, 0xe6, 0xd9, 0x13,
	0x3a, 0xc8, 0xc1, 0x01, 0x33, 0x56, 0x3b, 0x1d, 0x5d, 0x74, 0x3c, 0x0b, 0x3c, 0xfb, 0xc1, 0xb3,
	0x9e, 0x9f, 0x9e, 0x06, 0x17, 0x30, 0x92, 0xd7, 0x31, 0x17, 0xda, 0x22, 0xf7, 0xd7, 0xc2, 0x8d,
	0xe9, 0x62, 0x9f, 0x1b, 0xe0, 0xa0, 0x94, 0x76, 0xe0, 0xa4, 0x56, 0x15, 0xaf, 0x50, 0x55, 0xd2,
	0xc9, 0xba, 0xdf, 0x38, 0x7b, 0x25, 0xf4, 0x64, 0x25, 0x4b, 0x4c, 0x7c, 0x84, 0xbb, 0x3e, 0xc1,
	0x75, 0xe7, 0x19, 0xad, 0xe8, 0xb8, 0x42, 0x61, 0xd1, 0x65, 0x5e, 0x9c, 0x90, 0x05, 0x59, 0x8e,
	0xe3, 0x19, 0x0b, 0xb9, 0xc0, 0x48, 0x56, 0xc7, 0xcc, 0x3b, 0xb2, 0x1b, 0x70, 0x70, 0xaf, 0xf7,
	0x56, 0x60, 0x32, 0x7c, 0x7f, 0x7b, 0x7e, 0x19, 0x90, 0x35, 0x0d, 0xa4, 0x9f, 0x44, 0x33, 0x4a,
	0x77, 0x08, 0x39, 0xda, 0xec, 0x11, 0x9b, 0xc9, 0x60, 0x41, 0x96, 0xa3, 0xf5, 0x28, 0x28, 0xb7,
	0xd8, 0x44, 0xe7, 0xf4, 0xa8, 0x1f, 0x1b, 0x8b, 0x85, 0x3c, 0x4c, 0xfe, 0x75, 0x1b, 0xff, 0x83,
	0x98, 0x76, 0x5a, 0xf2, 0xf0, 0xd1, 0xce, 0xc9, 0x67, 0x3b, 0x27, 0x5f, 0xed, 0x9c, 0xd0, 0x58,
	0xea, 0x10, 0xc3, 0x58, 0x7d, 0x68, 0xd8, 0x5f, 0x9a, 0x4a, 0x8e, 0x7f, 0xbd, 0x99, 0xfa, 0x06,
	0x52, 0xb2, 0x19, 0x76, 0x55, 0x5c, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x0c, 0x2f, 0x8c, 0xd7,
	0xb2, 0x01, 0x00, 0x00,
}

func (m *FileBasedMetadataConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FileBasedMetadataConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FileBasedMetadataConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.HeaderPrefix) > 0 {
		i -= len(m.HeaderPrefix)
		copy(dAtA[i:], m.HeaderPrefix)
		i = encodeVarintFileBasedMetadata(dAtA, i, uint64(len(m.HeaderPrefix)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.HeaderKey) > 0 {
		i -= len(m.HeaderKey)
		copy(dAtA[i:], m.HeaderKey)
		i = encodeVarintFileBasedMetadata(dAtA, i, uint64(len(m.HeaderKey)))
		i--
		dAtA[i] = 0x12
	}
	if m.SecretData != nil {
		{
			size, err := m.SecretData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFileBasedMetadata(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFileBasedMetadata(dAtA []byte, offset int, v uint64) int {
	offset -= sovFileBasedMetadata(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FileBasedMetadataConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SecretData != nil {
		l = m.SecretData.Size()
		n += 1 + l + sovFileBasedMetadata(uint64(l))
	}
	l = len(m.HeaderKey)
	if l > 0 {
		n += 1 + l + sovFileBasedMetadata(uint64(l))
	}
	l = len(m.HeaderPrefix)
	if l > 0 {
		n += 1 + l + sovFileBasedMetadata(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovFileBasedMetadata(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFileBasedMetadata(x uint64) (n int) {
	return sovFileBasedMetadata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FileBasedMetadataConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFileBasedMetadata
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
			return fmt.Errorf("proto: FileBasedMetadataConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FileBasedMetadataConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SecretData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileBasedMetadata
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
				return ErrInvalidLengthFileBasedMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFileBasedMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SecretData == nil {
				m.SecretData = &core.DataSource{}
			}
			if err := m.SecretData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeaderKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileBasedMetadata
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
				return ErrInvalidLengthFileBasedMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFileBasedMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HeaderKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeaderPrefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileBasedMetadata
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
				return ErrInvalidLengthFileBasedMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFileBasedMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HeaderPrefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFileBasedMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFileBasedMetadata
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFileBasedMetadata
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipFileBasedMetadata(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFileBasedMetadata
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
					return 0, ErrIntOverflowFileBasedMetadata
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
					return 0, ErrIntOverflowFileBasedMetadata
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
				return 0, ErrInvalidLengthFileBasedMetadata
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFileBasedMetadata
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFileBasedMetadata
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFileBasedMetadata        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFileBasedMetadata          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFileBasedMetadata = fmt.Errorf("proto: unexpected end of group")
)
