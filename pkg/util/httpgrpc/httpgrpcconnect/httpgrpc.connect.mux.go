// Code generated by protoc-gen-connect-go-mux. DO NOT EDIT.
//
// Source: util/httpgrpc/httpgrpc.proto

package httpgrpcconnect

import (
	connect "connectrpc.com/connect"
	mux "github.com/gorilla/mux"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

// RegisterHTTPHandler register an HTTP handler to a mux.Router from the service implementation.
func RegisterHTTPHandler(mux *mux.Router, svc HTTPHandler, opts ...connect.HandlerOption) {
	mux.Handle("/httpgrpc.HTTP/Handle", connect.NewUnaryHandler(
		"/httpgrpc.HTTP/Handle",
		svc.Handle,
		opts...,
	))
}
