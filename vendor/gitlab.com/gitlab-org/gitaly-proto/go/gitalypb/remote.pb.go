// Code generated by protoc-gen-go. DO NOT EDIT.
// source: remote.proto

package gitalypb

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

type AddRemoteRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	Name       string      `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Url        string      `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
	// If any, the remote is configured as a mirror with those mappings
	MirrorRefmaps []string `protobuf:"bytes,5,rep,name=mirror_refmaps,json=mirrorRefmaps" json:"mirror_refmaps,omitempty"`
}

func (m *AddRemoteRequest) Reset()                    { *m = AddRemoteRequest{} }
func (m *AddRemoteRequest) String() string            { return proto.CompactTextString(m) }
func (*AddRemoteRequest) ProtoMessage()               {}
func (*AddRemoteRequest) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *AddRemoteRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *AddRemoteRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddRemoteRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *AddRemoteRequest) GetMirrorRefmaps() []string {
	if m != nil {
		return m.MirrorRefmaps
	}
	return nil
}

type AddRemoteResponse struct {
}

func (m *AddRemoteResponse) Reset()                    { *m = AddRemoteResponse{} }
func (m *AddRemoteResponse) String() string            { return proto.CompactTextString(m) }
func (*AddRemoteResponse) ProtoMessage()               {}
func (*AddRemoteResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

type RemoveRemoteRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	Name       string      `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *RemoveRemoteRequest) Reset()                    { *m = RemoveRemoteRequest{} }
func (m *RemoveRemoteRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveRemoteRequest) ProtoMessage()               {}
func (*RemoveRemoteRequest) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{2} }

func (m *RemoveRemoteRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *RemoveRemoteRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RemoveRemoteResponse struct {
	Result bool `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
}

func (m *RemoveRemoteResponse) Reset()                    { *m = RemoveRemoteResponse{} }
func (m *RemoveRemoteResponse) String() string            { return proto.CompactTextString(m) }
func (*RemoveRemoteResponse) ProtoMessage()               {}
func (*RemoveRemoteResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{3} }

func (m *RemoveRemoteResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type FetchInternalRemoteRequest struct {
	Repository       *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	RemoteRepository *Repository `protobuf:"bytes,2,opt,name=remote_repository,json=remoteRepository" json:"remote_repository,omitempty"`
}

func (m *FetchInternalRemoteRequest) Reset()                    { *m = FetchInternalRemoteRequest{} }
func (m *FetchInternalRemoteRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchInternalRemoteRequest) ProtoMessage()               {}
func (*FetchInternalRemoteRequest) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{4} }

func (m *FetchInternalRemoteRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *FetchInternalRemoteRequest) GetRemoteRepository() *Repository {
	if m != nil {
		return m.RemoteRepository
	}
	return nil
}

