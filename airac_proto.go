// +build proto

/*
 * Copyright (c) 2017 Johannes Kohnen <wjkohnen@users.noreply.github.com>
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

//go:generate protoc --go_out=./ airac.proto

import (
	"math"
)

// FromProto converts an AIRAC protobuffer message to an AIRAC value.
func FromProto(a AiracMessage) Airac {
	if a.Airac19010110 > math.MaxUint16 {
		return 0
	}
	return Airac(a.Airac19010110)
}

// ToProto converts an AIRAC value to an AIRAC protobuffer message.
func (a Airac) ToProto() AiracMessage {
	return AiracMessage{uint32(a)}
}
