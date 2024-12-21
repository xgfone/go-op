// Copyright 2024 xgfone
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package op

import "testing"

func TestPageSize(t *testing.T) {
	page := PageSize(1, 20)

	type Limiter interface {
		Limit() int
	}

	if ps, ok := page.(Limiter); ok {
		t.Errorf("unexpect Limiter instance, limit=%d", ps.Limit())
	}

	if ps, ok := page.Op().Val.(Limiter); !ok {
		t.Error("expect Limiter instance, but got nil")
	} else if limit := ps.Limit(); limit != 20 {
		t.Errorf("expect limit %d, but got %d", 20, limit)
	}
}
