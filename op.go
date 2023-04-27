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

// Package op provides a common operation about condition and setter.
package op

// Oper is a common operation interface.
type Oper interface {
	GetValue() interface{}
	GetKey() string
	GetOp() string
}

// ContainsKey reports whether the operations contains the key.
func ContainsKey[S ~[]E, E Oper](ops S, key string) bool {
	for _, op := range ops {
		if op.GetKey() == key {
			return true
		}
	}
	return false
}

// Contains reports whether the operations contains the operation by the key,
// which is equal to ContainsKey(ops, op.GetKey()).
func Contains[S ~[]E1, E1, E2 Oper](ops S, op E2) bool {
	return ContainsKey(ops, op.GetKey())
}

/// ----------------------------------------------------------------------- ///

var _ Oper = Op{}

// Op represents an operation, which may be a condtion or Op.
type Op struct {
	Op  string
	Key string
	Val interface{}
}

// Key is equal to NewOp("", key, nil).
func Key(key string) Op { return NewOp("", key, nil) }

// NewOp returns a new Op.
func NewOp(op, key string, value interface{}) Op {
	return Op{Op: op, Key: key, Val: value}
}

// GetOp implements the interface Op to return the operation.
func (o Op) GetOp() string { return o.Op }

// GetKey implements the interface Op to return the operation key.
func (o Op) GetKey() string { return o.Key }

// GetValue implements the interface Op to return the operation value.
func (o Op) GetValue() interface{} { return o.Val }

// Operation returns the condition operation information.
func (o Op) Operation() (op, key string, value interface{}) {
	return o.Op, o.Key, o.Val
}
