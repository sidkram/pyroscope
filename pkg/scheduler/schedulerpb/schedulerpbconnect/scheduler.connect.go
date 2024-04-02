// SPDX-License-Identifier: AGPL-3.0-only
// Provenance-includes-location: https://github.com/cortexproject/cortex/blob/master/pkg/scheduler/schedulerpb/scheduler.proto
// Provenance-includes-license: Apache-2.0
// Provenance-includes-copyright: The Cortex Authors.

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: scheduler/schedulerpb/scheduler.proto

package schedulerpbconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	schedulerpb "github.com/grafana/pyroscope/pkg/scheduler/schedulerpb"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// SchedulerForQuerierName is the fully-qualified name of the SchedulerForQuerier service.
	SchedulerForQuerierName = "schedulerpb.SchedulerForQuerier"
	// SchedulerForFrontendName is the fully-qualified name of the SchedulerForFrontend service.
	SchedulerForFrontendName = "schedulerpb.SchedulerForFrontend"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// SchedulerForQuerierQuerierLoopProcedure is the fully-qualified name of the SchedulerForQuerier's
	// QuerierLoop RPC.
	SchedulerForQuerierQuerierLoopProcedure = "/schedulerpb.SchedulerForQuerier/QuerierLoop"
	// SchedulerForQuerierNotifyQuerierShutdownProcedure is the fully-qualified name of the
	// SchedulerForQuerier's NotifyQuerierShutdown RPC.
	SchedulerForQuerierNotifyQuerierShutdownProcedure = "/schedulerpb.SchedulerForQuerier/NotifyQuerierShutdown"
	// SchedulerForFrontendFrontendLoopProcedure is the fully-qualified name of the
	// SchedulerForFrontend's FrontendLoop RPC.
	SchedulerForFrontendFrontendLoopProcedure = "/schedulerpb.SchedulerForFrontend/FrontendLoop"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	schedulerForQuerierServiceDescriptor                     = schedulerpb.File_scheduler_schedulerpb_scheduler_proto.Services().ByName("SchedulerForQuerier")
	schedulerForQuerierQuerierLoopMethodDescriptor           = schedulerForQuerierServiceDescriptor.Methods().ByName("QuerierLoop")
	schedulerForQuerierNotifyQuerierShutdownMethodDescriptor = schedulerForQuerierServiceDescriptor.Methods().ByName("NotifyQuerierShutdown")
	schedulerForFrontendServiceDescriptor                    = schedulerpb.File_scheduler_schedulerpb_scheduler_proto.Services().ByName("SchedulerForFrontend")
	schedulerForFrontendFrontendLoopMethodDescriptor         = schedulerForFrontendServiceDescriptor.Methods().ByName("FrontendLoop")
)

// SchedulerForQuerierClient is a client for the schedulerpb.SchedulerForQuerier service.
type SchedulerForQuerierClient interface {
	// After calling this method, both Querier and Scheduler enter a loop, in which querier waits for
	// "SchedulerToQuerier" messages containing HTTP requests and processes them. After processing the request,
	// querier signals that it is ready to accept another one by sending empty QuerierToScheduler message.
	//
	// Long-running loop is used to detect broken connection between scheduler and querier. This is important
	// for scheduler to keep a list of connected queriers up-to-date.
	QuerierLoop(context.Context) *connect.BidiStreamForClient[schedulerpb.QuerierToScheduler, schedulerpb.SchedulerToQuerier]
	// The querier notifies the query-scheduler that it started a graceful shutdown.
	NotifyQuerierShutdown(context.Context, *connect.Request[schedulerpb.NotifyQuerierShutdownRequest]) (*connect.Response[schedulerpb.NotifyQuerierShutdownResponse], error)
}

