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

// Equal is equal to Key(key).Equal(value).
func Equal(key string, value interface{}) Condition {
	return Key(key).Equal(value)
}

// NotEqual is equal to Key(key).NotEqual(value).
func NotEqual(key string, value interface{}) Condition {
	return Key(key).NotEqual(value)
}

// Less is equal to Key(key).Less(value).
func Less(key string, value interface{}) Condition {
	return Key(key).Less(value)
}

// LessEqual is equal to Key(key).LessEqual(value).
func LessEqual(key string, value interface{}) Condition {
	return Key(key).LessEqual(value)
}

// Greater is equal to Key(key).Greater(value).
func Greater(key string, value interface{}) Condition {
	return Key(key).Greater(value)
}

// GreaterEqual is equal to Key(key).GreaterEqual(value).
func GreaterEqual(key string, value interface{}) Condition {
	return Key(key).GreaterEqual(value)
}

// In is equal to Key(key).In(values).
func In(key string, values interface{}) Condition {
	return Key(key).In(values)
}

// NotIn is equal to Key(key).NotIn(values).
func NotIn(key string, values interface{}) Condition {
	return Key(key).NotIn(values)
}

// IsNull is equal to Key(key).IsNull().
func IsNull(key string) Condition {
	return Key(key).IsNull()
}

// IsNotNull is equal to Key(key).IsNotNull().
func IsNotNull(key string) Condition {
	return Key(key).IsNotNull()
}

// Like is equal to Key(key).Like(value).
func Like(key string, value string) Condition {
	return Key(key).Like(value)
}

// NotLike is equal to Key(key).NotLike(value).
func NotLike(key string, value string) Condition {
	return Key(key).NotLike(value)
}

// Between is equal to Key(key).Between(lower, upper).
func Between(key string, lower, upper interface{}) Condition {
	return Key(key).Between(lower, upper)
}

