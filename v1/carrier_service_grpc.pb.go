// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: carrier_service.proto

package v1

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

// CarrierServiceClient is the client API for CarrierService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CarrierServiceClient interface {
	// get business
	GetBusinessById(ctx context.Context, in *BusinessRequest, opts ...grpc.CallOption) (*DynamoBusiness, error)
	GetAllBusinesses(ctx context.Context, in *AllBusinessRequest, opts ...grpc.CallOption) (*Ok, error)
	// location
	GetLocations(ctx context.Context, in *BusinessId, opts ...grpc.CallOption) (*ListsOfLocation, error)
	CreateLocation(ctx context.Context, in *Location, opts ...grpc.CallOption) (*Location, error)
	UpdateLocation(ctx context.Context, in *Location, opts ...grpc.CallOption) (*Ok, error)
	DeleteLocation(ctx context.Context, in *Location, opts ...grpc.CallOption) (*Ok, error)
	// health
	PingCarrierService(ctx context.Context, in *CarrierServicePing, opts ...grpc.CallOption) (*CarrierServicePing, error)
	// quote
	GetAllQuotes(ctx context.Context, in *FetchQuotes, opts ...grpc.CallOption) (*QuoteEntities, error)
	GetNewQuotes(ctx context.Context, in *QuoteEntity, opts ...grpc.CallOption) (*QuoteBids, error)
	UpdateQuote(ctx context.Context, in *QuoteEntity, opts ...grpc.CallOption) (*QuoteBids, error)
	DeleteQuote(ctx context.Context, in *DeleteQuote, opts ...grpc.CallOption) (*Ok, error)
	// booking
	BookQuote(ctx context.Context, in *BookingData, opts ...grpc.CallOption) (*BookingResponse, error)
	GetBookingHistory(ctx context.Context, in *FetchBookingsRequest, opts ...grpc.CallOption) (*ListOfBooking, error)
	GetBookingById(ctx context.Context, in *FetchBookingsRequest, opts ...grpc.CallOption) (*ListOfBooking, error)
}

type carrierServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCarrierServiceClient(cc grpc.ClientConnInterface) CarrierServiceClient {
	return &carrierServiceClient{cc}
}

