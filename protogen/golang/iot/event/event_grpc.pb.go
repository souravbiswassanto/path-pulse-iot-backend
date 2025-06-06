// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: iot/event/event.proto

package event

import (
	context "context"
	group "github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/group"
	user "github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/user"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EventManagerClient is the client API for EventManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventManagerClient interface {
	AddEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*user.Empty, error)
	UpdateEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*user.Empty, error)
	DeleteEvent(ctx context.Context, in *EventId, opts ...grpc.CallOption) (*user.Empty, error)
	GetSingleEventDetails(ctx context.Context, in *EventId, opts ...grpc.CallOption) (*Event, error)
	ListEventsOfSingleUser(ctx context.Context, in *user.UserID, opts ...grpc.CallOption) (*EventList, error)
	ListEventsOfSingleGroup(ctx context.Context, in *group.GroupId, opts ...grpc.CallOption) (*EventList, error)
}

type eventManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewEventManagerClient(cc grpc.ClientConnInterface) EventManagerClient {
	return &eventManagerClient{cc}
}

func (c *eventManagerClient) AddEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*user.Empty, error) {
	out := new(user.Empty)
	err := c.cc.Invoke(ctx, "/EventManager/AddEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventManagerClient) UpdateEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*user.Empty, error) {
	out := new(user.Empty)
	err := c.cc.Invoke(ctx, "/EventManager/UpdateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventManagerClient) DeleteEvent(ctx context.Context, in *EventId, opts ...grpc.CallOption) (*user.Empty, error) {
	out := new(user.Empty)
	err := c.cc.Invoke(ctx, "/EventManager/DeleteEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventManagerClient) GetSingleEventDetails(ctx context.Context, in *EventId, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/EventManager/GetSingleEventDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventManagerClient) ListEventsOfSingleUser(ctx context.Context, in *user.UserID, opts ...grpc.CallOption) (*EventList, error) {
	out := new(EventList)
	err := c.cc.Invoke(ctx, "/EventManager/ListEventsOfSingleUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventManagerClient) ListEventsOfSingleGroup(ctx context.Context, in *group.GroupId, opts ...grpc.CallOption) (*EventList, error) {
	out := new(EventList)
	err := c.cc.Invoke(ctx, "/EventManager/ListEventsOfSingleGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventManagerServer is the server API for EventManager service.
// All implementations must embed UnimplementedEventManagerServer
// for forward compatibility
type EventManagerServer interface {
	AddEvent(context.Context, *Event) (*user.Empty, error)
	UpdateEvent(context.Context, *Event) (*user.Empty, error)
	DeleteEvent(context.Context, *EventId) (*user.Empty, error)
	GetSingleEventDetails(context.Context, *EventId) (*Event, error)
	ListEventsOfSingleUser(context.Context, *user.UserID) (*EventList, error)
	ListEventsOfSingleGroup(context.Context, *group.GroupId) (*EventList, error)
	mustEmbedUnimplementedEventManagerServer()
}

// UnimplementedEventManagerServer must be embedded to have forward compatible implementations.
type UnimplementedEventManagerServer struct {
}

func (UnimplementedEventManagerServer) AddEvent(context.Context, *Event) (*user.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEvent not implemented")
}
func (UnimplementedEventManagerServer) UpdateEvent(context.Context, *Event) (*user.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEvent not implemented")
}
func (UnimplementedEventManagerServer) DeleteEvent(context.Context, *EventId) (*user.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEvent not implemented")
}
func (UnimplementedEventManagerServer) GetSingleEventDetails(context.Context, *EventId) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSingleEventDetails not implemented")
}
func (UnimplementedEventManagerServer) ListEventsOfSingleUser(context.Context, *user.UserID) (*EventList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEventsOfSingleUser not implemented")
}
func (UnimplementedEventManagerServer) ListEventsOfSingleGroup(context.Context, *group.GroupId) (*EventList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEventsOfSingleGroup not implemented")
}
func (UnimplementedEventManagerServer) mustEmbedUnimplementedEventManagerServer() {}

// UnsafeEventManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventManagerServer will
// result in compilation errors.
type UnsafeEventManagerServer interface {
	mustEmbedUnimplementedEventManagerServer()
}

func RegisterEventManagerServer(s grpc.ServiceRegistrar, srv EventManagerServer) {
	s.RegisterService(&EventManager_ServiceDesc, srv)
}

func _EventManager_AddEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventManagerServer).AddEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EventManager/AddEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventManagerServer).AddEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventManager_UpdateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventManagerServer).UpdateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EventManager/UpdateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventManagerServer).UpdateEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventManager_DeleteEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventManagerServer).DeleteEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EventManager/DeleteEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventManagerServer).DeleteEvent(ctx, req.(*EventId))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventManager_GetSingleEventDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventManagerServer).GetSingleEventDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EventManager/GetSingleEventDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventManagerServer).GetSingleEventDetails(ctx, req.(*EventId))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventManager_ListEventsOfSingleUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(user.UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventManagerServer).ListEventsOfSingleUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EventManager/ListEventsOfSingleUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventManagerServer).ListEventsOfSingleUser(ctx, req.(*user.UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventManager_ListEventsOfSingleGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(group.GroupId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventManagerServer).ListEventsOfSingleGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EventManager/ListEventsOfSingleGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventManagerServer).ListEventsOfSingleGroup(ctx, req.(*group.GroupId))
	}
	return interceptor(ctx, in, info, handler)
}

// EventManager_ServiceDesc is the grpc.ServiceDesc for EventManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "EventManager",
	HandlerType: (*EventManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddEvent",
			Handler:    _EventManager_AddEvent_Handler,
		},
		{
			MethodName: "UpdateEvent",
			Handler:    _EventManager_UpdateEvent_Handler,
		},
		{
			MethodName: "DeleteEvent",
			Handler:    _EventManager_DeleteEvent_Handler,
		},
		{
			MethodName: "GetSingleEventDetails",
			Handler:    _EventManager_GetSingleEventDetails_Handler,
		},
		{
			MethodName: "ListEventsOfSingleUser",
			Handler:    _EventManager_ListEventsOfSingleUser_Handler,
		},
		{
			MethodName: "ListEventsOfSingleGroup",
			Handler:    _EventManager_ListEventsOfSingleGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iot/event/event.proto",
}
