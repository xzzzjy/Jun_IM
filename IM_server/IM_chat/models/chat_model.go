package models

import "IM_server/common/models"

type ChatModel struct {
	models.Model
	SendUserID     uint    `json:"sendUserId"`     // 发送验证方
	RecvUserID     uint    `json:"recvUserId"`     // 接受验证方
	MessageType    int     `json:"messageType"`    // 消息类型 1 文本 2 图片 3 视频 4 文件 5 语音 6 语音通话 7 视频通话 8 撤回消息 9 回复消息 10 引用消息
	MessagePreview string  `json:"messagePreview"` // 消息预览
	Message        Message `json:"message"`        // 消息
}

type Message struct {
	Content          *string           `json:"content"`          // 为1的时候使用
	ImgMessage       *ImgMessage       `json:"imgMessage"`       // 为2的时候使用
	VideoMessage     *VideoMessage     `json:"VideoMessage"`     // 为3的时候使用
	FileMessage      *FileMessage      `json:"FileMessage"`      // 为4的时候使用
	VoiceMessage     *VoiceMessage     `json:"VoiceMessage"`     // 为5的时候使用
	VoiceCallMessage *VoiceCallMessage `json:"VoiceCallMessage"` // 为6的时候使用
	VideoCallMessage *VideoCallMessage `json:"VideoCallMessage"` // 为7的时候使用
	WithdrawMessage  *WithdrawMessage  `json:"WithdrawMessage"`  // 为8的时候使用
	ReplyMessage     *ReplyMessage     `json:"ReplyMessage"`     // 为9的时候使用
	QuoteMessage     *QuoteMessage     `json:"QuoteMessage"`     // 为10的时候使用
}

type ImgMessage struct {
}

type VideoMessage struct {
}

type FileMessage struct {
}

type VoiceMessage struct {
}

type VoiceCallMessage struct {
}

type VideoCallMessage struct {
}

type WithdrawMessage struct {
}

type ReplyMessage struct {
}

type QuoteMessage struct {
}
