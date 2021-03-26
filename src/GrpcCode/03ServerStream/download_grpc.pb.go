// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DownloadClient is the client API for Download service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DownloadClient interface {
	Download(ctx context.Context, in *DownloadReq, opts ...grpc.CallOption) (Download_DownloadClient, error)
}

type downloadClient struct {
	cc grpc.ClientConnInterface
}

func NewDownloadClient(cc grpc.ClientConnInterface) DownloadClient {
	return &downloadClient{cc}
}

func (c *downloadClient) Download(ctx context.Context, in *DownloadReq, opts ...grpc.CallOption) (Download_DownloadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Download_serviceDesc.Streams[0], "/Download/Download", opts...)
	if err != nil {
		return nil, err
	}
	x := &downloadDownloadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Download_DownloadClient interface {
	Recv() (*DownloadRes, error)
	grpc.ClientStream
}

type downloadDownloadClient struct {
	grpc.ClientStream
}

func (x *downloadDownloadClient) Recv() (*DownloadRes, error) {
	m := new(DownloadRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DownloadServer is the server API for Download service.
// All implementations should embed UnimplementedDownloadServer
// for forward compatibility
type DownloadServer interface {
	Download(*DownloadReq, Download_DownloadServer) error
}

// UnimplementedDownloadServer should be embedded to have forward compatible implementations.
type UnimplementedDownloadServer struct {
}

func (*UnimplementedDownloadServer) Download(*DownloadReq, Download_DownloadServer) error {
	return status.Errorf(codes.Unimplemented, "method Download not implemented")
}

func RegisterDownloadServer(s *grpc.Server, srv DownloadServer) {
	s.RegisterService(&_Download_serviceDesc, srv)
}

func _Download_Download_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DownloadServer).Download(m, &downloadDownloadServer{stream})
}

type Download_DownloadServer interface {
	Send(*DownloadRes) error
	grpc.ServerStream
}

type downloadDownloadServer struct {
	grpc.ServerStream
}

func (x *downloadDownloadServer) Send(m *DownloadRes) error {
	return x.ServerStream.SendMsg(m)
}

var _Download_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Download",
	HandlerType: (*DownloadServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Download",
			Handler:       _Download_Download_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "download.proto",
}