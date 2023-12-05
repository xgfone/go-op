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

	// For the interface Updater
	ups := []Updater{Add("key1", 123), Inc("key2"), Set("key", "value")}
	set1 := Sub("nok", "value")
	set2 := Mul("key", 456)
	fmt.Println(Contains(ups, set1))
	fmt.Println(Contains(ups, set2))

	// For the interface Oper
	ops := []Oper{Eq("key1", "value1"), Set("key2", "value2")}
	op1 := NotEq("nok", "value1")
	op2 := Add("key2", 123)
	fmt.Println(Contains(ops, op1))
	fmt.Println(Contains(ops, op2))

	// Output:
	// false
	// true
	// false
	// true
	// false
	// true
}

func ExampleOp() {
	// Manage the global op builders.
	builders := make(map[string]func(Op) string, 4)
	buildOper := func(op Oper) string { return builders[op.Op().Op](op.Op()) }
	register := func(op string, f func(Op) string) { builders[op] = f }

	// Register the Op builders.
	buildSignEq := func(op Op) string {
		if s, ok := op.Val.(string); ok {
			return fmt.Sprintf("`%s`='%s'", op.Key, s)
		}
		return fmt.Sprintf("`%s`=%v", op.Key, op.Val)
	}
	register(CondOpNotEqual, buildSignEq)
	register(CondOpEqual, buildSignEq)
	register(UpdateOpAdd, buildSignEq)
	register(UpdateOpSet, buildSignEq)

	// Define a UPDATE sql builder.
	buildUpdateSQL := func(table string, updaters []Updater, conds []Condition) string {
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
	ColumnID := Key("id")
	ColumnAge := Key("age")
	sql := buildUpdateSQL("user",
		[]Updater{ColumnAge.Add(1), Set("name", "Aaron")},
		[]Condition{ColumnID.Eq(123), NotEq("is_deleted", false)},
	)

	fmt.Println(sql)

	// Output:
	// UPDATE `user` SET `age`=1, `name`='Aaron' WHERE `id`=123 AND `is_deleted`=false
}

func ExampleOp_Scope() {
	op0 := Key("").Scope("")
	op1 := Key("").Scope("T")
	op2 := Key("key").Scope("")
	op3 := Key("key").Scope("T")

	fmt.Println(op0.Key)
	fmt.Println(op1.Key)
	fmt.Println(op2.Key)
	fmt.Println(op3.Key)

	// Output:
	//
	// T
	// key
	// T.key
}
