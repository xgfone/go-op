# Go Condition and Setter Operation [![Build Status](https://github.com/xgfone/go-op/actions/workflows/go.yml/badge.svg)](https://github.com/xgfone/go-op/actions/workflows/go.yml) [![GoDoc](https://pkg.go.dev/badge/github.com/xgfone/go-op)](https://pkg.go.dev/github.com/xgfone/go-op) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg?style=flat-square)](https://raw.githubusercontent.com/xgfone/go-op/master/LICENSE)


Provide a common condition and setter operation supporting Go `1.18+`.


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
	// The common converter.
	convert := func(_, _, key string, value interface{}) interface{} {
		if s, ok := value.(string); ok {
			return fmt.Sprintf("`%s`='%s'", key, s)
		}
		return fmt.Sprintf("`%s`=%v", key, value)
	}

	// Register the condition and setter converters.
	op.RegisterConverter("sql", op.CondOpNotEqual, convert)
	op.RegisterConverter("sql", op.CondOpEqual, convert)
	op.RegisterConverter("sql", op.SetOpAdd, convert)
	op.RegisterConverter("sql", op.SetOpSet, convert)

	// Define a UPDATE sql builder.
	buildUpdateSQL := func(table string, setters []op.Setter, conds []op.Condition) string {
		_setters := make([]string, len(setters))
		for i, setter := range setters {
			_op, key, value := setter.Operation()
			_setters[i] = op.GetConverter("sql", _op)("sql", _op, key, value).(string)
		}

		_conds := make([]string, len(conds))
		for i, cond := range conds {
			_op, key, value := cond.Operation()
			_conds[i] = op.GetConverter("sql", _op)("sql", _op, key, value).(string)
		}

		return fmt.Sprintf("UPDATE `%s` SET %s WHERE %s",
			table, strings.Join(_setters, ", "),
			strings.Join(_conds, " AND "))
	}

	// build the UPDATE sql.
	ColumnID := op.Key("id")
	ColumnAge := op.Key("age")
	sql := buildUpdateSQL("user",
		[]op.Setter{ColumnAge.Add(1), op.Set("name", "Aaron")},
		[]op.Condition{ColumnID.Eq(123), op.NotEq("deleted", false)},
	)

	fmt.Println(sql)
	// Output:
	// UPDATE `user` SET `age`=1, `name`='Aaron' WHERE `id`=123 AND `deleted`=false
}
```
