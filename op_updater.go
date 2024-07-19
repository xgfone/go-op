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

// Pre-define some update operations.
const (
	KindUpdate = "Update"

	UpdateOpBatch = "Batch"
	UpdateOpSet   = "Set"
	UpdateOpInc   = "Inc"
	UpdateOpDec   = "Dec"
	UpdateOpAdd   = "Add"
	UpdateOpSub   = "Sub"
	UpdateOpMul   = "Mul"
	UpdateOpDiv   = "Div"
)

// KV represents a key-value pair.
type KV struct {
	Key string
	Val any
}

// Updater represents a update operation.
type Updater interface {
	update()
	Oper
}

type updater struct{ oper }

func (u updater) update() {}

// Updater converts itself to Updater.
func (o Op) Updater() Updater { return updater{oper{o.WithKind(KindUpdate)}} }

/// ---------------------------------------------------------------------- ///

// Batch is equal to New(UpdateOpBatch, key, nil).Updater().
func Batch(ups ...Updater) Updater {
	switch len(ups) {
	case 0:
		return nil
	case 1:
		return ups[0]
	default:
		return New(UpdateOpBatch, "", ups).Updater()
	}
}

// Inc is equal to Key(key).Inc().
func Inc(key string) Updater {
	return Key(key).Inc()
}

// Dec is equal to Key(key).Dec().
func Dec(key string) Updater {
	return Key(key).Dec()
}

// Add is equal to Key(key).Add(value).
func Add(key string, value any) Updater {
	return Key(key).Add(value)
}

// Sub is equal to Key(key).Sub(value).
func Sub(key string, value any) Updater {
	return Key(key).Sub(value)
}

// Mul is equal to Key(key).Mul(value).
func Mul(key string, value any) Updater {
	return Key(key).Mul(value)
}

// Div is equal to Key(key).Div(value).
func Div(key string, value any) Updater {
	return Key(key).Div(value)
}

// Set is equal to Key(key).Set(value).
func Set(key string, value any) Updater {
	return Key(key).Set(value)
}

// Inc is equal to o.WithOp(UpdateOpInc).Updater().
func (o Op) Inc() Updater {
	return o.WithOp(UpdateOpInc).Updater()
}

// Dec is equal to o.WithOp(UpdateOpDec).Updater().
func (o Op) Dec() Updater {
	return o.WithOp(UpdateOpDec).Updater()
}

// Add is equal to o.WithOp(UpdateOpAdd).WithValue(value).Updater().
func (o Op) Add(value any) Updater {
	return o.WithOp(UpdateOpAdd).WithValue(value).Updater()
}

// Sub is equal to o.WithOp(UpdateOpSub).WithValue(value).Updater().
func (o Op) Sub(value any) Updater {
	return o.WithOp(UpdateOpSub).WithValue(value).Updater()
}

// Mul is equal to o.WithOp(UpdateOpMul).WithValue(value).Updater().
func (o Op) Mul(value any) Updater {
	return o.WithOp(UpdateOpMul).WithValue(value).Updater()
}

// Div is equal to o.WithOp(UpdateOpDiv).WithValue(value).Updater().
func (o Op) Div(value any) Updater {
	return o.WithOp(UpdateOpDiv).WithValue(value).Updater()
}

// Set is equal to o.WithOp(UpdateOpSet).WithValue(value).Updater().
func (o Op) Set(value any) Updater {
	return o.WithOp(UpdateOpSet).WithValue(value).Updater()
}

// AddKey is equal to o.WithOp(UpdateOpAdd).WithValue(KV{Key: key, Val: value}).Updater().
func (o Op) AddKey(key string, value any) Updater {
	return o.WithOp(UpdateOpAdd).WithValue(KV{Key: key, Val: value}).Updater()
}

// SubKey is equal to o.WithOp(UpdateOpSub).WithValue(KV{Key: key, Val: value}).Updater().
func (o Op) SubKey(key string, value any) Updater {
	return o.WithOp(UpdateOpSub).WithValue(KV{Key: key, Val: value}).Updater()
}

// MulKey is equal to o.WithOp(UpdateOpMul).WithValue(KV{Key: key, Val: value}).Updater().
func (o Op) MulKey(key string, value any) Updater {
	return o.WithOp(UpdateOpMul).WithValue(KV{Key: key, Val: value}).Updater()
}

// DivKey is equal to o.WithOp(UpdateOpDiv).WithValue(KV{Key: key, Val: value}).Updater().
func (o Op) DivKey(key string, value any) Updater {
	return o.WithOp(UpdateOpDiv).WithValue(KV{Key: key, Val: value}).Updater()
}
