package models

import "IM_server/common/models"

// FriendModel 好友表
type FriendModel struct {
	models.Model
	SendUserID uint   `json:"sendUserId"` // 发送验证方
	RecvUserID uint   `json:"recvUserId"` // 接受验证方
	Notice     string `json:"notice"`     // 备注
}
