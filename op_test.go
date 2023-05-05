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

import (
	"fmt"
	"strings"
)

func ExampleContains() {
	// For the interface Condition
	conds := []Condition{NotEq("key", "value"), Le("k1", 123), GtEq("k2", 456)}
	cond1 := Equal("nok", "value")
	cond2 := Equal("key", "value")
	fmt.Println(Contains(conds, cond1))
	fmt.Println(Contains(conds, cond2))

	// For the interface Setter
	sets := []Setter{Add("key1", 123), Inc("key2"), Set("key", "value")}
	set1 := Sub("nok", "value")
	set2 := Mul("key", 456)
	fmt.Println(Contains(sets, set1))
	fmt.Println(Contains(sets, set2))

	// For the interface Oper
	ops := []Oper{Eq("key1", "value1"), Set("key2", "value2")}
	op1 := NotEq("nok", "value1")
	op2 := Add("key2", 123)
	fmt.Println(Contains(ops, op1))
	fmt.Println(Contains(ops, op2))

	// For the struct Op
	opss := []Op{Eq("key1", "value1"), Set("key2", "value2")}
	ops1 := NotEq("nok", "value1")
	ops2 := Add("key2", 123)
	fmt.Println(Contains(opss, ops1))
	fmt.Println(Contains(opss, ops2))

	// Output:
	// false
	// true
	// false
	// true
	// false
	// true
	// false
	// true
}

func ExampleRegisterConverter() {
	// The common converter.
	convert := func(_, _, key string, value interface{}) interface{} {
		if s, ok := value.(string); ok {
			return fmt.Sprintf("`%s`='%s'", key, s)
		}
		return fmt.Sprintf("`%s`=%v", key, value)
	}

	// Register the condition and setter converters.
	RegisterConverter("sql", CondOpNotEqual, convert)
	RegisterConverter("sql", CondOpEqual, convert)
	RegisterConverter("sql", SetOpAdd, convert)
	RegisterConverter("sql", SetOpSet, convert)

	// Define a UPDATE sql builder.
	buildUpdateSQL := func(table string, setters []Setter, conds []Condition) string {
		_setters := make([]string, len(setters))
		for i, setter := range setters {
			op, key, value := setter.Setter()
			_setters[i] = GetConverter("sql", op)("sql", op, key, value).(string)
		}

		_conds := make([]string, len(conds))
		for i, cond := range conds {
			op, key, value := cond.Condition()
			_conds[i] = GetConverter("sql", op)("sql", op, key, value).(string)
		}

		return fmt.Sprintf("UPDATE `%s` SET %s WHERE %s", table, strings.Join(_setters, ", "), strings.Join(_conds, " AND "))
	}

	// build the UPDATE sql.
	ColumnID := Key("id")
	ColumnAge := Key("age")
	sql := buildUpdateSQL("user",
		[]Setter{ColumnAge.Add(1), Set("name", "Aaron")},
		[]Condition{ColumnID.Eq(123), NotEq("deleted", false)},
	)

	fmt.Println(sql)

	// Output:
	// UPDATE `user` SET `age`=1, `name`='Aaron' WHERE `id`=123 AND `deleted`=false
}
