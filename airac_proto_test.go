// +build proto

/*
 * Copyright (c) 2018 Johannes Kohnen <jwkohnen-github@ko-sys.com>
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

import (
	"math"
	"testing"

	"github.com/jwkohnen/airac/proto"
)

func TestProto(t *testing.T) {
	for want := AIRAC(0); want < FromStringMust("9213"); want++ {
		p := want.Proto()
		got := FromProto(p)
		if want != got {
			t.Errorf("Want %v, got %x", want, got)
		}
	}
}

func TestProtoOverflow(t *testing.T) {
	want := AIRAC(0)
	p := proto.AiracMessage{Airac19010110: math.MaxUint16 + 1}
	got := FromProto(p)
	if got != want {
		t.Errorf("Want %s, got %s", want, got)
	}
}
