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
// Source: buf/alpha/registry/v1alpha1/plugin_curation.proto

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
	// PluginCurationServiceName is the fully-qualified name of the PluginCurationService service.
	PluginCurationServiceName = "buf.alpha.registry.v1alpha1.PluginCurationService"
	// CodeGenerationServiceName is the fully-qualified name of the CodeGenerationService service.
	CodeGenerationServiceName = "buf.alpha.registry.v1alpha1.CodeGenerationService"
)

// PluginCurationServiceClient is a client for the buf.alpha.registry.v1alpha1.PluginCurationService
// service.
type PluginCurationServiceClient interface {
	// ListCuratedPlugins returns all the curated plugins available.
	ListCuratedPlugins(context.Context, *connect_go.Request[v1alpha1.ListCuratedPluginsRequest]) (*connect_go.Response[v1alpha1.ListCuratedPluginsResponse], error)
	// CreateCuratedPlugin creates a new curated plugin.
	CreateCuratedPlugin(context.Context, *connect_go.Request[v1alpha1.CreateCuratedPluginRequest]) (*connect_go.Response[v1alpha1.CreateCuratedPluginResponse], error)
	// GetLatestCuratedPlugin returns the latest version of a plugin matching given parameters.
	GetLatestCuratedPlugin(context.Context, *connect_go.Request[v1alpha1.GetLatestCuratedPluginRequest]) (*connect_go.Response[v1alpha1.GetLatestCuratedPluginResponse], error)
}

// NewPluginCurationServiceClient constructs a client for the
// buf.alpha.registry.v1alpha1.PluginCurationService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPluginCurationServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) PluginCurationServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &pluginCurationServiceClient{
		listCuratedPlugins: connect_go.NewClient[v1alpha1.ListCuratedPluginsRequest, v1alpha1.ListCuratedPluginsResponse](
			httpClient,
			baseURL+"/buf.alpha.registry.v1alpha1.PluginCurationService/ListCuratedPlugins",
			opts...,
		),
		createCuratedPlugin: connect_go.NewClient[v1alpha1.CreateCuratedPluginRequest, v1alpha1.CreateCuratedPluginResponse](
			httpClient,
			baseURL+"/buf.alpha.registry.v1alpha1.PluginCurationService/CreateCuratedPlugin",
			opts...,
		),
		getLatestCuratedPlugin: connect_go.NewClient[v1alpha1.GetLatestCuratedPluginRequest, v1alpha1.GetLatestCuratedPluginResponse](
			httpClient,
			baseURL+"/buf.alpha.registry.v1alpha1.PluginCurationService/GetLatestCuratedPlugin",
			opts...,
		),
	}
}

// pluginCurationServiceClient implements PluginCurationServiceClient.
type pluginCurationServiceClient struct {
	listCuratedPlugins     *connect_go.Client[v1alpha1.ListCuratedPluginsRequest, v1alpha1.ListCuratedPluginsResponse]
	createCuratedPlugin    *connect_go.Client[v1alpha1.CreateCuratedPluginRequest, v1alpha1.CreateCuratedPluginResponse]
	getLatestCuratedPlugin *connect_go.Client[v1alpha1.GetLatestCuratedPluginRequest, v1alpha1.GetLatestCuratedPluginResponse]
}

// ListCuratedPlugins calls buf.alpha.registry.v1alpha1.PluginCurationService.ListCuratedPlugins.
func (c *pluginCurationServiceClient) ListCuratedPlugins(ctx context.Context, req *connect_go.Request[v1alpha1.ListCuratedPluginsRequest]) (*connect_go.Response[v1alpha1.ListCuratedPluginsResponse], error) {
	return c.listCuratedPlugins.CallUnary(ctx, req)
}

// CreateCuratedPlugin calls buf.alpha.registry.v1alpha1.PluginCurationService.CreateCuratedPlugin.
func (c *pluginCurationServiceClient) CreateCuratedPlugin(ctx context.Context, req *connect_go.Request[v1alpha1.CreateCuratedPluginRequest]) (*connect_go.Response[v1alpha1.CreateCuratedPluginResponse], error) {
	return c.createCuratedPlugin.CallUnary(ctx, req)
}

// GetLatestCuratedPlugin calls
// buf.alpha.registry.v1alpha1.PluginCurationService.GetLatestCuratedPlugin.
func (c *pluginCurationServiceClient) GetLatestCuratedPlugin(ctx context.Context, req *connect_go.Request[v1alpha1.GetLatestCuratedPluginRequest]) (*connect_go.Response[v1alpha1.GetLatestCuratedPluginResponse], error) {
	return c.getLatestCuratedPlugin.CallUnary(ctx, req)
}

// PluginCurationServiceHandler is an implementation of the
// buf.alpha.registry.v1alpha1.PluginCurationService service.
type PluginCurationServiceHandler interface {
	// ListCuratedPlugins returns all the curated plugins available.
	ListCuratedPlugins(context.Context, *connect_go.Request[v1alpha1.ListCuratedPluginsRequest]) (*connect_go.Response[v1alpha1.ListCuratedPluginsResponse], error)
	// CreateCuratedPlugin creates a new curated plugin.
	CreateCuratedPlugin(context.Context, *connect_go.Request[v1alpha1.CreateCuratedPluginRequest]) (*connect_go.Response[v1alpha1.CreateCuratedPluginResponse], error)
	// GetLatestCuratedPlugin returns the latest version of a plugin matching given parameters.
	GetLatestCuratedPlugin(context.Context, *connect_go.Request[v1alpha1.GetLatestCuratedPluginRequest]) (*connect_go.Response[v1alpha1.GetLatestCuratedPluginResponse], error)
}

// NewPluginCurationServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPluginCurationServiceHandler(svc PluginCurationServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/buf.alpha.registry.v1alpha1.PluginCurationService/ListCuratedPlugins", connect_go.NewUnaryHandler(
		"/buf.alpha.registry.v1alpha1.PluginCurationService/ListCuratedPlugins",
		svc.ListCuratedPlugins,
		opts...,
	))
	mux.Handle("/buf.alpha.registry.v1alpha1.PluginCurationService/CreateCuratedPlugin", connect_go.NewUnaryHandler(
		"/buf.alpha.registry.v1alpha1.PluginCurationService/CreateCuratedPlugin",
		svc.CreateCuratedPlugin,
		opts...,
	))
	mux.Handle("/buf.alpha.registry.v1alpha1.PluginCurationService/GetLatestCuratedPlugin", connect_go.NewUnaryHandler(
		"/buf.alpha.registry.v1alpha1.PluginCurationService/GetLatestCuratedPlugin",
		svc.GetLatestCuratedPlugin,
		opts...,
	))
	return "/buf.alpha.registry.v1alpha1.PluginCurationService/", mux
}

// UnimplementedPluginCurationServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPluginCurationServiceHandler struct{}

func (UnimplementedPluginCurationServiceHandler) ListCuratedPlugins(context.Context, *connect_go.Request[v1alpha1.ListCuratedPluginsRequest]) (*connect_go.Response[v1alpha1.ListCuratedPluginsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("buf.alpha.registry.v1alpha1.PluginCurationService.ListCuratedPlugins is not implemented"))
}

func (UnimplementedPluginCurationServiceHandler) CreateCuratedPlugin(context.Context, *connect_go.Request[v1alpha1.CreateCuratedPluginRequest]) (*connect_go.Response[v1alpha1.CreateCuratedPluginResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("buf.alpha.registry.v1alpha1.PluginCurationService.CreateCuratedPlugin is not implemented"))
}

func (UnimplementedPluginCurationServiceHandler) GetLatestCuratedPlugin(context.Context, *connect_go.Request[v1alpha1.GetLatestCuratedPluginRequest]) (*connect_go.Response[v1alpha1.GetLatestCuratedPluginResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("buf.alpha.registry.v1alpha1.PluginCurationService.GetLatestCuratedPlugin is not implemented"))
}

// CodeGenerationServiceClient is a client for the buf.alpha.registry.v1alpha1.CodeGenerationService
// service.
type CodeGenerationServiceClient interface {
	// GenerateCode generates code using the specified remote plugins.
	GenerateCode(context.Context, *connect_go.Request[v1alpha1.GenerateCodeRequest]) (*connect_go.Response[v1alpha1.GenerateCodeResponse], error)
}

// NewCodeGenerationServiceClient constructs a client for the
// buf.alpha.registry.v1alpha1.CodeGenerationService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewCodeGenerationServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) CodeGenerationServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &codeGenerationServiceClient{
		generateCode: connect_go.NewClient[v1alpha1.GenerateCodeRequest, v1alpha1.GenerateCodeResponse](
			httpClient,
			baseURL+"/buf.alpha.registry.v1alpha1.CodeGenerationService/GenerateCode",
			opts...,
		),
	}
}

// codeGenerationServiceClient implements CodeGenerationServiceClient.
type codeGenerationServiceClient struct {
	generateCode *connect_go.Client[v1alpha1.GenerateCodeRequest, v1alpha1.GenerateCodeResponse]
}

// GenerateCode calls buf.alpha.registry.v1alpha1.CodeGenerationService.GenerateCode.
func (c *codeGenerationServiceClient) GenerateCode(ctx context.Context, req *connect_go.Request[v1alpha1.GenerateCodeRequest]) (*connect_go.Response[v1alpha1.GenerateCodeResponse], error) {
	return c.generateCode.CallUnary(ctx, req)
}

// CodeGenerationServiceHandler is an implementation of the
// buf.alpha.registry.v1alpha1.CodeGenerationService service.
type CodeGenerationServiceHandler interface {
	// GenerateCode generates code using the specified remote plugins.
	GenerateCode(context.Context, *connect_go.Request[v1alpha1.GenerateCodeRequest]) (*connect_go.Response[v1alpha1.GenerateCodeResponse], error)
}

// NewCodeGenerationServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewCodeGenerationServiceHandler(svc CodeGenerationServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/buf.alpha.registry.v1alpha1.CodeGenerationService/GenerateCode", connect_go.NewUnaryHandler(
		"/buf.alpha.registry.v1alpha1.CodeGenerationService/GenerateCode",
		svc.GenerateCode,
		opts...,
	))
	return "/buf.alpha.registry.v1alpha1.CodeGenerationService/", mux
}

// UnimplementedCodeGenerationServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedCodeGenerationServiceHandler struct{}

func (UnimplementedCodeGenerationServiceHandler) GenerateCode(context.Context, *connect_go.Request[v1alpha1.GenerateCodeRequest]) (*connect_go.Response[v1alpha1.GenerateCodeResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("buf.alpha.registry.v1alpha1.CodeGenerationService.GenerateCode is not implemented"))
}
