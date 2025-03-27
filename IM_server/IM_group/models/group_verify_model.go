package models

import (
	"IM_server/common/models"
	"IM_server/common/models/ctype"
)

type GroupVerifyModel struct {
	models.Model
	GroupID              uint                        `json:"groupId"`              // 群ID
	UserID               uint                        `json:"userId"`               // 加或退的用户ID
	AdditionMessage      string                      `json:"additionMessage"`      // 附加消息
	VerificationQuestion *ctype.VerificationQuestion `json:"verificationQuestion"` // 验证问题 为3和4的时候才需要
	Status               int                         `json:"status"`               // 状态 0 未处理 1 同意 2 拒绝 3 忽略
	Type                 int                         `json:"type"`                 // 类型 1 加群 2 退群
}
