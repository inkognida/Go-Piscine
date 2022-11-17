// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: armydevice.proto

package armydevice

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ArmyDeviceClient is the client API for ArmyDevice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArmyDeviceClient interface {
	GetArmyDeviceInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ArmyDevice_GetArmyDeviceInfoClient, error)
}

type armyDeviceClient struct {
	cc grpc.ClientConnInterface
}

func NewArmyDeviceClient(cc grpc.ClientConnInterface) ArmyDeviceClient {
	return &armyDeviceClient{cc}
}

func (c *armyDeviceClient) GetArmyDeviceInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ArmyDevice_GetArmyDeviceInfoClient, error) {
	stream, err := c.cc.NewStream(ctx, &ArmyDevice_ServiceDesc.Streams[0], "/armydevice.ArmyDevice/GetArmyDeviceInfo", opts...)
	if err != nil {
		return nil, err
	}
	x := &armyDeviceGetArmyDeviceInfoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ArmyDevice_GetArmyDeviceInfoClient interface {
	Recv() (*AddResponse, error)
	grpc.ClientStream
}

type armyDeviceGetArmyDeviceInfoClient struct {
	grpc.ClientStream
}

func (x *armyDeviceGetArmyDeviceInfoClient) Recv() (*AddResponse, error) {
	m := new(AddResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ArmyDeviceServer is the server API for ArmyDevice service.
// All implementations must embed UnimplementedArmyDeviceServer
// for forward compatibility
type ArmyDeviceServer interface {
	GetArmyDeviceInfo(*Empty, ArmyDevice_GetArmyDeviceInfoServer) error
	mustEmbedUnimplementedArmyDeviceServer()
}

// UnimplementedArmyDeviceServer must be embedded to have forward compatible implementations.
type UnimplementedArmyDeviceServer struct {
}

func (UnimplementedArmyDeviceServer) GetArmyDeviceInfo(*Empty, ArmyDevice_GetArmyDeviceInfoServer) error {
	return status.Errorf(codes.Unimplemented, "method GetArmyDeviceInfo not implemented")
}
func (UnimplementedArmyDeviceServer) mustEmbedUnimplementedArmyDeviceServer() {}

// UnsafeArmyDeviceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArmyDeviceServer will
// result in compilation errors.
type UnsafeArmyDeviceServer interface {
	mustEmbedUnimplementedArmyDeviceServer()
}

func RegisterArmyDeviceServer(s grpc.ServiceRegistrar, srv ArmyDeviceServer) {
	s.RegisterService(&ArmyDevice_ServiceDesc, srv)
}

func _ArmyDevice_GetArmyDeviceInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ArmyDeviceServer).GetArmyDeviceInfo(m, &armyDeviceGetArmyDeviceInfoServer{stream})
}

type ArmyDevice_GetArmyDeviceInfoServer interface {
	Send(*AddResponse) error
	grpc.ServerStream
}

type armyDeviceGetArmyDeviceInfoServer struct {
	grpc.ServerStream
}

func (x *armyDeviceGetArmyDeviceInfoServer) Send(m *AddResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ArmyDevice_ServiceDesc is the grpc.ServiceDesc for ArmyDevice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArmyDevice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "armydevice.ArmyDevice",
	HandlerType: (*ArmyDeviceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetArmyDeviceInfo",
			Handler:       _ArmyDevice_GetArmyDeviceInfo_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "armydevice.proto",
}
