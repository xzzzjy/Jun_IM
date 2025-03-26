package models

import "IM_server/common/models"

// FriendVerifyModel 好友验证表
type FriendVerifyModel struct {
	models.Model
	SendUserID           uint                  `json:"sendUserId"`           // 发送验证方
	RecvUserID           uint                  `json:"recvUserId"`           // 接受验证方
	Status               int                   `json:"status"`               // 状态 0 未处理 1 同意 2 拒绝 3 忽略
	AdditionMessage      string                `json:"additionMessage"`      // 附加消息
	VerificationQuestion *VerificationQuestion `json:"verificationQuestion"` // 验证问题 为3和4的时候才需要
}
