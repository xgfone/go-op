// Copyright 2023 xgfone
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

import "fmt"

var converters = make(map[string]Converter)

// Converter is used to convert an operation to other value.
type Converter func(optype string, oper Oper) interface{}

// RegisterConverter registers the converter of the operation belonging to the type.
//
// If the operation belonging to the type has existed, override it.
func RegisterConverter(_type, op string, convert Converter) {
	if _type == "" {
		panic("RegisterConverter: the type must not be empty")
	}
	if op == "" {
		panic("RegisterConverter: the operation must not be empty")
	}
	if convert == nil {
		panic("RegisterConverter: the operation converter must not be nil")
	}
	converters[fmt.Sprintf("%s:%s", _type, op)] = convert
}

// GetConverter returns the converter of the operation belonging to the type.
//
// If the operation belonging to the type does not exist, return nil instead.
func GetConverter(_type, op string) Converter {
	if _type == "" {
		panic("GetConverter: the type must not be empty")
	}
	if op == "" {
		panic("GetConverter: the operation must not be empty")
	}
	return converters[fmt.Sprintf("%s:%s", _type, op)]
}
