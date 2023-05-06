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

// Pre-define some condition operations.
const (
	CondOpEqual    = "Equal"
	CondOpNotEqual = "NotEqual"

	CondOpLess      = "Less"
	CondOpLessEqual = "LessEqual"

	CondOpGreater      = "Greater"
	CondOpGreaterEqual = "GreaterEqual"

	CondOpIn    = "In"
	CondOpNotIn = "NotIn"

	CondOpIsNull    = "IsNull"
	CondOpIsNotNull = "IsNotNull"

	CondOpLike    = "Like"
	CondOpNotLike = "NotLike"

	CondOpBetween    = "Between"
	CondOpNotBetween = "NotBetween"

	// Key is compared with other key.
	CondOpEqualKey        = "EqualKey"
	CondOpNotEqualKey     = "NotEqualKey"
	CondOpLessKey         = "LessKey"
	CondOpLessEqualKey    = "LessEqualKey"
	CondOpGreaterKey      = "GreaterKey"
	CondOpGreaterEqualKey = "GreaterEqualKey"
)

// Condition represents a condition operation, which contains a meaningless
// method condition that has no arguments and returns and is just used to
// distinguish it from Setter, like
//
//	condition()
type Condition interface {
	condition()
	Oper
}

var (
	_ Condition = Op{}
	_ Condition = OpFunc(nil)
)

func (o Op) condition()     {}
func (f OpFunc) condition() {}

/// ---------------------------------------------------------------------- ///

// Equal is equal to New(CondOpEqual, key, value).
func Equal(key string, value interface{}) Op {
	return New(CondOpEqual, key, value)
}

// NotEqual is equal to New(CondOpNotEqual, key, value).
func NotEqual(key string, value interface{}) Op {
	return New(CondOpNotEqual, key, value)
}

// Less is equal to New(CondOpLess, key, value).
func Less(key string, value interface{}) Op {
	return New(CondOpLess, key, value)
}

// LessEqual is equal to New(CondOpLessEqual, key, value).
func LessEqual(key string, value interface{}) Op {
	return New(CondOpLessEqual, key, value)
}

// Greater is equal to New(CondOpGreater, key, value).
func Greater(key string, value interface{}) Op {
	return New(CondOpGreater, key, value)
}

// GreaterEqual is equal to New(CondOpGreaterEqual, key, value).
func GreaterEqual(key string, value interface{}) Op {
	return New(CondOpGreaterEqual, key, value)
}

// In is equal to New(CondOpIn, key, values).
func In(key string, values ...interface{}) Op {
	return New(CondOpIn, key, values)
}

// NotIn is equal to New(CondOpNotIn, key, values).
func NotIn(key string, values ...interface{}) Op {
	return New(CondOpNotIn, key, values)
}

// IsNull is equal to New(CondOpIsNull, key, nil).
func IsNull(key string) Op {
	return New(CondOpIsNull, key, nil)
}

// IsNotNull is equal to New(CondOpIsNotNull, key, nil).
func IsNotNull(key string) Op {
	return New(CondOpIsNotNull, key, nil)
}

// Like is equal to New(CondOpLike, key, value).
func Like(key string, value string) Op {
	return New(CondOpLike, key, value)
}

// NotLike is equal to New(CondOpNotLike, key, value).
func NotLike(key string, value string) Op {
	return New(CondOpNotLike, key, value)
}

// Between is equal to New(CondOpBetween, key, []interface{}{lower, upper}).
func Between(key string, lower, upper interface{}) Op {
	return New(CondOpBetween, key, []interface{}{lower, upper})
}

// NotBetween is equal to New(CondOpNotBetween, key, []interface{}{lower, upper}).
func NotBetween(key string, lower, upper interface{}) Op {
	return New(CondOpNotBetween, key, []interface{}{lower, upper})
}

// Eq is short for Equal.
func Eq(key string, value interface{}) Op { return Equal(key, value) }

// NotEq is short for NotEqual.
func NotEq(key string, value interface{}) Op { return NotEqual(key, value) }

// Le is short for Less.
func Le(key string, value interface{}) Op { return Less(key, value) }

// LeEq is short for LessEqual.
func LeEq(key string, value interface{}) Op { return LessEqual(key, value) }

// Gt is short for Greater.
func Gt(key string, value interface{}) Op { return Greater(key, value) }

// GtEq is short for GreaterEqual.
func GtEq(key string, value interface{}) Op { return GreaterEqual(key, value) }

/// ---------------------------------------------------------------------- ///

// Equal is equal to Equal(o.Key, value).
func (o Op) Equal(value interface{}) Op {
	return Equal(o.Key, value)
}

// NotEqual is equal to NotEqual(o.Key, value).
func (o Op) NotEqual(value interface{}) Op {
	return NotEqual(o.Key, value)
}

// Greater is equal to Greater(o.Key, value).
func (o Op) Greater(value interface{}) Op {
	return Greater(o.Key, value)
}

