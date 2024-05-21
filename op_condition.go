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
	KindCondition = "Condition"

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

	// The composite conditions.
	CondOpAnd = "And"
	CondOpOr  = "Or"
)

// Condition represents a condition operation.
type Condition interface {
	condition()
	Oper
}

type condition struct{ oper }

func (c condition) condition() {}

// Condition converts itself to Condition.
func (o Op) Condition() Condition { return condition{oper{o.WithKind(KindCondition)}} }

/// ---------------------------------------------------------------------- ///

// Equal is equal to New(CondOpEqual, key, value).Condition().
func Equal(key string, value interface{}) Condition {
	return New(CondOpEqual, key, value).Condition()
}

// NotEqual is equal to New(CondOpNotEqual, key, value).Condition().
func NotEqual(key string, value interface{}) Condition {
	return New(CondOpNotEqual, key, value).Condition()
}

// Less is equal to New(CondOpLess, key, value).Condition().
func Less(key string, value interface{}) Condition {
	return New(CondOpLess, key, value).Condition()
}

// LessEqual is equal to New(CondOpLessEqual, key, value).Condition().
func LessEqual(key string, value interface{}) Condition {
	return New(CondOpLessEqual, key, value).Condition()
}

// Greater is equal to New(CondOpGreater, key, value).Condition().
func Greater(key string, value interface{}) Condition {
	return New(CondOpGreater, key, value).Condition()
}

// GreaterEqual is equal to New(CondOpGreaterEqual, key, value).Condition().
func GreaterEqual(key string, value interface{}) Condition {
	return New(CondOpGreaterEqual, key, value).Condition()
}

// In is equal to New(CondOpIn, key, values).Condition().
func In(key string, values ...interface{}) Condition {
	return New(CondOpIn, key, values).Condition()
}

// NotIn is equal to New(CondOpNotIn, key, values).Condition().
func NotIn(key string, values ...interface{}) Condition {
	return New(CondOpNotIn, key, values).Condition()
}

// IsNull is equal to New(CondOpIsNull, key, nil).Condition().
func IsNull(key string) Condition {
	return New(CondOpIsNull, key, nil).Condition()
}

// IsNotNull is equal to New(CondOpIsNotNull, key, nil).Condition().
func IsNotNull(key string) Condition {
	return New(CondOpIsNotNull, key, nil).Condition()
}

// Like is equal to New(CondOpLike, key, value).Condition().
func Like(key string, value string) Condition {
	return New(CondOpLike, key, value).Condition()
}

// NotLike is equal to New(CondOpNotLike, key, value).Condition().
func NotLike(key string, value string) Condition {
	return New(CondOpNotLike, key, value).Condition()
}

// Between is equal to New(CondOpBetween, key, []interface{}{lower, upper}).Condition().
func Between(key string, lower, upper interface{}) Condition {
	return New(CondOpBetween, key, []interface{}{lower, upper}).Condition()
}

// NotBetween is equal to New(CondOpNotBetween, key, []interface{}{lower, upper}).Condition().
func NotBetween(key string, lower, upper interface{}) Condition {
	return New(CondOpNotBetween, key, []interface{}{lower, upper}).Condition()
}

// And is equal to New(CondOpAnd, "", ops).Condition().
func And(ops ...Condition) Condition {
	return New(CondOpAnd, "", ops).Condition()
}

// Or is equal to New(CondOpOr, "", ops).Condition().
func Or(ops ...Condition) Condition {
	return New(CondOpOr, "", ops).Condition()
}

// Eq is short for Equal.
func Eq(key string, value interface{}) Condition { return Equal(key, value) }

// NotEq is short for NotEqual.
func NotEq(key string, value interface{}) Condition { return NotEqual(key, value) }

// Le is short for Less.
func Le(key string, value interface{}) Condition { return Less(key, value) }

// LeEq is short for LessEqual.
func LeEq(key string, value interface{}) Condition { return LessEqual(key, value) }

// Gt is short for Greater.
func Gt(key string, value interface{}) Condition { return Greater(key, value) }

// GtEq is short for GreaterEqual.
func GtEq(key string, value interface{}) Condition { return GreaterEqual(key, value) }

/// ---------------------------------------------------------------------- ///

// Equal is equal to Equal(o.Key, value).
func (o Op) Equal(value interface{}) Condition {
	return Equal(o.Key, value)
}

// NotEqual is equal to NotEqual(o.Key, value).
func (o Op) NotEqual(value interface{}) Condition {
	return NotEqual(o.Key, value)
}

// Greater is equal to Greater(o.Key, value).
func (o Op) Greater(value interface{}) Condition {
	return Greater(o.Key, value)
}

// GreaterEqual is equal to GreaterEqual(o.Key, value).
func (o Op) GreaterEqual(value interface{}) Condition {
	return GreaterEqual(o.Key, value)
}

// Less is equal to Less(o.Key, value).
func (o Op) Less(value interface{}) Condition {
	return Less(o.Key, value)
}

