// Code generated by protoc-gen-go. DO NOT EDIT.
// source: werewolf.proto

package werewolf

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// user input some information to generate their uid
type Signature struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Displayname          string   `protobuf:"bytes,2,opt,name=displayname,proto3" json:"displayname,omitempty"`
	Randomcode           int32    `protobuf:"varint,3,opt,name=randomcode,proto3" json:"randomcode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Signature) Reset()         { *m = Signature{} }
func (m *Signature) String() string { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()    {}
func (*Signature) Descriptor() ([]byte, []int) {
	return fileDescriptor_b09fab9d3579aa92, []int{0}
}

func (m *Signature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Signature.Unmarshal(m, b)
}
func (m *Signature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Signature.Marshal(b, m, deterministic)
}
func (m *Signature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Signature.Merge(m, src)
}
func (m *Signature) XXX_Size() int {
	return xxx_messageInfo_Signature.Size(m)
}
func (m *Signature) XXX_DiscardUnknown() {
	xxx_messageInfo_Signature.DiscardUnknown(m)
}

var xxx_messageInfo_Signature proto.InternalMessageInfo

func (m *Signature) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Signature) GetDisplayname() string {
	if m != nil {
		return m.Displayname
	}
	return ""
}

func (m *Signature) GetRandomcode() int32 {
	if m != nil {
		return m.Randomcode
	}
	return 0
}

// Player is the message descripe a player
// it contains the display name as a string and an unique id
type Player struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Player) Reset()         { *m = Player{} }
func (m *Player) String() string { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()    {}
func (*Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_b09fab9d3579aa92, []int{1}
}

func (m *Player) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Player.Unmarshal(m, b)
}
func (m *Player) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Player.Marshal(b, m, deterministic)
}
func (m *Player) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Player.Merge(m, src)
}
func (m *Player) XXX_Size() int {
	return xxx_messageInfo_Player.Size(m)
}
func (m *Player) XXX_DiscardUnknown() {
	xxx_messageInfo_Player.DiscardUnknown(m)
}

var xxx_messageInfo_Player proto.InternalMessageInfo

func (m *Player) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Player) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

// the request contains the game command
type Request struct {
	Sender               *Player   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Action               string    `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	Receiver             []*Player `protobuf:"bytes,3,rep,name=receiver,proto3" json:"receiver,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_b09fab9d3579aa92, []int{2}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetSender() *Player {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *Request) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *Request) GetReceiver() []*Player {
	if m != nil {
		return m.Receiver
	}
	return nil
}

// the response of the game command
type Response struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Results              []string `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_b09fab9d3579aa92, []int{3}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Response) GetResults() []string {
	if m != nil {
		return m.Results
	}
	return nil
}

func init() {
	proto.RegisterType((*Signature)(nil), "werewolf.Signature")
	proto.RegisterType((*Player)(nil), "werewolf.Player")
	proto.RegisterType((*Request)(nil), "werewolf.Request")
	proto.RegisterType((*Response)(nil), "werewolf.Response")
}

func init() { proto.RegisterFile("werewolf.proto", fileDescriptor_b09fab9d3579aa92) }

