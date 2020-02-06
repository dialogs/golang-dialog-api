// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: device_info.proto

package dialog

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

type PlatformType int32

const (
	PLATFORMTYPE_UNKNOWN PlatformType = 0
	PLATFORMTYPE_ANDROID PlatformType = 1
	PLATFORMTYPE_IOS     PlatformType = 2
	PLATFORMTYPE_WEB     PlatformType = 3
	PLATFORMTYPE_CLC     PlatformType = 4
	PLATFORMTYPE_TESTS   PlatformType = 42
)

var PlatformType_name = map[int32]string{
	0:  "PLATFORMTYPE_UNKNOWN",
	1:  "PLATFORMTYPE_ANDROID",
	2:  "PLATFORMTYPE_IOS",
	3:  "PLATFORMTYPE_WEB",
	4:  "PLATFORMTYPE_CLC",
	42: "PLATFORMTYPE_TESTS",
}

var PlatformType_value = map[string]int32{
	"PLATFORMTYPE_UNKNOWN": 0,
	"PLATFORMTYPE_ANDROID": 1,
	"PLATFORMTYPE_IOS":     2,
	"PLATFORMTYPE_WEB":     3,
	"PLATFORMTYPE_CLC":     4,
	"PLATFORMTYPE_TESTS":   42,
}

func (PlatformType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f9d543922a5df507, []int{0}
}

// Notifying about device information
type RequestNotifyAboutDeviceInfo struct {
	//* First language from this array will be used for some notifications from server *
	PreferredLanguages []string `protobuf:"bytes,1,rep,name=preferred_languages,json=preferredLanguages,proto3" json:"preferred_languages,omitempty"`
	//* Your timezone *
	TimeZone *types.StringValue `protobuf:"bytes,2,opt,name=time_zone,json=timeZone,proto3" json:"time_zone,omitempty"`
}

