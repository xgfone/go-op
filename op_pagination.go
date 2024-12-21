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

// Pre-define some pagination operations.
const (
	KindPagination = "Pagination"

	PaginationOpPageSize = "PageSize"
)

// Limiter represents a number limiter of the objects.
type Limiter interface {
	Limit() int
}

// Paginator represents a pagination operation.
type Paginator interface {
	paginate()
	Oper
}

type paginator struct{ oper }

func (p paginator) paginate() {}

// Paginator converts itself to Paginator.
func (o Op) Paginator() Paginator { return paginator{oper{o.WithKind(KindPagination)}} }

// GetLimitFromPaginator extracts the limit from the pagination operation.
//
// If p is nil or the pagination operation has not implemented Limiter, return 0.
func GetLimitFromPaginator(p Paginator) (limit int) {
	if p == nil {
		return
	}

	if ps, ok := p.(Limiter); ok {
		limit = ps.Limit()
	} else if ps, ok := p.Op().Val.(Limiter); ok {
		limit = ps.Limit()
	}

	return
}

/// ---------------------------------------------------------------------- ///

// PageSizer is a paginator based on page and size.
type PageSizer struct {
	Page int64 // Start with 1
	Size int64
}

// Limit implements the interface Limiter.
func (p PageSizer) Limit() int { return int(p.Size) }

// PageSize is used to new a paginator based on page and size.
//
// The key is empty, and the value is a PageSizer instance.
func PageSize(page, size int64) Paginator {
	return New(PaginationOpPageSize, "", PageSizer{Page: page, Size: size}).Paginator()
}
