package models

import (
	"IM_server/common/models"
	"IM_server/common/models/ctype"
)

type GroupChatModel struct {
	models.Model
	SendUserID     uint                 `json:"sendUserId"`     // 发送验证方
	MessageType    int                  `json:"messageType"`    // 消息类型 1 文本 2 图片 3 视频 4 文件 5 语音 6 语音通话 7 视频通话 8 撤回消息 9 回复消息 10 引用消息 11艾特用户
	MessagePreview string               `json:"messagePreview"` // 消息预览
	Message        ctype.Message        `json:"message"`        // 消息内容
	SystemMessage  *ctype.SystemMessage `json:"systemMessage"`  // 消息提示
}
