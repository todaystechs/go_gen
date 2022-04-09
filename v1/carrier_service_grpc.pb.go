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
	CreateBusiness(ctx context.Context, in *BusinessData, opts ...grpc.CallOption) (*Ok, error)
	GetBusinessById(ctx context.Context, in *GetBusinessRequest, opts ...grpc.CallOption) (*BusinessData, error)
	GetBusinesses(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetBusinessesResponse, error)
	CloseBusiness(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Ok, error)
	// location
	GetLocations(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListsOfLocation, error)
	GetLocation(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Location, error)
	CreateLocation(ctx context.Context, in *Location, opts ...grpc.CallOption) (*Ok, error)
	UpdateLocation(ctx context.Context, in *Location, opts ...grpc.CallOption) (*Ok, error)
	DeleteLocation(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Ok, error)
	DeleteLocations(ctx context.Context, in *Ids, opts ...grpc.CallOption) (*Ok, error)
	// health
	PingCarrierService(ctx context.Context, in *CarrierServicePing, opts ...grpc.CallOption) (*CarrierServicePing, error)
	// booking
	BookAQuote(ctx context.Context, in *BookBids, opts ...grpc.CallOption) (*BookingResponseData, error)
	BookQuotes(ctx context.Context, in *BookBids, opts ...grpc.CallOption) (*ListOfBookingResponse, error)
	GetBookingHistory(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListOfBookingResponse, error)
	GetBookingById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*ListOfBookingResponse, error)
	// quote
	GetAllQuotes(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*QuoteDatas, error)
	GetNewQuotes(ctx context.Context, in *QuoteData, opts ...grpc.CallOption) (*BidDatas, error)
	UpdateQuote(ctx context.Context, in *QuoteData, opts ...grpc.CallOption) (*BidDatas, error)
	DeleteQuote(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Ok, error)
	DeleteQuotes(ctx context.Context, in *Ids, opts ...grpc.CallOption) (*Ok, error)
}

type carrierServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCarrierServiceClient(cc grpc.ClientConnInterface) CarrierServiceClient {
	return &carrierServiceClient{cc}
}