func (m *RequestNotifyAboutDeviceInfo) Reset()      { *m = RequestNotifyAboutDeviceInfo{} }
func (*RequestNotifyAboutDeviceInfo) ProtoMessage() {}
func (*RequestNotifyAboutDeviceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9d543922a5df507, []int{0}
}
func (m *RequestNotifyAboutDeviceInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestNotifyAboutDeviceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestNotifyAboutDeviceInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestNotifyAboutDeviceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestNotifyAboutDeviceInfo.Merge(m, src)
}
func (m *RequestNotifyAboutDeviceInfo) XXX_Size() int {
	return m.Size()
}
func (m *RequestNotifyAboutDeviceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestNotifyAboutDeviceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RequestNotifyAboutDeviceInfo proto.InternalMessageInfo

func (m *RequestNotifyAboutDeviceInfo) GetPreferredLanguages() []string {
	if m != nil {
		return m.PreferredLanguages
	}
	return nil
}

func (m *RequestNotifyAboutDeviceInfo) GetTimeZone() *types.StringValue {
	if m != nil {
		return m.TimeZone
	}
	return nil
}

// Generic client info, containing information about platform type, version, sdk etc
type ClientInfo struct {
	/// The platform enum. Can be either Android, Web or iOS
	Platform PlatformType `protobuf:"varint,1,opt,name=platform,proto3,enum=dialog.PlatformType" json:"platform,omitempty"`
	/// For android: vendor and model; For iOS: model; For Web: platform and user agent
	DeviceName string `protobuf:"bytes,2,opt,name=device_name,json=deviceName,proto3" json:"device_name,omitempty"`
	/// Optional application name
	AppName string `protobuf:"bytes,3,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
	/// Application version
	AppVersion *types.StringValue `protobuf:"bytes,4,opt,name=app_version,json=appVersion,proto3" json:"app_version,omitempty"`
	/// Optional SDK version
	SdkVersion *types.StringValue `protobuf:"bytes,5,opt,name=sdk_version,json=sdkVersion,proto3" json:"sdk_version,omitempty"`
	/// Optional ISO-639 language code and ISO-3166 country code: ru-RU
	PreferredLanguages []string `protobuf:"bytes,6,rep,name=preferred_languages,json=preferredLanguages,proto3" json:"preferred_languages,omitempty"`
	/// Optional TimeZone id
	TimeZone *types.StringValue `protobuf:"bytes,7,opt,name=time_zone,json=timeZone,proto3" json:"time_zone,omitempty"`
}

func (m *ClientInfo) Reset()      { *m = ClientInfo{} }
func (*ClientInfo) ProtoMessage() {}
func (*ClientInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9d543922a5df507, []int{1}
}
func (m *ClientInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClientInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClientInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClientInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientInfo.Merge(m, src)
}
func (m *ClientInfo) XXX_Size() int {
	return m.Size()
}
func (m *ClientInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ClientInfo proto.InternalMessageInfo

func (m *ClientInfo) GetPlatform() PlatformType {
	if m != nil {
		return m.Platform
	}
	return PLATFORMTYPE_UNKNOWN
}

func (m *ClientInfo) GetDeviceName() string {
	if m != nil {
		return m.DeviceName
	}
	return ""
}

func (m *ClientInfo) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *ClientInfo) GetAppVersion() *types.StringValue {
	if m != nil {
		return m.AppVersion
	}
	return nil
}

func (m *ClientInfo) GetSdkVersion() *types.StringValue {
	if m != nil {
		return m.SdkVersion
	}
	return nil
}

func (m *ClientInfo) GetPreferredLanguages() []string {
	if m != nil {
		return m.PreferredLanguages
	}
	return nil
}

func (m *ClientInfo) GetTimeZone() *types.StringValue {
	if m != nil {
		return m.TimeZone
	}
	return nil
}

func init() {
	proto.RegisterEnum("dialog.PlatformType", PlatformType_name, PlatformType_value)
	proto.RegisterType((*RequestNotifyAboutDeviceInfo)(nil), "dialog.RequestNotifyAboutDeviceInfo")
	proto.RegisterType((*ClientInfo)(nil), "dialog.ClientInfo")
}

func init() { proto.RegisterFile("device_info.proto", fileDescriptor_f9d543922a5df507) }

var fileDescriptor_f9d543922a5df507 = []byte{
	// 622 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x3f, 0x6f, 0xd3, 0x4e,
	0x18, 0xc7, 0x7d, 0x4d, 0x7f, 0xfd, 0x73, 0xfd, 0x81, 0x5c, 0x37, 0x2d, 0xa1, 0x8a, 0x4e, 0x55,
	0xc4, 0x10, 0x3a, 0xd8, 0x50, 0x98, 0x3a, 0x50, 0xf5, 0x1f, 0x50, 0x11, 0x9c, 0x2a, 0x49, 0x5b,
	0xd1, 0x25, 0xba, 0xc4, 0x67, 0xeb, 0x54, 0xfb, 0xee, 0xb8, 0xb3, 0x83, 0xca, 0x84, 0xd8, 0x60,
	0x42, 0x62, 0x61, 0xe0, 0x05, 0xf0, 0x4a, 0x10, 0x0b, 0x52, 0x25, 0x96, 0x8e, 0xd4, 0x65, 0x40,
	0x4c, 0x7d, 0x09, 0x28, 0xb1, 0xdd, 0xa6, 0xa4, 0x20, 0x51, 0xa6, 0x48, 0xdf, 0xe7, 0xf3, 0x7c,
	0xf3, 0xdc, 0x7d, 0xef, 0x31, 0x9c, 0x74, 0x48, 0x87, 0xb6, 0x49, 0x93, 0x32, 0x97, 0x9b, 0x42,
	0xf2, 0x90, 0x1b, 0x23, 0x0e, 0xc5, 0x3e, 0xf7, 0x66, 0x91, 0xc7, 0xb9, 0xe7, 0x13, 0xab, 0xa7,
	0xb6, 0x22, 0xd7, 0x7a, 0x26, 0xb1, 0x10, 0x44, 0xaa, 0x84, 0x9b, 0x2d, 0xa6, 0x75, 0x2c, 0xa8,
	0x85, 0x19, 0xe3, 0x21, 0x0e, 0x29, 0x67, 0x59, 0x75, 0xd2, 0x21, 0x2e, 0x65, 0xb4, 0x5f, 0x9a,
	0x0a, 0xa8, 0x6a, 0x13, 0xdf, 0xc7, 0x8c, 0xf0, 0x28, 0x13, 0xa7, 0x55, 0x1b, 0xfb, 0x58, 0xb4,
	0xac, 0xf4, 0x37, 0x91, 0x4b, 0x1f, 0x01, 0x2c, 0xd6, 0xc8, 0xd3, 0x88, 0xa8, 0xd0, 0xe6, 0x21,
	0x75, 0xf7, 0x97, 0x5b, 0x3c, 0x0a, 0xd7, 0x7a, 0xc3, 0x6e, 0x30, 0x97, 0x1b, 0xf7, 0xe0, 0x94,
	0x90, 0xc4, 0x25, 0x52, 0x12, 0xa7, 0xe9, 0x63, 0xe6, 0x45, 0xd8, 0x23, 0xaa, 0x00, 0xe6, 0x72,
	0xe5, 0xf1, 0x95, 0x2b, 0xaf, 0x7f, 0xdc, 0x1a, 0x87, 0xa3, 0x1d, 0xaa, 0x68, 0xcb, 0x27, 0x35,
	0xe3, 0x94, 0xac, 0x64, 0xa0, 0xf1, 0x10, 0x8e, 0x87, 0x34, 0x20, 0xcd, 0xe7, 0x9c, 0x91, 0xc2,
	0xd0, 0x1c, 0x28, 0x4f, 0x2c, 0x14, 0xcd, 0xe4, 0x44, 0x66, 0x76, 0x62, 0xb3, 0x1e, 0x4a, 0xca,
	0xbc, 0x6d, 0xec, 0x47, 0xe4, 0x57, 0xcf, 0xb1, 0x6e, 0xf7, 0x2e, 0x67, 0x64, 0xb1, 0x18, 0x2f,
	0x5d, 0x87, 0xd7, 0x68, 0x60, 0x3a, 0xbe, 0x67, 0x7a, 0x52, 0xb4, 0xcd, 0x07, 0x52, 0xb4, 0xd3,
	0xe1, 0x4b, 0x9f, 0x73, 0x10, 0xae, 0xfa, 0x94, 0xb0, 0xb0, 0x37, 0xf6, 0x12, 0x1c, 0x13, 0x3e,
	0x0e, 0x5d, 0x2e, 0x83, 0x02, 0x98, 0x03, 0xe5, 0xab, 0x0b, 0x79, 0x33, 0xb9, 0x6f, 0x73, 0x33,
	0xd5, 0x1b, 0xfb, 0x62, 0xf0, 0xdf, 0xb2, 0x26, 0xc3, 0x84, 0x13, 0x69, 0x64, 0x0c, 0x07, 0xc9,
	0xe4, 0x03, 0xe7, 0x85, 0x09, 0x61, 0xe3, 0x80, 0x18, 0x65, 0x38, 0x86, 0x85, 0x48, 0xe0, 0xdc,
	0x45, 0xf0, 0x28, 0x16, 0xa2, 0x47, 0x56, 0xe0, 0x44, 0x97, 0xec, 0x10, 0xa9, 0x28, 0x67, 0x85,
	0xe1, 0xbf, 0xbf, 0x13, 0x88, 0x85, 0xd8, 0x4e, 0xda, 0xbb, 0x6e, 0xca, 0xd9, 0x3b, 0x75, 0xfb,
	0xef, 0x12, 0x6e, 0xca, 0xd9, 0xcb, 0xdc, 0x7e, 0x93, 0xf6, 0xc8, 0xa5, 0xd2, 0x1e, 0xfd, 0x87,
	0xb4, 0xe7, 0xdf, 0x03, 0xf8, 0x7f, 0x7f, 0x52, 0x46, 0x01, 0xe6, 0x37, 0x2b, 0xcb, 0x8d, 0xfb,
	0xd5, 0xda, 0xe3, 0xc6, 0x93, 0xcd, 0xf5, 0xe6, 0x96, 0xfd, 0xc8, 0xae, 0xee, 0xd8, 0xba, 0x36,
	0x50, 0x59, 0xb6, 0xd7, 0x6a, 0xd5, 0x8d, 0x35, 0x1d, 0x18, 0x79, 0xa8, 0x9f, 0xab, 0x6c, 0x54,
	0xeb, 0xfa, 0xd0, 0x80, 0xba, 0xb3, 0xbe, 0xa2, 0xe7, 0x06, 0xd4, 0xd5, 0xca, 0xaa, 0x3e, 0x6c,
	0xcc, 0x40, 0xe3, 0x9c, 0xda, 0x58, 0xaf, 0x37, 0xea, 0xfa, 0xfc, 0xc2, 0x3b, 0x00, 0x61, 0xdf,
	0x96, 0xbc, 0x02, 0x70, 0xfa, 0xe2, 0xfd, 0xb9, 0x91, 0x3d, 0xbb, 0x3f, 0x6d, 0xd9, 0x6c, 0xfe,
	0x8c, 0x52, 0x82, 0x33, 0x45, 0xb6, 0x39, 0x75, 0x4a, 0x77, 0x5f, 0x7e, 0xf9, 0xf6, 0x76, 0xc8,
	0x2c, 0xdd, 0xb4, 0x3a, 0xb7, 0xad, 0xee, 0x8b, 0xb7, 0xce, 0x5a, 0xac, 0x0b, 0x8d, 0x16, 0xc1,
	0xfc, 0xca, 0x56, 0xbc, 0x34, 0x03, 0xf3, 0xfd, 0x7b, 0xa2, 0x88, 0xec, 0x02, 0xea, 0xe0, 0x08,
	0x69, 0x87, 0x47, 0x48, 0x3b, 0x39, 0x42, 0xe0, 0x45, 0x8c, 0xc0, 0x87, 0x18, 0x81, 0x4f, 0x31,
	0x02, 0x07, 0x31, 0x02, 0x5f, 0x63, 0x04, 0xbe, 0xc7, 0x48, 0x3b, 0x89, 0x11, 0x78, 0x73, 0x8c,
	0xb4, 0x83, 0x63, 0xa4, 0x1d, 0x1e, 0x23, 0x6d, 0x37, 0xfd, 0x4c, 0xb5, 0x46, 0x7a, 0xf9, 0xdd,
	0xf9, 0x19, 0x00, 0x00, 0xff, 0xff, 0xc0, 0xaa, 0x22, 0x50, 0xca, 0x04, 0x00, 0x00,
}

func (x PlatformType) String() string {
	s, ok := PlatformType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *RequestNotifyAboutDeviceInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RequestNotifyAboutDeviceInfo)
	if !ok {
		that2, ok := that.(RequestNotifyAboutDeviceInfo)
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
	if len(this.PreferredLanguages) != len(that1.PreferredLanguages) {
		return false
	}
	for i := range this.PreferredLanguages {
		if this.PreferredLanguages[i] != that1.PreferredLanguages[i] {
			return false
		}
	}
	if !this.TimeZone.Equal(that1.TimeZone) {
		return false
	}
	return true
}
func (this *ClientInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ClientInfo)
	if !ok {
		that2, ok := that.(ClientInfo)
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
	if this.Platform != that1.Platform {
		return false
	}
	if this.DeviceName != that1.DeviceName {
		return false
	}
	if this.AppName != that1.AppName {
		return false
	}
	if !this.AppVersion.Equal(that1.AppVersion) {
		return false
	}
	if !this.SdkVersion.Equal(that1.SdkVersion) {
		return false
	}
	if len(this.PreferredLanguages) != len(that1.PreferredLanguages) {
		return false
	}
	for i := range this.PreferredLanguages {
		if this.PreferredLanguages[i] != that1.PreferredLanguages[i] {
			return false
		}
	}
	if !this.TimeZone.Equal(that1.TimeZone) {
		return false
	}
	return true
}
func (this *RequestNotifyAboutDeviceInfo) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&dialog.RequestNotifyAboutDeviceInfo{")
	s = append(s, "PreferredLanguages: "+fmt.Sprintf("%#v", this.PreferredLanguages)+",\n")
	if this.TimeZone != nil {
		s = append(s, "TimeZone: "+fmt.Sprintf("%#v", this.TimeZone)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ClientInfo) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 11)
	s = append(s, "&dialog.ClientInfo{")
	s = append(s, "Platform: "+fmt.Sprintf("%#v", this.Platform)+",\n")
	s = append(s, "DeviceName: "+fmt.Sprintf("%#v", this.DeviceName)+",\n")
	s = append(s, "AppName: "+fmt.Sprintf("%#v", this.AppName)+",\n")
	if this.AppVersion != nil {
		s = append(s, "AppVersion: "+fmt.Sprintf("%#v", this.AppVersion)+",\n")
	}
	if this.SdkVersion != nil {
		s = append(s, "SdkVersion: "+fmt.Sprintf("%#v", this.SdkVersion)+",\n")
	}
	s = append(s, "PreferredLanguages: "+fmt.Sprintf("%#v", this.PreferredLanguages)+",\n")
	if this.TimeZone != nil {
		s = append(s, "TimeZone: "+fmt.Sprintf("%#v", this.TimeZone)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringDeviceInfo(v interface{}, typ string) string {
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

// DeviceInfoClient is the client API for DeviceInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeviceInfoClient interface {
	/// Set info about current device
	NotifyAboutDeviceInfo(ctx context.Context, in *RequestNotifyAboutDeviceInfo, opts ...grpc.CallOption) (*ResponseVoid, error)
}

type deviceInfoClient struct {
	cc *grpc.ClientConn
}

func NewDeviceInfoClient(cc *grpc.ClientConn) DeviceInfoClient {
	return &deviceInfoClient{cc}
}

func (c *deviceInfoClient) NotifyAboutDeviceInfo(ctx context.Context, in *RequestNotifyAboutDeviceInfo, opts ...grpc.CallOption) (*ResponseVoid, error) {
	out := new(ResponseVoid)
	err := c.cc.Invoke(ctx, "/dialog.DeviceInfo/NotifyAboutDeviceInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceInfoServer is the server API for DeviceInfo service.
type DeviceInfoServer interface {
	/// Set info about current device
	NotifyAboutDeviceInfo(context.Context, *RequestNotifyAboutDeviceInfo) (*ResponseVoid, error)
}

// UnimplementedDeviceInfoServer can be embedded to have forward compatible implementations.
type UnimplementedDeviceInfoServer struct {
}

func (*UnimplementedDeviceInfoServer) NotifyAboutDeviceInfo(ctx context.Context, req *RequestNotifyAboutDeviceInfo) (*ResponseVoid, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyAboutDeviceInfo not implemented")
}

func RegisterDeviceInfoServer(s *grpc.Server, srv DeviceInfoServer) {
	s.RegisterService(&_DeviceInfo_serviceDesc, srv)
}

func _DeviceInfo_NotifyAboutDeviceInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestNotifyAboutDeviceInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceInfoServer).NotifyAboutDeviceInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dialog.DeviceInfo/NotifyAboutDeviceInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceInfoServer).NotifyAboutDeviceInfo(ctx, req.(*RequestNotifyAboutDeviceInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeviceInfo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dialog.DeviceInfo",
	HandlerType: (*DeviceInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NotifyAboutDeviceInfo",
			Handler:    _DeviceInfo_NotifyAboutDeviceInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "device_info.proto",
}

func (m *RequestNotifyAboutDeviceInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestNotifyAboutDeviceInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestNotifyAboutDeviceInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TimeZone != nil {
		{
			size, err := m.TimeZone.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDeviceInfo(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.PreferredLanguages) > 0 {
		for iNdEx := len(m.PreferredLanguages) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PreferredLanguages[iNdEx])
			copy(dAtA[i:], m.PreferredLanguages[iNdEx])
			i = encodeVarintDeviceInfo(dAtA, i, uint64(len(m.PreferredLanguages[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ClientInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClientInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TimeZone != nil {
		{
			size, err := m.TimeZone.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDeviceInfo(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if len(m.PreferredLanguages) > 0 {
		for iNdEx := len(m.PreferredLanguages) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PreferredLanguages[iNdEx])
			copy(dAtA[i:], m.PreferredLanguages[iNdEx])
			i = encodeVarintDeviceInfo(dAtA, i, uint64(len(m.PreferredLanguages[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if m.SdkVersion != nil {
		{
			size, err := m.SdkVersion.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDeviceInfo(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.AppVersion != nil {
		{
			size, err := m.AppVersion.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDeviceInfo(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.AppName) > 0 {
		i -= len(m.AppName)
		copy(dAtA[i:], m.AppName)
		i = encodeVarintDeviceInfo(dAtA, i, uint64(len(m.AppName)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DeviceName) > 0 {
		i -= len(m.DeviceName)
		copy(dAtA[i:], m.DeviceName)
		i = encodeVarintDeviceInfo(dAtA, i, uint64(len(m.DeviceName)))
		i--
		dAtA[i] = 0x12
	}
	if m.Platform != 0 {
		i = encodeVarintDeviceInfo(dAtA, i, uint64(m.Platform))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintDeviceInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovDeviceInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RequestNotifyAboutDeviceInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PreferredLanguages) > 0 {
		for _, s := range m.PreferredLanguages {
			l = len(s)
			n += 1 + l + sovDeviceInfo(uint64(l))
		}
	}
	if m.TimeZone != nil {
		l = m.TimeZone.Size()
		n += 1 + l + sovDeviceInfo(uint64(l))
	}
	return n
}

func (m *ClientInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Platform != 0 {
		n += 1 + sovDeviceInfo(uint64(m.Platform))
	}
	l = len(m.DeviceName)
	if l > 0 {
		n += 1 + l + sovDeviceInfo(uint64(l))
	}
	l = len(m.AppName)
	if l > 0 {
		n += 1 + l + sovDeviceInfo(uint64(l))
	}
	if m.AppVersion != nil {
		l = m.AppVersion.Size()
		n += 1 + l + sovDeviceInfo(uint64(l))
	}
	if m.SdkVersion != nil {
		l = m.SdkVersion.Size()
		n += 1 + l + sovDeviceInfo(uint64(l))
	}
	if len(m.PreferredLanguages) > 0 {
		for _, s := range m.PreferredLanguages {
			l = len(s)
			n += 1 + l + sovDeviceInfo(uint64(l))
		}
	}
	if m.TimeZone != nil {
		l = m.TimeZone.Size()
		n += 1 + l + sovDeviceInfo(uint64(l))
	}
	return n
}

func sovDeviceInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDeviceInfo(x uint64) (n int) {
	return sovDeviceInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *RequestNotifyAboutDeviceInfo) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RequestNotifyAboutDeviceInfo{`,
		`PreferredLanguages:` + fmt.Sprintf("%v", this.PreferredLanguages) + `,`,
		`TimeZone:` + strings.Replace(fmt.Sprintf("%v", this.TimeZone), "StringValue", "types.StringValue", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ClientInfo) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ClientInfo{`,
		`Platform:` + fmt.Sprintf("%v", this.Platform) + `,`,
		`DeviceName:` + fmt.Sprintf("%v", this.DeviceName) + `,`,
		`AppName:` + fmt.Sprintf("%v", this.AppName) + `,`,
		`AppVersion:` + strings.Replace(fmt.Sprintf("%v", this.AppVersion), "StringValue", "types.StringValue", 1) + `,`,
		`SdkVersion:` + strings.Replace(fmt.Sprintf("%v", this.SdkVersion), "StringValue", "types.StringValue", 1) + `,`,
		`PreferredLanguages:` + fmt.Sprintf("%v", this.PreferredLanguages) + `,`,
		`TimeZone:` + strings.Replace(fmt.Sprintf("%v", this.TimeZone), "StringValue", "types.StringValue", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringDeviceInfo(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *RequestNotifyAboutDeviceInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeviceInfo
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
			return fmt.Errorf("proto: RequestNotifyAboutDeviceInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestNotifyAboutDeviceInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreferredLanguages", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceInfo
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
				return ErrInvalidLengthDeviceInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PreferredLanguages = append(m.PreferredLanguages, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeZone", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceInfo
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
				return ErrInvalidLengthDeviceInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TimeZone == nil {
				m.TimeZone = &types.StringValue{}
			}
			if err := m.TimeZone.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeviceInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeviceInfo
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDeviceInfo
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
func (m *ClientInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeviceInfo
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
			return fmt.Errorf("proto: ClientInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Platform", wireType)
			}
			m.Platform = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Platform |= PlatformType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceInfo
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
				return ErrInvalidLengthDeviceInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeviceName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceInfo
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
				return ErrInvalidLengthDeviceInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppVersion", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceInfo
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
				return ErrInvalidLengthDeviceInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AppVersion == nil {
				m.AppVersion = &types.StringValue{}
			}
			if err := m.AppVersion.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SdkVersion", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceInfo
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
				return ErrInvalidLengthDeviceInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SdkVersion == nil {
				m.SdkVersion = &types.StringValue{}
			}
			if err := m.SdkVersion.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreferredLanguages", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceInfo
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
				return ErrInvalidLengthDeviceInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PreferredLanguages = append(m.PreferredLanguages, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeZone", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceInfo
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
				return ErrInvalidLengthDeviceInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TimeZone == nil {
				m.TimeZone = &types.StringValue{}
			}
			if err := m.TimeZone.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeviceInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeviceInfo
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDeviceInfo
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
func skipDeviceInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDeviceInfo
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
					return 0, ErrIntOverflowDeviceInfo
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
					return 0, ErrIntOverflowDeviceInfo
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
				return 0, ErrInvalidLengthDeviceInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDeviceInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDeviceInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDeviceInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDeviceInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDeviceInfo = fmt.Errorf("proto: unexpected end of group")
)