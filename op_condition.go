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

// Equal is the same as Equal(o.Key, value).
func (o Op) Equal(value interface{}) Condition {
	return o.WithOp(CondOpEqual).WithValue(value).Condition()
}

// NotEqual is the same as NotEqual(o.Key, value).
func (o Op) NotEqual(value interface{}) Condition {
	return o.WithOp(CondOpNotEqual).WithValue(value).Condition()
}

// Greater is the same as Greater(o.Key, value).
func (o Op) Greater(value interface{}) Condition {
	return o.WithOp(CondOpGreater).WithValue(value).Condition()
}

// GreaterEqual is the same as GreaterEqual(o.Key, value).
func (o Op) GreaterEqual(value interface{}) Condition {
	return o.WithOp(CondOpGreaterEqual).WithValue(value).Condition()
}

// Less is the same as Less(o.Key, value).
func (o Op) Less(value interface{}) Condition {
	return o.WithOp(CondOpLess).WithValue(value).Condition()
}

// LessEqual is the same as LessEqual(o.Key, value).
func (o Op) LessEqual(value interface{}) Condition {
	return o.WithOp(CondOpLessEqual).WithValue(value).Condition()
}

// In is the same as In(o.Key, values...).
func (o Op) In(values ...interface{}) Condition {
	return o.WithOp(CondOpIn).WithValue(values).Condition()
}

// NotIn is the same as NotIn(o.Key, value).
func (o Op) NotIn(values ...interface{}) Condition {
	return o.WithOp(CondOpNotIn).WithValue(values).Condition()
}

// IsNull is the same as IsNull(o.Key).
func (o Op) IsNull() Condition {
	return o.WithOp(CondOpIsNull).Condition()
}

// IsNotNull is the same as IsNotNull(o.Key).
func (o Op) IsNotNull() Condition {
	return o.WithOp(CondOpIsNotNull).Condition()
}

// Like is the same as Like(o.Key, value).
func (o Op) Like(value string) Condition {
	return o.WithOp(CondOpLike).WithValue(value).Condition()
}

// NotLike is the same as NotLike(o.Key, value).
func (o Op) NotLike(value string) Condition {
	return o.WithOp(CondOpNotLike).WithValue(value).Condition()
}

// Between is the same as Between(o.Key, lower, upper).
func (o Op) Between(lower, upper interface{}) Condition {
	return o.WithOp(CondOpBetween).WithValue([]interface{}{lower, upper}).Condition()
}

// NotBetween is the same as NotBetween(o.Key, lower, upper).
func (o Op) NotBetween(lower, upper interface{}) Condition {
	return o.WithOp(CondOpNotBetween).WithValue([]interface{}{lower, upper}).Condition()
}

// Eq is the same as Eq(o.Key, value).
func (o Op) Eq(value interface{}) Condition {
	return o.Equal(value)
}

// Le is the same as Le(o.Key, value).
func (o Op) Le(value interface{}) Condition {
	return o.Less(value)
}

// Gt is the same as Gt(o.Key, value).
func (o Op) Gt(value interface{}) Condition {
	return o.Greater(value)
}

// NotEq is the same as NotEq(o.Key, value).
func (o Op) NotEq(value interface{}) Condition {
	return o.NotEqual(value)
}

// LeEq is the same as LeEq(o.Key, value).
func (o Op) LeEq(value interface{}) Condition {
	return o.LessEqual(value)
}

// GtEq is the same as GtEq(o.Key, value).
func (o Op) GtEq(value interface{}) Condition {
	return o.GreaterEqual(value)
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

// EqualKey is the same as EqualKey(o.Key, otherKey).
func (o Op) EqualKey(otherKey string) Condition {
	return o.WithOp(CondOpEqualKey).WithValue(otherKey).Condition()
}

// NotEqualKey is the same as NotEqualKey(o.Key, otherKey).
func (o Op) NotEqualKey(otherKey string) Condition {
	return o.WithOp(CondOpNotEqualKey).WithValue(otherKey).Condition()
}

// LessKey is the same as LessKey(o.Key, otherKey).
func (o Op) LessKey(otherKey string) Condition {
	return o.WithOp(CondOpLessKey).WithValue(otherKey).Condition()
}

// LessEqualKey is the same as LessEqualKey(o.Key, otherKey).
func (o Op) LessEqualKey(otherKey string) Condition {
	return o.WithOp(CondOpLessEqualKey).WithValue(otherKey).Condition()
}

// GreaterKey is the same as GreaterKey(o.Key, otherKey).
func (o Op) GreaterKey(otherKey string) Condition {
	return o.WithOp(CondOpGreaterKey).WithValue(otherKey).Condition()
}

// GreaterEqualKey is the same as GreaterEqualKey(o.Key, otherKey).
func (o Op) GreaterEqualKey(otherKey string) Condition {
	return o.WithOp(CondOpGreaterEqualKey).WithValue(otherKey).Condition()
}