// LessEqual is equal to LessEqual(o.Key, value).
func (o Op) LessEqual(value interface{}) Condition {
	return LessEqual(o.Key, value)
}

// In is equal to In(o.Key, values...).
func (o Op) In(values ...interface{}) Condition {
	return In(o.Key, values...)
}

// NotIn is equal to NotIn(o.Key, value).
func (o Op) NotIn(values ...interface{}) Condition {
	return NotIn(o.Key, values...)
}

// IsNull is equal to IsNull(o.Key).
func (o Op) IsNull() Condition {
	return IsNull(o.Key)
}

// IsNotNull is equal to IsNotNull(o.Key).
func (o Op) IsNotNull() Condition {
	return IsNotNull(o.Key)
}

// Like is equal to Like(o.Key, value).
func (o Op) Like(value string) Condition {
	return Like(o.Key, value)
}

// NotLike is equal to NotLike(o.Key, value).
func (o Op) NotLike(value string) Condition {
	return NotLike(o.Key, value)
}

// Between is equal to Between(o.Key, lower, upper).
func (o Op) Between(lower, upper interface{}) Condition {
	return Between(o.Key, lower, upper)
}

// NotBetween is equal to NotBetween(o.Key, lower, upper).
func (o Op) NotBetween(lower, upper interface{}) Condition {
	return NotBetween(o.Key, lower, upper)
}

// Eq is equal to Eq(o.Key, value).
func (o Op) Eq(value interface{}) Condition {
	return Eq(o.Key, value)
}

// Le is equal to Le(o.Key, value).
func (o Op) Le(value interface{}) Condition {
	return Le(o.Key, value)
}

// Gt is equal to Gt(o.Key, value).
func (o Op) Gt(value interface{}) Condition {
	return Gt(o.Key, value)
}

// NotEq is equal to NotEq(o.Key, value).
func (o Op) NotEq(value interface{}) Condition {
	return NotEq(o.Key, value)
}

// LeEq is equal to LeEq(o.Key, value).
func (o Op) LeEq(value interface{}) Condition {
	return LeEq(o.Key, value)
}

// GtEq is equal to GtEq(o.Key, value).
func (o Op) GtEq(value interface{}) Condition {
	return GtEq(o.Key, value)
}

/// ---------------------------------------------------------------------- ///

// EqualKey is equal to New(CondOpEqualKey, leftKey, rightKey).Condition().
func EqualKey(leftKey, rightKey string) Condition {
	return New(CondOpEqualKey, leftKey, rightKey).Condition()
}

// NotEqualKey is equal to New(CondOpNotEqualKey, leftKey, rightKey).Condition().
func NotEqualKey(leftKey, rightKey string) Condition {
	return New(CondOpNotEqualKey, leftKey, rightKey).Condition()
}

// LessKey is equal to New(CondOpLessKey, leftKey, rightKey).Condition().
func LessKey(leftKey, rightKey string) Condition {
	return New(CondOpLessKey, leftKey, rightKey).Condition()
}

// LessEqualKey is equal to New(CondOpLessEqualKey, leftKey, rightKey).Condition().
func LessEqualKey(leftKey, rightKey string) Condition {
	return New(CondOpLessEqualKey, leftKey, rightKey).Condition()
}

// GreaterKey is equal to New(CondOpGreaterKey, leftKey, rightKey).Condition().
func GreaterKey(leftKey, rightKey string) Condition {
	return New(CondOpGreaterKey, leftKey, rightKey).Condition()
}

// GreaterEqualKey is equal to New(CondOpGreaterEqualKey, leftKey, rightKey).Condition().
func GreaterEqualKey(leftKey, rightKey string) Condition {
	return New(CondOpGreaterEqualKey, leftKey, rightKey).Condition()
}

// EqualKey is equal to EqualKey(o.Key, otherKey).
func (o Op) EqualKey(otherKey string) Condition {
	return EqualKey(o.Key, otherKey)
}

// NotEqualKey is equal to NotEqualKey(o.Key, otherKey).
func (o Op) NotEqualKey(otherKey string) Condition {
	return NotEqualKey(o.Key, otherKey)
}

// LessKey is equal to LessKey(o.Key, otherKey).
func (o Op) LessKey(otherKey string) Condition {
	return LessKey(o.Key, otherKey)
}

// LessEqualKey is equal to LessEqualKey(o.Key, otherKey).
func (o Op) LessEqualKey(otherKey string) Condition {
	return LessEqualKey(o.Key, otherKey)
}

// GreaterKey is equal to GreaterKey(o.Key, otherKey).
func (o Op) GreaterKey(otherKey string) Condition {
	return GreaterKey(o.Key, otherKey)
}

// GreaterEqualKey is equal to GreaterEqualKey(o.Key, otherKey).
func (o Op) GreaterEqualKey(otherKey string) Condition {
	return GreaterEqualKey(o.Key, otherKey)
}
