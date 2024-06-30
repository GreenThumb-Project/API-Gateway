// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: gardenManagement.proto

package gardenManagement

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

// GardenManagementClient is the client API for GardenManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GardenManagementClient interface {
	CreateGarden(ctx context.Context, in *CreateGardenRequest, opts ...grpc.CallOption) (*CreateGardenResponse, error)
	ViewGarden(ctx context.Context, in *ViewGardenRequest, opts ...grpc.CallOption) (*ViewGardenResponse, error)
	UpdateGarden(ctx context.Context, in *UpdateGardenRequest, opts ...grpc.CallOption) (*UpdateGardenResponse, error)
	DeleteGarden(ctx context.Context, in *DeleteGardenRequest, opts ...grpc.CallOption) (*DeleteGardenResponse, error)
	ViewUserGardens(ctx context.Context, in *ViewUserGardensRequest, opts ...grpc.CallOption) (*ViewUserGardensResponse, error)
	AddPlanttoGarden(ctx context.Context, in *AddPlanttoGardenRequest, opts ...grpc.CallOption) (*AddPlanttoGardenResponse, error)
	ViewGardenPlants(ctx context.Context, in *ViewGardenPlantsRequest, opts ...grpc.CallOption) (*ViewGardenPlantsResponse, error)
	UpdatePlant(ctx context.Context, in *UpdatePlantRequest, opts ...grpc.CallOption) (*UpdatePlantResponse, error)
	DeletePlant(ctx context.Context, in *DeletePlantRequest, opts ...grpc.CallOption) (*DeletePlantResponse, error)
	AddPlantCareLog(ctx context.Context, in *AddPlantCareLogResquest, opts ...grpc.CallOption) (*AddPlantCareLogResponse, error)
	ViewPlantCareLogs(ctx context.Context, in *ViewPlantCareLogsRequest, opts ...grpc.CallOption) (*ViewPlantCareLogsResponse, error)
}

type gardenManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewGardenManagementClient(cc grpc.ClientConnInterface) GardenManagementClient {
	return &gardenManagementClient{cc}
}

