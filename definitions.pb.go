// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: definitions.proto

package dialog

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	_ "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

type UUIDValue struct {
	Msb int64 `protobuf:"varint,1,opt,name=msb,proto3" json:"msb,omitempty"`
	Lsb int64 `protobuf:"varint,2,opt,name=lsb,proto3" json:"lsb,omitempty"`
}

func (m *UUIDValue) Reset()      { *m = UUIDValue{} }
func (*UUIDValue) ProtoMessage() {}
func (*UUIDValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_bef01fe2be18d2f0, []int{0}
}
func (m *UUIDValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UUIDValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UUIDValue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UUIDValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UUIDValue.Merge(m, src)
}
func (m *UUIDValue) XXX_Size() int {
	return m.Size()
}
func (m *UUIDValue) XXX_DiscardUnknown() {
	xxx_messageInfo_UUIDValue.DiscardUnknown(m)
}

var xxx_messageInfo_UUIDValue proto.InternalMessageInfo

func (m *UUIDValue) GetMsb() int64 {
	if m != nil {
		return m.Msb
	}
	return 0
}

func (m *UUIDValue) GetLsb() int64 {
	if m != nil {
		return m.Lsb
	}
	return 0
}

type DialogOptions struct {
	Log string `protobuf:"bytes,1,opt,name=log,proto3" json:"log,omitempty"`
}

func (m *DialogOptions) Reset()      { *m = DialogOptions{} }
func (*DialogOptions) ProtoMessage() {}
func (*DialogOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_bef01fe2be18d2f0, []int{1}
}
func (m *DialogOptions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DialogOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DialogOptions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DialogOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DialogOptions.Merge(m, src)
}
func (m *DialogOptions) XXX_Size() int {
	return m.Size()
}
func (m *DialogOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_DialogOptions.DiscardUnknown(m)
}

var xxx_messageInfo_DialogOptions proto.InternalMessageInfo

func (m *DialogOptions) GetLog() string {
	if m != nil {
		return m.Log
	}
	return ""
}

/// server timestamp when state was created or modified
type DataClock struct {
	Clock int64 `protobuf:"varint,1,opt,name=clock,proto3" json:"clock,omitempty"`
}

func (m *DataClock) Reset()      { *m = DataClock{} }
func (*DataClock) ProtoMessage() {}
func (*DataClock) Descriptor() ([]byte, []int) {
	return fileDescriptor_bef01fe2be18d2f0, []int{2}
}
func (m *DataClock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DataClock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DataClock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DataClock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataClock.Merge(m, src)
}
func (m *DataClock) XXX_Size() int {
	return m.Size()
}
func (m *DataClock) XXX_DiscardUnknown() {
	xxx_messageInfo_DataClock.DiscardUnknown(m)
}

var xxx_messageInfo_DataClock proto.InternalMessageInfo

func (m *DataClock) GetClock() int64 {
	if m != nil {
		return m.Clock
	}
	return 0
}

var E_Dlg = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*DialogOptions)(nil),
	Field:         100001,
	Name:          "dialog.dlg",
	Tag:           "bytes,100001,opt,name=dlg",
	Filename:      "definitions.proto",
}

func init() {
	proto.RegisterType((*UUIDValue)(nil), "dialog.UUIDValue")
	proto.RegisterType((*DialogOptions)(nil), "dialog.DialogOptions")
	proto.RegisterType((*DataClock)(nil), "dialog.DataClock")
	proto.RegisterExtension(E_Dlg)
}

func init() { proto.RegisterFile("definitions.proto", fileDescriptor_bef01fe2be18d2f0) }

