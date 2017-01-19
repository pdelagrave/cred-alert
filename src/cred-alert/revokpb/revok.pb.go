// Code generated by protoc-gen-go.
// source: revok.proto
// DO NOT EDIT!

/*
Package revokpb is a generated protocol buffer package.

It is generated from these files:
	revok.proto

It has these top-level messages:
	CredentialCountRequest
	OrganizationCredentialCount
	CredentialCountResponse
	OrganizationCredentialCountRequest
	RepositoryCredentialCount
	OrganizationCredentialCountResponse
	RepositoryCredentialCountRequest
	BranchCredentialCount
	RepositoryCredentialCountResponse
	SearchQuery
	SearchResult
	SourceLocation
	Repository
*/
package revokpb

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

type CredentialCountRequest struct {
}

func (m *CredentialCountRequest) Reset()                    { *m = CredentialCountRequest{} }
func (m *CredentialCountRequest) String() string            { return proto.CompactTextString(m) }
func (*CredentialCountRequest) ProtoMessage()               {}
func (*CredentialCountRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type OrganizationCredentialCount struct {
	Owner string `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Count int64  `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
}

func (m *OrganizationCredentialCount) Reset()                    { *m = OrganizationCredentialCount{} }
func (m *OrganizationCredentialCount) String() string            { return proto.CompactTextString(m) }
func (*OrganizationCredentialCount) ProtoMessage()               {}
func (*OrganizationCredentialCount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *OrganizationCredentialCount) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *OrganizationCredentialCount) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type CredentialCountResponse struct {
	CredentialCounts []*OrganizationCredentialCount `protobuf:"bytes,1,rep,name=credentialCounts" json:"credentialCounts,omitempty"`
}

func (m *CredentialCountResponse) Reset()                    { *m = CredentialCountResponse{} }
func (m *CredentialCountResponse) String() string            { return proto.CompactTextString(m) }
func (*CredentialCountResponse) ProtoMessage()               {}
func (*CredentialCountResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CredentialCountResponse) GetCredentialCounts() []*OrganizationCredentialCount {
	if m != nil {
		return m.CredentialCounts
	}
	return nil
}

type OrganizationCredentialCountRequest struct {
	Owner string `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
}

func (m *OrganizationCredentialCountRequest) Reset()         { *m = OrganizationCredentialCountRequest{} }
func (m *OrganizationCredentialCountRequest) String() string { return proto.CompactTextString(m) }
func (*OrganizationCredentialCountRequest) ProtoMessage()    {}
func (*OrganizationCredentialCountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3}
}

func (m *OrganizationCredentialCountRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type RepositoryCredentialCount struct {
	Owner string `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Count int64  `protobuf:"varint,3,opt,name=count" json:"count,omitempty"`
}

func (m *RepositoryCredentialCount) Reset()                    { *m = RepositoryCredentialCount{} }
func (m *RepositoryCredentialCount) String() string            { return proto.CompactTextString(m) }
func (*RepositoryCredentialCount) ProtoMessage()               {}
func (*RepositoryCredentialCount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RepositoryCredentialCount) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *RepositoryCredentialCount) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RepositoryCredentialCount) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type OrganizationCredentialCountResponse struct {
	CredentialCounts []*RepositoryCredentialCount `protobuf:"bytes,1,rep,name=credentialCounts" json:"credentialCounts,omitempty"`
}

func (m *OrganizationCredentialCountResponse) Reset()         { *m = OrganizationCredentialCountResponse{} }
func (m *OrganizationCredentialCountResponse) String() string { return proto.CompactTextString(m) }
func (*OrganizationCredentialCountResponse) ProtoMessage()    {}
func (*OrganizationCredentialCountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5}
}

func (m *OrganizationCredentialCountResponse) GetCredentialCounts() []*RepositoryCredentialCount {
	if m != nil {
		return m.CredentialCounts
	}
	return nil
}

type RepositoryCredentialCountRequest struct {
	Owner string `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *RepositoryCredentialCountRequest) Reset()         { *m = RepositoryCredentialCountRequest{} }
func (m *RepositoryCredentialCountRequest) String() string { return proto.CompactTextString(m) }
func (*RepositoryCredentialCountRequest) ProtoMessage()    {}
func (*RepositoryCredentialCountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{6}
}

func (m *RepositoryCredentialCountRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *RepositoryCredentialCountRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type BranchCredentialCount struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Count int64  `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
}

