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

// KeyValue represents a key-value pair.
type KeyValue struct {
	Key string
	Val interface{}
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

// Inc is equal to New(UpdateOpInc, key, nil).Updater().
func Inc(key string) Updater {
	return New(UpdateOpInc, key, nil).Updater()
}

// Dec is equal to New(UpdateOpDec, key, nil).Updater().
func Dec(key string) Updater {
	return New(UpdateOpDec, key, nil).Updater()
}

// Add is equal to New(UpdateOpAdd, key, value).Updater().
func Add(key string, value interface{}) Updater {
	return New(UpdateOpAdd, key, value).Updater()
}

// Sub is equal to New(UpdateOpSub, key, value).Updater().
func Sub(key string, value interface{}) Updater {
	return New(UpdateOpSub, key, value).Updater()
}

// Mul is equal to New(UpdateOpMul, key, value).Updater().
func Mul(key string, value interface{}) Updater {
	return New(UpdateOpMul, key, value).Updater()
}

// Div is equal to New(UpdateOpDiv, key, value).Updater().
func Div(key string, value interface{}) Updater {
	return New(UpdateOpDiv, key, value).Updater()
}

// Set is equal to New(UpdateOpSet, key, value).Updater().
func Set(key string, value interface{}) Updater {
	return New(UpdateOpSet, key, value).Updater()
}

// Inc is the same as Inc(o.Key).
func (o Op) Inc() Updater {
	return o.WithOp(UpdateOpInc).Updater()
}

// Dec is the same as Dec(o.Key).
func (o Op) Dec() Updater {
	return o.WithOp(UpdateOpDec).Updater()
}

// Add is the same as Add(o.Key, value).
func (o Op) Add(value interface{}) Updater {
	return o.WithOp(UpdateOpAdd).WithValue(value).Updater()
}

// Sub is the same as Sub(o.Key, value).
func (o Op) Sub(value interface{}) Updater {
	return o.WithOp(UpdateOpSub).WithValue(value).Updater()
}

// Mul is the same as Mul(o.Key, value).
func (o Op) Mul(value interface{}) Updater {
	return o.WithOp(UpdateOpMul).WithValue(value).Updater()
}

// Div is the same as Div(o.Key, value).
func (o Op) Div(value interface{}) Updater {
	return o.WithOp(UpdateOpDiv).WithValue(value).Updater()
}

// Set is the same as Set(o.Key, value).
func (o Op) Set(value interface{}) Updater {
	return o.WithOp(UpdateOpSet).WithValue(value).Updater()
}

// AddKey is the same as Add(o.Key, KeyValue{Key: key, Val: value}).
func (o Op) AddKey(key string, value interface{}) Updater {
	return o.WithOp(UpdateOpAdd).WithValue(KeyValue{Key: key, Val: value}).Updater()
}

// SubKey is the same as Sub(o.Key, KeyValue{Key: key, Val: value}).
func (o Op) SubKey(key string, value interface{}) Updater {
	return o.WithOp(UpdateOpSub).WithValue(KeyValue{Key: key, Val: value}).Updater()
}

// MulKey is the same as Mul(o.Key, KeyValue{Key: key, Val: value}).
func (o Op) MulKey(key string, value interface{}) Updater {
	return o.WithOp(UpdateOpMul).WithValue(KeyValue{Key: key, Val: value}).Updater()
}

// DivKey is the same as Div(o.Key, KeyValue{Key: key, Val: value}).
func (o Op) DivKey(key string, value interface{}) Updater {
	return o.WithOp(UpdateOpDiv).WithValue(KeyValue{Key: key, Val: value}).Updater()
}
