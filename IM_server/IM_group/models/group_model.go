package models

import (
	"IM_server/common/models"
	"IM_server/common/models/ctype"
)

type GroupModel struct {
	models.Model
	Title                string                      `json:"title"`                // 群名称
	Abstract             string                      `json:"abstract"`             // 群简介
	Avatar               string                      `json:"avatar"`               // 群头像
	Creator              uint                        `json:"creator"`              // 群主
	IsSearch             bool                        `json:"isSearch"`             // 是否允许被搜索
	Verification         int                         `json:"groupVerification"`    // 群验证    0 不允许任何人添加 1 允许任何人添加 2 需要验证消息 3 需要回答问题 4 需要正确回答问题
	VerificationQuestion *ctype.VerificationQuestion `json:"verificationQuestion"` // 验证问题 为3和4的时候才需要
	IsInvite             bool                        `json:"isInvite"`             // 是否允许邀请
	IsTemporarySession   bool                        `json:"isTemporarySession"`   // 是否开启临时会话
	IsProhibition        bool                        `json:"isProhibition"`        // 是否开启禁言
	Size                 int                         `json:"size"`                 // 群规模
}
