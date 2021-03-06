// Code generated by protoc-gen-go.
// source: gameShipRpc.proto
// DO NOT EDIT!

/*
Package gameShipRpc is a generated protocol buffer package.

It is generated from these files:
	gameShipRpc.proto

It has these top-level messages:
	NewGameRequest
	NewGameResponse
	GameStartRequest
	GameStartResponse
	PlayerJoinRequest
	PlayerJoinResponse
	PlayerStrikeBoxRequest
	PlayerStrikeBoxResponse
	PlayerBingoRequest
	PlayerBingoResponse
*/
package gameShipRpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type NewGameRequest struct {
	CreatorId string `protobuf:"bytes,1,opt,name=creatorId" json:"creatorId,omitempty"`
}

func (m *NewGameRequest) Reset()                    { *m = NewGameRequest{} }
func (m *NewGameRequest) String() string            { return proto.CompactTextString(m) }
func (*NewGameRequest) ProtoMessage()               {}
func (*NewGameRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *NewGameRequest) GetCreatorId() string {
	if m != nil {
		return m.CreatorId
	}
	return ""
}

type NewGameResponse struct {
	GameId    string `protobuf:"bytes,1,opt,name=gameId" json:"gameId,omitempty"`
	CreatorId string `protobuf:"bytes,2,opt,name=creatorId" json:"creatorId,omitempty"`
}

func (m *NewGameResponse) Reset()                    { *m = NewGameResponse{} }
func (m *NewGameResponse) String() string            { return proto.CompactTextString(m) }
func (*NewGameResponse) ProtoMessage()               {}
func (*NewGameResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *NewGameResponse) GetGameId() string {
	if m != nil {
		return m.GameId
	}
	return ""
}

func (m *NewGameResponse) GetCreatorId() string {
	if m != nil {
		return m.CreatorId
	}
	return ""
}

type GameStartRequest struct {
	GameId string `protobuf:"bytes,1,opt,name=gameId" json:"gameId,omitempty"`
}

func (m *GameStartRequest) Reset()                    { *m = GameStartRequest{} }
func (m *GameStartRequest) String() string            { return proto.CompactTextString(m) }
func (*GameStartRequest) ProtoMessage()               {}
func (*GameStartRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GameStartRequest) GetGameId() string {
	if m != nil {
		return m.GameId
	}
	return ""
}

type GameStartResponse struct {
	GameId   string `protobuf:"bytes,1,opt,name=gameId" json:"gameId,omitempty"`
	PlayerId string `protobuf:"bytes,2,opt,name=playerId" json:"playerId,omitempty"`
}

func (m *GameStartResponse) Reset()                    { *m = GameStartResponse{} }
func (m *GameStartResponse) String() string            { return proto.CompactTextString(m) }
func (*GameStartResponse) ProtoMessage()               {}
func (*GameStartResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GameStartResponse) GetGameId() string {
	if m != nil {
		return m.GameId
	}
	return ""
}

func (m *GameStartResponse) GetPlayerId() string {
	if m != nil {
		return m.PlayerId
	}
	return ""
}

type PlayerJoinRequest struct {
	PlayerId string `protobuf:"bytes,1,opt,name=playerId" json:"playerId,omitempty"`
	GameId   string `protobuf:"bytes,2,opt,name=gameId" json:"gameId,omitempty"`
}

func (m *PlayerJoinRequest) Reset()                    { *m = PlayerJoinRequest{} }
func (m *PlayerJoinRequest) String() string            { return proto.CompactTextString(m) }
func (*PlayerJoinRequest) ProtoMessage()               {}
func (*PlayerJoinRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PlayerJoinRequest) GetPlayerId() string {
	if m != nil {
		return m.PlayerId
	}
	return ""
}

func (m *PlayerJoinRequest) GetGameId() string {
	if m != nil {
		return m.GameId
	}
	return ""
}

type PlayerJoinResponse struct {
	PlayerId string `protobuf:"bytes,1,opt,name=playerId" json:"playerId,omitempty"`
	GameId   string `protobuf:"bytes,2,opt,name=gameId" json:"gameId,omitempty"`
}

