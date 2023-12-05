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

// Package op provides a common operation, such as Condition and Updater.
package op

import (
	"fmt"
	"strings"
)

// Oper is a common operation interface.
type Oper interface {
	Op() Op
}

type oper struct{ op Op }

func (o oper) Op() Op { return o.op }

// ContainsKey reports whether the operations contains the key.
func ContainsKey[S ~[]E, E Oper](ops S, key string) bool {
	for _, op := range ops {
		if op.Op().Key == key {
			return true
		}
	}
	return false
}

// Contains reports whether the operations contains the operation by the key.
func Contains[S ~[]E1, E1, E2 Oper](ops S, op E2) bool {
	return ContainsKey(ops, op.Op().Key)
}

/// ----------------------------------------------------------------------- ///

// Op represents an operation.
type Op struct {
	// Required
	Op  string
	Key string
	Val interface{}

	// Optional
	Kind string
	Tags map[string]string
}

// Key is equal to New("", key, nil).
func Key(key string) Op { return New("", key, nil) }

// New returns a new Op.
func New(op, key string, value interface{}) Op {
	return Op{Op: op, Key: key, Val: value}
}

// Oper converts itself to an Oper.
func (o Op) Oper() Oper { return oper{o} }

// Prefix is short for KeyPrefix.
func (o Op) Prefix(prefix string) Op { return o.KeyPrefix(prefix) }

// Suffix is short for KeySuffix.
func (o Op) Suffix(suffix string) Op { return o.KeySuffix(suffix) }

// Scope is equal to o.Prefix(name + ".").
func (o Op) Scope(name string) Op {
	switch {
	case len(name) == 0:
		return o

	case len(o.Key) == 0:
		return o.WithKey(name)

	default:
		o.Key = strings.Join([]string{name, o.Key}, ".")
		return o
	}
}

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

// WithOp replaces the op with the new and returns a new Op.
func (o Op) WithOp(op string) Op {
	o.Op = op
	return o
}

// WithKey replaces the key with the new and returns a new Op.
func (o Op) WithKey(key string) Op {
	o.Key = key
	return o
}

// WithValue replaces the value with the new and returns a new Op.
func (o Op) WithValue(value interface{}) Op {
	o.Val = value
	return o
}

// WithKind replaces the kind with the new and returns a new Op.
func (o Op) WithKind(kind string) Op {
	o.Kind = kind
	return o
}

// WithTags replaces the tags with the new and returns a new Op.
func (o Op) WithTags(tags map[string]string) Op {
	o.Tags = tags
	return o
}

// WithTag clones the tags and appends the new key-value tag, and returns a new Op.
func (o Op) WithTag(tkey, tvalue string) Op {
	if o.Tags == nil {
		o.Tags = map[string]string{tkey: tvalue}
	} else {
		tags := make(map[string]string, len(o.Tags)+1)
		for k, v := range o.Tags {
			tags[k] = v
		}
		o.Tags = tags
	}
	return o
}

// AppendTag appends the new key-value tag into the tags, and returns a new Op.
func (o Op) AppendTag(tkey, tvalue string) Op {
	if o.Tags == nil {
		o.Tags = map[string]string{tkey: tvalue}
	} else {
		o.Tags[tkey] = tvalue
	}
	return o
}

// Name returns the name of the operation.
func (o Op) Name(name string) string {
	if name := o.Tags[name]; name != "" {
		return name
	}
	return o.Key
}

func (o Op) String() string {
	if o.Kind == "" {
		return fmt.Sprintf("Op(key=%s, op=%s, value=%v)", o.Key, o.Op, o.Val)
	}
	return fmt.Sprintf("Op(kind=%s, key=%s, op=%s, value=%v)", o.Kind, o.Key, o.Op, o.Val)
}

// IsOp reports whether the operation is equal to op.
func (o Op) IsOp(op string) bool { return o.Op == op }

// IsKind reports whether the operation kind is equal to kind.
func (o Op) IsKind(kind string) bool { return o.Kind == kind }
