package models

import (
	"IM_server/common/models"
	"IM_server/common/models/ctype"
)

// UserConfigModel 用户配置表
type UserConfigModel struct {
	models.Model
	UserId               uint                        `json:"userId"`               // 用户ID
	RecallMessage        *string                     `json:"recallMessage"`        // 撤回消息的提示内容
	FriendOnline         bool                        `json:"friendOnline"`         // 好友上线提醒
	Sound                bool                        `json:"sound"`                // 声音提醒
	SecureLink           bool                        `json:"secureLink"`           // 安全链接
	SavePassword         bool                        `json:"savePassword"`         // 保存密码
	SearchUser           int                         `json:"searchUser"`           // 别人查找到你的方式 0 不允许 1 通过用户号 2 通过搜索昵称
	Verification         int                         `json:"Verification"`         // 好友验证    0 不允许任何人添加 1 允许任何人添加 2 需要验证消息 3 需要回答问题 4 需要正确回答问题
	VerificationQuestion *ctype.VerificationQuestion `json:"verificationQuestion"` // 验证问题 为3和4的时候才需要
	Online               bool                        `json:"online"`               // 在线状态
}
