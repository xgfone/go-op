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

// Pre-define some sort operations.
const (
	KindSort     = "Sort"
	SortOpOrder  = "Order"
	SortOpOrders = "Orders"
)

// Pre-define some sort methods.
const (
	SortAsc  = "Asc"
	SortDesc = "Desc"
)

// Sort represents a sort operation
type Sort interface {
	sort()
	Oper
}

type sort struct{ oper }

func (s sort) sort() {}

// NewSort converts an op to Sort.
func NewSort(op Op) Sort { return op.Sort() }

// Sort converts itself to Sort.
func (o Op) Sort() Sort { return sort{oper{o.WithKind(KindSort)}} }

/// ---------------------------------------------------------------------- ///

// Order is equal to New(SortOpOrder, key, order).Sort().
//
// order may be SortAsc or SortDesc.
func Order(key, order string) Sort {
	return New(SortOpOrder, key, order).Sort()
}

// Orders is equal to New(SortOpOrders, "", orders).Sort().
func Orders(orders ...Sort) Sort {
	return New(SortOpOrders, "", orders).Sort()
}

func (o Op) OrderAsc() Sort  { return o.Order(SortAsc) }
func (o Op) OrderDesc() Sort { return o.Order(SortDesc) }

// Order is equal to Order(o.Key, order).
func (o Op) Order(order string) Sort { return Order(o.Key, order) }

// Orders is equal to Orders(orders...).
func (o Op) Orders(orders ...Sort) Sort { return Orders(orders...) }
