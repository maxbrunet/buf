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

// Package appprotoexec provides protoc plugin handling and execution.
//
// Note this is currently implicitly tested through buf's protoc command.
// If this were split out into a separate package, testing would need to be moved to this package.
package appprotoexec

import (
	"context"
	"fmt"

	"github.com/bufbuild/buf/private/pkg/app"
	"github.com/bufbuild/buf/private/pkg/app/appproto"
	"github.com/bufbuild/buf/private/pkg/command"
	"github.com/bufbuild/buf/private/pkg/storage/storageos"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/pluginpb"
)

const (
	// Note on versions: while Protobuf is on vX.0 where X >=21, and we
	// download protoc vX.0, the version reported by protoc --version is 3.X.0.
	// This is what we want to report here.

	// DefaultMajorVersion is the default major version.
	defaultMajorVersion = 3
	// DefaultMinorVersion is the default minor version.
	defaultMinorVersion = 21
	// DefaultPatchVersion is the default patch version.
	defaultPatchVersion = 7
	// DefaultSuffixVersion is the default suffix version.
	defaultSuffixVersion = ""
)

var (
	// ProtocProxyPluginNames are the names of the plugins that should be proxied through protoc
	// in the absence of a binary.
	ProtocProxyPluginNames = map[string]struct{}{
		"cpp":    {},
		"csharp": {},
		"java":   {},
		"js":     {},
		"objc":   {},
		"php":    {},
		"python": {},
		"pyi":    {},
		"ruby":   {},
		"kotlin": {},
	}

	// DefaultVersion represents the default version to use as compiler version for codegen requests.
	DefaultVersion = newVersion(
		defaultMajorVersion,
		defaultMinorVersion,
		defaultPatchVersion,
		defaultSuffixVersion,
	)
)

// Generator is used to generate code with plugins found on the local filesystem.
type Generator interface {
	// Generate generates a CodeGeneratorResponse for the given pluginName. The
	// pluginName must be available on the system's PATH or one of the plugins
	// built-in to protoc. The plugin path can be overridden via the
	// GenerateWithPluginPath option.
	Generate(
		ctx context.Context,
		container app.EnvStderrContainer,
		pluginName string,
		requests []*pluginpb.CodeGeneratorRequest,
		options ...GenerateOption,
	) (*pluginpb.CodeGeneratorResponse, error)
}

// NewGenerator returns a new Generator.
func NewGenerator(
	logger *zap.Logger,
	storageosProvider storageos.Provider,
	runner command.Runner,
) Generator {
	return newGenerator(logger, storageosProvider, runner)
}

// GenerateOption is an option for Generate.
type GenerateOption func(*generateOptions)

// GenerateWithPluginPath returns a new GenerateOption that uses the given
// path to the plugin.
func GenerateWithPluginPath(pluginPath string) GenerateOption {
	return func(generateOptions *generateOptions) {
		generateOptions.pluginPath = pluginPath
	}
}

// NewHandler returns a new Handler based on the plugin name and optional path.
//
// protocPath and pluginPath are optional.
//
//   - If the plugin path is set, this returns a new binary handler for that path.
//   - If the plugin path is unset, this does exec.LookPath for a binary named protoc-gen-pluginName,
//     and if one is found, a new binary handler is returned for this.
//   - Else, if the name is in ProtocProxyPluginNames, this returns a new protoc proxy handler.
//   - Else, this returns error.
func NewHandler(
	logger *zap.Logger,
	storageosProvider storageos.Provider,
	runner command.Runner,
	pluginName string,
	options ...HandlerOption,
) (appproto.Handler, error) {
	handlerOptions := newHandlerOptions()
	for _, option := range options {
		option(handlerOptions)
	}
	if handlerOptions.pluginPath != "" {
		pluginPath, err := unsafeLookPath(handlerOptions.pluginPath)
		if err != nil {
			return nil, err
		}
		return newBinaryHandler(logger, runner, pluginPath), nil
	}
	pluginPath, err := unsafeLookPath("protoc-gen-" + pluginName)
	if err == nil {
		return newBinaryHandler(logger, runner, pluginPath), nil
	}
	// we always look for protoc-gen-X first, but if not, check the builtins
	if _, ok := ProtocProxyPluginNames[pluginName]; ok {
		if handlerOptions.protocPath == "" {
			handlerOptions.protocPath = "protoc"
		}
		protocPath, err := unsafeLookPath(handlerOptions.protocPath)
		if err != nil {
			return nil, err
		}
		return newProtocProxyHandler(logger, storageosProvider, runner, protocPath, pluginName), nil
	}
	return nil, fmt.Errorf(
		"could not find protoc plugin for name %s - please make sure protoc-gen-%s is installed and present on your $PATH",
		pluginName,
		pluginName,
	)
}

// HandlerOption is an option for a new Handler.
type HandlerOption func(*handlerOptions)

// HandlerWithProtocPath returns a new HandlerOption that sets the path to the protoc binary.
//
// The default is to do exec.LookPath on "protoc".
func HandlerWithProtocPath(protocPath string) HandlerOption {
	return func(handlerOptions *handlerOptions) {
		handlerOptions.protocPath = protocPath
	}
}

// HandlerWithPluginPath returns a new HandlerOption that sets the path to the plugin binary.
//
// The default is to do exec.LookPath on "protoc-gen-" + pluginName.
func HandlerWithPluginPath(pluginPath string) HandlerOption {
	return func(handlerOptions *handlerOptions) {
		handlerOptions.pluginPath = pluginPath
	}
}

type handlerOptions struct {
	protocPath string
	pluginPath string
}

func newHandlerOptions() *handlerOptions {
	return &handlerOptions{}
}
