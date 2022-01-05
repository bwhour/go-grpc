//go:build !linux || appengine
// +build !linux appengine

/*
 *
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package service

import (
	channelzpb "github.com/bwhour/go-grpc/lib/grpc/channelz/grpc_channelz_v1"
	"github.com/bwhour/go-grpc/lib/grpc/internal/channelz"
)

func sockoptToProto(skopts *channelz.SocketOptionData) []*channelzpb.SocketOption {
	return nil
}