// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: permissions.proto

package dialog

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strconv "strconv"
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

// Permissions definition
type Permission int32

const (
	PERMISSION_UNKNOWN       Permission = 0
	PERMISSION_SEARCH        Permission = 1
	PERMISSION_CREATE_GROUPS Permission = 2
)

var Permission_name = map[int32]string{
	0: "PERMISSION_UNKNOWN",
	1: "PERMISSION_SEARCH",
	2: "PERMISSION_CREATE_GROUPS",
}

var Permission_value = map[string]int32{
	"PERMISSION_UNKNOWN":       0,
	"PERMISSION_SEARCH":        1,
	"PERMISSION_CREATE_GROUPS": 2,
}

func (Permission) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_46cca66312ac1c30, []int{0}
}

type RequestGetPermissions struct {
	Clock int64 `protobuf:"varint,1,opt,name=clock,proto3" json:"clock,omitempty"`
}

func (m *RequestGetPermissions) Reset()      { *m = RequestGetPermissions{} }
func (*RequestGetPermissions) ProtoMessage() {}
func (*RequestGetPermissions) Descriptor() ([]byte, []int) {
	return fileDescriptor_46cca66312ac1c30, []int{0}
}
func (m *RequestGetPermissions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestGetPermissions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestGetPermissions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestGetPermissions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestGetPermissions.Merge(m, src)
}
func (m *RequestGetPermissions) XXX_Size() int {
	return m.Size()
}
func (m *RequestGetPermissions) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestGetPermissions.DiscardUnknown(m)
}

var xxx_messageInfo_RequestGetPermissions proto.InternalMessageInfo

func (m *RequestGetPermissions) GetClock() int64 {
	if m != nil {
		return m.Clock
	}
	return 0
}

type ResponseGetPermissions struct {
	Permissions         []Permission `protobuf:"varint,1,rep,packed,name=permissions,proto3,enum=dialog.Permission" json:"permissions,omitempty"`
	AllKnownPermissions []Permission `protobuf:"varint,2,rep,packed,name=all_known_permissions,json=allKnownPermissions,proto3,enum=dialog.Permission" json:"all_known_permissions,omitempty"`
	Clock               int64        `protobuf:"varint,3,opt,name=clock,proto3" json:"clock,omitempty"`
}

func (m *ResponseGetPermissions) Reset()      { *m = ResponseGetPermissions{} }
func (*ResponseGetPermissions) ProtoMessage() {}
func (*ResponseGetPermissions) Descriptor() ([]byte, []int) {
	return fileDescriptor_46cca66312ac1c30, []int{1}
}
func (m *ResponseGetPermissions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseGetPermissions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseGetPermissions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseGetPermissions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseGetPermissions.Merge(m, src)
}
func (m *ResponseGetPermissions) XXX_Size() int {
	return m.Size()
}
func (m *ResponseGetPermissions) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseGetPermissions.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseGetPermissions proto.InternalMessageInfo

func (m *ResponseGetPermissions) GetPermissions() []Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *ResponseGetPermissions) GetAllKnownPermissions() []Permission {
	if m != nil {
		return m.AllKnownPermissions
	}
	return nil
}

func (m *ResponseGetPermissions) GetClock() int64 {
	if m != nil {
		return m.Clock
	}
	return 0
}

type UpdatePermissionsChange struct {
	Permissions []Permission `protobuf:"varint,1,rep,packed,name=permissions,proto3,enum=dialog.Permission" json:"permissions,omitempty"`
	Clock       int64        `protobuf:"varint,2,opt,name=clock,proto3" json:"clock,omitempty"`
}

func (m *UpdatePermissionsChange) Reset()      { *m = UpdatePermissionsChange{} }
func (*UpdatePermissionsChange) ProtoMessage() {}
func (*UpdatePermissionsChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_46cca66312ac1c30, []int{2}
}
func (m *UpdatePermissionsChange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdatePermissionsChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdatePermissionsChange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdatePermissionsChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePermissionsChange.Merge(m, src)
}
func (m *UpdatePermissionsChange) XXX_Size() int {
	return m.Size()
}
func (m *UpdatePermissionsChange) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePermissionsChange.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePermissionsChange proto.InternalMessageInfo