var fileDescriptor_b09fab9d3579aa92 = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x4b, 0x33, 0x31,
	0x10, 0xc5, 0xbb, 0xdd, 0x7e, 0xdb, 0xed, 0x14, 0x3e, 0xea, 0x08, 0xb2, 0xf4, 0x20, 0x4b, 0x4e,
	0x7b, 0x90, 0xa2, 0xed, 0xd5, 0x9b, 0x82, 0x78, 0x93, 0x78, 0xf0, 0x1c, 0xbb, 0xa3, 0x04, 0xb6,
	0x49, 0x9d, 0x24, 0x96, 0xfe, 0xf7, 0xd2, 0x6d, 0xba, 0x5d, 0xd0, 0xdb, 0xbc, 0x79, 0xc9, 0x6f,
	0x26, 0x2f, 0xf0, 0x7f, 0x47, 0x4c, 0x3b, 0xdb, 0x7c, 0x2c, 0xb6, 0x6c, 0xbd, 0xc5, 0xfc, 0xa4,
	0x85, 0x86, 0xc9, 0xab, 0xfe, 0x34, 0xca, 0x07, 0x26, 0x9c, 0x43, 0x1e, 0x1c, 0xb1, 0x51, 0x1b,
	0x2a, 0x92, 0x32, 0xa9, 0x26, 0xb2, 0xd3, 0x58, 0xc2, 0xb4, 0xd6, 0x6e, 0xdb, 0xa8, 0x7d, 0x6b,
	0x0f, 0x5b, 0xbb, 0xdf, 0xc2, 0x6b, 0x00, 0x56, 0xa6, 0xb6, 0x9b, 0xb5, 0xad, 0xa9, 0x48, 0xcb,
	0xa4, 0xfa, 0x27, 0x7b, 0x1d, 0xb1, 0x80, 0xec, 0xa5, 0x51, 0x7b, 0x62, 0x44, 0x18, 0xf5, 0x66,
	0xb4, 0x35, 0xce, 0x20, 0x0d, 0xba, 0x8e, 0xdc, 0x43, 0x29, 0xf6, 0x30, 0x96, 0xf4, 0x15, 0xc8,
	0x79, 0xac, 0x20, 0x73, 0x64, 0x6a, 0xe2, 0xf6, 0xca, 0x74, 0x39, 0x5b, 0x74, 0x0f, 0x3a, 0x22,
	0x65, 0xf4, 0xf1, 0x0a, 0x32, 0xb5, 0xf6, 0xda, 0x9a, 0x48, 0x8a, 0x0a, 0x6f, 0x20, 0x67, 0x5a,
	0x93, 0xfe, 0x26, 0x2e, 0xd2, 0x32, 0xfd, 0x93, 0xd1, 0x9d, 0x10, 0xf7, 0x90, 0x4b, 0x72, 0x5b,
	0x6b, 0x1c, 0x1d, 0x88, 0xce, 0x2b, 0x1f, 0x5c, 0x5c, 0x37, 0x2a, 0x2c, 0x60, 0xcc, 0xe4, 0x42,
	0xe3, 0x5d, 0x31, 0x2c, 0xd3, 0x6a, 0x22, 0x4f, 0x72, 0xc9, 0x90, 0xbf, 0x45, 0x34, 0xae, 0x20,
	0x7f, 0x60, 0x52, 0x9e, 0x9e, 0x1f, 0xf1, 0xf2, 0x3c, 0xb1, 0xcb, 0x7c, 0xfe, 0x6b, 0x0d, 0x31,
	0xc0, 0x3b, 0x18, 0x3d, 0x1d, 0x32, 0xb9, 0x38, 0x7b, 0x31, 0x89, 0x39, 0xf6, 0x5b, 0xc7, 0x0d,
	0xc5, 0xe0, 0x36, 0x79, 0xcf, 0xda, 0x8f, 0x5d, 0xfd, 0x04, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x78,
	0xfa, 0xbe, 0xea, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WerewolfClient is the client API for Werewolf service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WerewolfClient interface {
	// Use some user input information, to generated a unique ID to identify user
	// in the system. Uid can be used to reconnect to the game.
	CreateID(ctx context.Context, in *Signature, opts ...grpc.CallOption) (*Player, error)
	// Game service serve the communication between client and server for game
	// commands. Using server side streaming to ensure the game process flow
	// smoothly.
	Game(ctx context.Context, in *Request, opts ...grpc.CallOption) (Werewolf_GameClient, error)
}

type werewolfClient struct {
	cc *grpc.ClientConn
}

func NewWerewolfClient(cc *grpc.ClientConn) WerewolfClient {
	return &werewolfClient{cc}
}

func (c *werewolfClient) CreateID(ctx context.Context, in *Signature, opts ...grpc.CallOption) (*Player, error) {
	out := new(Player)
	err := c.cc.Invoke(ctx, "/werewolf.Werewolf/CreateID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *werewolfClient) Game(ctx context.Context, in *Request, opts ...grpc.CallOption) (Werewolf_GameClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Werewolf_serviceDesc.Streams[0], "/werewolf.Werewolf/Game", opts...)
	if err != nil {
		return nil, err
	}
	x := &werewolfGameClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Werewolf_GameClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type werewolfGameClient struct {
	grpc.ClientStream
}

func (x *werewolfGameClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WerewolfServer is the server API for Werewolf service.
type WerewolfServer interface {
	// Use some user input information, to generated a unique ID to identify user
	// in the system. Uid can be used to reconnect to the game.
	CreateID(context.Context, *Signature) (*Player, error)
	// Game service serve the communication between client and server for game
	// commands. Using server side streaming to ensure the game process flow
	// smoothly.
	Game(*Request, Werewolf_GameServer) error
}

// UnimplementedWerewolfServer can be embedded to have forward compatible implementations.
type UnimplementedWerewolfServer struct {
}

func (*UnimplementedWerewolfServer) CreateID(ctx context.Context, req *Signature) (*Player, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateID not implemented")
}
func (*UnimplementedWerewolfServer) Game(req *Request, srv Werewolf_GameServer) error {
	return status.Errorf(codes.Unimplemented, "method Game not implemented")
}

func RegisterWerewolfServer(s *grpc.Server, srv WerewolfServer) {
	s.RegisterService(&_Werewolf_serviceDesc, srv)
}

func _Werewolf_CreateID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Signature)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WerewolfServer).CreateID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/werewolf.Werewolf/CreateID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WerewolfServer).CreateID(ctx, req.(*Signature))
	}
	return interceptor(ctx, in, info, handler)
}

func _Werewolf_Game_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WerewolfServer).Game(m, &werewolfGameServer{stream})
}

type Werewolf_GameServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type werewolfGameServer struct {
	grpc.ServerStream
}

func (x *werewolfGameServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

var _Werewolf_serviceDesc = grpc.ServiceDesc{
	ServiceName: "werewolf.Werewolf",
	HandlerType: (*WerewolfServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateID",
			Handler:    _Werewolf_CreateID_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Game",
			Handler:       _Werewolf_Game_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "werewolf.proto",
}
