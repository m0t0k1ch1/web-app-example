// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: app/v1/task.proto

package appv1connect

import (
	v1 "app/gen/buf/app/v1"
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// TaskServiceName is the fully-qualified name of the TaskService service.
	TaskServiceName = "app.v1.TaskService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// TaskServiceCreateProcedure is the fully-qualified name of the TaskService's Create RPC.
	TaskServiceCreateProcedure = "/app.v1.TaskService/Create"
	// TaskServiceGetProcedure is the fully-qualified name of the TaskService's Get RPC.
	TaskServiceGetProcedure = "/app.v1.TaskService/Get"
	// TaskServiceListProcedure is the fully-qualified name of the TaskService's List RPC.
	TaskServiceListProcedure = "/app.v1.TaskService/List"
	// TaskServiceUpdateProcedure is the fully-qualified name of the TaskService's Update RPC.
	TaskServiceUpdateProcedure = "/app.v1.TaskService/Update"
	// TaskServiceDeleteProcedure is the fully-qualified name of the TaskService's Delete RPC.
	TaskServiceDeleteProcedure = "/app.v1.TaskService/Delete"
)

// TaskServiceClient is a client for the app.v1.TaskService service.
type TaskServiceClient interface {
	Create(context.Context, *connect.Request[v1.TaskServiceCreateRequest]) (*connect.Response[v1.TaskServiceCreateResponse], error)
	Get(context.Context, *connect.Request[v1.TaskServiceGetRequest]) (*connect.Response[v1.TaskServiceGetResponse], error)
	List(context.Context, *connect.Request[v1.TaskServiceListRequest]) (*connect.Response[v1.TaskServiceListResponse], error)
	Update(context.Context, *connect.Request[v1.TaskServiceUpdateRequest]) (*connect.Response[v1.TaskServiceUpdateResponse], error)
	Delete(context.Context, *connect.Request[v1.TaskServiceDeleteRequest]) (*connect.Response[v1.TaskServiceDeleteResponse], error)
}

// NewTaskServiceClient constructs a client for the app.v1.TaskService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTaskServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) TaskServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &taskServiceClient{
		create: connect.NewClient[v1.TaskServiceCreateRequest, v1.TaskServiceCreateResponse](
			httpClient,
			baseURL+TaskServiceCreateProcedure,
			opts...,
		),
		get: connect.NewClient[v1.TaskServiceGetRequest, v1.TaskServiceGetResponse](
			httpClient,
			baseURL+TaskServiceGetProcedure,
			opts...,
		),
		list: connect.NewClient[v1.TaskServiceListRequest, v1.TaskServiceListResponse](
			httpClient,
			baseURL+TaskServiceListProcedure,
			opts...,
		),
		update: connect.NewClient[v1.TaskServiceUpdateRequest, v1.TaskServiceUpdateResponse](
			httpClient,
			baseURL+TaskServiceUpdateProcedure,
			opts...,
		),
		delete: connect.NewClient[v1.TaskServiceDeleteRequest, v1.TaskServiceDeleteResponse](
			httpClient,
			baseURL+TaskServiceDeleteProcedure,
			opts...,
		),
	}
}

// taskServiceClient implements TaskServiceClient.
type taskServiceClient struct {
	create *connect.Client[v1.TaskServiceCreateRequest, v1.TaskServiceCreateResponse]
	get    *connect.Client[v1.TaskServiceGetRequest, v1.TaskServiceGetResponse]
	list   *connect.Client[v1.TaskServiceListRequest, v1.TaskServiceListResponse]
	update *connect.Client[v1.TaskServiceUpdateRequest, v1.TaskServiceUpdateResponse]
	delete *connect.Client[v1.TaskServiceDeleteRequest, v1.TaskServiceDeleteResponse]
}

// Create calls app.v1.TaskService.Create.
func (c *taskServiceClient) Create(ctx context.Context, req *connect.Request[v1.TaskServiceCreateRequest]) (*connect.Response[v1.TaskServiceCreateResponse], error) {
	return c.create.CallUnary(ctx, req)
}

// Get calls app.v1.TaskService.Get.
func (c *taskServiceClient) Get(ctx context.Context, req *connect.Request[v1.TaskServiceGetRequest]) (*connect.Response[v1.TaskServiceGetResponse], error) {
	return c.get.CallUnary(ctx, req)
}