func (m *PlayerJoinResponse) Reset()                    { *m = PlayerJoinResponse{} }
func (m *PlayerJoinResponse) String() string            { return proto.CompactTextString(m) }
func (*PlayerJoinResponse) ProtoMessage()               {}
func (*PlayerJoinResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PlayerJoinResponse) GetPlayerId() string {
	if m != nil {
		return m.PlayerId
	}
	return ""
}

func (m *PlayerJoinResponse) GetGameId() string {
	if m != nil {
		return m.GameId
	}
	return ""
}

type PlayerStrikeBoxRequest struct {
	PlayerId string `protobuf:"bytes,1,opt,name=playerId" json:"playerId,omitempty"`
	GameId   string `protobuf:"bytes,2,opt,name=gameId" json:"gameId,omitempty"`
	Row      int32  `protobuf:"varint,3,opt,name=row" json:"row,omitempty"`
	Column   int32  `protobuf:"varint,4,opt,name=column" json:"column,omitempty"`
}

func (m *PlayerStrikeBoxRequest) Reset()                    { *m = PlayerStrikeBoxRequest{} }
func (m *PlayerStrikeBoxRequest) String() string            { return proto.CompactTextString(m) }
func (*PlayerStrikeBoxRequest) ProtoMessage()               {}
func (*PlayerStrikeBoxRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *PlayerStrikeBoxRequest) GetPlayerId() string {
	if m != nil {
		return m.PlayerId
	}
	return ""
}

func (m *PlayerStrikeBoxRequest) GetGameId() string {
	if m != nil {
		return m.GameId
	}
	return ""
}

func (m *PlayerStrikeBoxRequest) GetRow() int32 {
	if m != nil {
		return m.Row
	}
	return 0
}

func (m *PlayerStrikeBoxRequest) GetColumn() int32 {
	if m != nil {
		return m.Column
	}
	return 0
}

type PlayerStrikeBoxResponse struct {
	PlayerId string `protobuf:"bytes,1,opt,name=playerId" json:"playerId,omitempty"`
	GameId   string `protobuf:"bytes,2,opt,name=gameId" json:"gameId,omitempty"`
}

func (m *PlayerStrikeBoxResponse) Reset()                    { *m = PlayerStrikeBoxResponse{} }
func (m *PlayerStrikeBoxResponse) String() string            { return proto.CompactTextString(m) }
func (*PlayerStrikeBoxResponse) ProtoMessage()               {}
func (*PlayerStrikeBoxResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *PlayerStrikeBoxResponse) GetPlayerId() string {
	if m != nil {
		return m.PlayerId
	}
	return ""
}

func (m *PlayerStrikeBoxResponse) GetGameId() string {
	if m != nil {
		return m.GameId
	}
	return ""
}

type PlayerBingoRequest struct {
	PlayerId string `protobuf:"bytes,1,opt,name=playerId" json:"playerId,omitempty"`
	GameId   string `protobuf:"bytes,2,opt,name=gameId" json:"gameId,omitempty"`
}

func (m *PlayerBingoRequest) Reset()                    { *m = PlayerBingoRequest{} }
func (m *PlayerBingoRequest) String() string            { return proto.CompactTextString(m) }
func (*PlayerBingoRequest) ProtoMessage()               {}
func (*PlayerBingoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *PlayerBingoRequest) GetPlayerId() string {
	if m != nil {
		return m.PlayerId
	}
	return ""
}

func (m *PlayerBingoRequest) GetGameId() string {
	if m != nil {
		return m.GameId
	}
	return ""
}

type PlayerBingoResponse struct {
	PlayerId string `protobuf:"bytes,1,opt,name=playerId" json:"playerId,omitempty"`
	GameId   string `protobuf:"bytes,2,opt,name=gameId" json:"gameId,omitempty"`
	Position int32  `protobuf:"varint,3,opt,name=Position" json:"Position,omitempty"`
}