// NewSchedulerForQuerierClient constructs a client for the schedulerpb.SchedulerForQuerier service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewSchedulerForQuerierClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) SchedulerForQuerierClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &schedulerForQuerierClient{
		querierLoop: connect.NewClient[schedulerpb.QuerierToScheduler, schedulerpb.SchedulerToQuerier](
			httpClient,
			baseURL+SchedulerForQuerierQuerierLoopProcedure,
			connect.WithSchema(schedulerForQuerierQuerierLoopMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		notifyQuerierShutdown: connect.NewClient[schedulerpb.NotifyQuerierShutdownRequest, schedulerpb.NotifyQuerierShutdownResponse](
			httpClient,
			baseURL+SchedulerForQuerierNotifyQuerierShutdownProcedure,
			connect.WithSchema(schedulerForQuerierNotifyQuerierShutdownMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// schedulerForQuerierClient implements SchedulerForQuerierClient.
type schedulerForQuerierClient struct {
	querierLoop           *connect.Client[schedulerpb.QuerierToScheduler, schedulerpb.SchedulerToQuerier]
	notifyQuerierShutdown *connect.Client[schedulerpb.NotifyQuerierShutdownRequest, schedulerpb.NotifyQuerierShutdownResponse]
}

// QuerierLoop calls schedulerpb.SchedulerForQuerier.QuerierLoop.
func (c *schedulerForQuerierClient) QuerierLoop(ctx context.Context) *connect.BidiStreamForClient[schedulerpb.QuerierToScheduler, schedulerpb.SchedulerToQuerier] {
	return c.querierLoop.CallBidiStream(ctx)
}

// NotifyQuerierShutdown calls schedulerpb.SchedulerForQuerier.NotifyQuerierShutdown.
func (c *schedulerForQuerierClient) NotifyQuerierShutdown(ctx context.Context, req *connect.Request[schedulerpb.NotifyQuerierShutdownRequest]) (*connect.Response[schedulerpb.NotifyQuerierShutdownResponse], error) {
	return c.notifyQuerierShutdown.CallUnary(ctx, req)
}

// SchedulerForQuerierHandler is an implementation of the schedulerpb.SchedulerForQuerier service.
type SchedulerForQuerierHandler interface {
	// After calling this method, both Querier and Scheduler enter a loop, in which querier waits for
	// "SchedulerToQuerier" messages containing HTTP requests and processes them. After processing the request,
	// querier signals that it is ready to accept another one by sending empty QuerierToScheduler message.
	//
	// Long-running loop is used to detect broken connection between scheduler and querier. This is important
	// for scheduler to keep a list of connected queriers up-to-date.
	QuerierLoop(context.Context, *connect.BidiStream[schedulerpb.QuerierToScheduler, schedulerpb.SchedulerToQuerier]) error
	// The querier notifies the query-scheduler that it started a graceful shutdown.
	NotifyQuerierShutdown(context.Context, *connect.Request[schedulerpb.NotifyQuerierShutdownRequest]) (*connect.Response[schedulerpb.NotifyQuerierShutdownResponse], error)
}

// NewSchedulerForQuerierHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewSchedulerForQuerierHandler(svc SchedulerForQuerierHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	schedulerForQuerierQuerierLoopHandler := connect.NewBidiStreamHandler(
		SchedulerForQuerierQuerierLoopProcedure,
		svc.QuerierLoop,
		connect.WithSchema(schedulerForQuerierQuerierLoopMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	schedulerForQuerierNotifyQuerierShutdownHandler := connect.NewUnaryHandler(
		SchedulerForQuerierNotifyQuerierShutdownProcedure,
		svc.NotifyQuerierShutdown,
		connect.WithSchema(schedulerForQuerierNotifyQuerierShutdownMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/schedulerpb.SchedulerForQuerier/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case SchedulerForQuerierQuerierLoopProcedure:
			schedulerForQuerierQuerierLoopHandler.ServeHTTP(w, r)
		case SchedulerForQuerierNotifyQuerierShutdownProcedure:
			schedulerForQuerierNotifyQuerierShutdownHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedSchedulerForQuerierHandler returns CodeUnimplemented from all methods.
type UnimplementedSchedulerForQuerierHandler struct{}

func (UnimplementedSchedulerForQuerierHandler) QuerierLoop(context.Context, *connect.BidiStream[schedulerpb.QuerierToScheduler, schedulerpb.SchedulerToQuerier]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("schedulerpb.SchedulerForQuerier.QuerierLoop is not implemented"))
}

func (UnimplementedSchedulerForQuerierHandler) NotifyQuerierShutdown(context.Context, *connect.Request[schedulerpb.NotifyQuerierShutdownRequest]) (*connect.Response[schedulerpb.NotifyQuerierShutdownResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("schedulerpb.SchedulerForQuerier.NotifyQuerierShutdown is not implemented"))
}

// SchedulerForFrontendClient is a client for the schedulerpb.SchedulerForFrontend service.
type SchedulerForFrontendClient interface {
	// After calling this method, both Frontend and Scheduler enter a loop. Frontend will keep sending ENQUEUE and
	// CANCEL requests, and scheduler is expected to process them. Scheduler returns one response for each request.
	//
	// Long-running loop is used to detect broken connection between frontend and scheduler. This is important for both
	// parties... if connection breaks, frontend can cancel (and possibly retry on different scheduler) all pending
	// requests sent to this scheduler, while scheduler can cancel queued requests from given frontend.
	FrontendLoop(context.Context) *connect.BidiStreamForClient[schedulerpb.FrontendToScheduler, schedulerpb.SchedulerToFrontend]
}

// NewSchedulerForFrontendClient constructs a client for the schedulerpb.SchedulerForFrontend
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewSchedulerForFrontendClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) SchedulerForFrontendClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &schedulerForFrontendClient{
		frontendLoop: connect.NewClient[schedulerpb.FrontendToScheduler, schedulerpb.SchedulerToFrontend](
			httpClient,
			baseURL+SchedulerForFrontendFrontendLoopProcedure,
			connect.WithSchema(schedulerForFrontendFrontendLoopMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// schedulerForFrontendClient implements SchedulerForFrontendClient.
type schedulerForFrontendClient struct {
	frontendLoop *connect.Client[schedulerpb.FrontendToScheduler, schedulerpb.SchedulerToFrontend]
}

// FrontendLoop calls schedulerpb.SchedulerForFrontend.FrontendLoop.
func (c *schedulerForFrontendClient) FrontendLoop(ctx context.Context) *connect.BidiStreamForClient[schedulerpb.FrontendToScheduler, schedulerpb.SchedulerToFrontend] {
	return c.frontendLoop.CallBidiStream(ctx)
}

// SchedulerForFrontendHandler is an implementation of the schedulerpb.SchedulerForFrontend service.
type SchedulerForFrontendHandler interface {
	// After calling this method, both Frontend and Scheduler enter a loop. Frontend will keep sending ENQUEUE and
	// CANCEL requests, and scheduler is expected to process them. Scheduler returns one response for each request.
	//
	// Long-running loop is used to detect broken connection between frontend and scheduler. This is important for both
	// parties... if connection breaks, frontend can cancel (and possibly retry on different scheduler) all pending
	// requests sent to this scheduler, while scheduler can cancel queued requests from given frontend.
	FrontendLoop(context.Context, *connect.BidiStream[schedulerpb.FrontendToScheduler, schedulerpb.SchedulerToFrontend]) error
}

// NewSchedulerForFrontendHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewSchedulerForFrontendHandler(svc SchedulerForFrontendHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	schedulerForFrontendFrontendLoopHandler := connect.NewBidiStreamHandler(
		SchedulerForFrontendFrontendLoopProcedure,
		svc.FrontendLoop,
		connect.WithSchema(schedulerForFrontendFrontendLoopMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/schedulerpb.SchedulerForFrontend/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case SchedulerForFrontendFrontendLoopProcedure:
			schedulerForFrontendFrontendLoopHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedSchedulerForFrontendHandler returns CodeUnimplemented from all methods.
type UnimplementedSchedulerForFrontendHandler struct{}

func (UnimplementedSchedulerForFrontendHandler) FrontendLoop(context.Context, *connect.BidiStream[schedulerpb.FrontendToScheduler, schedulerpb.SchedulerToFrontend]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("schedulerpb.SchedulerForFrontend.FrontendLoop is not implemented"))
}
