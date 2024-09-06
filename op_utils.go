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

import "unicode/utf8"

// StrBytesLen returns a lazy function to limit the byte length of a string to n.
func StrBytesLen(n int) Lazy {
	if n <= 0 {
		panic("op.StrBytesLen: n must be a positive negative")
	}

	return func(o Op) Op {
		if s := o.Val.(string); len(s) > n {
			o.Val = s[:n]
		}
		return o
	}
}

// StrCharsLen returns a lazy function to limit the character length of a string to n.
func StrCharsLen(n int) Lazy {
	if n <= 0 {
		panic("op.StrCharsLen: n must be a positive negative")
	}

	return func(o Op) Op {
		if s := o.Val.(string); len(s) > n && utf8.RuneCountInString(s) > n {
			var m int
			for i := range s {
				if m < n {
					m++
				} else {
					o.Val = s[:i]
					break
				}
			}
		}
		return o
	}
}