func (c *carrierServiceClient) CreateBusiness(ctx context.Context, in *BusinessData, opts ...grpc.CallOption) (*Ok, error) {
	out := new(Ok)
	err := c.cc.Invoke(ctx, "/user.CarrierService/CreateBusiness", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carrierServiceClient) GetBusinessById(ctx context.Context, in *GetBusinessRequest, opts ...grpc.CallOption) (*BusinessData, error) {
	out := new(BusinessData)
	err := c.cc.Invoke(ctx, "/user.CarrierService/GetBusinessById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carrierServiceClient) GetBusinesses(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetBusinessesResponse, error) {
	out := new(GetBusinessesResponse)
	err := c.cc.Invoke(ctx, "/user.CarrierService/GetBusinesses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carrierServiceClient) CloseBusiness(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Ok, error) {
	out := new(Ok)
	err := c.cc.Invoke(ctx, "/user.CarrierService/CloseBusiness", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carrierServiceClient) GetLocations(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListsOfLocation, error) {
	out := new(ListsOfLocation)
	err := c.cc.Invoke(ctx, "/user.CarrierService/GetLocations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carrierServiceClient) GetLocation(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Location, error) {
	out := new(Location)
	err := c.cc.Invoke(ctx, "/user.CarrierService/GetLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carrierServiceClient) CreateLocation(ctx context.Context, in *Location, opts ...grpc.CallOption) (*Ok, error) {
	out := new(Ok)
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

func (c *carrierServiceClient) DeleteLocation(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Ok, error) {
	out := new(Ok)
	err := c.cc.Invoke(ctx, "/user.CarrierService/DeleteLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carrierServiceClient) DeleteLocations(ctx context.Context, in *Ids, opts ...grpc.CallOption) (*Ok, error) {
	out := new(Ok)
	err := c.cc.Invoke(ctx, "/user.CarrierService/DeleteLocations", in, out, opts...)
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

func (c *carrierServiceClient) BookAQuote(ctx context.Context, in *BookBids, opts ...grpc.CallOption) (*BookingResponseData, error) {
	out := new(BookingResponseData)
	err := c.cc.Invoke(ctx, "/user.CarrierService/BookAQuote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carrierServiceClient) BookQuotes(ctx context.Context, in *BookBids, opts ...grpc.CallOption) (*ListOfBookingResponse, error) {
	out := new(ListOfBookingResponse)
	err := c.cc.Invoke(ctx, "/user.CarrierService/BookQuotes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carrierServiceClient) GetBookingHistory(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListOfBookingResponse, error) {
	out := new(ListOfBookingResponse)
	err := c.cc.Invoke(ctx, "/user.CarrierService/GetBookingHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carrierServiceClient) GetBookingById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*ListOfBookingResponse, error) {
	out := new(ListOfBookingResponse)
	err := c.cc.Invoke(ctx, "/user.CarrierService/GetBookingById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carrierServiceClient) GetAllQuotes(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*QuoteDatas, error) {
	out := new(QuoteDatas)
	err := c.cc.Invoke(ctx, "/user.CarrierService/GetAllQuotes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carrierServiceClient) GetNewQuotes(ctx context.Context, in *QuoteData, opts ...grpc.CallOption) (*BidDatas, error) {
	out := new(BidDatas)
	err := c.cc.Invoke(ctx, "/user.CarrierService/GetNewQuotes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carrierServiceClient) UpdateQuote(ctx context.Context, in *QuoteData, opts ...grpc.CallOption) (*BidDatas, error) {
	out := new(BidDatas)
	err := c.cc.Invoke(ctx, "/user.CarrierService/UpdateQuote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carrierServiceClient) DeleteQuote(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Ok, error) {
	out := new(Ok)
	err := c.cc.Invoke(ctx, "/user.CarrierService/DeleteQuote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carrierServiceClient) DeleteQuotes(ctx context.Context, in *Ids, opts ...grpc.CallOption) (*Ok, error) {
	out := new(Ok)
	err := c.cc.Invoke(ctx, "/user.CarrierService/DeleteQuotes", in, out, opts...)
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
	CreateBusiness(context.Context, *BusinessData) (*Ok, error)
	GetBusinessById(context.Context, *GetBusinessRequest) (*BusinessData, error)
	GetBusinesses(context.Context, *Empty) (*GetBusinessesResponse, error)
	CloseBusiness(context.Context, *Id) (*Ok, error)
	// location
	GetLocations(context.Context, *Empty) (*ListsOfLocation, error)
	GetLocation(context.Context, *Empty) (*Location, error)
	CreateLocation(context.Context, *Location) (*Ok, error)
	UpdateLocation(context.Context, *Location) (*Ok, error)
	DeleteLocation(context.Context, *Id) (*Ok, error)
	DeleteLocations(context.Context, *Ids) (*Ok, error)
	// health
	PingCarrierService(context.Context, *CarrierServicePing) (*CarrierServicePing, error)
	// booking
	BookAQuote(context.Context, *BookBids) (*BookingResponseData, error)
	BookQuotes(context.Context, *BookBids) (*ListOfBookingResponse, error)
	GetBookingHistory(context.Context, *Empty) (*ListOfBookingResponse, error)
	GetBookingById(context.Context, *Id) (*ListOfBookingResponse, error)
	// quote
	GetAllQuotes(context.Context, *Empty) (*QuoteDatas, error)
	GetNewQuotes(context.Context, *QuoteData) (*BidDatas, error)
	UpdateQuote(context.Context, *QuoteData) (*BidDatas, error)
	DeleteQuote(context.Context, *Id) (*Ok, error)
	DeleteQuotes(context.Context, *Ids) (*Ok, error)
}

// UnimplementedCarrierServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCarrierServiceServer struct {
}

func (UnimplementedCarrierServiceServer) CreateBusiness(context.Context, *BusinessData) (*Ok, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBusiness not implemented")
}
func (UnimplementedCarrierServiceServer) GetBusinessById(context.Context, *GetBusinessRequest) (*BusinessData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBusinessById not implemented")
}
func (UnimplementedCarrierServiceServer) GetBusinesses(context.Context, *Empty) (*GetBusinessesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBusinesses not implemented")
}
func (UnimplementedCarrierServiceServer) CloseBusiness(context.Context, *Id) (*Ok, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseBusiness not implemented")
}
func (UnimplementedCarrierServiceServer) GetLocations(context.Context, *Empty) (*ListsOfLocation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocations not implemented")
}
func (UnimplementedCarrierServiceServer) GetLocation(context.Context, *Empty) (*Location, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocation not implemented")
}
func (UnimplementedCarrierServiceServer) CreateLocation(context.Context, *Location) (*Ok, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLocation not implemented")
}
func (UnimplementedCarrierServiceServer) UpdateLocation(context.Context, *Location) (*Ok, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLocation not implemented")
}
func (UnimplementedCarrierServiceServer) DeleteLocation(context.Context, *Id) (*Ok, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLocation not implemented")
}
func (UnimplementedCarrierServiceServer) DeleteLocations(context.Context, *Ids) (*Ok, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLocations not implemented")
}
func (UnimplementedCarrierServiceServer) PingCarrierService(context.Context, *CarrierServicePing) (*CarrierServicePing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingCarrierService not implemented")
}
func (UnimplementedCarrierServiceServer) BookAQuote(context.Context, *BookBids) (*BookingResponseData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BookAQuote not implemented")
}
func (UnimplementedCarrierServiceServer) BookQuotes(context.Context, *BookBids) (*ListOfBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BookQuotes not implemented")
}
func (UnimplementedCarrierServiceServer) GetBookingHistory(context.Context, *Empty) (*ListOfBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookingHistory not implemented")
}
func (UnimplementedCarrierServiceServer) GetBookingById(context.Context, *Id) (*ListOfBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookingById not implemented")
}
func (UnimplementedCarrierServiceServer) GetAllQuotes(context.Context, *Empty) (*QuoteDatas, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllQuotes not implemented")
}
func (UnimplementedCarrierServiceServer) GetNewQuotes(context.Context, *QuoteData) (*BidDatas, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNewQuotes not implemented")
}
func (UnimplementedCarrierServiceServer) UpdateQuote(context.Context, *QuoteData) (*BidDatas, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQuote not implemented")
}
func (UnimplementedCarrierServiceServer) DeleteQuote(context.Context, *Id) (*Ok, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQuote not implemented")
}
func (UnimplementedCarrierServiceServer) DeleteQuotes(context.Context, *Ids) (*Ok, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQuotes not implemented")
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

func _CarrierService_CreateBusiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BusinessData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarrierServiceServer).CreateBusiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.CarrierService/CreateBusiness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarrierServiceServer).CreateBusiness(ctx, req.(*BusinessData))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarrierService_GetBusinessById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBusinessRequest)
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
		return srv.(CarrierServiceServer).GetBusinessById(ctx, req.(*GetBusinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarrierService_GetBusinesses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarrierServiceServer).GetBusinesses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.CarrierService/GetBusinesses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarrierServiceServer).GetBusinesses(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarrierService_CloseBusiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarrierServiceServer).CloseBusiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.CarrierService/CloseBusiness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarrierServiceServer).CloseBusiness(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarrierService_GetLocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
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
		return srv.(CarrierServiceServer).GetLocations(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarrierService_GetLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarrierServiceServer).GetLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.CarrierService/GetLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarrierServiceServer).GetLocation(ctx, req.(*Empty))
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
	in := new(Id)
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
		return srv.(CarrierServiceServer).DeleteLocation(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarrierService_DeleteLocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ids)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarrierServiceServer).DeleteLocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.CarrierService/DeleteLocations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarrierServiceServer).DeleteLocations(ctx, req.(*Ids))
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

func _CarrierService_BookAQuote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookBids)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarrierServiceServer).BookAQuote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.CarrierService/BookAQuote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarrierServiceServer).BookAQuote(ctx, req.(*BookBids))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarrierService_BookQuotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookBids)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarrierServiceServer).BookQuotes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.CarrierService/BookQuotes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarrierServiceServer).BookQuotes(ctx, req.(*BookBids))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarrierService_GetBookingHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
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
		return srv.(CarrierServiceServer).GetBookingHistory(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarrierService_GetBookingById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
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
		return srv.(CarrierServiceServer).GetBookingById(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarrierService_GetAllQuotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
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
		return srv.(CarrierServiceServer).GetAllQuotes(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarrierService_GetNewQuotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuoteData)
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
		return srv.(CarrierServiceServer).GetNewQuotes(ctx, req.(*QuoteData))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarrierService_UpdateQuote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuoteData)
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
		return srv.(CarrierServiceServer).UpdateQuote(ctx, req.(*QuoteData))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarrierService_DeleteQuote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
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
		return srv.(CarrierServiceServer).DeleteQuote(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarrierService_DeleteQuotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ids)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarrierServiceServer).DeleteQuotes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.CarrierService/DeleteQuotes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarrierServiceServer).DeleteQuotes(ctx, req.(*Ids))
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
			MethodName: "CreateBusiness",
			Handler:    _CarrierService_CreateBusiness_Handler,
		},
		{
			MethodName: "GetBusinessById",
			Handler:    _CarrierService_GetBusinessById_Handler,
		},
		{
			MethodName: "GetBusinesses",
			Handler:    _CarrierService_GetBusinesses_Handler,
		},
		{
			MethodName: "CloseBusiness",
			Handler:    _CarrierService_CloseBusiness_Handler,
		},
		{
			MethodName: "GetLocations",
			Handler:    _CarrierService_GetLocations_Handler,
		},
		{
			MethodName: "GetLocation",
			Handler:    _CarrierService_GetLocation_Handler,
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
			MethodName: "DeleteLocations",
			Handler:    _CarrierService_DeleteLocations_Handler,
		},
		{
			MethodName: "PingCarrierService",
			Handler:    _CarrierService_PingCarrierService_Handler,
		},
		{
			MethodName: "BookAQuote",
			Handler:    _CarrierService_BookAQuote_Handler,
		},
		{
			MethodName: "BookQuotes",
			Handler:    _CarrierService_BookQuotes_Handler,
		},
		{
			MethodName: "GetBookingHistory",
			Handler:    _CarrierService_GetBookingHistory_Handler,
		},
		{
			MethodName: "GetBookingById",
			Handler:    _CarrierService_GetBookingById_Handler,
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
			MethodName: "DeleteQuotes",
			Handler:    _CarrierService_DeleteQuotes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "carrier_service.proto",
}
