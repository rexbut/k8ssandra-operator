// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: pkg/medusa/medusa.proto

package medusa

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

const (
	Medusa_Backup_FullMethodName         = "/Medusa/Backup"
	Medusa_AsyncBackup_FullMethodName    = "/Medusa/AsyncBackup"
	Medusa_BackupStatus_FullMethodName   = "/Medusa/BackupStatus"
	Medusa_DeleteBackup_FullMethodName   = "/Medusa/DeleteBackup"
	Medusa_GetBackups_FullMethodName     = "/Medusa/GetBackups"
	Medusa_PurgeBackups_FullMethodName   = "/Medusa/PurgeBackups"
	Medusa_PrepareRestore_FullMethodName = "/Medusa/PrepareRestore"
)

// MedusaClient is the client API for Medusa service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MedusaClient interface {
	Backup(ctx context.Context, in *BackupRequest, opts ...grpc.CallOption) (*BackupResponse, error)
	AsyncBackup(ctx context.Context, in *BackupRequest, opts ...grpc.CallOption) (*BackupResponse, error)
	BackupStatus(ctx context.Context, in *BackupStatusRequest, opts ...grpc.CallOption) (*BackupStatusResponse, error)
	DeleteBackup(ctx context.Context, in *DeleteBackupRequest, opts ...grpc.CallOption) (*DeleteBackupResponse, error)
	GetBackups(ctx context.Context, in *GetBackupsRequest, opts ...grpc.CallOption) (*GetBackupsResponse, error)
	PurgeBackups(ctx context.Context, in *PurgeBackupsRequest, opts ...grpc.CallOption) (*PurgeBackupsResponse, error)
	PrepareRestore(ctx context.Context, in *PrepareRestoreRequest, opts ...grpc.CallOption) (*PrepareRestoreResponse, error)
}

type medusaClient struct {
	cc grpc.ClientConnInterface
}

func NewMedusaClient(cc grpc.ClientConnInterface) MedusaClient {
	return &medusaClient{cc}
}

