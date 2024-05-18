// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: main.proto

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

const (
	TeacherService_RegisterTeacher_FullMethodName = "/proto.TeacherService/RegisterTeacher"
	TeacherService_AddSession_FullMethodName      = "/proto.TeacherService/AddSession"
	TeacherService_GetSession_FullMethodName      = "/proto.TeacherService/GetSession"
)

// TeacherServiceClient is the client API for TeacherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TeacherServiceClient interface {
	RegisterTeacher(ctx context.Context, in *RegisterTeacherRequest, opts ...grpc.CallOption) (*RegisterTeacherResponse, error)
	AddSession(ctx context.Context, in *AddSessionRequest, opts ...grpc.CallOption) (*AddSessionResponse, error)
	GetSession(ctx context.Context, in *GetSessionRequest, opts ...grpc.CallOption) (*GetSessionResponse, error)
}

type teacherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTeacherServiceClient(cc grpc.ClientConnInterface) TeacherServiceClient {
	return &teacherServiceClient{cc}
}

func (c *teacherServiceClient) RegisterTeacher(ctx context.Context, in *RegisterTeacherRequest, opts ...grpc.CallOption) (*RegisterTeacherResponse, error) {
	out := new(RegisterTeacherResponse)
	err := c.cc.Invoke(ctx, TeacherService_RegisterTeacher_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teacherServiceClient) AddSession(ctx context.Context, in *AddSessionRequest, opts ...grpc.CallOption) (*AddSessionResponse, error) {
	out := new(AddSessionResponse)
	err := c.cc.Invoke(ctx, TeacherService_AddSession_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teacherServiceClient) GetSession(ctx context.Context, in *GetSessionRequest, opts ...grpc.CallOption) (*GetSessionResponse, error) {
	out := new(GetSessionResponse)
	err := c.cc.Invoke(ctx, TeacherService_GetSession_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TeacherServiceServer is the server API for TeacherService service.
// All implementations must embed UnimplementedTeacherServiceServer
// for forward compatibility
type TeacherServiceServer interface {
	RegisterTeacher(context.Context, *RegisterTeacherRequest) (*RegisterTeacherResponse, error)
	AddSession(context.Context, *AddSessionRequest) (*AddSessionResponse, error)
	GetSession(context.Context, *GetSessionRequest) (*GetSessionResponse, error)
	mustEmbedUnimplementedTeacherServiceServer()
}

// UnimplementedTeacherServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTeacherServiceServer struct {
}

func (UnimplementedTeacherServiceServer) RegisterTeacher(context.Context, *RegisterTeacherRequest) (*RegisterTeacherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterTeacher not implemented")
}
func (UnimplementedTeacherServiceServer) AddSession(context.Context, *AddSessionRequest) (*AddSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSession not implemented")
}
func (UnimplementedTeacherServiceServer) GetSession(context.Context, *GetSessionRequest) (*GetSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSession not implemented")
}
func (UnimplementedTeacherServiceServer) mustEmbedUnimplementedTeacherServiceServer() {}

// UnsafeTeacherServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TeacherServiceServer will
// result in compilation errors.
type UnsafeTeacherServiceServer interface {
	mustEmbedUnimplementedTeacherServiceServer()
}

func RegisterTeacherServiceServer(s grpc.ServiceRegistrar, srv TeacherServiceServer) {
	s.RegisterService(&TeacherService_ServiceDesc, srv)
}

func _TeacherService_RegisterTeacher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterTeacherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeacherServiceServer).RegisterTeacher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeacherService_RegisterTeacher_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeacherServiceServer).RegisterTeacher(ctx, req.(*RegisterTeacherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeacherService_AddSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeacherServiceServer).AddSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeacherService_AddSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeacherServiceServer).AddSession(ctx, req.(*AddSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeacherService_GetSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeacherServiceServer).GetSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeacherService_GetSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeacherServiceServer).GetSession(ctx, req.(*GetSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TeacherService_ServiceDesc is the grpc.ServiceDesc for TeacherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TeacherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TeacherService",
	HandlerType: (*TeacherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterTeacher",
			Handler:    _TeacherService_RegisterTeacher_Handler,
		},
		{
			MethodName: "AddSession",
			Handler:    _TeacherService_AddSession_Handler,
		},
		{
			MethodName: "GetSession",
			Handler:    _TeacherService_GetSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "main.proto",
}

const (
	StudentService_AddAttendance_FullMethodName   = "/proto.StudentService/AddAttendance"
	StudentService_AddStudent_FullMethodName      = "/proto.StudentService/AddStudent"
	StudentService_GetStudentNames_FullMethodName = "/proto.StudentService/GetStudentNames"
	StudentService_UploadImages_FullMethodName    = "/proto.StudentService/UploadImages"
	StudentService_GetImages_FullMethodName       = "/proto.StudentService/GetImages"
)

// StudentServiceClient is the client API for StudentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StudentServiceClient interface {
	AddAttendance(ctx context.Context, in *AttendanceRequest, opts ...grpc.CallOption) (*AttendanceResponse, error)
	AddStudent(ctx context.Context, in *AddStudentRequest, opts ...grpc.CallOption) (*AddStudentResponse, error)
	GetStudentNames(ctx context.Context, in *GetStudentNamesRequest, opts ...grpc.CallOption) (*GetStudentNamesResponse, error)
	UploadImages(ctx context.Context, in *UploadImagesRequest, opts ...grpc.CallOption) (*UploadImagesResponse, error)
	GetImages(ctx context.Context, in *GetImagesRequest, opts ...grpc.CallOption) (*GetImagesResponse, error)
}

type studentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStudentServiceClient(cc grpc.ClientConnInterface) StudentServiceClient {
	return &studentServiceClient{cc}
}

func (c *studentServiceClient) AddAttendance(ctx context.Context, in *AttendanceRequest, opts ...grpc.CallOption) (*AttendanceResponse, error) {
	out := new(AttendanceResponse)
	err := c.cc.Invoke(ctx, StudentService_AddAttendance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) AddStudent(ctx context.Context, in *AddStudentRequest, opts ...grpc.CallOption) (*AddStudentResponse, error) {
	out := new(AddStudentResponse)
	err := c.cc.Invoke(ctx, StudentService_AddStudent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) GetStudentNames(ctx context.Context, in *GetStudentNamesRequest, opts ...grpc.CallOption) (*GetStudentNamesResponse, error) {
	out := new(GetStudentNamesResponse)
	err := c.cc.Invoke(ctx, StudentService_GetStudentNames_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) UploadImages(ctx context.Context, in *UploadImagesRequest, opts ...grpc.CallOption) (*UploadImagesResponse, error) {
	out := new(UploadImagesResponse)
	err := c.cc.Invoke(ctx, StudentService_UploadImages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) GetImages(ctx context.Context, in *GetImagesRequest, opts ...grpc.CallOption) (*GetImagesResponse, error) {
	out := new(GetImagesResponse)
	err := c.cc.Invoke(ctx, StudentService_GetImages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudentServiceServer is the server API for StudentService service.
// All implementations must embed UnimplementedStudentServiceServer
// for forward compatibility
type StudentServiceServer interface {
	AddAttendance(context.Context, *AttendanceRequest) (*AttendanceResponse, error)
	AddStudent(context.Context, *AddStudentRequest) (*AddStudentResponse, error)
	GetStudentNames(context.Context, *GetStudentNamesRequest) (*GetStudentNamesResponse, error)
	UploadImages(context.Context, *UploadImagesRequest) (*UploadImagesResponse, error)
	GetImages(context.Context, *GetImagesRequest) (*GetImagesResponse, error)
	mustEmbedUnimplementedStudentServiceServer()
}

// UnimplementedStudentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStudentServiceServer struct {
}

func (UnimplementedStudentServiceServer) AddAttendance(context.Context, *AttendanceRequest) (*AttendanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAttendance not implemented")
}
func (UnimplementedStudentServiceServer) AddStudent(context.Context, *AddStudentRequest) (*AddStudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddStudent not implemented")
}
func (UnimplementedStudentServiceServer) GetStudentNames(context.Context, *GetStudentNamesRequest) (*GetStudentNamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudentNames not implemented")
}
func (UnimplementedStudentServiceServer) UploadImages(context.Context, *UploadImagesRequest) (*UploadImagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadImages not implemented")
}
func (UnimplementedStudentServiceServer) GetImages(context.Context, *GetImagesRequest) (*GetImagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImages not implemented")
}
func (UnimplementedStudentServiceServer) mustEmbedUnimplementedStudentServiceServer() {}

// UnsafeStudentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StudentServiceServer will
// result in compilation errors.
type UnsafeStudentServiceServer interface {
	mustEmbedUnimplementedStudentServiceServer()
}

func RegisterStudentServiceServer(s grpc.ServiceRegistrar, srv StudentServiceServer) {
	s.RegisterService(&StudentService_ServiceDesc, srv)
}

func _StudentService_AddAttendance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttendanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).AddAttendance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_AddAttendance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).AddAttendance(ctx, req.(*AttendanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_AddStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).AddStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_AddStudent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).AddStudent(ctx, req.(*AddStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_GetStudentNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStudentNamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).GetStudentNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_GetStudentNames_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).GetStudentNames(ctx, req.(*GetStudentNamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_UploadImages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadImagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).UploadImages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_UploadImages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).UploadImages(ctx, req.(*UploadImagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_GetImages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetImagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).GetImages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_GetImages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).GetImages(ctx, req.(*GetImagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StudentService_ServiceDesc is the grpc.ServiceDesc for StudentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StudentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.StudentService",
	HandlerType: (*StudentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddAttendance",
			Handler:    _StudentService_AddAttendance_Handler,
		},
		{
			MethodName: "AddStudent",
			Handler:    _StudentService_AddStudent_Handler,
		},
		{
			MethodName: "GetStudentNames",
			Handler:    _StudentService_GetStudentNames_Handler,
		},
		{
			MethodName: "UploadImages",
			Handler:    _StudentService_UploadImages_Handler,
		},
		{
			MethodName: "GetImages",
			Handler:    _StudentService_GetImages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "main.proto",
}

const (
	AdminService_AddStudentToCourse_FullMethodName = "/proto.AdminService/AddStudentToCourse"
	AdminService_DeleteStudent_FullMethodName      = "/proto.AdminService/DeleteStudent"
	AdminService_DeleteTeacher_FullMethodName      = "/proto.AdminService/DeleteTeacher"
	AdminService_AddCourse_FullMethodName          = "/proto.AdminService/AddCourse"
	AdminService_DeleteCourse_FullMethodName       = "/proto.AdminService/DeleteCourse"
	AdminService_GetCourses_FullMethodName         = "/proto.AdminService/GetCourses"
	AdminService_GetStudents_FullMethodName        = "/proto.AdminService/GetStudents"
	AdminService_GetTeachers_FullMethodName        = "/proto.AdminService/GetTeachers"
)

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminServiceClient interface {
	AddStudentToCourse(ctx context.Context, in *AddStudentToCourseRequest, opts ...grpc.CallOption) (*AddStudentToCourseResponse, error)
	DeleteStudent(ctx context.Context, in *DeleteStudentRequest, opts ...grpc.CallOption) (*DeleteStudentResponse, error)
	DeleteTeacher(ctx context.Context, in *DeleteTeacherRequest, opts ...grpc.CallOption) (*DeleteTeacherResponse, error)
	AddCourse(ctx context.Context, in *AddCourseRequest, opts ...grpc.CallOption) (*AddCourseResponse, error)
	DeleteCourse(ctx context.Context, in *DeleteCourseRequest, opts ...grpc.CallOption) (*DeleteCourseResponse, error)
	GetCourses(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetCoursesResponse, error)
	GetStudents(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetStudentsResponse, error)
	GetTeachers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetTeachersResponse, error)
}

type adminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminServiceClient(cc grpc.ClientConnInterface) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) AddStudentToCourse(ctx context.Context, in *AddStudentToCourseRequest, opts ...grpc.CallOption) (*AddStudentToCourseResponse, error) {
	out := new(AddStudentToCourseResponse)
	err := c.cc.Invoke(ctx, AdminService_AddStudentToCourse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) DeleteStudent(ctx context.Context, in *DeleteStudentRequest, opts ...grpc.CallOption) (*DeleteStudentResponse, error) {
	out := new(DeleteStudentResponse)
	err := c.cc.Invoke(ctx, AdminService_DeleteStudent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) DeleteTeacher(ctx context.Context, in *DeleteTeacherRequest, opts ...grpc.CallOption) (*DeleteTeacherResponse, error) {
	out := new(DeleteTeacherResponse)
	err := c.cc.Invoke(ctx, AdminService_DeleteTeacher_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) AddCourse(ctx context.Context, in *AddCourseRequest, opts ...grpc.CallOption) (*AddCourseResponse, error) {
	out := new(AddCourseResponse)
	err := c.cc.Invoke(ctx, AdminService_AddCourse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) DeleteCourse(ctx context.Context, in *DeleteCourseRequest, opts ...grpc.CallOption) (*DeleteCourseResponse, error) {
	out := new(DeleteCourseResponse)
	err := c.cc.Invoke(ctx, AdminService_DeleteCourse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) GetCourses(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetCoursesResponse, error) {
	out := new(GetCoursesResponse)
	err := c.cc.Invoke(ctx, AdminService_GetCourses_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) GetStudents(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetStudentsResponse, error) {
	out := new(GetStudentsResponse)
	err := c.cc.Invoke(ctx, AdminService_GetStudents_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) GetTeachers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetTeachersResponse, error) {
	out := new(GetTeachersResponse)
	err := c.cc.Invoke(ctx, AdminService_GetTeachers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
// All implementations must embed UnimplementedAdminServiceServer
// for forward compatibility
type AdminServiceServer interface {
	AddStudentToCourse(context.Context, *AddStudentToCourseRequest) (*AddStudentToCourseResponse, error)
	DeleteStudent(context.Context, *DeleteStudentRequest) (*DeleteStudentResponse, error)
	DeleteTeacher(context.Context, *DeleteTeacherRequest) (*DeleteTeacherResponse, error)
	AddCourse(context.Context, *AddCourseRequest) (*AddCourseResponse, error)
	DeleteCourse(context.Context, *DeleteCourseRequest) (*DeleteCourseResponse, error)
	GetCourses(context.Context, *Empty) (*GetCoursesResponse, error)
	GetStudents(context.Context, *Empty) (*GetStudentsResponse, error)
	GetTeachers(context.Context, *Empty) (*GetTeachersResponse, error)
	mustEmbedUnimplementedAdminServiceServer()
}

// UnimplementedAdminServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServiceServer struct {
}

func (UnimplementedAdminServiceServer) AddStudentToCourse(context.Context, *AddStudentToCourseRequest) (*AddStudentToCourseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddStudentToCourse not implemented")
}
func (UnimplementedAdminServiceServer) DeleteStudent(context.Context, *DeleteStudentRequest) (*DeleteStudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStudent not implemented")
}
func (UnimplementedAdminServiceServer) DeleteTeacher(context.Context, *DeleteTeacherRequest) (*DeleteTeacherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTeacher not implemented")
}
func (UnimplementedAdminServiceServer) AddCourse(context.Context, *AddCourseRequest) (*AddCourseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCourse not implemented")
}
func (UnimplementedAdminServiceServer) DeleteCourse(context.Context, *DeleteCourseRequest) (*DeleteCourseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCourse not implemented")
}
func (UnimplementedAdminServiceServer) GetCourses(context.Context, *Empty) (*GetCoursesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourses not implemented")
}
func (UnimplementedAdminServiceServer) GetStudents(context.Context, *Empty) (*GetStudentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudents not implemented")
}
func (UnimplementedAdminServiceServer) GetTeachers(context.Context, *Empty) (*GetTeachersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeachers not implemented")
}
func (UnimplementedAdminServiceServer) mustEmbedUnimplementedAdminServiceServer() {}

// UnsafeAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServiceServer will
// result in compilation errors.
type UnsafeAdminServiceServer interface {
	mustEmbedUnimplementedAdminServiceServer()
}

func RegisterAdminServiceServer(s grpc.ServiceRegistrar, srv AdminServiceServer) {
	s.RegisterService(&AdminService_ServiceDesc, srv)
}

func _AdminService_AddStudentToCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddStudentToCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AddStudentToCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_AddStudentToCourse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AddStudentToCourse(ctx, req.(*AddStudentToCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_DeleteStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).DeleteStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_DeleteStudent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).DeleteStudent(ctx, req.(*DeleteStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_DeleteTeacher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTeacherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).DeleteTeacher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_DeleteTeacher_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).DeleteTeacher(ctx, req.(*DeleteTeacherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_AddCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AddCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_AddCourse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AddCourse(ctx, req.(*AddCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_DeleteCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).DeleteCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_DeleteCourse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).DeleteCourse(ctx, req.(*DeleteCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_GetCourses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetCourses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_GetCourses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetCourses(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_GetStudents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetStudents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_GetStudents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetStudents(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_GetTeachers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetTeachers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_GetTeachers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetTeachers(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminService_ServiceDesc is the grpc.ServiceDesc for AdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddStudentToCourse",
			Handler:    _AdminService_AddStudentToCourse_Handler,
		},
		{
			MethodName: "DeleteStudent",
			Handler:    _AdminService_DeleteStudent_Handler,
		},
		{
			MethodName: "DeleteTeacher",
			Handler:    _AdminService_DeleteTeacher_Handler,
		},
		{
			MethodName: "AddCourse",
			Handler:    _AdminService_AddCourse_Handler,
		},
		{
			MethodName: "DeleteCourse",
			Handler:    _AdminService_DeleteCourse_Handler,
		},
		{
			MethodName: "GetCourses",
			Handler:    _AdminService_GetCourses_Handler,
		},
		{
			MethodName: "GetStudents",
			Handler:    _AdminService_GetStudents_Handler,
		},
		{
			MethodName: "GetTeachers",
			Handler:    _AdminService_GetTeachers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "main.proto",
}