func (c *carrierServiceClient) GetBusinessById(ctx context.Context, in *BusinessRequest, opts ...grpc.CallOption) (*DynamoBusiness, error) {
	out := new(DynamoBusiness)
	err := c.cc.Invoke(ctx, "/user.CarrierService/GetBusinessById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carrierServiceClient) GetAllBusinesses(ctx context.Context, in *AllBusinessRequest, opts ...grpc.CallOption) (*Ok, error) {
	out := new(Ok)
	err := c.cc.Invoke(ctx, "/user.CarrierService/GetAllBusinesses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carrierServiceClient) GetLocations(ctx context.Context, in *BusinessId, opts ...grpc.CallOption) (*ListsOfLocation, error) {
	out := new(ListsOfLocation)
	err := c.cc.Invoke(ctx, "/user.CarrierService/GetLocations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carrierServiceClient) CreateLocation(ctx context.Context, in *Location, opts ...grpc.CallOption) (*Location, error) {
	out := new(Location)
	err := c.cc.Invoke(ctx, "/user.CarrierService/CreateLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carrierServiceClient) UpdateLocation(ctx context.Context, in *Location, opts ...grpc.CallOption) (*Ok, error) {
	out := new(Ok)
	err := c.cc.Invoke(ctx, "/user.CarrierService/UpdateLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carrierServiceClient) DeleteLocation(ctx context.Context, in *Location, opts ...grpc.CallOption) (*Ok, error) {
	out := new(Ok)
	err := c.cc.Invoke(ctx, "/user.CarrierService/DeleteLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carrierServiceClient) PingCarrierService(ctx context.Context, in *CarrierServicePing, opts ...grpc.CallOption) (*CarrierServicePing, error) {
	out := new(CarrierServicePing)
	err := c.cc.Invoke(ctx, "/user.CarrierService/PingCarrierService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carrierServiceClient) GetAllQuotes(ctx context.Context, in *FetchQuotes, opts ...grpc.CallOption) (*QuoteEntities, error) {
	out := new(QuoteEntities)
	err := c.cc.Invoke(ctx, "/user.CarrierService/GetAllQuotes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carrierServiceClient) GetNewQuotes(ctx context.Context, in *QuoteEntity, opts ...grpc.CallOption) (*QuoteBids, error) {
	out := new(QuoteBids)
	err := c.cc.Invoke(ctx, "/user.CarrierService/GetNewQuotes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carrierServiceClient) UpdateQuote(ctx context.Context, in *QuoteEntity, opts ...grpc.CallOption) (*QuoteBids, error) {
	out := new(QuoteBids)
	err := c.cc.Invoke(ctx, "/user.CarrierService/UpdateQuote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carrierServiceClient) DeleteQuote(ctx context.Context, in *DeleteQuote, opts ...grpc.CallOption) (*Ok, error) {
	out := new(Ok)
	err := c.cc.Invoke(ctx, "/user.CarrierService/DeleteQuote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carrierServiceClient) BookQuote(ctx context.Context, in *BookingData, opts ...grpc.CallOption) (*BookingResponse, error) {
	out := new(BookingResponse)
	err := c.cc.Invoke(ctx, "/user.CarrierService/BookQuote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carrierServiceClient) GetBookingHistory(ctx context.Context, in *FetchBookingsRequest, opts ...grpc.CallOption) (*ListOfBooking, error) {
	out := new(ListOfBooking)
	err := c.cc.Invoke(ctx, "/user.CarrierService/GetBookingHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carrierServiceClient) GetBookingById(ctx context.Context, in *FetchBookingsRequest, opts ...grpc.CallOption) (*ListOfBooking, error) {
	out := new(ListOfBooking)
	err := c.cc.Invoke(ctx, "/user.CarrierService/GetBookingById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CarrierServiceServer is the server API for CarrierService service.
// All implementations should embed UnimplementedCarrierServiceServer
// for forward compatibility
type CarrierServiceServer interface {
	// get business
	GetBusinessById(context.Context, *BusinessRequest) (*DynamoBusiness, error)
	GetAllBusinesses(context.Context, *AllBusinessRequest) (*Ok, error)
	// location
	GetLocations(context.Context, *BusinessId) (*ListsOfLocation, error)
	CreateLocation(context.Context, *Location) (*Location, error)
	UpdateLocation(context.Context, *Location) (*Ok, error)
	DeleteLocation(context.Context, *Location) (*Ok, error)
	// health
	PingCarrierService(context.Context, *CarrierServicePing) (*CarrierServicePing, error)
	// quote
	GetAllQuotes(context.Context, *FetchQuotes) (*QuoteEntities, error)
	GetNewQuotes(context.Context, *QuoteEntity) (*QuoteBids, error)
	UpdateQuote(context.Context, *QuoteEntity) (*QuoteBids, error)
	DeleteQuote(context.Context, *DeleteQuote) (*Ok, error)
	// booking
	BookQuote(context.Context, *BookingData) (*BookingResponse, error)
	GetBookingHistory(context.Context, *FetchBookingsRequest) (*ListOfBooking, error)
	GetBookingById(context.Context, *FetchBookingsRequest) (*ListOfBooking, error)
}

// UnimplementedCarrierServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCarrierServiceServer struct {
}

func (UnimplementedCarrierServiceServer) GetBusinessById(context.Context, *BusinessRequest) (*DynamoBusiness, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBusinessById not implemented")
}
func (UnimplementedCarrierServiceServer) GetAllBusinesses(context.Context, *AllBusinessRequest) (*Ok, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllBusinesses not implemented")
}
func (UnimplementedCarrierServiceServer) GetLocations(context.Context, *BusinessId) (*ListsOfLocation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocations not implemented")
}
func (UnimplementedCarrierServiceServer) CreateLocation(context.Context, *Location) (*Location, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLocation not implemented")
}
func (UnimplementedCarrierServiceServer) UpdateLocation(context.Context, *Location) (*Ok, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLocation not implemented")
}
func (UnimplementedCarrierServiceServer) DeleteLocation(context.Context, *Location) (*Ok, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLocation not implemented")
}
func (UnimplementedCarrierServiceServer) PingCarrierService(context.Context, *CarrierServicePing) (*CarrierServicePing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingCarrierService not implemented")
}
func (UnimplementedCarrierServiceServer) GetAllQuotes(context.Context, *FetchQuotes) (*QuoteEntities, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllQuotes not implemented")
}
func (UnimplementedCarrierServiceServer) GetNewQuotes(context.Context, *QuoteEntity) (*QuoteBids, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNewQuotes not implemented")
}
func (UnimplementedCarrierServiceServer) UpdateQuote(context.Context, *QuoteEntity) (*QuoteBids, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQuote not implemented")
}
func (UnimplementedCarrierServiceServer) DeleteQuote(context.Context, *DeleteQuote) (*Ok, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQuote not implemented")
}
func (UnimplementedCarrierServiceServer) BookQuote(context.Context, *BookingData) (*BookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BookQuote not implemented")
}
func (UnimplementedCarrierServiceServer) GetBookingHistory(context.Context, *FetchBookingsRequest) (*ListOfBooking, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookingHistory not implemented")
}
func (UnimplementedCarrierServiceServer) GetBookingById(context.Context, *FetchBookingsRequest) (*ListOfBooking, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookingById not implemented")
}

// UnsafeCarrierServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CarrierServiceServer will
// result in compilation errors.
type UnsafeCarrierServiceServer interface {
	mustEmbedUnimplementedCarrierServiceServer()
}

func RegisterCarrierServiceServer(s grpc.ServiceRegistrar, srv CarrierServiceServer) {
	s.RegisterService(&CarrierService_ServiceDesc, srv)
}

func _CarrierService_GetBusinessById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BusinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarrierServiceServer).GetBusinessById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.CarrierService/GetBusinessById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarrierServiceServer).GetBusinessById(ctx, req.(*BusinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarrierService_GetAllBusinesses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllBusinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarrierServiceServer).GetAllBusinesses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.CarrierService/GetAllBusinesses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarrierServiceServer).GetAllBusinesses(ctx, req.(*AllBusinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarrierService_GetLocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BusinessId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarrierServiceServer).GetLocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.CarrierService/GetLocations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarrierServiceServer).GetLocations(ctx, req.(*BusinessId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarrierService_CreateLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Location)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarrierServiceServer).CreateLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.CarrierService/CreateLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarrierServiceServer).CreateLocation(ctx, req.(*Location))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarrierService_UpdateLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Location)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarrierServiceServer).UpdateLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.CarrierService/UpdateLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarrierServiceServer).UpdateLocation(ctx, req.(*Location))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarrierService_DeleteLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Location)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarrierServiceServer).DeleteLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.CarrierService/DeleteLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarrierServiceServer).DeleteLocation(ctx, req.(*Location))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarrierService_PingCarrierService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CarrierServicePing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarrierServiceServer).PingCarrierService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.CarrierService/PingCarrierService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarrierServiceServer).PingCarrierService(ctx, req.(*CarrierServicePing))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarrierService_GetAllQuotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchQuotes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarrierServiceServer).GetAllQuotes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.CarrierService/GetAllQuotes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarrierServiceServer).GetAllQuotes(ctx, req.(*FetchQuotes))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarrierService_GetNewQuotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuoteEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarrierServiceServer).GetNewQuotes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.CarrierService/GetNewQuotes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarrierServiceServer).GetNewQuotes(ctx, req.(*QuoteEntity))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarrierService_UpdateQuote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuoteEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarrierServiceServer).UpdateQuote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.CarrierService/UpdateQuote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarrierServiceServer).UpdateQuote(ctx, req.(*QuoteEntity))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarrierService_DeleteQuote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteQuote)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarrierServiceServer).DeleteQuote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.CarrierService/DeleteQuote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarrierServiceServer).DeleteQuote(ctx, req.(*DeleteQuote))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarrierService_BookQuote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookingData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarrierServiceServer).BookQuote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.CarrierService/BookQuote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarrierServiceServer).BookQuote(ctx, req.(*BookingData))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarrierService_GetBookingHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchBookingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarrierServiceServer).GetBookingHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.CarrierService/GetBookingHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarrierServiceServer).GetBookingHistory(ctx, req.(*FetchBookingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarrierService_GetBookingById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchBookingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarrierServiceServer).GetBookingById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.CarrierService/GetBookingById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarrierServiceServer).GetBookingById(ctx, req.(*FetchBookingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CarrierService_ServiceDesc is the grpc.ServiceDesc for CarrierService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CarrierService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.CarrierService",
	HandlerType: (*CarrierServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBusinessById",
			Handler:    _CarrierService_GetBusinessById_Handler,
		},
		{
			MethodName: "GetAllBusinesses",
			Handler:    _CarrierService_GetAllBusinesses_Handler,
		},
		{
			MethodName: "GetLocations",
			Handler:    _CarrierService_GetLocations_Handler,
		},
		{
			MethodName: "CreateLocation",
			Handler:    _CarrierService_CreateLocation_Handler,
		},
		{
			MethodName: "UpdateLocation",
			Handler:    _CarrierService_UpdateLocation_Handler,
		},
		{
			MethodName: "DeleteLocation",
			Handler:    _CarrierService_DeleteLocation_Handler,
		},
		{
			MethodName: "PingCarrierService",
			Handler:    _CarrierService_PingCarrierService_Handler,
		},
		{
			MethodName: "GetAllQuotes",
			Handler:    _CarrierService_GetAllQuotes_Handler,
		},
		{
			MethodName: "GetNewQuotes",
			Handler:    _CarrierService_GetNewQuotes_Handler,
		},
		{
			MethodName: "UpdateQuote",
			Handler:    _CarrierService_UpdateQuote_Handler,
		},
		{
			MethodName: "DeleteQuote",
			Handler:    _CarrierService_DeleteQuote_Handler,
		},
		{
			MethodName: "BookQuote",
			Handler:    _CarrierService_BookQuote_Handler,
		},
		{
			MethodName: "GetBookingHistory",
			Handler:    _CarrierService_GetBookingHistory_Handler,
		},
		{
			MethodName: "GetBookingById",
			Handler:    _CarrierService_GetBookingById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "carrier_service.proto",
}
