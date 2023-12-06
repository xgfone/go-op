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

// NewUpdater converts an op to Updater.
func NewUpdater(op Op) Updater { return op.Updater() }

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

// Inc is equal to Inc(o.Key).
func (o Op) Inc() Updater {
	return Inc(o.Key)
}

// Dec is equal to Dec(o.Key).
func (o Op) Dec() Updater {
	return Dec(o.Key)
}

// Add is equal to Add(o.Key, value).
func (o Op) Add(value interface{}) Updater {
	return Add(o.Key, value)
}

// Sub is equal to Sub(o.Key, value).
func (o Op) Sub(value interface{}) Updater {
	return Sub(o.Key, value)
}

// Mul is equal to Mul(o.Key, value).
func (o Op) Mul(value interface{}) Updater {
	return Mul(o.Key, value)
}

// Div is equal to Div(o.Key, value).
func (o Op) Div(value interface{}) Updater {
	return Div(o.Key, value)
}

// Set is equal to Set(o.Key, value).
func (o Op) Set(value interface{}) Updater {
	return Set(o.Key, value)
}

// AddKey is equal to Add(o.Key, KeyValue{Key: key, Val: value}).
func (o Op) AddKey(key string, value interface{}) Updater {
	return Add(o.Key, KeyValue{Key: key, Val: value})
}

// SubKey is equal to Sub(o.Key, KeyValue{Key: key, Val: value}).
func (o Op) SubKey(key string, value interface{}) Updater {
	return Sub(o.Key, KeyValue{Key: key, Val: value})
}

// MulKey is equal to Mul(o.Key, KeyValue{Key: key, Val: value}).
func (o Op) MulKey(key string, value interface{}) Updater {
	return Mul(o.Key, KeyValue{Key: key, Val: value})
}

// DivKey is equal to Div(o.Key, KeyValue{Key: key, Val: value}).
func (o Op) DivKey(key string, value interface{}) Updater {
	return Div(o.Key, KeyValue{Key: key, Val: value})
}
