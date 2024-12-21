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
	KindPagination   = "Pagination"
	PaginationOpPage = "Page"
)

// Paginator represents a pagination operation.
type Paginator interface {
	paginate()
	Oper
}

type paginator struct{ oper }

func (p paginator) paginate() {}

// Paginator converts itself to Paginator.
func (o Op) Paginator() Paginator { return paginator{oper{o.WithKind(KindPagination)}} }

/// ---------------------------------------------------------------------- ///

type PageSize struct {
	Page int64 // Start with 1
	Size int64
}

func (p PageSize) Limit() int { return int(p.Size) }

// Page is short of Paginate.
//
// DEPRECATED!!! Please use Paginate instead.
func Page(page, size int64) Paginator {
	return Paginate(page, size)
}

// Paginate is equal to New(PaginationOpPage, "", PageSize{Page: page, Size: size}).Paginator().
func Paginate(page, size int64) Paginator {
	return New(PaginationOpPage, "", PageSize{Page: page, Size: size}).Paginator()
}
