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
	KeyAddrId = Key("addr_id")

	KeyOrderId  = Key("order_id")
	KeyTransId  = Key("trans_id")
	KeyCouponId = Key("coupon_id")

	KeyTmplId  = Key("tmpl_id")
	KeyModelId = Key("model_id")
	KeyCoverId = Key("cover_id")
	KeyImageId = Key("image_id")
	KeyAudioId = Key("audio_id")
	KeyMusicId = Key("music_id")
	KeyVideoId = Key("video_id")
	KeyStyleId = Key("style_id")

	KeyUnionId   = Key("union_id")
	KeyGroupId   = Key("group_id")
	KeyMemberId  = Key("member_id")
	KeyInviterId = Key("inviter_id")

	KeyClientId = Key("client_id")
	KeyServerId = Key("server_id")

	KeyItemId   = Key("item_id")
	KeyItemType = Key("item_type")

	KeyTaskerId   = Key("tasker_id")
	KeyTaskerAddr = Key("tasker_addr")
	KeyTaskerInfo = Key("tasker_info")

	KeyOpCode = Key("op_code")
	KeyOpData = Key("op_data")
	KeyOpType = Key("op_type")
	KeyOpInfo = Key("op_info")
	KeyOpNote = Key("op_note")

	KeyArg  = Key("arg")
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
	KeyOmit = Key("omit")
	KeyRead = Key("read")
	KeySign = Key("sign")
	KeyTmpl = Key("tmpl")
	KeyArgs = Key("args")

	KeySize  = Key("size")
	KeyUsed  = Key("used")
	KeyRest  = Key("rest")
	KeyRate  = Key("rate")
	KeyAvail = Key("avail")
	KeyCount = Key("count")
	KeyLimit = Key("limit")
	KeyTotal = Key("total")
	KeyValue = Key("value")
	KeyIndex = Key("index")

	KeyStyle = Key("style")
	KeyTitle = Key("title")
	KeyCover = Key("cover")
	KeyImage = Key("image")
	KeyMusic = Key("music")
	KeyVideo = Key("video")
	KeyModel = Key("model")
	KeyAudio = Key("audio")
	KeyLyric = Key("lyric")
	KeyTheme = Key("theme")

	KeyTmplUrl = Key("tmpl_url")
	KeyOrigUrl = Key("orig_url")
	KeyMarkUrl = Key("mark_url")
	KeyTmbnUrl = Key("tmbn_url")
	KeyTmbUrl  = Key("tmb_url")
	KeyMaxUrl  = Key("max_url")

	KeyCoverUrl = Key("cover_url")
	KeyImageUrl = Key("image_url")
	KeyAudioUrl = Key("audio_url")
	KeyMusicUrl = Key("music_url")
	KeyVideoUrl = Key("video_url")

	KeyClientUrl  = Key("client_url")
	KeyServerUrl  = Key("server_url")
	KeyClientAddr = Key("client_addr")
	KeyServerAddr = Key("server_addr")

	KeyIp     = Key("ip")
	KeyNet    = Key("net")
	KeyUrl    = Key("url")
	KeyCidr   = Key("cidr")
	KeyHost   = Key("host")
	KeyPort   = Key("port")
	KeyScheme = Key("scheme")
	KeyMethod = Key("method")
	KeySource = Key("source")
	KeyTarget = Key("target")
	KeyClient = Key("client")
	KeyServer = Key("servet")

	KeyPrice   = Key("price")
	KeyAmount  = Key("amount")
	KeyCoupon  = Key("coupon")
	KeyNumber  = Key("number")
	KeyBalance = Key("balance")
	KeyStorage = Key("storage")
	KeyIllegal = Key("illegal")
	KeyFailure = Key("failure")
	KeySuccess = Key("success")
	KeyResult  = Key("result")

	KeyDiscount    = Key("discount")
	KeyPayType     = Key("pay_type")
	KeyPayPrice    = Key("pay_price")
	KeyPayAmount   = Key("pay_amount")
	KeyTransType   = Key("trans_type")
	KeyTransAmount = Key("trans_amount")
	KeyOrigPrice   = Key("orig_price")
	KeyOffPrice    = Key("off_price")

	KeyPType       = Key("ptype")
	KeyPConf       = Key("pconf")
	KeyProductType = Key("product_type")
	KeyProductConf = Key("product_conf")

	KeyWeight   = Key("weight")
	KeyHorizon  = Key("horizon")
	KeyVertical = Key("vertical")
	KeyTemplate = Key("template")

	KeyPrompt         = Key("prompt")
	KeyNPrompt        = Key("nprompt")
	KeyPosPrompt      = Key("pos_prompt")
	KeyNegPrompt      = Key("neg_prompt")
	KeyPositivePrompt = Key("positive_prompt")
	KeyNegativePrompt = Key("negative_prompt")

	KeyVoted  = Key("voted")
	KeyVotes  = Key("votes")
	KeyPrize  = Key("prize")
	KeyPrizes = Key("prizes")

	KeyAlias    = Key("alias")
	KeyAppid    = Key("appid")
	KeyEmail    = Key("email")
	KeyPhone    = Key("phone")
	KeyGrade    = Key("grade")
	KeyAvatar   = Key("avatar")
	KeyApiKey   = Key("apikey")
	KeyPubKey   = Key("pubkey")
	KeyPriKey   = Key("prikey")
	KeySecret   = Key("secret")
	KeyQrCode   = Key("qrcode")
	KeyContact  = Key("contact")
	KeyAddress  = Key("address")
	KeyPassword = Key("password")

	KeyCbUrl      = Key("cburl")
	KeyEvent      = Key("event")
	KeyReason     = Key("reason")
	KeyStatus     = Key("status")
	KeyRemain     = Key("remain")
	KeySecond     = Key("second")
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
	KeyTaskType = Key("task_type")

	KeyReqName = Key("req_name")
	KeyReqBody = Key("req_body")
	KeyReqInfo = Key("req_info")
	KeyReqData = Key("req_data")
	KeyReqType = Key("req_type")

	KeyResCode = Key("res_code")
	KeyResBody = Key("res_body")
	KeyResInfo = Key("res_info")
	KeyResData = Key("res_data")
	KeyResType = Key("res_type")

	KeyPubAt     = Key("pub_at")
	KeySubAt     = Key("sub_at")
	KeyDoneAt    = Key("done_at")
	KeySchdAt    = Key("schd_at")
	KeyRefundAt  = Key("refund_at")
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