func (m *UpdatePermissionsChange) GetPermissions() []Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *UpdatePermissionsChange) GetClock() int64 {
	if m != nil {
		return m.Clock
	}
	return 0
}

func init() {
	proto.RegisterEnum("dialog.Permission", Permission_name, Permission_value)
	proto.RegisterType((*RequestGetPermissions)(nil), "dialog.RequestGetPermissions")
	proto.RegisterType((*ResponseGetPermissions)(nil), "dialog.ResponseGetPermissions")
	proto.RegisterType((*UpdatePermissionsChange)(nil), "dialog.UpdatePermissionsChange")
}

func init() { proto.RegisterFile("permissions.proto", fileDescriptor_46cca66312ac1c30) }

var fileDescriptor_46cca66312ac1c30 = []byte{
	// 451 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x41, 0x6b, 0x13, 0x41,
	0x18, 0xdd, 0x49, 0xb0, 0xe2, 0x57, 0x2c, 0xc9, 0x68, 0x62, 0x08, 0x75, 0x28, 0xe9, 0x25, 0xf4,
	0xb0, 0xab, 0xf5, 0x26, 0x82, 0xb4, 0x21, 0xd4, 0x52, 0x4c, 0xc2, 0xc6, 0x45, 0xf4, 0xb2, 0x4c,
	0x76, 0xc7, 0x75, 0xe8, 0x74, 0x66, 0xdd, 0x59, 0xe3, 0x51, 0xd1, 0x9b, 0x07, 0x11, 0xfc, 0x13,
	0xfe, 0x14, 0x8f, 0x01, 0x2f, 0xf5, 0x66, 0x26, 0x1e, 0xc4, 0x53, 0x7f, 0x82, 0xd4, 0x54, 0x77,
	0x0c, 0x51, 0x3c, 0x78, 0x5a, 0xf8, 0xde, 0xf7, 0xde, 0xbe, 0x37, 0xf3, 0x06, 0xaa, 0x29, 0xcb,
	0x8e, 0xb8, 0xd6, 0x5c, 0x49, 0xed, 0xa6, 0x99, 0xca, 0x15, 0x5e, 0x89, 0x39, 0x15, 0x2a, 0x69,
	0x56, 0x63, 0xf6, 0x88, 0x4b, 0x9e, 0x17, 0x50, 0x73, 0x3d, 0x51, 0x2a, 0x11, 0xcc, 0xa3, 0x29,
	0xf7, 0xa8, 0x94, 0x2a, 0xa7, 0x36, 0x5a, 0xd3, 0x11, 0x15, 0x34, 0x1d, 0x79, 0x67, 0xdf, 0xf9,
	0xb8, 0x75, 0x0b, 0x6a, 0x3e, 0x7b, 0xf2, 0x94, 0xe9, 0x7c, 0x8f, 0xe5, 0x83, 0xe2, 0x77, 0x78,
	0x13, 0xce, 0x45, 0x42, 0x45, 0x87, 0x0d, 0xb4, 0x81, 0xda, 0xe5, 0xdd, 0x8b, 0xaf, 0xbf, 0x5d,
	0xbb, 0x00, 0xe7, 0xc7, 0x5c, 0xf3, 0x91, 0x60, 0xfe, 0x1c, 0x6b, 0x7d, 0x42, 0x50, 0xf7, 0x99,
	0x4e, 0x95, 0xd4, 0x6c, 0x81, 0xdf, 0x81, 0x55, 0xcb, 0x7d, 0x03, 0x6d, 0x94, 0xdb, 0x6b, 0xdb,
	0xd8, 0x9d, 0xdb, 0x77, 0x8b, 0xcd, 0x45, 0x65, 0x9b, 0x85, 0x03, 0xa8, 0x51, 0x21, 0xc2, 0x43,
	0xa9, 0x9e, 0xc9, 0xd0, 0x96, 0x2b, 0xfd, 0xab, 0xdc, 0x25, 0x2a, 0xc4, 0xc1, 0x29, 0x7d, 0x69,
	0xb6, 0xf2, 0x5f, 0xb2, 0xbd, 0x42, 0x70, 0x25, 0x48, 0x63, 0x9a, 0x33, 0x8b, 0xda, 0x79, 0x4c,
	0x65, 0xc2, 0xfe, 0x4f, 0xb8, 0x5f, 0x2e, 0x4a, 0x7f, 0x76, 0xb1, 0xf5, 0x00, 0xa0, 0x90, 0xc3,
	0x75, 0xc0, 0x83, 0xae, 0x7f, 0x77, 0x7f, 0x38, 0xdc, 0xef, 0xf7, 0xc2, 0xa0, 0x77, 0xd0, 0xeb,
	0xdf, 0xef, 0x55, 0x1c, 0x5c, 0x83, 0xaa, 0x35, 0x1f, 0x76, 0x77, 0xfc, 0xce, 0x9d, 0x0a, 0xc2,
	0xeb, 0xd0, 0xb0, 0xc6, 0x1d, 0xbf, 0xbb, 0x73, 0xaf, 0x1b, 0xee, 0xf9, 0xfd, 0x60, 0x30, 0xac,
	0x94, 0xb6, 0xdf, 0x20, 0x58, 0xb5, 0x4f, 0xe5, 0x39, 0xac, 0x2d, 0xdc, 0xe1, 0xd5, 0x9f, 0x89,
	0x96, 0x56, 0xa4, 0x49, 0x0a, 0x78, 0x59, 0x05, 0x5a, 0xee, 0xcb, 0x8f, 0x5f, 0xde, 0x95, 0xda,
	0xad, 0x4d, 0x6f, 0x7c, 0xdd, 0x4b, 0xb2, 0x34, 0xf2, 0x2c, 0xd4, 0xfb, 0x7d, 0xf9, 0x26, 0xda,
	0xda, 0x0d, 0xcc, 0xed, 0x3a, 0x5c, 0xe6, 0x47, 0x6e, 0x2c, 0x12, 0xf7, 0x74, 0xdf, 0xd5, 0x2c,
	0x1b, 0xf3, 0x88, 0xe9, 0xc9, 0x94, 0x38, 0xc7, 0x53, 0xe2, 0x9c, 0x4c, 0x09, 0x7a, 0x61, 0x08,
	0x7a, 0x6f, 0x08, 0xfa, 0x60, 0x08, 0x9a, 0x18, 0x82, 0x3e, 0x1b, 0x82, 0xbe, 0x1a, 0xe2, 0x9c,
	0x18, 0x82, 0xde, 0xce, 0x88, 0x33, 0x99, 0x11, 0xe7, 0x78, 0x46, 0x9c, 0x87, 0x67, 0x4f, 0x65,
	0xb4, 0xf2, 0xa3, 0xe9, 0x37, 0xbe, 0x07, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x78, 0x06, 0x6f, 0x4e,
	0x03, 0x00, 0x00,
}

