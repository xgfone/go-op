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
	KeyId     = Key("id")
	KeyUserId = Key("user_id")
	KeyOpenId = Key("open_id")
	KeyTaskId = Key("task_id")

	KeyUnionId  = Key("union_id")
	KeyGroupId  = Key("group_id")
	KeyMemberId = Key("member_id")

	KeyItemId   = Key("item_id")
	KeyItemType = Key("item_type")

	KeyApp  = Key("app")
	KeyDay  = Key("day")
	KeyNum  = Key("num")
	KeySrc  = Key("src")
	KeyDst  = Key("dst")
	KeyKey  = Key("key")
	KeyTmb  = Key("tmb")
	KeyTag  = Key("tag")
	KeyTags = Key("tags")
	KeyAuth = Key("auth")
	KeyAbbr = Key("abbr")
	KeyAddr = Key("addr")
	KeyCity = Key("city")
	KeyName = Key("name")
	KeyNote = Key("note")
	KeyDesc = Key("desc")
	KeyCode = Key("code")
	KeyConf = Key("conf")
	KeyCost = Key("cost")
	KeyData = Key("data")
	KeyDate = Key("date")
	KeyTime = Key("time")
	KeySort = Key("sort")
	KeyType = Key("type")
	KeyStep = Key("step")
	KeyFlag = Key("flag")
	KeyFrom = Key("from")

	KeySize  = Key("size")
	KeyUsed  = Key("used")
	KeyRest  = Key("rest")
	KeyRate  = Key("rate")
	KeyAvail = Key("avail")
	KeyCount = Key("count")
	KeyLimit = Key("limit")
	KeyTotal = Key("total")
	KeyValue = Key("value")

	KeyStyle = Key("style")
	KeyTitle = Key("title")
	KeyCover = Key("cover")
	KeyImage = Key("image")
	KeyMusic = Key("music")
	KeyVideo = Key("video")
	KeyModel = Key("model")

	KeyIp     = Key("ip")
	KeyUrl    = Key("url")
	KeyCidr   = Key("cidr")
	KeyHost   = Key("host")
	KeyPort   = Key("port")
	KeyScheme = Key("scheme")
	KeyMethod = Key("method")
	KeySource = Key("source")
	KeyTarget = Key("target")

	KeyPrice   = Key("price")
	KeyAmount  = Key("amount")
	KeyCoupon  = Key("coupon")
	KeyNumber  = Key("number")
	KeyBalance = Key("balance")
	KeyStorage = Key("storage")

	KeyAlias    = Key("alias")
	KeyAppid    = Key("appid")
	KeyEmail    = Key("email")
	KeyPhone    = Key("phone")
	KeyGrade    = Key("grade")
	KeyAvatar   = Key("avatar")
	KeyApiKey   = Key("apikey")
	KeySecret   = Key("secret")
	KeyAddress  = Key("address")
	KeyPassword = Key("password")

	KeyEvent      = Key("event")
	KeyReason     = Key("reason")
	KeyStatus     = Key("status")
	KeyRemain     = Key("remain")
	KeyPercent    = Key("percent")
	KeyVersion    = Key("version")
	KeySeconds    = Key("seconds")
	KeyInterval   = Key("interval")
	KeyCallback   = Key("callback")
	KeyIsHidden   = Key("is_hidden")
	KeyIsEnabled  = Key("is_enabled")
	KeyIsDisabled = Key("is_disabled")

	KeyTaskName = Key("task_name")
	KeyTaskInfo = Key("task_info")
	KeyTaskData = Key("task_data")

	KeyDoneAt    = Key("done_at")
	KeySchdAt    = Key("schd_at")
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
