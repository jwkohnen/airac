// +build jwkohnen_airac_proto

/*
 * Copyright (c) 2020 Johannes Kohnen <jwkohnen-github@ko-sys.com>
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
 */

package airac

//go:generate protoc --go_out=./ proto/airac.proto

import (
	"math"

	"github.com/jwkohnen/airac/proto"
)

// FromProto converts an AIRAC protobuffer message to an AIRAC value.
func FromProto(a proto.AiracMessage) AIRAC {
	if a.Airac19010110 > math.MaxUint16 {
		return 0
	}
	return AIRAC(a.Airac19010110)
}

// Proto converts an AIRAC value to an AIRAC protobuffer message.
func (a AIRAC) Proto() proto.AiracMessage {
	return proto.AiracMessage{Airac19010110: uint32(a)}
}