type FetchInternalRemoteResponse struct {
	Result bool `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
}

func (m *FetchInternalRemoteResponse) Reset()                    { *m = FetchInternalRemoteResponse{} }
func (m *FetchInternalRemoteResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchInternalRemoteResponse) ProtoMessage()               {}
func (*FetchInternalRemoteResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{5} }

func (m *FetchInternalRemoteResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type UpdateRemoteMirrorRequest struct {
	Repository           *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	RefName              string      `protobuf:"bytes,2,opt,name=ref_name,json=refName" json:"ref_name,omitempty"`
	OnlyBranchesMatching [][]byte    `protobuf:"bytes,3,rep,name=only_branches_matching,json=onlyBranchesMatching,proto3" json:"only_branches_matching,omitempty"`
}

func (m *UpdateRemoteMirrorRequest) Reset()                    { *m = UpdateRemoteMirrorRequest{} }
func (m *UpdateRemoteMirrorRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRemoteMirrorRequest) ProtoMessage()               {}
func (*UpdateRemoteMirrorRequest) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{6} }

func (m *UpdateRemoteMirrorRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *UpdateRemoteMirrorRequest) GetRefName() string {
	if m != nil {
		return m.RefName
	}
	return ""
}

func (m *UpdateRemoteMirrorRequest) GetOnlyBranchesMatching() [][]byte {
	if m != nil {
		return m.OnlyBranchesMatching
	}
	return nil
}

type UpdateRemoteMirrorResponse struct {
}

func (m *UpdateRemoteMirrorResponse) Reset()                    { *m = UpdateRemoteMirrorResponse{} }
func (m *UpdateRemoteMirrorResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateRemoteMirrorResponse) ProtoMessage()               {}
func (*UpdateRemoteMirrorResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{7} }

type FindRemoteRepositoryRequest struct {
	Remote string `protobuf:"bytes,1,opt,name=remote" json:"remote,omitempty"`
}

func (m *FindRemoteRepositoryRequest) Reset()                    { *m = FindRemoteRepositoryRequest{} }
func (m *FindRemoteRepositoryRequest) String() string            { return proto.CompactTextString(m) }
func (*FindRemoteRepositoryRequest) ProtoMessage()               {}
func (*FindRemoteRepositoryRequest) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{8} }

func (m *FindRemoteRepositoryRequest) GetRemote() string {
	if m != nil {
		return m.Remote
	}
	return ""
}

// This migth throw a GRPC Unavailable code, to signal the request failure
// is transient.
type FindRemoteRepositoryResponse struct {
	Exists bool `protobuf:"varint,1,opt,name=exists" json:"exists,omitempty"`
}

func (m *FindRemoteRepositoryResponse) Reset()                    { *m = FindRemoteRepositoryResponse{} }
func (m *FindRemoteRepositoryResponse) String() string            { return proto.CompactTextString(m) }
func (*FindRemoteRepositoryResponse) ProtoMessage()               {}
func (*FindRemoteRepositoryResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{9} }

func (m *FindRemoteRepositoryResponse) GetExists() bool {
	if m != nil {
		return m.Exists
	}
	return false
}

type FindRemoteRootRefRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	Remote     string      `protobuf:"bytes,2,opt,name=remote" json:"remote,omitempty"`
}

func (m *FindRemoteRootRefRequest) Reset()                    { *m = FindRemoteRootRefRequest{} }
func (m *FindRemoteRootRefRequest) String() string            { return proto.CompactTextString(m) }
func (*FindRemoteRootRefRequest) ProtoMessage()               {}
func (*FindRemoteRootRefRequest) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{10} }

func (m *FindRemoteRootRefRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *FindRemoteRootRefRequest) GetRemote() string {
	if m != nil {
		return m.Remote
	}
	return ""
}

type FindRemoteRootRefResponse struct {
	Ref string `protobuf:"bytes,1,opt,name=ref" json:"ref,omitempty"`
}

func (m *FindRemoteRootRefResponse) Reset()                    { *m = FindRemoteRootRefResponse{} }
func (m *FindRemoteRootRefResponse) String() string            { return proto.CompactTextString(m) }
func (*FindRemoteRootRefResponse) ProtoMessage()               {}
func (*FindRemoteRootRefResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{11} }

func (m *FindRemoteRootRefResponse) GetRef() string {
	if m != nil {
		return m.Ref
	}
	return ""
}

func init() {
	proto.RegisterType((*AddRemoteRequest)(nil), "gitaly.AddRemoteRequest")
	proto.RegisterType((*AddRemoteResponse)(nil), "gitaly.AddRemoteResponse")
	proto.RegisterType((*RemoveRemoteRequest)(nil), "gitaly.RemoveRemoteRequest")
	proto.RegisterType((*RemoveRemoteResponse)(nil), "gitaly.RemoveRemoteResponse")
	proto.RegisterType((*FetchInternalRemoteRequest)(nil), "gitaly.FetchInternalRemoteRequest")
	proto.RegisterType((*FetchInternalRemoteResponse)(nil), "gitaly.FetchInternalRemoteResponse")
	proto.RegisterType((*UpdateRemoteMirrorRequest)(nil), "gitaly.UpdateRemoteMirrorRequest")
	proto.RegisterType((*UpdateRemoteMirrorResponse)(nil), "gitaly.UpdateRemoteMirrorResponse")
	proto.RegisterType((*FindRemoteRepositoryRequest)(nil), "gitaly.FindRemoteRepositoryRequest")
	proto.RegisterType((*FindRemoteRepositoryResponse)(nil), "gitaly.FindRemoteRepositoryResponse")
	proto.RegisterType((*FindRemoteRootRefRequest)(nil), "gitaly.FindRemoteRootRefRequest")
	proto.RegisterType((*FindRemoteRootRefResponse)(nil), "gitaly.FindRemoteRootRefResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RemoteService service

type RemoteServiceClient interface {
	AddRemote(ctx context.Context, in *AddRemoteRequest, opts ...grpc.CallOption) (*AddRemoteResponse, error)
	FetchInternalRemote(ctx context.Context, in *FetchInternalRemoteRequest, opts ...grpc.CallOption) (*FetchInternalRemoteResponse, error)
	RemoveRemote(ctx context.Context, in *RemoveRemoteRequest, opts ...grpc.CallOption) (*RemoveRemoteResponse, error)
	UpdateRemoteMirror(ctx context.Context, opts ...grpc.CallOption) (RemoteService_UpdateRemoteMirrorClient, error)
	FindRemoteRepository(ctx context.Context, in *FindRemoteRepositoryRequest, opts ...grpc.CallOption) (*FindRemoteRepositoryResponse, error)
	FindRemoteRootRef(ctx context.Context, in *FindRemoteRootRefRequest, opts ...grpc.CallOption) (*FindRemoteRootRefResponse, error)
}

type remoteServiceClient struct {
	cc *grpc.ClientConn
}

func NewRemoteServiceClient(cc *grpc.ClientConn) RemoteServiceClient {
	return &remoteServiceClient{cc}
}

func (c *remoteServiceClient) AddRemote(ctx context.Context, in *AddRemoteRequest, opts ...grpc.CallOption) (*AddRemoteResponse, error) {
	out := new(AddRemoteResponse)
	err := grpc.Invoke(ctx, "/gitaly.RemoteService/AddRemote", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteServiceClient) FetchInternalRemote(ctx context.Context, in *FetchInternalRemoteRequest, opts ...grpc.CallOption) (*FetchInternalRemoteResponse, error) {
	out := new(FetchInternalRemoteResponse)
	err := grpc.Invoke(ctx, "/gitaly.RemoteService/FetchInternalRemote", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteServiceClient) RemoveRemote(ctx context.Context, in *RemoveRemoteRequest, opts ...grpc.CallOption) (*RemoveRemoteResponse, error) {
	out := new(RemoveRemoteResponse)
	err := grpc.Invoke(ctx, "/gitaly.RemoteService/RemoveRemote", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteServiceClient) UpdateRemoteMirror(ctx context.Context, opts ...grpc.CallOption) (RemoteService_UpdateRemoteMirrorClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RemoteService_serviceDesc.Streams[0], c.cc, "/gitaly.RemoteService/UpdateRemoteMirror", opts...)
	if err != nil {
		return nil, err
	}
	x := &remoteServiceUpdateRemoteMirrorClient{stream}
	return x, nil
}

type RemoteService_UpdateRemoteMirrorClient interface {
	Send(*UpdateRemoteMirrorRequest) error
	CloseAndRecv() (*UpdateRemoteMirrorResponse, error)
	grpc.ClientStream
}

type remoteServiceUpdateRemoteMirrorClient struct {
	grpc.ClientStream
}

func (x *remoteServiceUpdateRemoteMirrorClient) Send(m *UpdateRemoteMirrorRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *remoteServiceUpdateRemoteMirrorClient) CloseAndRecv() (*UpdateRemoteMirrorResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UpdateRemoteMirrorResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *remoteServiceClient) FindRemoteRepository(ctx context.Context, in *FindRemoteRepositoryRequest, opts ...grpc.CallOption) (*FindRemoteRepositoryResponse, error) {
	out := new(FindRemoteRepositoryResponse)
	err := grpc.Invoke(ctx, "/gitaly.RemoteService/FindRemoteRepository", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteServiceClient) FindRemoteRootRef(ctx context.Context, in *FindRemoteRootRefRequest, opts ...grpc.CallOption) (*FindRemoteRootRefResponse, error) {
	out := new(FindRemoteRootRefResponse)
	err := grpc.Invoke(ctx, "/gitaly.RemoteService/FindRemoteRootRef", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RemoteService service

type RemoteServiceServer interface {
	AddRemote(context.Context, *AddRemoteRequest) (*AddRemoteResponse, error)
	FetchInternalRemote(context.Context, *FetchInternalRemoteRequest) (*FetchInternalRemoteResponse, error)
	RemoveRemote(context.Context, *RemoveRemoteRequest) (*RemoveRemoteResponse, error)
	UpdateRemoteMirror(RemoteService_UpdateRemoteMirrorServer) error
	FindRemoteRepository(context.Context, *FindRemoteRepositoryRequest) (*FindRemoteRepositoryResponse, error)
	FindRemoteRootRef(context.Context, *FindRemoteRootRefRequest) (*FindRemoteRootRefResponse, error)
}

func RegisterRemoteServiceServer(s *grpc.Server, srv RemoteServiceServer) {
	s.RegisterService(&_RemoteService_serviceDesc, srv)
}

func _RemoteService_AddRemote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRemoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteServiceServer).AddRemote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.RemoteService/AddRemote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteServiceServer).AddRemote(ctx, req.(*AddRemoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteService_FetchInternalRemote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchInternalRemoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteServiceServer).FetchInternalRemote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.RemoteService/FetchInternalRemote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteServiceServer).FetchInternalRemote(ctx, req.(*FetchInternalRemoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteService_RemoveRemote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRemoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteServiceServer).RemoveRemote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.RemoteService/RemoveRemote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteServiceServer).RemoveRemote(ctx, req.(*RemoveRemoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteService_UpdateRemoteMirror_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RemoteServiceServer).UpdateRemoteMirror(&remoteServiceUpdateRemoteMirrorServer{stream})
}

type RemoteService_UpdateRemoteMirrorServer interface {
	SendAndClose(*UpdateRemoteMirrorResponse) error
	Recv() (*UpdateRemoteMirrorRequest, error)
	grpc.ServerStream
}

type remoteServiceUpdateRemoteMirrorServer struct {
	grpc.ServerStream
}

func (x *remoteServiceUpdateRemoteMirrorServer) SendAndClose(m *UpdateRemoteMirrorResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *remoteServiceUpdateRemoteMirrorServer) Recv() (*UpdateRemoteMirrorRequest, error) {
	m := new(UpdateRemoteMirrorRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RemoteService_FindRemoteRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRemoteRepositoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteServiceServer).FindRemoteRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.RemoteService/FindRemoteRepository",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteServiceServer).FindRemoteRepository(ctx, req.(*FindRemoteRepositoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteService_FindRemoteRootRef_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRemoteRootRefRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteServiceServer).FindRemoteRootRef(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.RemoteService/FindRemoteRootRef",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteServiceServer).FindRemoteRootRef(ctx, req.(*FindRemoteRootRefRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RemoteService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gitaly.RemoteService",
	HandlerType: (*RemoteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddRemote",
			Handler:    _RemoteService_AddRemote_Handler,
		},
		{
			MethodName: "FetchInternalRemote",
			Handler:    _RemoteService_FetchInternalRemote_Handler,
		},
		{
			MethodName: "RemoveRemote",
			Handler:    _RemoteService_RemoveRemote_Handler,
		},
		{
			MethodName: "FindRemoteRepository",
			Handler:    _RemoteService_FindRemoteRepository_Handler,
		},
		{
			MethodName: "FindRemoteRootRef",
			Handler:    _RemoteService_FindRemoteRootRef_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UpdateRemoteMirror",
			Handler:       _RemoteService_UpdateRemoteMirror_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "remote.proto",
}

func init() { proto.RegisterFile("remote.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 536 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x4f, 0x8f, 0xd3, 0x3e,
	0x10, 0xfd, 0xa5, 0xe9, 0xf6, 0xd7, 0x0e, 0x5d, 0xd4, 0xba, 0xd5, 0x2a, 0xcd, 0xf6, 0xd0, 0x35,
	0x20, 0xe5, 0x42, 0x0f, 0xe5, 0xcf, 0x15, 0xb1, 0x07, 0x24, 0x40, 0xcb, 0xc1, 0x88, 0x0b, 0x12,
	0x0a, 0xd9, 0x74, 0xb2, 0x8d, 0x94, 0xc4, 0xc1, 0x76, 0x57, 0xf4, 0x63, 0xf0, 0x0d, 0x38, 0x70,
	0xe0, 0x63, 0xa2, 0x24, 0x4e, 0x9a, 0xd2, 0xb4, 0x48, 0xac, 0xb8, 0xd9, 0x33, 0xf3, 0xc6, 0xef,
	0x8d, 0x9f, 0x13, 0xe8, 0x0b, 0x8c, 0xb9, 0xc2, 0x79, 0x2a, 0xb8, 0xe2, 0xa4, 0x73, 0x13, 0x2a,
	0x2f, 0xda, 0xd8, 0x7d, 0xb9, 0xf2, 0x04, 0x2e, 0x8b, 0x28, 0xfd, 0x69, 0xc0, 0xe0, 0xe5, 0x72,
	0xc9, 0xf2, 0x4a, 0x86, 0x5f, 0xd6, 0x28, 0x15, 0x59, 0x00, 0x08, 0x4c, 0xb9, 0x0c, 0x15, 0x17,
	0x1b, 0xcb, 0x98, 0x19, 0xce, 0xbd, 0x05, 0x99, 0x17, 0xf8, 0x39, 0xab, 0x32, 0xac, 0x56, 0x45,
	0x08, 0xb4, 0x13, 0x2f, 0x46, 0xab, 0x35, 0x33, 0x9c, 0x1e, 0xcb, 0xd7, 0x64, 0x00, 0xe6, 0x5a,
	0x44, 0x96, 0x99, 0x87, 0xb2, 0x25, 0x79, 0x04, 0xf7, 0xe3, 0x50, 0x08, 0x2e, 0x5c, 0x81, 0x41,
	0xec, 0xa5, 0xd2, 0x3a, 0x99, 0x99, 0x4e, 0x8f, 0x9d, 0x16, 0x51, 0x56, 0x04, 0xdf, 0xb4, 0xbb,
	0xed, 0xc1, 0x49, 0x19, 0xd4, 0xa5, 0x74, 0x04, 0xc3, 0x1a, 0x53, 0x99, 0xf2, 0x44, 0x22, 0xfd,
	0x04, 0xa3, 0x2c, 0x72, 0x8b, 0xff, 0x44, 0x01, 0x9d, 0xc3, 0x78, 0xb7, 0x7d, 0x71, 0x2c, 0x39,
	0x83, 0x8e, 0x40, 0xb9, 0x8e, 0x54, 0xde, 0xbb, 0xcb, 0xf4, 0x8e, 0x7e, 0x33, 0xc0, 0x7e, 0x85,
	0xca, 0x5f, 0xbd, 0x4e, 0x14, 0x8a, 0xc4, 0x8b, 0xee, 0x4e, 0xeb, 0x05, 0x0c, 0x8b, 0x7b, 0x74,
	0x6b, 0xd0, 0xd6, 0x41, 0xe8, 0x40, 0xe8, 0x13, 0xcb, 0x08, 0x7d, 0x06, 0xe7, 0x8d, 0x94, 0xfe,
	0x20, 0xe5, 0xbb, 0x01, 0x93, 0x0f, 0xe9, 0xd2, 0x53, 0x5a, 0xfb, 0x95, 0xbe, 0xa1, 0xbf, 0x57,
	0x32, 0x81, 0xae, 0xc0, 0xc0, 0xad, 0x0d, 0xf9, 0x7f, 0x81, 0xc1, 0xbb, 0xcc, 0x29, 0x4f, 0xe1,
	0x8c, 0x27, 0xd1, 0xc6, 0xbd, 0x16, 0x5e, 0xe2, 0xaf, 0x50, 0xba, 0xb1, 0xa7, 0xfc, 0x55, 0x98,
	0xdc, 0x58, 0xe6, 0xcc, 0x74, 0xfa, 0x6c, 0x9c, 0x65, 0x2f, 0x75, 0xf2, 0x4a, 0xe7, 0xe8, 0x14,
	0xec, 0x26, 0x86, 0xda, 0x1a, 0x99, 0xee, 0x30, 0xa9, 0x0c, 0x53, 0x51, 0xd2, 0x0a, 0x72, 0xdd,
	0x59, 0x2a, 0x67, 0xdf, 0x63, 0x7a, 0x47, 0x9f, 0xc3, 0xb4, 0x19, 0xb6, 0x9d, 0x17, 0x7e, 0x0d,
	0xa5, 0x92, 0xe5, 0xbc, 0x8a, 0x1d, 0x0d, 0xc0, 0xaa, 0xe1, 0x38, 0x57, 0x0c, 0x83, 0xbb, 0x4c,
	0x6b, 0xcb, 0xaf, 0xb5, 0xc3, 0xef, 0x31, 0x4c, 0x1a, 0xce, 0xd1, 0xe4, 0x06, 0x60, 0x0a, 0x0c,
	0xb4, 0xa2, 0x6c, 0xb9, 0xf8, 0xd1, 0x86, 0xd3, 0xa2, 0xf6, 0x3d, 0x8a, 0xdb, 0xd0, 0x47, 0x72,
	0x09, 0xbd, 0xea, 0x1d, 0x11, 0xab, 0x64, 0xf1, 0xfb, 0x47, 0xc0, 0x9e, 0x34, 0x64, 0xf4, 0x64,
	0xff, 0x23, 0x9f, 0x61, 0xd4, 0xe0, 0x29, 0x42, 0x4b, 0xcc, 0xe1, 0x37, 0x60, 0x3f, 0x38, 0x5a,
	0x53, 0x9d, 0xf0, 0x16, 0xfa, 0xf5, 0x97, 0x47, 0xce, 0xb7, 0xe3, 0xda, 0x7b, 0xee, 0xf6, 0xb4,
	0x39, 0x59, 0x35, 0x73, 0x81, 0xec, 0x1b, 0x85, 0x5c, 0x94, 0xa8, 0x83, 0x36, 0xb7, 0xe9, 0xb1,
	0x92, 0xb2, 0xbd, 0x63, 0x10, 0x1f, 0xc6, 0x4d, 0xa6, 0x21, 0x5b, 0xb1, 0x87, 0x9d, 0x68, 0x3f,
	0x3c, 0x5e, 0x54, 0xa9, 0xf8, 0x08, 0xc3, 0xbd, 0x9b, 0x27, 0xb3, 0x06, 0xf0, 0x8e, 0xf9, 0xec,
	0x8b, 0x23, 0x15, 0x65, 0xef, 0xeb, 0x4e, 0xfe, 0x3b, 0x78, 0xf2, 0x2b, 0x00, 0x00, 0xff, 0xff,
	0x6e, 0x61, 0xf7, 0x2a, 0x34, 0x06, 0x00, 0x00,
}
