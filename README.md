# Go Operation [![Build Status](https://github.com/xgfone/go-op/actions/workflows/go.yml/badge.svg)](https://github.com/xgfone/go-op/actions/workflows/go.yml) [![GoDoc](https://pkg.go.dev/badge/github.com/xgfone/go-op)](https://pkg.go.dev/github.com/xgfone/go-op) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg?style=flat-square)](https://raw.githubusercontent.com/xgfone/go-op/master/LICENSE)


Provide a common operation, such as `Condition` and `Updater`, supporting Go `1.18+`.


## Install
```shell
$ go get -u github.com/xgfone/go-op
```


## Example
```go
package main

import (
	"fmt"
	"strings"

	"github.com/xgfone/go-op"
)

func main() {
	// Manage the global op builders.
	builders := make(map[string]func(op.Op) string, 4)
	buildOper := func(op op.Oper) string { return builders[op.Op().Op](op.Op()) }
	register := func(op string, f func(op.Op) string) { builders[op] = f }

	// Register the Op builders.
	buildSignEq := func(op op.Op) string {
		if s, ok := op.Val.(string); ok {
			return fmt.Sprintf("`%s`='%s'", op.Key, s)
		}
		return fmt.Sprintf("`%s`=%v", op.Key, op.Val)
	}
	register(op.CondOpNotEqual, buildSignEq)
	register(op.CondOpEqual, buildSignEq)
	register(op.UpdateOpAdd, buildSignEq)
	register(op.UpdateOpSet, buildSignEq)

	// Define a UPDATE sql builder.
	buildUpdateSQL := func(table string, updaters []op.Updater, conds []op.Condition) string {
		sets := make([]string, len(updaters))
		for i, up := range updaters {
			sets[i] = buildOper(up)
		}

		wheres := make([]string, len(conds))
		for i, cond := range conds {
			wheres[i] = buildOper(cond)
		}

		return fmt.Sprintf("UPDATE `%s` SET %s WHERE %s",
			table, strings.Join(sets, ", "),
			strings.Join(wheres, " AND "))
	}

	// build the UPDATE sql.
	ColumnID := op.Key("id")
	ColumnAge := op.Key("age")
	sql := buildUpdateSQL("user",
		[]op.Updater{ColumnAge.Add(1), op.Set("name", "Aaron")},
		[]op.Condition{ColumnID.Eq(123), op.NotEq("is_deleted", false)},
	)

	fmt.Println(sql)

	// Output:
	// UPDATE `user` SET `age`=1, `name`='Aaron' WHERE `id`=123 AND `is_deleted`=false
}
```
