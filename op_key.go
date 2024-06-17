// Copyright 2024 xgfone
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

import "strings"

// Pre-define some operations.
var (
	KeyId      = Key("id")
	KeyUserId  = Key("user_id")
	KetTaskId  = Key("task_id")
	KetGroupId = Key("group_id")

	KeyAddr = Key("addr")
	KeyName = Key("name")
	KeyNote = Key("note")
	KeyDesc = Key("desc")
	KeyCode = Key("code")
	KeyConf = Key("conf")
	KeyData = Key("data")
	KeyDate = Key("date")
	KeySort = Key("sort")
	KeyType = Key("type")

	KeyEmail    = Key("email")
	KeyPhone    = Key("phone")
	KeyAvatar   = Key("avatar")
	KeyPassword = Key("password")

	KeyTotal    = Key("total")
	KeyItemId   = Key("item_id")
	KeyItemType = Key("item_type")

	KeyStatus     = Key("status")
	KeyIsEnabled  = Key("is_enabled")
	KeyIsDisabled = Key("is_disabled")

	KeyCreatedAt = Key("created_at")
	KeyUpdatedAt = Key("updated_at")
	KeyDeletedAt = Key("deleted_at")
	KeyExpiredAt = Key("expired_at")
)

// Pre-define some field condition operations.
var (
	IsDeletedCond    = KeyDeletedAt.NotEq("0000-00-00 00:00:00")
	IsNotDeletedCond = KeyDeletedAt.Eq("0000-00-00 00:00:00")
)

// IsNotDeletedCondWithTable returns a condition operation that is not deleted
// with table prefix.
func IsNotDeletedCondWithTable(table string) Condition {
	op := IsNotDeletedCond.Op()
	key := strings.Join([]string{table, op.Key}, Sep)
	return op.WithKey(key).Condition()
}
