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
type Sorter interface {
	sort()
	Oper
}

type sort struct{ oper }

func (s sort) sort() {}

// Sorter converts itself to Sorter.
func (o Op) Sorter() Sorter { return sort{oper{o.WithKind(KindSort)}} }

/// ---------------------------------------------------------------------- ///

// Order is equal to New(SortOpOrder, key, order).Sorter().
//
// order may be SortAsc or SortDesc.
func Order(key, order string) Sorter {
	return New(SortOpOrder, key, order).Sorter()
}

// Orders is equal to New(SortOpOrders, "", orders).Sorter().
func Orders(orders ...Sorter) Sorter {
	return New(SortOpOrders, "", orders).Sorter()
}

func (o Op) OrderAsc() Sorter  { return o.Order(SortAsc) }
func (o Op) OrderDesc() Sorter { return o.Order(SortDesc) }

// Order is equal to Order(o.Key, order).
func (o Op) Order(order string) Sorter { return Order(o.Key, order) }

// Orders is equal to Orders(orders...).
func (o Op) Orders(orders ...Sorter) Sorter { return Orders(orders...) }