func (m *BranchCredentialCount) Reset()                    { *m = BranchCredentialCount{} }
func (m *BranchCredentialCount) String() string            { return proto.CompactTextString(m) }
func (*BranchCredentialCount) ProtoMessage()               {}
func (*BranchCredentialCount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *BranchCredentialCount) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BranchCredentialCount) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type RepositoryCredentialCountResponse struct {
	CredentialCounts []*BranchCredentialCount `protobuf:"bytes,1,rep,name=credentialCounts" json:"credentialCounts,omitempty"`
}

func (m *RepositoryCredentialCountResponse) Reset()         { *m = RepositoryCredentialCountResponse{} }
func (m *RepositoryCredentialCountResponse) String() string { return proto.CompactTextString(m) }
func (*RepositoryCredentialCountResponse) ProtoMessage()    {}
func (*RepositoryCredentialCountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{8}
}

func (m *RepositoryCredentialCountResponse) GetCredentialCounts() []*BranchCredentialCount {
	if m != nil {
		return m.CredentialCounts
	}
	return nil
}

type SearchQuery struct {
	Regex string `protobuf:"bytes,1,opt,name=regex" json:"regex,omitempty"`
}

func (m *SearchQuery) Reset()                    { *m = SearchQuery{} }
func (m *SearchQuery) String() string            { return proto.CompactTextString(m) }
func (*SearchQuery) ProtoMessage()               {}
func (*SearchQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *SearchQuery) GetRegex() string {
	if m != nil {
		return m.Regex
	}
	return ""
}

type SearchResult struct {
	Location *SourceLocation `protobuf:"bytes,1,opt,name=location" json:"location,omitempty"`
	Content  []byte          `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *SearchResult) Reset()                    { *m = SearchResult{} }
func (m *SearchResult) String() string            { return proto.CompactTextString(m) }
func (*SearchResult) ProtoMessage()               {}
func (*SearchResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *SearchResult) GetLocation() *SourceLocation {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *SearchResult) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type SourceLocation struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	Revision   string      `protobuf:"bytes,2,opt,name=revision" json:"revision,omitempty"`
	Path       string      `protobuf:"bytes,3,opt,name=path" json:"path,omitempty"`
	LineNumber uint32      `protobuf:"varint,4,opt,name=line_number,json=lineNumber" json:"line_number,omitempty"`
	Location   uint32      `protobuf:"varint,5,opt,name=location" json:"location,omitempty"`
	Length     uint32      `protobuf:"varint,6,opt,name=length" json:"length,omitempty"`
}

func (m *SourceLocation) Reset()                    { *m = SourceLocation{} }
func (m *SourceLocation) String() string            { return proto.CompactTextString(m) }
func (*SourceLocation) ProtoMessage()               {}
func (*SourceLocation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *SourceLocation) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *SourceLocation) GetRevision() string {
	if m != nil {
		return m.Revision
	}
	return ""
}

func (m *SourceLocation) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *SourceLocation) GetLineNumber() uint32 {
	if m != nil {
		return m.LineNumber
	}
	return 0
}

func (m *SourceLocation) GetLocation() uint32 {
	if m != nil {
		return m.Location
	}
	return 0
}

func (m *SourceLocation) GetLength() uint32 {
	if m != nil {
		return m.Length
	}
	return 0
}

type Repository struct {
	Owner string `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *Repository) Reset()                    { *m = Repository{} }
func (m *Repository) String() string            { return proto.CompactTextString(m) }
func (*Repository) ProtoMessage()               {}
func (*Repository) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *Repository) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Repository) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*CredentialCountRequest)(nil), "revokpb.CredentialCountRequest")
	proto.RegisterType((*OrganizationCredentialCount)(nil), "revokpb.OrganizationCredentialCount")
	proto.RegisterType((*CredentialCountResponse)(nil), "revokpb.CredentialCountResponse")
	proto.RegisterType((*OrganizationCredentialCountRequest)(nil), "revokpb.OrganizationCredentialCountRequest")
	proto.RegisterType((*RepositoryCredentialCount)(nil), "revokpb.RepositoryCredentialCount")
	proto.RegisterType((*OrganizationCredentialCountResponse)(nil), "revokpb.OrganizationCredentialCountResponse")
	proto.RegisterType((*RepositoryCredentialCountRequest)(nil), "revokpb.RepositoryCredentialCountRequest")
	proto.RegisterType((*BranchCredentialCount)(nil), "revokpb.BranchCredentialCount")
	proto.RegisterType((*RepositoryCredentialCountResponse)(nil), "revokpb.RepositoryCredentialCountResponse")
	proto.RegisterType((*SearchQuery)(nil), "revokpb.SearchQuery")
	proto.RegisterType((*SearchResult)(nil), "revokpb.SearchResult")
	proto.RegisterType((*SourceLocation)(nil), "revokpb.SourceLocation")
	proto.RegisterType((*Repository)(nil), "revokpb.Repository")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Revok service