var fileDescriptor_bef01fe2be18d2f0 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x3f, 0x4b, 0x73, 0x31,
	0x14, 0xc6, 0x6f, 0xde, 0xf2, 0x16, 0x1a, 0x11, 0xf4, 0xd2, 0x4a, 0x29, 0x78, 0x68, 0x3b, 0x75,
	0xca, 0x05, 0xdd, 0x5c, 0x04, 0x2d, 0xa2, 0x93, 0x50, 0xa8, 0x83, 0x5b, 0xfe, 0x35, 0x04, 0xd3,
	0xe6, 0x72, 0x73, 0xeb, 0xec, 0x17, 0x10, 0xfc, 0x0a, 0x6e, 0x7e, 0x14, 0xc7, 0x8e, 0x1d, 0x6d,
	0xee, 0xe2, 0xd8, 0x8f, 0x20, 0x37, 0xb9, 0x0e, 0x4e, 0x79, 0xce, 0xc9, 0xc9, 0xf3, 0xfc, 0x72,
	0xf0, 0xb1, 0x90, 0x0b, 0xbd, 0xd2, 0xa5, 0xb6, 0x2b, 0x47, 0xf2, 0xc2, 0x96, 0x36, 0x6d, 0x0b,
	0x4d, 0x8d, 0x55, 0x83, 0x9e, 0xe3, 0xd4, 0xd0, 0x9c, 0x65, 0xcd, 0x19, 0xaf, 0x07, 0x43, 0x65,
	0xad, 0x32, 0x32, 0x0b, 0x15, 0x5b, 0x2f, 0x32, 0x21, 0x1d, 0x2f, 0x74, 0x5e, 0xda, 0x22, 0x4e,
	0x8c, 0x33, 0xdc, 0x99, 0xcf, 0xef, 0xa6, 0x0f, 0xd4, 0xac, 0x65, 0x7a, 0x84, 0x5b, 0x4b, 0xc7,
	0xfa, 0x68, 0x88, 0x26, 0xad, 0x59, 0x2d, 0xeb, 0x8e, 0x71, 0xac, 0xff, 0x2f, 0x76, 0x8c, 0x63,
	0xe3, 0x11, 0x3e, 0x9c, 0x86, 0xcc, 0xfb, 0x3c, 0x80, 0x84, 0x11, 0xab, 0xc2, 0xa3, 0xce, 0xac,
	0x96, 0xe3, 0x11, 0xee, 0x4c, 0x69, 0x49, 0xaf, 0x8d, 0xe5, 0x4f, 0x69, 0x17, 0xff, 0xe7, 0xb5,
	0x68, 0x5c, 0x63, 0x71, 0x71, 0x8b, 0x5b, 0xc2, 0xa8, 0xf4, 0x94, 0x44, 0x40, 0xf2, 0x0b, 0x48,
	0x6e, 0xb4, 0x34, 0xa2, 0xb1, 0xee, 0xbf, 0xbf, 0xb6, 0x87, 0x68, 0x72, 0x70, 0xd6, 0x23, 0xf1,
	0x9b, 0xe4, 0x4f, 0xf2, 0xac, 0xb6, 0xb8, 0x9a, 0xfb, 0xcb, 0x13, 0xdc, 0xd5, 0x4b, 0x22, 0x8c,
	0x22, 0xaa, 0xc8, 0x39, 0x71, 0xb2, 0x78, 0xd6, 0x5c, 0xba, 0xcd, 0x0e, 0x92, 0xed, 0x0e, 0x92,
	0xfd, 0x0e, 0xd0, 0x8b, 0x07, 0xf4, 0xe1, 0x01, 0x7d, 0x7a, 0x40, 0x1b, 0x0f, 0xe8, 0xcb, 0x03,
	0xfa, 0xf6, 0x90, 0xec, 0x3d, 0xa0, 0xb7, 0x0a, 0x92, 0x4d, 0x05, 0xc9, 0xb6, 0x82, 0xe4, 0xb1,
	0x59, 0x28, 0x6b, 0x07, 0xa2, 0xf3, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x53, 0x7b, 0x73, 0x4b,
	0x74, 0x01, 0x00, 0x00,
}

