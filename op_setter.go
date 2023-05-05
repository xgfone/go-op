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

// Pre-define some setter operations.
const (
	SetOpInc = "Inc"
	SetOpDec = "Dec"
	SetOpAdd = "Add"
	SetOpSub = "Sub"
	SetOpMul = "Mul"
	SetOpDiv = "Div"
	SetOpSet = "Set"
)

// Setter represents a setter operation.
type Setter interface {
	setter() // Meaningless, just distinguish Setter from Condition.
	Oper
}

var _ Setter = Op{}

func (o Op) setter() {}

/// ---------------------------------------------------------------------- ///

// Inc is equal to New(SetOpInc, key, nil).
func Inc(key string) Op {
	return New(SetOpInc, key, nil)
}

// Dec is equal to New(SetOpDec, key, nil).
func Dec(key string) Op {
	return New(SetOpDec, key, nil)
}

// Add is equal to New(SetOpAdd, key, value).
func Add(key string, value interface{}) Op {
	return New(SetOpAdd, key, value)
}

// Sub is equal to New(SetOpSub, key, value).
func Sub(key string, value interface{}) Op {
	return New(SetOpSub, key, value)
}

// Mul is equal to New(SetOpMul, key, value).
func Mul(key string, value interface{}) Op {
	return New(SetOpMul, key, value)
}

// Div is equal to New(SetOpDiv, key, value).
func Div(key string, value interface{}) Op {
	return New(SetOpDiv, key, value)
}

// Set is equal to New(SetOpSet, key, value).
func Set(key string, value interface{}) Op {
	return New(SetOpSet, key, value)
}

// Inc is equal to Inc(o.Key).
func (o Op) Inc() Op {
	return Inc(o.Key)
}

// Dec is equal to Dec(o.Key).
func (o Op) Dec() Op {
	return Dec(o.Key)
}

// Add is equal to Add(o.Key, value).
func (o Op) Add(value interface{}) Op {
	return Add(o.Key, value)
}

// Sub is equal to Sub(o.Key, value).
func (o Op) Sub(value interface{}) Op {
	return Sub(o.Key, value)
}

// Mul is equal to Mul(o.Key, value).
func (o Op) Mul(value interface{}) Op {
	return Mul(o.Key, value)
}

// Div is equal to Div(o.Key, value).
func (o Op) Div(value interface{}) Op {
	return Div(o.Key, value)
}

// Set is equal to Set(o.Key, value).
func (o Op) Set(value interface{}) Op {
	return Set(o.Key, value)
}
