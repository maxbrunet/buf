// Copyright 2020-2022 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: buf/alpha/registry/v1alpha1/schema.proto

package registryv1alpha1connect

import (
	context "context"
	errors "errors"
	v1alpha1 "github.com/bufbuild/buf/private/gen/proto/go/buf/alpha/registry/v1alpha1"
	connect_go "github.com/bufbuild/connect-go"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// SchemaServiceName is the fully-qualified name of the SchemaService service.
	SchemaServiceName = "buf.alpha.registry.v1alpha1.SchemaService"
)

// SchemaServiceClient is a client for the buf.alpha.registry.v1alpha1.SchemaService service.
type SchemaServiceClient interface {
	// GetSchema allows the caller to download a schema for one or more requested
	// types, RPC services, or RPC methods.
	GetSchema(context.Context, *connect_go.Request[v1alpha1.GetSchemaRequest]) (*connect_go.Response[v1alpha1.GetSchemaResponse], error)
	// ConvertMessage allows the caller to convert a given message data blob from
	// one format to another by referring to a type schema for the blob.
	ConvertMessage(context.Context, *connect_go.Request[v1alpha1.ConvertMessageRequest]) (*connect_go.Response[v1alpha1.ConvertMessageResponse], error)
}

// NewSchemaServiceClient constructs a client for the buf.alpha.registry.v1alpha1.SchemaService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewSchemaServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) SchemaServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &schemaServiceClient{
		getSchema: connect_go.NewClient[v1alpha1.GetSchemaRequest, v1alpha1.GetSchemaResponse](
			httpClient,
			baseURL+"/buf.alpha.registry.v1alpha1.SchemaService/GetSchema",
			opts...,
		),
		convertMessage: connect_go.NewClient[v1alpha1.ConvertMessageRequest, v1alpha1.ConvertMessageResponse](
			httpClient,
			baseURL+"/buf.alpha.registry.v1alpha1.SchemaService/ConvertMessage",
			opts...,
		),
	}
}

// schemaServiceClient implements SchemaServiceClient.
type schemaServiceClient struct {
	getSchema      *connect_go.Client[v1alpha1.GetSchemaRequest, v1alpha1.GetSchemaResponse]
	convertMessage *connect_go.Client[v1alpha1.ConvertMessageRequest, v1alpha1.ConvertMessageResponse]
}

// GetSchema calls buf.alpha.registry.v1alpha1.SchemaService.GetSchema.
func (c *schemaServiceClient) GetSchema(ctx context.Context, req *connect_go.Request[v1alpha1.GetSchemaRequest]) (*connect_go.Response[v1alpha1.GetSchemaResponse], error) {
	return c.getSchema.CallUnary(ctx, req)
}

// ConvertMessage calls buf.alpha.registry.v1alpha1.SchemaService.ConvertMessage.
func (c *schemaServiceClient) ConvertMessage(ctx context.Context, req *connect_go.Request[v1alpha1.ConvertMessageRequest]) (*connect_go.Response[v1alpha1.ConvertMessageResponse], error) {
	return c.convertMessage.CallUnary(ctx, req)
}

// SchemaServiceHandler is an implementation of the buf.alpha.registry.v1alpha1.SchemaService
// service.
type SchemaServiceHandler interface {
	// GetSchema allows the caller to download a schema for one or more requested
	// types, RPC services, or RPC methods.
	GetSchema(context.Context, *connect_go.Request[v1alpha1.GetSchemaRequest]) (*connect_go.Response[v1alpha1.GetSchemaResponse], error)
	// ConvertMessage allows the caller to convert a given message data blob from
	// one format to another by referring to a type schema for the blob.
	ConvertMessage(context.Context, *connect_go.Request[v1alpha1.ConvertMessageRequest]) (*connect_go.Response[v1alpha1.ConvertMessageResponse], error)
}

// NewSchemaServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewSchemaServiceHandler(svc SchemaServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/buf.alpha.registry.v1alpha1.SchemaService/GetSchema", connect_go.NewUnaryHandler(
		"/buf.alpha.registry.v1alpha1.SchemaService/GetSchema",
		svc.GetSchema,
		opts...,
	))
	mux.Handle("/buf.alpha.registry.v1alpha1.SchemaService/ConvertMessage", connect_go.NewUnaryHandler(
		"/buf.alpha.registry.v1alpha1.SchemaService/ConvertMessage",
		svc.ConvertMessage,
		opts...,
	))
	return "/buf.alpha.registry.v1alpha1.SchemaService/", mux
}

// UnimplementedSchemaServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedSchemaServiceHandler struct{}

func (UnimplementedSchemaServiceHandler) GetSchema(context.Context, *connect_go.Request[v1alpha1.GetSchemaRequest]) (*connect_go.Response[v1alpha1.GetSchemaResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("buf.alpha.registry.v1alpha1.SchemaService.GetSchema is not implemented"))
}

func (UnimplementedSchemaServiceHandler) ConvertMessage(context.Context, *connect_go.Request[v1alpha1.ConvertMessageRequest]) (*connect_go.Response[v1alpha1.ConvertMessageResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("buf.alpha.registry.v1alpha1.SchemaService.ConvertMessage is not implemented"))
}