func (c *gardenManagementClient) CreateGarden(ctx context.Context, in *CreateGardenRequest, opts ...grpc.CallOption) (*CreateGardenResponse, error) {
	out := new(CreateGardenResponse)
	err := c.cc.Invoke(ctx, "/gardenManagement.gardenManagement/CreateGarden", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gardenManagementClient) ViewGarden(ctx context.Context, in *ViewGardenRequest, opts ...grpc.CallOption) (*ViewGardenResponse, error) {
	out := new(ViewGardenResponse)
	err := c.cc.Invoke(ctx, "/gardenManagement.gardenManagement/ViewGarden", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gardenManagementClient) UpdateGarden(ctx context.Context, in *UpdateGardenRequest, opts ...grpc.CallOption) (*UpdateGardenResponse, error) {
	out := new(UpdateGardenResponse)
	err := c.cc.Invoke(ctx, "/gardenManagement.gardenManagement/UpdateGarden", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gardenManagementClient) DeleteGarden(ctx context.Context, in *DeleteGardenRequest, opts ...grpc.CallOption) (*DeleteGardenResponse, error) {
	out := new(DeleteGardenResponse)
	err := c.cc.Invoke(ctx, "/gardenManagement.gardenManagement/DeleteGarden", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gardenManagementClient) ViewUserGardens(ctx context.Context, in *ViewUserGardensRequest, opts ...grpc.CallOption) (*ViewUserGardensResponse, error) {
	out := new(ViewUserGardensResponse)
	err := c.cc.Invoke(ctx, "/gardenManagement.gardenManagement/ViewUserGardens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gardenManagementClient) AddPlanttoGarden(ctx context.Context, in *AddPlanttoGardenRequest, opts ...grpc.CallOption) (*AddPlanttoGardenResponse, error) {
	out := new(AddPlanttoGardenResponse)
	err := c.cc.Invoke(ctx, "/gardenManagement.gardenManagement/AddPlanttoGarden", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gardenManagementClient) ViewGardenPlants(ctx context.Context, in *ViewGardenPlantsRequest, opts ...grpc.CallOption) (*ViewGardenPlantsResponse, error) {
	out := new(ViewGardenPlantsResponse)
	err := c.cc.Invoke(ctx, "/gardenManagement.gardenManagement/ViewGardenPlants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gardenManagementClient) UpdatePlant(ctx context.Context, in *UpdatePlantRequest, opts ...grpc.CallOption) (*UpdatePlantResponse, error) {
	out := new(UpdatePlantResponse)
	err := c.cc.Invoke(ctx, "/gardenManagement.gardenManagement/UpdatePlant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gardenManagementClient) DeletePlant(ctx context.Context, in *DeletePlantRequest, opts ...grpc.CallOption) (*DeletePlantResponse, error) {
	out := new(DeletePlantResponse)
	err := c.cc.Invoke(ctx, "/gardenManagement.gardenManagement/DeletePlant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gardenManagementClient) AddPlantCareLog(ctx context.Context, in *AddPlantCareLogResquest, opts ...grpc.CallOption) (*AddPlantCareLogResponse, error) {
	out := new(AddPlantCareLogResponse)
	err := c.cc.Invoke(ctx, "/gardenManagement.gardenManagement/AddPlantCareLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gardenManagementClient) ViewPlantCareLogs(ctx context.Context, in *ViewPlantCareLogsRequest, opts ...grpc.CallOption) (*ViewPlantCareLogsResponse, error) {
	out := new(ViewPlantCareLogsResponse)
	err := c.cc.Invoke(ctx, "/gardenManagement.gardenManagement/ViewPlantCareLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GardenManagementServer is the server API for GardenManagement service.
// All implementations must embed UnimplementedGardenManagementServer
// for forward compatibility
type GardenManagementServer interface {
	CreateGarden(context.Context, *CreateGardenRequest) (*CreateGardenResponse, error)
	ViewGarden(context.Context, *ViewGardenRequest) (*ViewGardenResponse, error)
	UpdateGarden(context.Context, *UpdateGardenRequest) (*UpdateGardenResponse, error)
	DeleteGarden(context.Context, *DeleteGardenRequest) (*DeleteGardenResponse, error)
	ViewUserGardens(context.Context, *ViewUserGardensRequest) (*ViewUserGardensResponse, error)
	AddPlanttoGarden(context.Context, *AddPlanttoGardenRequest) (*AddPlanttoGardenResponse, error)
	ViewGardenPlants(context.Context, *ViewGardenPlantsRequest) (*ViewGardenPlantsResponse, error)
	UpdatePlant(context.Context, *UpdatePlantRequest) (*UpdatePlantResponse, error)
	DeletePlant(context.Context, *DeletePlantRequest) (*DeletePlantResponse, error)
	AddPlantCareLog(context.Context, *AddPlantCareLogResquest) (*AddPlantCareLogResponse, error)
	ViewPlantCareLogs(context.Context, *ViewPlantCareLogsRequest) (*ViewPlantCareLogsResponse, error)
	mustEmbedUnimplementedGardenManagementServer()
}

// UnimplementedGardenManagementServer must be embedded to have forward compatible implementations.
type UnimplementedGardenManagementServer struct {
}

func (UnimplementedGardenManagementServer) CreateGarden(context.Context, *CreateGardenRequest) (*CreateGardenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGarden not implemented")
}
func (UnimplementedGardenManagementServer) ViewGarden(context.Context, *ViewGardenRequest) (*ViewGardenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewGarden not implemented")
}
func (UnimplementedGardenManagementServer) UpdateGarden(context.Context, *UpdateGardenRequest) (*UpdateGardenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGarden not implemented")
}
func (UnimplementedGardenManagementServer) DeleteGarden(context.Context, *DeleteGardenRequest) (*DeleteGardenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGarden not implemented")
}
func (UnimplementedGardenManagementServer) ViewUserGardens(context.Context, *ViewUserGardensRequest) (*ViewUserGardensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewUserGardens not implemented")
}
func (UnimplementedGardenManagementServer) AddPlanttoGarden(context.Context, *AddPlanttoGardenRequest) (*AddPlanttoGardenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPlanttoGarden not implemented")
}
func (UnimplementedGardenManagementServer) ViewGardenPlants(context.Context, *ViewGardenPlantsRequest) (*ViewGardenPlantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewGardenPlants not implemented")
}
func (UnimplementedGardenManagementServer) UpdatePlant(context.Context, *UpdatePlantRequest) (*UpdatePlantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePlant not implemented")
}
func (UnimplementedGardenManagementServer) DeletePlant(context.Context, *DeletePlantRequest) (*DeletePlantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePlant not implemented")
}
func (UnimplementedGardenManagementServer) AddPlantCareLog(context.Context, *AddPlantCareLogResquest) (*AddPlantCareLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPlantCareLog not implemented")
}
func (UnimplementedGardenManagementServer) ViewPlantCareLogs(context.Context, *ViewPlantCareLogsRequest) (*ViewPlantCareLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewPlantCareLogs not implemented")
}
func (UnimplementedGardenManagementServer) mustEmbedUnimplementedGardenManagementServer() {}

// UnsafeGardenManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GardenManagementServer will
// result in compilation errors.
type UnsafeGardenManagementServer interface {
	mustEmbedUnimplementedGardenManagementServer()
}

func RegisterGardenManagementServer(s grpc.ServiceRegistrar, srv GardenManagementServer) {
	s.RegisterService(&GardenManagement_ServiceDesc, srv)
}

func _GardenManagement_CreateGarden_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGardenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GardenManagementServer).CreateGarden(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gardenManagement.gardenManagement/CreateGarden",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GardenManagementServer).CreateGarden(ctx, req.(*CreateGardenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GardenManagement_ViewGarden_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewGardenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GardenManagementServer).ViewGarden(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gardenManagement.gardenManagement/ViewGarden",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GardenManagementServer).ViewGarden(ctx, req.(*ViewGardenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GardenManagement_UpdateGarden_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGardenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GardenManagementServer).UpdateGarden(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gardenManagement.gardenManagement/UpdateGarden",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GardenManagementServer).UpdateGarden(ctx, req.(*UpdateGardenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GardenManagement_DeleteGarden_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGardenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GardenManagementServer).DeleteGarden(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gardenManagement.gardenManagement/DeleteGarden",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GardenManagementServer).DeleteGarden(ctx, req.(*DeleteGardenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GardenManagement_ViewUserGardens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewUserGardensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GardenManagementServer).ViewUserGardens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gardenManagement.gardenManagement/ViewUserGardens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GardenManagementServer).ViewUserGardens(ctx, req.(*ViewUserGardensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GardenManagement_AddPlanttoGarden_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPlanttoGardenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GardenManagementServer).AddPlanttoGarden(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gardenManagement.gardenManagement/AddPlanttoGarden",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GardenManagementServer).AddPlanttoGarden(ctx, req.(*AddPlanttoGardenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GardenManagement_ViewGardenPlants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewGardenPlantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GardenManagementServer).ViewGardenPlants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gardenManagement.gardenManagement/ViewGardenPlants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GardenManagementServer).ViewGardenPlants(ctx, req.(*ViewGardenPlantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GardenManagement_UpdatePlant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePlantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GardenManagementServer).UpdatePlant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gardenManagement.gardenManagement/UpdatePlant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GardenManagementServer).UpdatePlant(ctx, req.(*UpdatePlantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GardenManagement_DeletePlant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePlantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GardenManagementServer).DeletePlant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gardenManagement.gardenManagement/DeletePlant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GardenManagementServer).DeletePlant(ctx, req.(*DeletePlantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GardenManagement_AddPlantCareLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPlantCareLogResquest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GardenManagementServer).AddPlantCareLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gardenManagement.gardenManagement/AddPlantCareLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GardenManagementServer).AddPlantCareLog(ctx, req.(*AddPlantCareLogResquest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GardenManagement_ViewPlantCareLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewPlantCareLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GardenManagementServer).ViewPlantCareLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gardenManagement.gardenManagement/ViewPlantCareLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GardenManagementServer).ViewPlantCareLogs(ctx, req.(*ViewPlantCareLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GardenManagement_ServiceDesc is the grpc.ServiceDesc for GardenManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GardenManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gardenManagement.gardenManagement",
	HandlerType: (*GardenManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGarden",
			Handler:    _GardenManagement_CreateGarden_Handler,
		},
		{
			MethodName: "ViewGarden",
			Handler:    _GardenManagement_ViewGarden_Handler,
		},
		{
			MethodName: "UpdateGarden",
			Handler:    _GardenManagement_UpdateGarden_Handler,
		},
		{
			MethodName: "DeleteGarden",
			Handler:    _GardenManagement_DeleteGarden_Handler,
		},
		{
			MethodName: "ViewUserGardens",
			Handler:    _GardenManagement_ViewUserGardens_Handler,
		},
		{
			MethodName: "AddPlanttoGarden",
			Handler:    _GardenManagement_AddPlanttoGarden_Handler,
		},
		{
			MethodName: "ViewGardenPlants",
			Handler:    _GardenManagement_ViewGardenPlants_Handler,
		},
		{
			MethodName: "UpdatePlant",
			Handler:    _GardenManagement_UpdatePlant_Handler,
		},
		{
			MethodName: "DeletePlant",
			Handler:    _GardenManagement_DeletePlant_Handler,
		},
		{
			MethodName: "AddPlantCareLog",
			Handler:    _GardenManagement_AddPlantCareLog_Handler,
		},
		{
			MethodName: "ViewPlantCareLogs",
			Handler:    _GardenManagement_ViewPlantCareLogs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gardenManagement.proto",
}