type RevokClient interface {
	GetCredentialCounts(ctx context.Context, in *CredentialCountRequest, opts ...grpc.CallOption) (*CredentialCountResponse, error)
	GetOrganizationCredentialCounts(ctx context.Context, in *OrganizationCredentialCountRequest, opts ...grpc.CallOption) (*OrganizationCredentialCountResponse, error)
	GetRepositoryCredentialCounts(ctx context.Context, in *RepositoryCredentialCountRequest, opts ...grpc.CallOption) (*RepositoryCredentialCountResponse, error)
	Search(ctx context.Context, in *SearchQuery, opts ...grpc.CallOption) (Revok_SearchClient, error)
}

type revokClient struct {
	cc *grpc.ClientConn
}

func NewRevokClient(cc *grpc.ClientConn) RevokClient {
	return &revokClient{cc}
}

func (c *revokClient) GetCredentialCounts(ctx context.Context, in *CredentialCountRequest, opts ...grpc.CallOption) (*CredentialCountResponse, error) {
	out := new(CredentialCountResponse)
	err := grpc.Invoke(ctx, "/revokpb.Revok/GetCredentialCounts", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *revokClient) GetOrganizationCredentialCounts(ctx context.Context, in *OrganizationCredentialCountRequest, opts ...grpc.CallOption) (*OrganizationCredentialCountResponse, error) {
	out := new(OrganizationCredentialCountResponse)
	err := grpc.Invoke(ctx, "/revokpb.Revok/GetOrganizationCredentialCounts", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *revokClient) GetRepositoryCredentialCounts(ctx context.Context, in *RepositoryCredentialCountRequest, opts ...grpc.CallOption) (*RepositoryCredentialCountResponse, error) {
	out := new(RepositoryCredentialCountResponse)
	err := grpc.Invoke(ctx, "/revokpb.Revok/GetRepositoryCredentialCounts", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *revokClient) Search(ctx context.Context, in *SearchQuery, opts ...grpc.CallOption) (Revok_SearchClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Revok_serviceDesc.Streams[0], c.cc, "/revokpb.Revok/Search", opts...)
	if err != nil {
		return nil, err
	}
	x := &revokSearchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Revok_SearchClient interface {
	Recv() (*SearchResult, error)
	grpc.ClientStream
}

type revokSearchClient struct {
	grpc.ClientStream
}

func (x *revokSearchClient) Recv() (*SearchResult, error) {
	m := new(SearchResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Revok service

type RevokServer interface {
	GetCredentialCounts(context.Context, *CredentialCountRequest) (*CredentialCountResponse, error)
	GetOrganizationCredentialCounts(context.Context, *OrganizationCredentialCountRequest) (*OrganizationCredentialCountResponse, error)
	GetRepositoryCredentialCounts(context.Context, *RepositoryCredentialCountRequest) (*RepositoryCredentialCountResponse, error)
	Search(*SearchQuery, Revok_SearchServer) error
}

func RegisterRevokServer(s *grpc.Server, srv RevokServer) {
	s.RegisterService(&_Revok_serviceDesc, srv)
}

func _Revok_GetCredentialCounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CredentialCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RevokServer).GetCredentialCounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/revokpb.Revok/GetCredentialCounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RevokServer).GetCredentialCounts(ctx, req.(*CredentialCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Revok_GetOrganizationCredentialCounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationCredentialCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RevokServer).GetOrganizationCredentialCounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/revokpb.Revok/GetOrganizationCredentialCounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RevokServer).GetOrganizationCredentialCounts(ctx, req.(*OrganizationCredentialCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Revok_GetRepositoryCredentialCounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepositoryCredentialCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RevokServer).GetRepositoryCredentialCounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/revokpb.Revok/GetRepositoryCredentialCounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RevokServer).GetRepositoryCredentialCounts(ctx, req.(*RepositoryCredentialCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Revok_Search_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchQuery)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RevokServer).Search(m, &revokSearchServer{stream})
}

type Revok_SearchServer interface {
	Send(*SearchResult) error
	grpc.ServerStream
}

type revokSearchServer struct {
	grpc.ServerStream
}

func (x *revokSearchServer) Send(m *SearchResult) error {
	return x.ServerStream.SendMsg(m)
}

var _Revok_serviceDesc = grpc.ServiceDesc{
	ServiceName: "revokpb.Revok",
	HandlerType: (*RevokServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCredentialCounts",
			Handler:    _Revok_GetCredentialCounts_Handler,
		},
		{
			MethodName: "GetOrganizationCredentialCounts",
			Handler:    _Revok_GetOrganizationCredentialCounts_Handler,
		},
		{
			MethodName: "GetRepositoryCredentialCounts",
			Handler:    _Revok_GetRepositoryCredentialCounts_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Search",
			Handler:       _Revok_Search_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "revok.proto",
}

func init() { proto.RegisterFile("revok.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 537 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x55, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xad, 0xbf, 0x34, 0x69, 0x73, 0xd3, 0x0f, 0xa1, 0xe9, 0x9f, 0x09, 0x82, 0x98, 0x29, 0x8b,
	0xf0, 0xa3, 0x08, 0x25, 0x12, 0x12, 0xec, 0xa0, 0x8b, 0x08, 0x54, 0x15, 0x98, 0xee, 0x40, 0x08,
	0x39, 0xe6, 0x2a, 0xb1, 0xea, 0xce, 0x98, 0x99, 0x71, 0x69, 0x59, 0xf0, 0x1c, 0x3c, 0x13, 0x4f,
	0x85, 0x3c, 0x76, 0xc6, 0x4d, 0x6a, 0x5b, 0xde, 0xf9, 0xce, 0x3d, 0x93, 0x7b, 0xce, 0xb9, 0xc7,
	0x31, 0xf4, 0x24, 0x5e, 0x8a, 0xf3, 0x51, 0x2c, 0x85, 0x16, 0x64, 0xcb, 0x14, 0xf1, 0x8c, 0xba,
	0x70, 0x70, 0x2c, 0xf1, 0x3b, 0x72, 0x1d, 0xfa, 0xd1, 0xb1, 0x48, 0xb8, 0x66, 0xf8, 0x23, 0x41,
	0xa5, 0xe9, 0x3b, 0xb8, 0xff, 0x41, 0xce, 0x7d, 0x1e, 0xfe, 0xf2, 0x75, 0x28, 0xf8, 0x1a, 0x8a,
	0xec, 0x41, 0x5b, 0xfc, 0xe4, 0x28, 0x5d, 0xc7, 0x73, 0x86, 0x5d, 0x96, 0x15, 0xe9, 0x69, 0x90,
	0xb6, 0xdd, 0xff, 0x3c, 0x67, 0xd8, 0x62, 0x59, 0x41, 0xcf, 0xe1, 0xf0, 0xd6, 0x10, 0x15, 0x0b,
	0xae, 0x90, 0x7c, 0x84, 0xbb, 0xc1, 0x6a, 0x4b, 0xb9, 0x8e, 0xd7, 0x1a, 0xf6, 0xc6, 0x8f, 0x47,
	0x39, 0xc7, 0x51, 0x0d, 0x0d, 0x76, 0xeb, 0x36, 0x7d, 0x0d, 0xb4, 0xee, 0x42, 0xa6, 0xae, 0x9c,
	0x3e, 0xfd, 0x02, 0xf7, 0x18, 0xc6, 0x42, 0x85, 0x5a, 0xc8, 0xeb, 0x66, 0x8a, 0x09, 0x6c, 0x72,
	0xff, 0x02, 0x8d, 0xe0, 0x2e, 0x33, 0xcf, 0x85, 0x0b, 0xad, 0x9b, 0x2e, 0x24, 0x70, 0x54, 0x4b,
	0x2c, 0x77, 0xe4, 0xb4, 0xd2, 0x11, 0x6a, 0x1d, 0xa9, 0x24, 0x59, 0xe2, 0xc7, 0x09, 0x78, 0xd5,
	0xf0, 0x3a, 0x37, 0xca, 0xa4, 0xd1, 0x37, 0xb0, 0xff, 0x56, 0xfa, 0x3c, 0x58, 0xac, 0xbb, 0xb3,
	0x04, 0x3b, 0x65, 0x3e, 0xac, 0xa4, 0x41, 0xc0, 0xa3, 0x1a, 0x42, 0xb9, 0x0b, 0xef, 0x2b, 0x5d,
	0x78, 0x68, 0x5d, 0x28, 0x25, 0x52, 0xe2, 0xc0, 0x11, 0xf4, 0xce, 0xd0, 0x97, 0xc1, 0xe2, 0x53,
	0x82, 0xf2, 0x3a, 0x65, 0x25, 0x71, 0x8e, 0x57, 0x4b, 0xb1, 0xa6, 0xa0, 0x5f, 0x61, 0x27, 0x03,
	0x31, 0x54, 0x49, 0xa4, 0xc9, 0x04, 0xb6, 0x23, 0x11, 0x98, 0x4d, 0x19, 0x60, 0x6f, 0x7c, 0x68,
	0x07, 0x9f, 0x89, 0x44, 0x06, 0x78, 0x92, 0xb7, 0x99, 0x05, 0x12, 0x17, 0xb6, 0x02, 0xc1, 0x35,
	0xe6, 0x92, 0x77, 0xd8, 0xb2, 0xa4, 0x7f, 0x1d, 0xb8, 0xb3, 0x7a, 0x8d, 0x4c, 0x00, 0xa4, 0xf5,
	0x21, 0x9f, 0xb1, 0x5b, 0xb2, 0x62, 0x76, 0x03, 0x46, 0xfa, 0xb0, 0x2d, 0xf1, 0x32, 0x54, 0x29,
	0xad, 0x6c, 0x2f, 0xb6, 0x4e, 0x57, 0x10, 0xfb, 0x7a, 0x61, 0x52, 0xd7, 0x65, 0xe6, 0x99, 0x0c,
	0xa0, 0x17, 0x85, 0x1c, 0xbf, 0xf1, 0xe4, 0x62, 0x86, 0xd2, 0xdd, 0xf4, 0x9c, 0xe1, 0xff, 0x0c,
	0xd2, 0xa3, 0x53, 0x73, 0x92, 0xfe, 0xa0, 0xd5, 0xd9, 0x36, 0xdd, 0x42, 0xce, 0x01, 0x74, 0x22,
	0xe4, 0x73, 0xbd, 0x70, 0x3b, 0xa6, 0x93, 0x57, 0xf4, 0x25, 0x40, 0x41, 0xaf, 0x79, 0x78, 0xc6,
	0x7f, 0x5a, 0xd0, 0x66, 0xa9, 0x3e, 0xf2, 0x19, 0x76, 0xa7, 0xa8, 0xd7, 0x56, 0xa7, 0xc8, 0xc0,
	0xca, 0x2f, 0x0f, 0x6a, 0xdf, 0xab, 0x06, 0x64, 0xc1, 0xa1, 0x1b, 0xe4, 0x37, 0x0c, 0xa6, 0xa8,
	0x6b, 0x5e, 0x35, 0x45, 0x9e, 0x35, 0xfa, 0x6f, 0xc9, 0x67, 0x3e, 0x6f, 0x06, 0xb6, 0xf3, 0xaf,
	0xe0, 0xc1, 0x14, 0x75, 0x65, 0xc4, 0x15, 0x79, 0xd2, 0xe0, 0x3d, 0xce, 0x67, 0x3f, 0x6d, 0x02,
	0xb5, 0x93, 0x5f, 0x41, 0x27, 0xcb, 0x30, 0xd9, 0x2b, 0xb2, 0x5a, 0x24, 0xbf, 0xbf, 0xbf, 0x76,
	0x9a, 0x45, 0x9d, 0x6e, 0xbc, 0x70, 0x66, 0x1d, 0xf3, 0x5d, 0x98, 0xfc, 0x0b, 0x00, 0x00, 0xff,
	0xff, 0x55, 0x09, 0xa1, 0x1f, 0x26, 0x06, 0x00, 0x00,
}
