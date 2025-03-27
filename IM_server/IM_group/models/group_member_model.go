package models

import "IM_server/common/models"

type GroupMemberModel struct {
	models.Model
	GroupID         uint   `json:"groupId"`         // 群ID
	UserID          uint   `json:"userId"`          // 用户ID
	MemberNickName  string `json:"memberNickName"`  // 群内昵称
	Role            int    `json:"role"`            // 1 群主 2 管理员 3 普通成员
	ProhibitionTime *int   `json:"prohibitionTime"` // 禁言时间
}
