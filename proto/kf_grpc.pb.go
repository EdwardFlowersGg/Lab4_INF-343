// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: proto/kf.proto

package proto

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

// KFserviceClient is the client API for KFservice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KFserviceClient interface {
	// SolicitarEntrarPiso es el método para solicitar entrar al piso.
	SolicitarEntrarPiso(ctx context.Context, in *EntrarPisoRequest, opts ...grpc.CallOption) (*EstadoPisoUpdate, error)
	// SolicitarMovimiento es el método para solicitar un movimiento en el piso.
	SolicitarMovimiento(ctx context.Context, in *MovimientoRequest, opts ...grpc.CallOption) (*MovimientoRequest, error)
	// ActualizarEstadoPiso es el método para actualizar el estado del piso.
	ActualizarEstadoPiso(ctx context.Context, in *EstadoPisoUpdate, opts ...grpc.CallOption) (*EstadoUpdateResponse, error)
	// SolicitarMontoDineroMercenarioDirector es el método para solicitar el monto de dinero actual entre el mercenario y el director.
	SolicitarMontoDineroMercenarioDirector(ctx context.Context, in *MontoDineroMercenarioDirectorRequest, opts ...grpc.CallOption) (*MontoDineroResponse, error)
	// SolicitarMontoDineroDirectorBanco es el método para solicitar el monto de dinero actual entre el director y el banco.
	SolicitarMontoDineroDirectorBanco(ctx context.Context, in *MontoDineroDirectorBancoRequest, opts ...grpc.CallOption) (*MontoDineroResponse, error)
	// SolicitarRegistroDataName es el método para solicitar registro de DataName.
	SolicitarRegistroDataName(ctx context.Context, in *RegistroDataNameRequest, opts ...grpc.CallOption) (*RegistroDataNameResponse, error)
	// SolicitarRegistroNodo es el método para solicitar registro a nodo.
	SolicitarRegistroNodo(ctx context.Context, in *RegistroNodoRequest, opts ...grpc.CallOption) (*RegistroNodoResponse, error)
}

type kFserviceClient struct {
	cc grpc.ClientConnInterface
}

func NewKFserviceClient(cc grpc.ClientConnInterface) KFserviceClient {
	return &kFserviceClient{cc}
}