func (x Permission) String() string {
	s, ok := Permission_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *RequestGetPermissions) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RequestGetPermissions)
	if !ok {
		that2, ok := that.(RequestGetPermissions)
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
func (this *ResponseGetPermissions) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ResponseGetPermissions)
	if !ok {
		that2, ok := that.(ResponseGetPermissions)
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
	if len(this.Permissions) != len(that1.Permissions) {
		return false
	}
	for i := range this.Permissions {
		if this.Permissions[i] != that1.Permissions[i] {
			return false
		}
	}
	if len(this.AllKnownPermissions) != len(that1.AllKnownPermissions) {
		return false
	}
	for i := range this.AllKnownPermissions {
		if this.AllKnownPermissions[i] != that1.AllKnownPermissions[i] {
			return false
		}
	}
	if this.Clock != that1.Clock {
		return false
	}
	return true
}
func (this *UpdatePermissionsChange) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpdatePermissionsChange)
	if !ok {
		that2, ok := that.(UpdatePermissionsChange)
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
	if len(this.Permissions) != len(that1.Permissions) {
		return false
	}
	for i := range this.Permissions {
		if this.Permissions[i] != that1.Permissions[i] {
			return false
		}
	}
	if this.Clock != that1.Clock {
		return false
	}
	return true
}
func (this *RequestGetPermissions) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&dialog.RequestGetPermissions{")
	s = append(s, "Clock: "+fmt.Sprintf("%#v", this.Clock)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ResponseGetPermissions) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&dialog.ResponseGetPermissions{")
	s = append(s, "Permissions: "+fmt.Sprintf("%#v", this.Permissions)+",\n")
	s = append(s, "AllKnownPermissions: "+fmt.Sprintf("%#v", this.AllKnownPermissions)+",\n")
	s = append(s, "Clock: "+fmt.Sprintf("%#v", this.Clock)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *UpdatePermissionsChange) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&dialog.UpdatePermissionsChange{")
	s = append(s, "Permissions: "+fmt.Sprintf("%#v", this.Permissions)+",\n")
	s = append(s, "Clock: "+fmt.Sprintf("%#v", this.Clock)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringPermissions(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PermissionsClient is the client API for Permissions service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PermissionsClient interface {
	GetPermissions(ctx context.Context, in *RequestGetPermissions, opts ...grpc.CallOption) (*ResponseGetPermissions, error)
}

type permissionsClient struct {
	cc *grpc.ClientConn
}

func NewPermissionsClient(cc *grpc.ClientConn) PermissionsClient {
	return &permissionsClient{cc}
}

func (c *permissionsClient) GetPermissions(ctx context.Context, in *RequestGetPermissions, opts ...grpc.CallOption) (*ResponseGetPermissions, error) {
	out := new(ResponseGetPermissions)
	err := c.cc.Invoke(ctx, "/dialog.Permissions/GetPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionsServer is the server API for Permissions service.
type PermissionsServer interface {
	GetPermissions(context.Context, *RequestGetPermissions) (*ResponseGetPermissions, error)
}

// UnimplementedPermissionsServer can be embedded to have forward compatible implementations.
type UnimplementedPermissionsServer struct {
}

func (*UnimplementedPermissionsServer) GetPermissions(ctx context.Context, req *RequestGetPermissions) (*ResponseGetPermissions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPermissions not implemented")
}

func RegisterPermissionsServer(s *grpc.Server, srv PermissionsServer) {
	s.RegisterService(&_Permissions_serviceDesc, srv)
}

func _Permissions_GetPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestGetPermissions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionsServer).GetPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dialog.Permissions/GetPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionsServer).GetPermissions(ctx, req.(*RequestGetPermissions))
	}
	return interceptor(ctx, in, info, handler)
}

var _Permissions_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dialog.Permissions",
	HandlerType: (*PermissionsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPermissions",
			Handler:    _Permissions_GetPermissions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "permissions.proto",
}

func (m *RequestGetPermissions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestGetPermissions) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestGetPermissions) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Clock != 0 {
		i = encodeVarintPermissions(dAtA, i, uint64(m.Clock))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ResponseGetPermissions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseGetPermissions) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseGetPermissions) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Clock != 0 {
		i = encodeVarintPermissions(dAtA, i, uint64(m.Clock))
		i--
		dAtA[i] = 0x18
	}
	if len(m.AllKnownPermissions) > 0 {
		dAtA2 := make([]byte, len(m.AllKnownPermissions)*10)
		var j1 int
		for _, num := range m.AllKnownPermissions {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintPermissions(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Permissions) > 0 {
		dAtA4 := make([]byte, len(m.Permissions)*10)
		var j3 int
		for _, num := range m.Permissions {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		i -= j3
		copy(dAtA[i:], dAtA4[:j3])
		i = encodeVarintPermissions(dAtA, i, uint64(j3))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UpdatePermissionsChange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdatePermissionsChange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdatePermissionsChange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Clock != 0 {
		i = encodeVarintPermissions(dAtA, i, uint64(m.Clock))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Permissions) > 0 {
		dAtA6 := make([]byte, len(m.Permissions)*10)
		var j5 int
		for _, num := range m.Permissions {
			for num >= 1<<7 {
				dAtA6[j5] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j5++
			}
			dAtA6[j5] = uint8(num)
			j5++
		}
		i -= j5
		copy(dAtA[i:], dAtA6[:j5])
		i = encodeVarintPermissions(dAtA, i, uint64(j5))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPermissions(dAtA []byte, offset int, v uint64) int {
	offset -= sovPermissions(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RequestGetPermissions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Clock != 0 {
		n += 1 + sovPermissions(uint64(m.Clock))
	}
	return n
}

func (m *ResponseGetPermissions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Permissions) > 0 {
		l = 0
		for _, e := range m.Permissions {
			l += sovPermissions(uint64(e))
		}
		n += 1 + sovPermissions(uint64(l)) + l
	}
	if len(m.AllKnownPermissions) > 0 {
		l = 0
		for _, e := range m.AllKnownPermissions {
			l += sovPermissions(uint64(e))
		}
		n += 1 + sovPermissions(uint64(l)) + l
	}
	if m.Clock != 0 {
		n += 1 + sovPermissions(uint64(m.Clock))
	}
	return n
}

func (m *UpdatePermissionsChange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Permissions) > 0 {
		l = 0
		for _, e := range m.Permissions {
			l += sovPermissions(uint64(e))
		}
		n += 1 + sovPermissions(uint64(l)) + l
	}
	if m.Clock != 0 {
		n += 1 + sovPermissions(uint64(m.Clock))
	}
	return n
}

func sovPermissions(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPermissions(x uint64) (n int) {
	return sovPermissions(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *RequestGetPermissions) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RequestGetPermissions{`,
		`Clock:` + fmt.Sprintf("%v", this.Clock) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ResponseGetPermissions) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ResponseGetPermissions{`,
		`Permissions:` + fmt.Sprintf("%v", this.Permissions) + `,`,
		`AllKnownPermissions:` + fmt.Sprintf("%v", this.AllKnownPermissions) + `,`,
		`Clock:` + fmt.Sprintf("%v", this.Clock) + `,`,
		`}`,
	}, "")
	return s
}
func (this *UpdatePermissionsChange) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&UpdatePermissionsChange{`,
		`Permissions:` + fmt.Sprintf("%v", this.Permissions) + `,`,
		`Clock:` + fmt.Sprintf("%v", this.Clock) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringPermissions(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *RequestGetPermissions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPermissions
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
			return fmt.Errorf("proto: RequestGetPermissions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestGetPermissions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Clock", wireType)
			}
			m.Clock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermissions
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
			skippy, err := skipPermissions(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPermissions
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPermissions
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
func (m *ResponseGetPermissions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPermissions
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
			return fmt.Errorf("proto: ResponseGetPermissions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseGetPermissions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v Permission
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPermissions
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= Permission(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Permissions = append(m.Permissions, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPermissions
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthPermissions
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthPermissions
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				if elementCount != 0 && len(m.Permissions) == 0 {
					m.Permissions = make([]Permission, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v Permission
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowPermissions
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= Permission(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Permissions = append(m.Permissions, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Permissions", wireType)
			}
		case 2:
			if wireType == 0 {
				var v Permission
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPermissions
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= Permission(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.AllKnownPermissions = append(m.AllKnownPermissions, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPermissions
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthPermissions
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthPermissions
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				if elementCount != 0 && len(m.AllKnownPermissions) == 0 {
					m.AllKnownPermissions = make([]Permission, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v Permission
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowPermissions
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= Permission(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.AllKnownPermissions = append(m.AllKnownPermissions, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field AllKnownPermissions", wireType)
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Clock", wireType)
			}
			m.Clock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermissions
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
			skippy, err := skipPermissions(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPermissions
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPermissions
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
func (m *UpdatePermissionsChange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPermissions
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
			return fmt.Errorf("proto: UpdatePermissionsChange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdatePermissionsChange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v Permission
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPermissions
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= Permission(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Permissions = append(m.Permissions, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPermissions
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthPermissions
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthPermissions
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				if elementCount != 0 && len(m.Permissions) == 0 {
					m.Permissions = make([]Permission, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v Permission
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowPermissions
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= Permission(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Permissions = append(m.Permissions, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Permissions", wireType)
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Clock", wireType)
			}
			m.Clock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermissions
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
			skippy, err := skipPermissions(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPermissions
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPermissions
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
func skipPermissions(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPermissions
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
					return 0, ErrIntOverflowPermissions
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
					return 0, ErrIntOverflowPermissions
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
				return 0, ErrInvalidLengthPermissions
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPermissions
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPermissions
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPermissions        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPermissions          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPermissions = fmt.Errorf("proto: unexpected end of group")
)