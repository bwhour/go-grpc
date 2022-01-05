/*
 *
 * Copyright 2019 gRPC authors.
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

// Package experimental contains xds implementation, still in experimental
// state. Users only need to import this package to get all xds functionality.
// Things are expected to change fast until we get to a stable state, at
// which point, all this will be moved to the xds package.
package experimental

import (
	"github.com/bwhour/go-grpc/lib/grpc/balancer"
	xdsbalancer "github.com/bwhour/go-grpc/lib/grpc/xds/internal/balancer"
)

func init() {
	balancer.Register(xdsbalancer.NewBalancerBuilder())
}