func (c *kFserviceClient) SolicitarEntrarPiso(ctx context.Context, in *EntrarPisoRequest, opts ...grpc.CallOption) (*EstadoPisoUpdate, error) {
	out := new(EstadoPisoUpdate)
	err := c.cc.Invoke(ctx, "/grpc.KFservice/SolicitarEntrarPiso", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kFserviceClient) SolicitarMovimiento(ctx context.Context, in *MovimientoRequest, opts ...grpc.CallOption) (*MovimientoRequest, error) {
	out := new(MovimientoRequest)
	err := c.cc.Invoke(ctx, "/grpc.KFservice/SolicitarMovimiento", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kFserviceClient) ActualizarEstadoPiso(ctx context.Context, in *EstadoPisoUpdate, opts ...grpc.CallOption) (*EstadoUpdateResponse, error) {
	out := new(EstadoUpdateResponse)
	err := c.cc.Invoke(ctx, "/grpc.KFservice/ActualizarEstadoPiso", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kFserviceClient) SolicitarMontoDineroMercenarioDirector(ctx context.Context, in *MontoDineroMercenarioDirectorRequest, opts ...grpc.CallOption) (*MontoDineroResponse, error) {
	out := new(MontoDineroResponse)
	err := c.cc.Invoke(ctx, "/grpc.KFservice/SolicitarMontoDineroMercenarioDirector", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kFserviceClient) SolicitarMontoDineroDirectorBanco(ctx context.Context, in *MontoDineroDirectorBancoRequest, opts ...grpc.CallOption) (*MontoDineroResponse, error) {
	out := new(MontoDineroResponse)
	err := c.cc.Invoke(ctx, "/grpc.KFservice/SolicitarMontoDineroDirectorBanco", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kFserviceClient) SolicitarRegistroDataName(ctx context.Context, in *RegistroDataNameRequest, opts ...grpc.CallOption) (*RegistroDataNameResponse, error) {
	out := new(RegistroDataNameResponse)
	err := c.cc.Invoke(ctx, "/grpc.KFservice/SolicitarRegistroDataName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kFserviceClient) SolicitarRegistroNodo(ctx context.Context, in *RegistroNodoRequest, opts ...grpc.CallOption) (*RegistroNodoResponse, error) {
	out := new(RegistroNodoResponse)
	err := c.cc.Invoke(ctx, "/grpc.KFservice/SolicitarRegistroNodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KFserviceServer is the server API for KFservice service.
// All implementations must embed UnimplementedKFserviceServer
// for forward compatibility
type KFserviceServer interface {
	// SolicitarEntrarPiso es el método para solicitar entrar al piso.
	SolicitarEntrarPiso(context.Context, *EntrarPisoRequest) (*EstadoPisoUpdate, error)
	// SolicitarMovimiento es el método para solicitar un movimiento en el piso.
	SolicitarMovimiento(context.Context, *MovimientoRequest) (*MovimientoRequest, error)
	// ActualizarEstadoPiso es el método para actualizar el estado del piso.
	ActualizarEstadoPiso(context.Context, *EstadoPisoUpdate) (*EstadoUpdateResponse, error)
	// SolicitarMontoDineroMercenarioDirector es el método para solicitar el monto de dinero actual entre el mercenario y el director.
	SolicitarMontoDineroMercenarioDirector(context.Context, *MontoDineroMercenarioDirectorRequest) (*MontoDineroResponse, error)
	// SolicitarMontoDineroDirectorBanco es el método para solicitar el monto de dinero actual entre el director y el banco.
	SolicitarMontoDineroDirectorBanco(context.Context, *MontoDineroDirectorBancoRequest) (*MontoDineroResponse, error)
	// SolicitarRegistroDataName es el método para solicitar registro de DataName.
	SolicitarRegistroDataName(context.Context, *RegistroDataNameRequest) (*RegistroDataNameResponse, error)
	// SolicitarRegistroNodo es el método para solicitar registro a nodo.
	SolicitarRegistroNodo(context.Context, *RegistroNodoRequest) (*RegistroNodoResponse, error)
	mustEmbedUnimplementedKFserviceServer()
}

// UnimplementedKFserviceServer must be embedded to have forward compatible implementations.
type UnimplementedKFserviceServer struct {
}

func (UnimplementedKFserviceServer) SolicitarEntrarPiso(context.Context, *EntrarPisoRequest) (*EstadoPisoUpdate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SolicitarEntrarPiso not implemented")
}
func (UnimplementedKFserviceServer) SolicitarMovimiento(context.Context, *MovimientoRequest) (*MovimientoRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SolicitarMovimiento not implemented")
}
func (UnimplementedKFserviceServer) ActualizarEstadoPiso(context.Context, *EstadoPisoUpdate) (*EstadoUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActualizarEstadoPiso not implemented")
}
func (UnimplementedKFserviceServer) SolicitarMontoDineroMercenarioDirector(context.Context, *MontoDineroMercenarioDirectorRequest) (*MontoDineroResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SolicitarMontoDineroMercenarioDirector not implemented")
}
func (UnimplementedKFserviceServer) SolicitarMontoDineroDirectorBanco(context.Context, *MontoDineroDirectorBancoRequest) (*MontoDineroResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SolicitarMontoDineroDirectorBanco not implemented")
}
func (UnimplementedKFserviceServer) SolicitarRegistroDataName(context.Context, *RegistroDataNameRequest) (*RegistroDataNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SolicitarRegistroDataName not implemented")
}
func (UnimplementedKFserviceServer) SolicitarRegistroNodo(context.Context, *RegistroNodoRequest) (*RegistroNodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SolicitarRegistroNodo not implemented")
}
func (UnimplementedKFserviceServer) mustEmbedUnimplementedKFserviceServer() {}

// UnsafeKFserviceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KFserviceServer will
// result in compilation errors.
type UnsafeKFserviceServer interface {
	mustEmbedUnimplementedKFserviceServer()
}

func RegisterKFserviceServer(s grpc.ServiceRegistrar, srv KFserviceServer) {
	s.RegisterService(&KFservice_ServiceDesc, srv)
}

func _KFservice_SolicitarEntrarPiso_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EntrarPisoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KFserviceServer).SolicitarEntrarPiso(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.KFservice/SolicitarEntrarPiso",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KFserviceServer).SolicitarEntrarPiso(ctx, req.(*EntrarPisoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KFservice_SolicitarMovimiento_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MovimientoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KFserviceServer).SolicitarMovimiento(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.KFservice/SolicitarMovimiento",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KFserviceServer).SolicitarMovimiento(ctx, req.(*MovimientoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KFservice_ActualizarEstadoPiso_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EstadoPisoUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KFserviceServer).ActualizarEstadoPiso(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.KFservice/ActualizarEstadoPiso",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KFserviceServer).ActualizarEstadoPiso(ctx, req.(*EstadoPisoUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _KFservice_SolicitarMontoDineroMercenarioDirector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MontoDineroMercenarioDirectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KFserviceServer).SolicitarMontoDineroMercenarioDirector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.KFservice/SolicitarMontoDineroMercenarioDirector",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KFserviceServer).SolicitarMontoDineroMercenarioDirector(ctx, req.(*MontoDineroMercenarioDirectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KFservice_SolicitarMontoDineroDirectorBanco_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MontoDineroDirectorBancoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KFserviceServer).SolicitarMontoDineroDirectorBanco(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.KFservice/SolicitarMontoDineroDirectorBanco",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KFserviceServer).SolicitarMontoDineroDirectorBanco(ctx, req.(*MontoDineroDirectorBancoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KFservice_SolicitarRegistroDataName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistroDataNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KFserviceServer).SolicitarRegistroDataName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.KFservice/SolicitarRegistroDataName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KFserviceServer).SolicitarRegistroDataName(ctx, req.(*RegistroDataNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KFservice_SolicitarRegistroNodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistroNodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KFserviceServer).SolicitarRegistroNodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.KFservice/SolicitarRegistroNodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KFserviceServer).SolicitarRegistroNodo(ctx, req.(*RegistroNodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KFservice_ServiceDesc is the grpc.ServiceDesc for KFservice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KFservice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.KFservice",
	HandlerType: (*KFserviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SolicitarEntrarPiso",
			Handler:    _KFservice_SolicitarEntrarPiso_Handler,
		},
		{
			MethodName: "SolicitarMovimiento",
			Handler:    _KFservice_SolicitarMovimiento_Handler,
		},
		{
			MethodName: "ActualizarEstadoPiso",
			Handler:    _KFservice_ActualizarEstadoPiso_Handler,
		},
		{
			MethodName: "SolicitarMontoDineroMercenarioDirector",
			Handler:    _KFservice_SolicitarMontoDineroMercenarioDirector_Handler,
		},
		{
			MethodName: "SolicitarMontoDineroDirectorBanco",
			Handler:    _KFservice_SolicitarMontoDineroDirectorBanco_Handler,
		},
		{
			MethodName: "SolicitarRegistroDataName",
			Handler:    _KFservice_SolicitarRegistroDataName_Handler,
		},
		{
			MethodName: "SolicitarRegistroNodo",
			Handler:    _KFservice_SolicitarRegistroNodo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/kf.proto",
}