// List calls app.v1.TaskService.List.
func (c *taskServiceClient) List(ctx context.Context, req *connect.Request[v1.TaskServiceListRequest]) (*connect.Response[v1.TaskServiceListResponse], error) {
	return c.list.CallUnary(ctx, req)
}

// Update calls app.v1.TaskService.Update.
func (c *taskServiceClient) Update(ctx context.Context, req *connect.Request[v1.TaskServiceUpdateRequest]) (*connect.Response[v1.TaskServiceUpdateResponse], error) {
	return c.update.CallUnary(ctx, req)
}

// Delete calls app.v1.TaskService.Delete.
func (c *taskServiceClient) Delete(ctx context.Context, req *connect.Request[v1.TaskServiceDeleteRequest]) (*connect.Response[v1.TaskServiceDeleteResponse], error) {
	return c.delete.CallUnary(ctx, req)
}

// TaskServiceHandler is an implementation of the app.v1.TaskService service.
type TaskServiceHandler interface {
	Create(context.Context, *connect.Request[v1.TaskServiceCreateRequest]) (*connect.Response[v1.TaskServiceCreateResponse], error)
	Get(context.Context, *connect.Request[v1.TaskServiceGetRequest]) (*connect.Response[v1.TaskServiceGetResponse], error)
	List(context.Context, *connect.Request[v1.TaskServiceListRequest]) (*connect.Response[v1.TaskServiceListResponse], error)
	Update(context.Context, *connect.Request[v1.TaskServiceUpdateRequest]) (*connect.Response[v1.TaskServiceUpdateResponse], error)
	Delete(context.Context, *connect.Request[v1.TaskServiceDeleteRequest]) (*connect.Response[v1.TaskServiceDeleteResponse], error)
}

// NewTaskServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTaskServiceHandler(svc TaskServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	taskServiceCreateHandler := connect.NewUnaryHandler(
		TaskServiceCreateProcedure,
		svc.Create,
		opts...,
	)
	taskServiceGetHandler := connect.NewUnaryHandler(
		TaskServiceGetProcedure,
		svc.Get,
		opts...,
	)
	taskServiceListHandler := connect.NewUnaryHandler(
		TaskServiceListProcedure,
		svc.List,
		opts...,
	)
	taskServiceUpdateHandler := connect.NewUnaryHandler(
		TaskServiceUpdateProcedure,
		svc.Update,
		opts...,
	)
	taskServiceDeleteHandler := connect.NewUnaryHandler(
		TaskServiceDeleteProcedure,
		svc.Delete,
		opts...,
	)
	return "/app.v1.TaskService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TaskServiceCreateProcedure:
			taskServiceCreateHandler.ServeHTTP(w, r)
		case TaskServiceGetProcedure:
			taskServiceGetHandler.ServeHTTP(w, r)
		case TaskServiceListProcedure:
			taskServiceListHandler.ServeHTTP(w, r)
		case TaskServiceUpdateProcedure:
			taskServiceUpdateHandler.ServeHTTP(w, r)
		case TaskServiceDeleteProcedure:
			taskServiceDeleteHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedTaskServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTaskServiceHandler struct{}

func (UnimplementedTaskServiceHandler) Create(context.Context, *connect.Request[v1.TaskServiceCreateRequest]) (*connect.Response[v1.TaskServiceCreateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("app.v1.TaskService.Create is not implemented"))
}

func (UnimplementedTaskServiceHandler) Get(context.Context, *connect.Request[v1.TaskServiceGetRequest]) (*connect.Response[v1.TaskServiceGetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("app.v1.TaskService.Get is not implemented"))
}

func (UnimplementedTaskServiceHandler) List(context.Context, *connect.Request[v1.TaskServiceListRequest]) (*connect.Response[v1.TaskServiceListResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("app.v1.TaskService.List is not implemented"))
}

func (UnimplementedTaskServiceHandler) Update(context.Context, *connect.Request[v1.TaskServiceUpdateRequest]) (*connect.Response[v1.TaskServiceUpdateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("app.v1.TaskService.Update is not implemented"))
}

func (UnimplementedTaskServiceHandler) Delete(context.Context, *connect.Request[v1.TaskServiceDeleteRequest]) (*connect.Response[v1.TaskServiceDeleteResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("app.v1.TaskService.Delete is not implemented"))
}
