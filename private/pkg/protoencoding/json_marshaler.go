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

package protoencoding

import (
	"bytes"
	"encoding/json"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type jsonMarshaler struct {
	resolver      Resolver
	indent        string
	useProtoNames bool
}

func newJSONMarshaler(resolver Resolver, indent string, useProtoNames bool) Marshaler {
	return &jsonMarshaler{
		resolver:      resolver,
		indent:        indent,
		useProtoNames: useProtoNames,
	}
}

func (m *jsonMarshaler) Marshal(message proto.Message) ([]byte, error) {
	if err := ReparseUnrecognized(m.resolver, message.ProtoReflect()); err != nil {
		return nil, err
	}
	options := protojson.MarshalOptions{
		Resolver:      m.resolver,
		Indent:        m.indent,
		UseProtoNames: m.useProtoNames,
	}
	data, err := options.Marshal(message)
	if err != nil {
		return nil, err
	}
	// This is needed due to the instability of protojson output.
	//
	// https://github.com/golang/protobuf/issues/1121
	// https://go-review.googlesource.com/c/protobuf/+/151340
	// https://developers.google.com/protocol-buffers/docs/reference/go/faq#unstable-json
	//
	// We may need to do a full encoding/json encode/decode in the future if protojson
	// produces non-deterministic output.
	buffer := bytes.NewBuffer(nil)
	if err := json.Compact(buffer, data); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