func (this *UUIDValue) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UUIDValue)
	if !ok {
		that2, ok := that.(UUIDValue)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Msb != that1.Msb {
		return false
	}
	if this.Lsb != that1.Lsb {
		return false
	}
	return true
}
func (this *DialogOptions) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DialogOptions)
	if !ok {
		that2, ok := that.(DialogOptions)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Log != that1.Log {
		return false
	}
	return true
}
func (this *DataClock) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DataClock)
	if !ok {
		that2, ok := that.(DataClock)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Clock != that1.Clock {
		return false
	}
	return true
}
func (this *UUIDValue) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&dialog.UUIDValue{")
	s = append(s, "Msb: "+fmt.Sprintf("%#v", this.Msb)+",\n")
	s = append(s, "Lsb: "+fmt.Sprintf("%#v", this.Lsb)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *DialogOptions) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&dialog.DialogOptions{")
	s = append(s, "Log: "+fmt.Sprintf("%#v", this.Log)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *DataClock) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&dialog.DataClock{")
	s = append(s, "Clock: "+fmt.Sprintf("%#v", this.Clock)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringDefinitions(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *UUIDValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UUIDValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UUIDValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Lsb != 0 {
		i = encodeVarintDefinitions(dAtA, i, uint64(m.Lsb))
		i--
		dAtA[i] = 0x10
	}
	if m.Msb != 0 {
		i = encodeVarintDefinitions(dAtA, i, uint64(m.Msb))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DialogOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DialogOptions) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DialogOptions) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Log) > 0 {
		i -= len(m.Log)
		copy(dAtA[i:], m.Log)
		i = encodeVarintDefinitions(dAtA, i, uint64(len(m.Log)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DataClock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataClock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DataClock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Clock != 0 {
		i = encodeVarintDefinitions(dAtA, i, uint64(m.Clock))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintDefinitions(dAtA []byte, offset int, v uint64) int {
	offset -= sovDefinitions(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UUIDValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Msb != 0 {
		n += 1 + sovDefinitions(uint64(m.Msb))
	}
	if m.Lsb != 0 {
		n += 1 + sovDefinitions(uint64(m.Lsb))
	}
	return n
}

func (m *DialogOptions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Log)
	if l > 0 {
		n += 1 + l + sovDefinitions(uint64(l))
	}
	return n
}

func (m *DataClock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Clock != 0 {
		n += 1 + sovDefinitions(uint64(m.Clock))
	}
	return n
}

func sovDefinitions(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDefinitions(x uint64) (n int) {
	return sovDefinitions(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *UUIDValue) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&UUIDValue{`,
		`Msb:` + fmt.Sprintf("%v", this.Msb) + `,`,
		`Lsb:` + fmt.Sprintf("%v", this.Lsb) + `,`,
		`}`,
	}, "")
	return s
}
func (this *DialogOptions) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DialogOptions{`,
		`Log:` + fmt.Sprintf("%v", this.Log) + `,`,
		`}`,
	}, "")
	return s
}
func (this *DataClock) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DataClock{`,
		`Clock:` + fmt.Sprintf("%v", this.Clock) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringDefinitions(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *UUIDValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDefinitions
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
			return fmt.Errorf("proto: UUIDValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UUIDValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msb", wireType)
			}
			m.Msb = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDefinitions
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Msb |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lsb", wireType)
			}
			m.Lsb = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDefinitions
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Lsb |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDefinitions(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDefinitions
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDefinitions
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
func (m *DialogOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDefinitions
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
			return fmt.Errorf("proto: DialogOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DialogOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Log", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDefinitions
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
				return ErrInvalidLengthDefinitions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDefinitions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Log = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDefinitions(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDefinitions
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDefinitions
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
func (m *DataClock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDefinitions
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
			return fmt.Errorf("proto: DataClock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataClock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Clock", wireType)
			}
			m.Clock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDefinitions
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Clock |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDefinitions(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDefinitions
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDefinitions
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
func skipDefinitions(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDefinitions
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
					return 0, ErrIntOverflowDefinitions
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
					return 0, ErrIntOverflowDefinitions
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
				return 0, ErrInvalidLengthDefinitions
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDefinitions
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDefinitions
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDefinitions        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDefinitions          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDefinitions = fmt.Errorf("proto: unexpected end of group")
)