// GreaterEqual is equal to GreaterEqual(o.Key, value).
func (o Op) GreaterEqual(value interface{}) Op {
	return GreaterEqual(o.Key, value)
}

// Less is equal to Less(o.Key, value).
func (o Op) Less(value interface{}) Op {
	return Less(o.Key, value)
}

// LessEqual is equal to LessEqual(o.Key, value).
func (o Op) LessEqual(value interface{}) Op {
	return LessEqual(o.Key, value)
}

// In is equal to In(o.Key, values...).
func (o Op) In(values ...interface{}) Op {
	return In(o.Key, values...)
}

// NotIn is equal to NotIn(o.Key, value).
func (o Op) NotIn(values ...interface{}) Op {
	return NotIn(o.Key, values...)
}

// IsNull is equal to IsNull(o.Key).
func (o Op) IsNull() Op {
	return IsNull(o.Key)
}

// IsNotNull is equal to IsNotNull(o.Key).
func (o Op) IsNotNull() Op {
	return IsNotNull(o.Key)
}

// Like is equal to Like(o.Key, value).
func (o Op) Like(value string) Op {
	return Like(o.Key, value)
}

// NotLike is equal to NotLike(o.Key, value).
func (o Op) NotLike(value string) Op {
	return NotLike(o.Key, value)
}

// Between is equal to Between(o.Key, lower, upper).
func (o Op) Between(lower, upper interface{}) Op {
	return Between(o.Key, lower, upper)
}

// NotBetween is equal to NotBetween(o.Key, lower, upper).
func (o Op) NotBetween(lower, upper interface{}) Op {
	return NotBetween(o.Key, lower, upper)
}

// Eq is equal to Eq(o.Key, value).
func (o Op) Eq(value interface{}) Op {
	return Eq(o.Key, value)
}

// Le is equal to Le(o.Key, value).
func (o Op) Le(value interface{}) Op {
	return Le(o.Key, value)
}

// Gt is equal to Gt(o.Key, value).
func (o Op) Gt(value interface{}) Op {
	return Gt(o.Key, value)
}

// NotEq is equal to NotEq(o.Key, value).
func (o Op) NotEq(value interface{}) Op {
	return NotEq(o.Key, value)
}

// LeEq is equal to LeEq(o.Key, value).
func (o Op) LeEq(value interface{}) Op {
	return LeEq(o.Key, value)
}

// GtEq is equal to GtEq(o.Key, value).
func (o Op) GtEq(value interface{}) Op {
	return GtEq(o.Key, value)
}

/// ---------------------------------------------------------------------- ///

// EqualKey is equal to New(CondOpEqualKey, key1, key2).
func EqualKey(key1, key2 string) Op {
	return New(CondOpEqualKey, key1, key2)
}

// NotEqualKey is equal to New(CondOpNotEqualKey, key1, key2).
func NotEqualKey(key1, key2 string) Op {
	return New(CondOpNotEqualKey, key1, key2)
}

// LessKey is equal to New(CondOpLessKey, key1, key2).
func LessKey(key1, key2 string) Op {
	return New(CondOpLessKey, key1, key2)
}

// LessEqualKey is equal to New(CondOpLessEqualKey, key1, key2).
func LessEqualKey(key1, key2 string) Op {
	return New(CondOpLessEqualKey, key1, key2)
}

// GreaterKey is equal to New(CondOpGreaterKey, key1, key2).
func GreaterKey(key1, key2 string) Op {
	return New(CondOpGreaterKey, key1, key2)
}

// GreaterEqualKey is equal to New(CondOpGreaterEqualKey, key1, key2).
func GreaterEqualKey(key1, key2 string) Op {
	return New(CondOpGreaterEqualKey, key1, key2)
}

// EqualKey is equal to EqualKey(o.Key, otherKey).
func (o Op) EqualKey(otherKey string) Op {
	return EqualKey(o.Key, otherKey)
}

// NotEqualKey is equal to NotEqualKey(o.Key, otherKey).
func (o Op) NotEqualKey(otherKey string) Op {
	return NotEqualKey(o.Key, otherKey)
}

// LessKey is equal to LessKey(o.Key, otherKey).
func (o Op) LessKey(otherKey string) Op {
	return LessKey(o.Key, otherKey)
}

// LessEqualKey is equal to LessEqualKey(o.Key, otherKey).
func (o Op) LessEqualKey(otherKey string) Op {
	return LessEqualKey(o.Key, otherKey)
}

// GreaterKey is equal to GreaterKey(o.Key, otherKey).
func (o Op) GreaterKey(otherKey string) Op {
	return GreaterKey(o.Key, otherKey)
}

// GreaterEqualKey is equal to GreaterEqualKey(o.Key, otherKey).
func (o Op) GreaterEqualKey(otherKey string) Op {
	return GreaterEqualKey(o.Key, otherKey)
}