func (m *PlayerBingoResponse) Reset()                    { *m = PlayerBingoResponse{} }
func (m *PlayerBingoResponse) String() string            { return proto.CompactTextString(m) }
func (*PlayerBingoResponse) ProtoMessage()               {}
func (*PlayerBingoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *PlayerBingoResponse) GetPlayerId() string {
	if m != nil {
		return m.PlayerId
	}
	return ""
}

func (m *PlayerBingoResponse) GetGameId() string {
	if m != nil {
		return m.GameId
	}
	return ""
}

func (m *PlayerBingoResponse) GetPosition() int32 {
	if m != nil {
		return m.Position
	}
	return 0
}

func init() {
	proto.RegisterType((*NewGameRequest)(nil), "gameShipRpc.NewGameRequest")
	proto.RegisterType((*NewGameResponse)(nil), "gameShipRpc.NewGameResponse")
	proto.RegisterType((*GameStartRequest)(nil), "gameShipRpc.GameStartRequest")
	proto.RegisterType((*GameStartResponse)(nil), "gameShipRpc.GameStartResponse")
	proto.RegisterType((*PlayerJoinRequest)(nil), "gameShipRpc.PlayerJoinRequest")
	proto.RegisterType((*PlayerJoinResponse)(nil), "gameShipRpc.PlayerJoinResponse")
	proto.RegisterType((*PlayerStrikeBoxRequest)(nil), "gameShipRpc.PlayerStrikeBoxRequest")
	proto.RegisterType((*PlayerStrikeBoxResponse)(nil), "gameShipRpc.PlayerStrikeBoxResponse")
	proto.RegisterType((*PlayerBingoRequest)(nil), "gameShipRpc.PlayerBingoRequest")
	proto.RegisterType((*PlayerBingoResponse)(nil), "gameShipRpc.PlayerBingoResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GameShipRpc service

type GameShipRpcClient interface {
	NewGame(ctx context.Context, in *NewGameRequest, opts ...grpc.CallOption) (*NewGameResponse, error)
	GameStart(ctx context.Context, in *GameStartRequest, opts ...grpc.CallOption) (*GameStartResponse, error)
	PlayerJoin(ctx context.Context, in *PlayerJoinRequest, opts ...grpc.CallOption) (*PlayerJoinResponse, error)
	PlayerStrikeBox(ctx context.Context, in *PlayerStrikeBoxRequest, opts ...grpc.CallOption) (*PlayerStrikeBoxResponse, error)
	PlayerBingo(ctx context.Context, in *PlayerBingoRequest, opts ...grpc.CallOption) (*PlayerBingoResponse, error)
}

type gameShipRpcClient struct {
	cc *grpc.ClientConn
}

func NewGameShipRpcClient(cc *grpc.ClientConn) GameShipRpcClient {
	return &gameShipRpcClient{cc}
}

func (c *gameShipRpcClient) NewGame(ctx context.Context, in *NewGameRequest, opts ...grpc.CallOption) (*NewGameResponse, error) {
	out := new(NewGameResponse)
	err := grpc.Invoke(ctx, "/gameShipRpc.gameShipRpc/NewGame", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameShipRpcClient) GameStart(ctx context.Context, in *GameStartRequest, opts ...grpc.CallOption) (*GameStartResponse, error) {
	out := new(GameStartResponse)
	err := grpc.Invoke(ctx, "/gameShipRpc.gameShipRpc/GameStart", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameShipRpcClient) PlayerJoin(ctx context.Context, in *PlayerJoinRequest, opts ...grpc.CallOption) (*PlayerJoinResponse, error) {
	out := new(PlayerJoinResponse)
	err := grpc.Invoke(ctx, "/gameShipRpc.gameShipRpc/PlayerJoin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameShipRpcClient) PlayerStrikeBox(ctx context.Context, in *PlayerStrikeBoxRequest, opts ...grpc.CallOption) (*PlayerStrikeBoxResponse, error) {
	out := new(PlayerStrikeBoxResponse)
	err := grpc.Invoke(ctx, "/gameShipRpc.gameShipRpc/PlayerStrikeBox", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameShipRpcClient) PlayerBingo(ctx context.Context, in *PlayerBingoRequest, opts ...grpc.CallOption) (*PlayerBingoResponse, error) {
	out := new(PlayerBingoResponse)
	err := grpc.Invoke(ctx, "/gameShipRpc.gameShipRpc/PlayerBingo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GameShipRpc service

type GameShipRpcServer interface {
	NewGame(context.Context, *NewGameRequest) (*NewGameResponse, error)
	GameStart(context.Context, *GameStartRequest) (*GameStartResponse, error)
	PlayerJoin(context.Context, *PlayerJoinRequest) (*PlayerJoinResponse, error)
	PlayerStrikeBox(context.Context, *PlayerStrikeBoxRequest) (*PlayerStrikeBoxResponse, error)
	PlayerBingo(context.Context, *PlayerBingoRequest) (*PlayerBingoResponse, error)
}

func RegisterGameShipRpcServer(s *grpc.Server, srv GameShipRpcServer) {
	s.RegisterService(&_GameShipRpc_serviceDesc, srv)
}

func _GameShipRpc_NewGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameShipRpcServer).NewGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gameShipRpc.gameShipRpc/NewGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameShipRpcServer).NewGame(ctx, req.(*NewGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameShipRpc_GameStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GameStartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameShipRpcServer).GameStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gameShipRpc.gameShipRpc/GameStart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameShipRpcServer).GameStart(ctx, req.(*GameStartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameShipRpc_PlayerJoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerJoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameShipRpcServer).PlayerJoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gameShipRpc.gameShipRpc/PlayerJoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameShipRpcServer).PlayerJoin(ctx, req.(*PlayerJoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameShipRpc_PlayerStrikeBox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerStrikeBoxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameShipRpcServer).PlayerStrikeBox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gameShipRpc.gameShipRpc/PlayerStrikeBox",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameShipRpcServer).PlayerStrikeBox(ctx, req.(*PlayerStrikeBoxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameShipRpc_PlayerBingo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerBingoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameShipRpcServer).PlayerBingo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gameShipRpc.gameShipRpc/PlayerBingo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameShipRpcServer).PlayerBingo(ctx, req.(*PlayerBingoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GameShipRpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gameShipRpc.gameShipRpc",
	HandlerType: (*GameShipRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewGame",
			Handler:    _GameShipRpc_NewGame_Handler,
		},
		{
			MethodName: "GameStart",
			Handler:    _GameShipRpc_GameStart_Handler,
		},
		{
			MethodName: "PlayerJoin",
			Handler:    _GameShipRpc_PlayerJoin_Handler,
		},
		{
			MethodName: "PlayerStrikeBox",
			Handler:    _GameShipRpc_PlayerStrikeBox_Handler,
		},
		{
			MethodName: "PlayerBingo",
			Handler:    _GameShipRpc_PlayerBingo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gameShipRpc.proto",
}

func init() { proto.RegisterFile("gameShipRpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x94, 0xd1, 0x4e, 0xf2, 0x30,
	0x14, 0xc7, 0x33, 0xf8, 0x3e, 0x94, 0x43, 0x22, 0x50, 0x13, 0x24, 0x15, 0x95, 0x4c, 0x2f, 0x88,
	0x17, 0x5c, 0xe8, 0x1b, 0x10, 0x13, 0x94, 0x04, 0x43, 0xc6, 0xad, 0x37, 0x13, 0x1b, 0x6c, 0x84,
	0x75, 0x76, 0x45, 0xf4, 0x65, 0x7c, 0x56, 0xd3, 0x51, 0xba, 0x76, 0x6c, 0x9a, 0xb0, 0xbb, 0x9e,
	0xf6, 0x7f, 0x7e, 0xe7, 0xac, 0xfd, 0x9f, 0x41, 0x73, 0xee, 0x2f, 0xc9, 0xf4, 0x95, 0x86, 0x5e,
	0x38, 0xeb, 0x87, 0x9c, 0x09, 0x86, 0x6a, 0xc6, 0x96, 0xdb, 0x87, 0xa3, 0x47, 0xb2, 0x1e, 0xfa,
	0x4b, 0xe2, 0x91, 0xf7, 0x15, 0x89, 0x04, 0xea, 0x40, 0x75, 0xc6, 0x89, 0x2f, 0x18, 0x7f, 0x78,
	0x69, 0x3b, 0x5d, 0xa7, 0x57, 0xf5, 0x92, 0x0d, 0x77, 0x08, 0x75, 0xad, 0x8f, 0x42, 0x16, 0x44,
	0x04, 0xb5, 0xa0, 0x22, 0x89, 0x5a, 0xad, 0x22, 0x1b, 0x54, 0x4a, 0x83, 0xae, 0xa1, 0x21, 0x29,
	0x53, 0xe1, 0x73, 0xb1, 0x2d, 0x9d, 0x43, 0x72, 0x87, 0xd0, 0x34, 0xb4, 0x7f, 0x94, 0xc5, 0x70,
	0x18, 0x2e, 0xfc, 0x2f, 0x92, 0x54, 0xd5, 0xb1, 0x04, 0x4d, 0xe2, 0xf5, 0x88, 0xd1, 0x60, 0x5b,
	0xd5, 0x4c, 0x70, 0xec, 0x04, 0xa3, 0x48, 0xc9, 0xea, 0xe8, 0x1e, 0x90, 0x09, 0x52, 0x2d, 0xed,
	0x43, 0xfa, 0x80, 0xd6, 0x86, 0x34, 0x15, 0x9c, 0xbe, 0x91, 0x01, 0xfb, 0x2c, 0xd0, 0x17, 0x6a,
	0x40, 0x99, 0xb3, 0x75, 0xbb, 0xdc, 0x75, 0x7a, 0xff, 0x3d, 0xb9, 0x94, 0xca, 0x19, 0x5b, 0xac,
	0x96, 0x41, 0xfb, 0x5f, 0xbc, 0xa9, 0x22, 0x77, 0x0c, 0x27, 0x3b, 0x75, 0x0b, 0x7c, 0x86, 0xbe,
	0x90, 0x01, 0x0d, 0xe6, 0xac, 0xc8, 0xd5, 0x12, 0x38, 0xb6, 0x48, 0xfb, 0x37, 0x25, 0x73, 0x26,
	0x2c, 0xa2, 0x82, 0xb2, 0x40, 0x5d, 0x89, 0x8e, 0x6f, 0xbe, 0xcb, 0x60, 0x0e, 0x02, 0xba, 0x83,
	0x03, 0x65, 0x6c, 0x74, 0xda, 0x37, 0x87, 0xc6, 0x1e, 0x0f, 0xdc, 0xc9, 0x3e, 0x54, 0x5d, 0x8e,
	0xa0, 0xaa, 0x9d, 0x8a, 0xce, 0x2c, 0x69, 0xda, 0xed, 0xf8, 0x3c, 0xef, 0x58, 0xb1, 0xc6, 0x00,
	0x89, 0xc7, 0x90, 0xad, 0xde, 0x71, 0x31, 0xbe, 0xc8, 0x3d, 0x57, 0xb8, 0x27, 0xa8, 0xa7, 0x1e,
	0x1c, 0x5d, 0x66, 0xe4, 0xa4, 0x6d, 0x88, 0xaf, 0x7e, 0x17, 0x29, 0xfa, 0x04, 0x6a, 0xc6, 0xab,
	0xa1, 0xac, 0x6e, 0x4c, 0x67, 0xe0, 0x6e, 0xbe, 0x60, 0x43, 0x7c, 0xae, 0xc4, 0x7f, 0xab, 0xdb,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x97, 0xfd, 0xc7, 0xc3, 0xc2, 0x04, 0x00, 0x00,
}
