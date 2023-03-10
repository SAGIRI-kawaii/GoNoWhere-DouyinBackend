// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.0--rc2
// source: interact.proto

package interact

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

// InteractClient is the client API for Interact service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InteractClient interface {
	FavoriteAction(ctx context.Context, in *DouyinFavoriteActionRequest, opts ...grpc.CallOption) (*DouyinFavoriteActionResponse, error)
	FavoriteList(ctx context.Context, in *DouyinFavoriteListRequest, opts ...grpc.CallOption) (*DouyinFavoriteListResponse, error)
	CommentAction(ctx context.Context, in *DouyinCommentActionRequest, opts ...grpc.CallOption) (*DouyinCommentActionResponse, error)
	CommentList(ctx context.Context, in *DouyinCommentListRequest, opts ...grpc.CallOption) (*DouyinCommentListResponse, error)
}

type interactClient struct {
	cc grpc.ClientConnInterface
}

func NewInteractClient(cc grpc.ClientConnInterface) InteractClient {
	return &interactClient{cc}
}

func (c *interactClient) FavoriteAction(ctx context.Context, in *DouyinFavoriteActionRequest, opts ...grpc.CallOption) (*DouyinFavoriteActionResponse, error) {
	out := new(DouyinFavoriteActionResponse)
	err := c.cc.Invoke(ctx, "/interactclient.Interact/favorite_action", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactClient) FavoriteList(ctx context.Context, in *DouyinFavoriteListRequest, opts ...grpc.CallOption) (*DouyinFavoriteListResponse, error) {
	out := new(DouyinFavoriteListResponse)
	err := c.cc.Invoke(ctx, "/interactclient.Interact/favorite_list", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactClient) CommentAction(ctx context.Context, in *DouyinCommentActionRequest, opts ...grpc.CallOption) (*DouyinCommentActionResponse, error) {
	out := new(DouyinCommentActionResponse)
	err := c.cc.Invoke(ctx, "/interactclient.Interact/comment_action", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactClient) CommentList(ctx context.Context, in *DouyinCommentListRequest, opts ...grpc.CallOption) (*DouyinCommentListResponse, error) {
	out := new(DouyinCommentListResponse)
	err := c.cc.Invoke(ctx, "/interactclient.Interact/comment_list", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InteractServer is the server API for Interact service.
// All implementations must embed UnimplementedInteractServer
// for forward compatibility
type InteractServer interface {
	FavoriteAction(context.Context, *DouyinFavoriteActionRequest) (*DouyinFavoriteActionResponse, error)
	FavoriteList(context.Context, *DouyinFavoriteListRequest) (*DouyinFavoriteListResponse, error)
	CommentAction(context.Context, *DouyinCommentActionRequest) (*DouyinCommentActionResponse, error)
	CommentList(context.Context, *DouyinCommentListRequest) (*DouyinCommentListResponse, error)
	mustEmbedUnimplementedInteractServer()
}

// UnimplementedInteractServer must be embedded to have forward compatible implementations.
type UnimplementedInteractServer struct {
}

func (UnimplementedInteractServer) FavoriteAction(context.Context, *DouyinFavoriteActionRequest) (*DouyinFavoriteActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteAction not implemented")
}
func (UnimplementedInteractServer) FavoriteList(context.Context, *DouyinFavoriteListRequest) (*DouyinFavoriteListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteList not implemented")
}
func (UnimplementedInteractServer) CommentAction(context.Context, *DouyinCommentActionRequest) (*DouyinCommentActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentAction not implemented")
}
func (UnimplementedInteractServer) CommentList(context.Context, *DouyinCommentListRequest) (*DouyinCommentListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentList not implemented")
}
func (UnimplementedInteractServer) mustEmbedUnimplementedInteractServer() {}

// UnsafeInteractServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InteractServer will
// result in compilation errors.
type UnsafeInteractServer interface {
	mustEmbedUnimplementedInteractServer()
}

func RegisterInteractServer(s grpc.ServiceRegistrar, srv InteractServer) {
	s.RegisterService(&Interact_ServiceDesc, srv)
}

func _Interact_FavoriteAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinFavoriteActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractServer).FavoriteAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interactclient.Interact/favorite_action",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractServer).FavoriteAction(ctx, req.(*DouyinFavoriteActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Interact_FavoriteList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinFavoriteListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractServer).FavoriteList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interactclient.Interact/favorite_list",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractServer).FavoriteList(ctx, req.(*DouyinFavoriteListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Interact_CommentAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinCommentActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractServer).CommentAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interactclient.Interact/comment_action",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractServer).CommentAction(ctx, req.(*DouyinCommentActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Interact_CommentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinCommentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractServer).CommentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interactclient.Interact/comment_list",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractServer).CommentList(ctx, req.(*DouyinCommentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Interact_ServiceDesc is the grpc.ServiceDesc for Interact service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Interact_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "interactclient.Interact",
	HandlerType: (*InteractServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "favorite_action",
			Handler:    _Interact_FavoriteAction_Handler,
		},
		{
			MethodName: "favorite_list",
			Handler:    _Interact_FavoriteList_Handler,
		},
		{
			MethodName: "comment_action",
			Handler:    _Interact_CommentAction_Handler,
		},
		{
			MethodName: "comment_list",
			Handler:    _Interact_CommentList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interact.proto",
}
