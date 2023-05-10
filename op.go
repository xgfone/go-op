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

import "fmt"

// Oper is a common operation interface.
type Oper interface {
	Operation() Op
}

var _ Oper = Func(nil)

// Func is an operation function, which implements the interface Oper,
// Setter and Condition.
type Func func() Op

// Operation implements the interface Oper.
func (f Func) Operation() Op { return f() }

// ContainsKey reports whether the operations contains the key.
func ContainsKey[S ~[]E, E Oper](ops S, key string) bool {
	for _, op := range ops {
		if op.Operation().Key == key {
			return true
		}
	}
	return false
}

// Contains reports whether the operations contains the operation by the key.
func Contains[S ~[]E1, E1, E2 Oper](ops S, op E2) bool {
	return ContainsKey(ops, op.Operation().Key)
}

/// ----------------------------------------------------------------------- ///

var _ Oper = Op{}

// Op represents an operation, which has implemented the interface Oper,
// Setter and Condition.
type Op struct {
	Op  string
	Key string
	Val interface{}
}

// Key is equal to New("", key, nil).
func Key(key string) Op { return New("", key, nil) }

// New returns a new Op.
func New(op, key string, value interface{}) Op {
	return Op{Op: op, Key: key, Val: value}
}

// Operation returns the condition operation information.
func (o Op) Operation() Op { return o }

// Prefix is short for KeyPrefix.
func (o Op) Prefix(prefix string) Op { return o.KeyPrefix(prefix) }

// Suffix is short for KeySuffix.
func (o Op) Suffix(suffix string) Op { return o.KeySuffix(suffix) }

// KeyPrefix returns a new Op, which uses prefix as the prefix of the key.
func (o Op) KeyPrefix(prefix string) Op {
	o.Key = prefix + o.Key
	return o
}

// KeySuffix returns a new Op, which uses suffix as the suffix of the key.
func (o Op) KeySuffix(suffix string) Op {
	o.Key += suffix
	return o
}

// WithKey replaces the key with the new key.
func (o Op) WithKey(key string) Op {
	o.Key = key
	return o
}

func (o Op) String() string {
	return fmt.Sprintf("Op(key=%s, op=%s, value=%v)", o.Key, o.Op, o.Val)
}