// NotBetween is equal to Key(key).NotBetween(lower, upper).
func NotBetween(key string, lower, upper interface{}) Condition {
	return Key(key).NotBetween(lower, upper)
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

// Equal is equal to o.WithOp(CondOpEqual).WithValue(value).Condition().
func (o Op) Equal(value interface{}) Condition {
	return o.WithOp(CondOpEqual).WithValue(value).Condition()
}

// NotEqual is equal to o.WithOp(CondOpNotEqual).WithValue(value).Condition().
func (o Op) NotEqual(value interface{}) Condition {
	return o.WithOp(CondOpNotEqual).WithValue(value).Condition()
}

// Greater is equal to o.WithOp(CondOpGreater).WithValue(value).Condition().
func (o Op) Greater(value interface{}) Condition {
	return o.WithOp(CondOpGreater).WithValue(value).Condition()
}

// GreaterEqual is equal to o.WithOp(CondOpGreaterEqual).WithValue(value).Condition().
func (o Op) GreaterEqual(value interface{}) Condition {
	return o.WithOp(CondOpGreaterEqual).WithValue(value).Condition()
}

// Less is equal to o.WithOp(CondOpLess).WithValue(value).Condition().
func (o Op) Less(value interface{}) Condition {
	return o.WithOp(CondOpLess).WithValue(value).Condition()
}

// LessEqual is equal to o.WithOp(CondOpLessEqual).WithValue(value).Condition().
func (o Op) LessEqual(value interface{}) Condition {
	return o.WithOp(CondOpLessEqual).WithValue(value).Condition()
}

// In is equal to o.WithOp(CondOpIn).WithValue(values).Condition().
func (o Op) In(values interface{}) Condition {
	return o.WithOp(CondOpIn).WithValue(values).Condition()
}

// NotIn is equal to o.WithOp(CondOpNotIn).WithValue(values).Condition().
func (o Op) NotIn(values interface{}) Condition {
	return o.WithOp(CondOpNotIn).WithValue(values).Condition()
}

// IsNull is equal to o.WithOp(CondOpIsNull).Condition().
func (o Op) IsNull() Condition {
	return o.WithOp(CondOpIsNull).Condition()
}

// IsNotNull is equal to o.WithOp(CondOpIsNotNull).Condition().
func (o Op) IsNotNull() Condition {
	return o.WithOp(CondOpIsNotNull).Condition()
}

// Like is equal to o.WithOp(CondOpLike).WithValue(value).Condition().
func (o Op) Like(value string) Condition {
	return o.WithOp(CondOpLike).WithValue(value).Condition()
}

// NotLike is equal to o.WithOp(CondOpNotLike).WithValue(value).Condition().
func (o Op) NotLike(value string) Condition {
	return o.WithOp(CondOpNotLike).WithValue(value).Condition()
}

// Between is equal to o.WithOp(CondOpBetween).WithValue([]interface{}{lower, upper}).Condition().
func (o Op) Between(lower, upper interface{}) Condition {
	return o.WithOp(CondOpBetween).WithValue([]interface{}{lower, upper}).Condition()
}

// NotBetween is equal to o.WithOp(CondOpNotBetween).WithValue([]interface{}{lower, upper}).Condition().
func (o Op) NotBetween(lower, upper interface{}) Condition {
	return o.WithOp(CondOpNotBetween).WithValue([]interface{}{lower, upper}).Condition()
}

// Eq is equal to o.Equal(value).
func (o Op) Eq(value interface{}) Condition {
	return o.Equal(value)
}

// Le is equal to o.Less(value).
func (o Op) Le(value interface{}) Condition {
	return o.Less(value)
}

// Gt is equal to o.Greater(value).
func (o Op) Gt(value interface{}) Condition {
	return o.Greater(value)
}

// NotEq is equal to o.NotEqual(value).
func (o Op) NotEq(value interface{}) Condition {
	return o.NotEqual(value)
}

// LeEq is equal to o.LessEqual(value).
func (o Op) LeEq(value interface{}) Condition {
	return o.LessEqual(value)
}

// GtEq is equal to o.GreaterEqual(value).
func (o Op) GtEq(value interface{}) Condition {
	return o.GreaterEqual(value)
}

/// ---------------------------------------------------------------------- ///

// EqualKey is equal to Key(leftKey).EqualKey(rightKey).
func EqualKey(leftKey, rightKey string) Condition {
	return Key(leftKey).EqualKey(rightKey)
}

// NotEqualKey is equal to Key(leftKey).NotEqualKey(rightKey).
func NotEqualKey(leftKey, rightKey string) Condition {
	return Key(leftKey).NotEqualKey(rightKey)
}

// LessKey is equal to Key(leftKey).LessKey(rightKey).
func LessKey(leftKey, rightKey string) Condition {
	return Key(leftKey).LessKey(rightKey)
}

// LessEqualKey is equal to Key(leftKey).LessEqualKey(rightKey).
func LessEqualKey(leftKey, rightKey string) Condition {
	return Key(leftKey).LessEqualKey(rightKey)
}

// GreaterKey is equal to Key(leftKey).GreaterKey(rightKey).
func GreaterKey(leftKey, rightKey string) Condition {
	return Key(leftKey).GreaterKey(rightKey)
}

// GreaterEqualKey is equal to Key(leftKey).GreaterEqualKey(rightKey).
func GreaterEqualKey(leftKey, rightKey string) Condition {
	return Key(leftKey).GreaterEqualKey(rightKey)
}

// EqualKey is equal to o.WithOp(CondOpEqualKey).WithValue(otherKey).Condition().
func (o Op) EqualKey(otherKey string) Condition {
	return o.WithOp(CondOpEqualKey).WithValue(otherKey).Condition()
}

// NotEqualKey is equal to o.WithOp(CondOpNotEqualKey).WithValue(otherKey).Condition().
func (o Op) NotEqualKey(otherKey string) Condition {
	return o.WithOp(CondOpNotEqualKey).WithValue(otherKey).Condition()
}

// LessKey is equal to o.WithOp(CondOpLessKey).WithValue(otherKey).Condition().
func (o Op) LessKey(otherKey string) Condition {
	return o.WithOp(CondOpLessKey).WithValue(otherKey).Condition()
}

// LessEqualKey is equal to o.WithOp(CondOpLessEqualKey).WithValue(otherKey).Condition().
func (o Op) LessEqualKey(otherKey string) Condition {
	return o.WithOp(CondOpLessEqualKey).WithValue(otherKey).Condition()
}

// GreaterKey is equal to o.WithOp(CondOpGreaterKey).WithValue(otherKey).Condition().
func (o Op) GreaterKey(otherKey string) Condition {
	return o.WithOp(CondOpGreaterKey).WithValue(otherKey).Condition()
}

// GreaterEqualKey is equal to o.WithOp(CondOpGreaterEqualKey).WithValue(otherKey).Condition().
func (o Op) GreaterEqualKey(otherKey string) Condition {
	return o.WithOp(CondOpGreaterEqualKey).WithValue(otherKey).Condition()
}