func (c *medusaClient) Backup(ctx context.Context, in *BackupRequest, opts ...grpc.CallOption) (*BackupResponse, error) {
	out := new(BackupResponse)
	err := c.cc.Invoke(ctx, Medusa_Backup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medusaClient) AsyncBackup(ctx context.Context, in *BackupRequest, opts ...grpc.CallOption) (*BackupResponse, error) {
	out := new(BackupResponse)
	err := c.cc.Invoke(ctx, Medusa_AsyncBackup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medusaClient) BackupStatus(ctx context.Context, in *BackupStatusRequest, opts ...grpc.CallOption) (*BackupStatusResponse, error) {
	out := new(BackupStatusResponse)
	err := c.cc.Invoke(ctx, Medusa_BackupStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medusaClient) DeleteBackup(ctx context.Context, in *DeleteBackupRequest, opts ...grpc.CallOption) (*DeleteBackupResponse, error) {
	out := new(DeleteBackupResponse)
	err := c.cc.Invoke(ctx, Medusa_DeleteBackup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medusaClient) GetBackups(ctx context.Context, in *GetBackupsRequest, opts ...grpc.CallOption) (*GetBackupsResponse, error) {
	out := new(GetBackupsResponse)
	err := c.cc.Invoke(ctx, Medusa_GetBackups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medusaClient) PurgeBackups(ctx context.Context, in *PurgeBackupsRequest, opts ...grpc.CallOption) (*PurgeBackupsResponse, error) {
	out := new(PurgeBackupsResponse)
	err := c.cc.Invoke(ctx, Medusa_PurgeBackups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medusaClient) PrepareRestore(ctx context.Context, in *PrepareRestoreRequest, opts ...grpc.CallOption) (*PrepareRestoreResponse, error) {
	out := new(PrepareRestoreResponse)
	err := c.cc.Invoke(ctx, Medusa_PrepareRestore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MedusaServer is the server API for Medusa service.
// All implementations must embed UnimplementedMedusaServer
// for forward compatibility
type MedusaServer interface {
	Backup(context.Context, *BackupRequest) (*BackupResponse, error)
	AsyncBackup(context.Context, *BackupRequest) (*BackupResponse, error)
	BackupStatus(context.Context, *BackupStatusRequest) (*BackupStatusResponse, error)
	DeleteBackup(context.Context, *DeleteBackupRequest) (*DeleteBackupResponse, error)
	GetBackups(context.Context, *GetBackupsRequest) (*GetBackupsResponse, error)
	PurgeBackups(context.Context, *PurgeBackupsRequest) (*PurgeBackupsResponse, error)
	PrepareRestore(context.Context, *PrepareRestoreRequest) (*PrepareRestoreResponse, error)
	mustEmbedUnimplementedMedusaServer()
}

// UnimplementedMedusaServer must be embedded to have forward compatible implementations.
type UnimplementedMedusaServer struct {
}

func (UnimplementedMedusaServer) Backup(context.Context, *BackupRequest) (*BackupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Backup not implemented")
}
func (UnimplementedMedusaServer) AsyncBackup(context.Context, *BackupRequest) (*BackupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AsyncBackup not implemented")
}
func (UnimplementedMedusaServer) BackupStatus(context.Context, *BackupStatusRequest) (*BackupStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BackupStatus not implemented")
}
func (UnimplementedMedusaServer) DeleteBackup(context.Context, *DeleteBackupRequest) (*DeleteBackupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBackup not implemented")
}
func (UnimplementedMedusaServer) GetBackups(context.Context, *GetBackupsRequest) (*GetBackupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBackups not implemented")
}
func (UnimplementedMedusaServer) PurgeBackups(context.Context, *PurgeBackupsRequest) (*PurgeBackupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PurgeBackups not implemented")
}
func (UnimplementedMedusaServer) PrepareRestore(context.Context, *PrepareRestoreRequest) (*PrepareRestoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrepareRestore not implemented")
}
func (UnimplementedMedusaServer) mustEmbedUnimplementedMedusaServer() {}

// UnsafeMedusaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MedusaServer will
// result in compilation errors.
type UnsafeMedusaServer interface {
	mustEmbedUnimplementedMedusaServer()
}

func RegisterMedusaServer(s grpc.ServiceRegistrar, srv MedusaServer) {
	s.RegisterService(&Medusa_ServiceDesc, srv)
}

func _Medusa_Backup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BackupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedusaServer).Backup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Medusa_Backup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedusaServer).Backup(ctx, req.(*BackupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Medusa_AsyncBackup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BackupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedusaServer).AsyncBackup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Medusa_AsyncBackup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedusaServer).AsyncBackup(ctx, req.(*BackupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Medusa_BackupStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BackupStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedusaServer).BackupStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Medusa_BackupStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedusaServer).BackupStatus(ctx, req.(*BackupStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Medusa_DeleteBackup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBackupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedusaServer).DeleteBackup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Medusa_DeleteBackup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedusaServer).DeleteBackup(ctx, req.(*DeleteBackupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Medusa_GetBackups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBackupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedusaServer).GetBackups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Medusa_GetBackups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedusaServer).GetBackups(ctx, req.(*GetBackupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Medusa_PurgeBackups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurgeBackupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedusaServer).PurgeBackups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Medusa_PurgeBackups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedusaServer).PurgeBackups(ctx, req.(*PurgeBackupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Medusa_PrepareRestore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareRestoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedusaServer).PrepareRestore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Medusa_PrepareRestore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedusaServer).PrepareRestore(ctx, req.(*PrepareRestoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Medusa_ServiceDesc is the grpc.ServiceDesc for Medusa service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Medusa_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Medusa",
	HandlerType: (*MedusaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Backup",
			Handler:    _Medusa_Backup_Handler,
		},
		{
			MethodName: "AsyncBackup",
			Handler:    _Medusa_AsyncBackup_Handler,
		},
		{
			MethodName: "BackupStatus",
			Handler:    _Medusa_BackupStatus_Handler,
		},
		{
			MethodName: "DeleteBackup",
			Handler:    _Medusa_DeleteBackup_Handler,
		},
		{
			MethodName: "GetBackups",
			Handler:    _Medusa_GetBackups_Handler,
		},
		{
			MethodName: "PurgeBackups",
			Handler:    _Medusa_PurgeBackups_Handler,
		},
		{
			MethodName: "PrepareRestore",
			Handler:    _Medusa_PrepareRestore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/medusa/medusa.proto",
}
