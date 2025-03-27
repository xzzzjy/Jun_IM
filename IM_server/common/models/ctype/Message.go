package ctype

import "time"

type SystemMessage struct {
	Type int `json:"type"` // 违规类型 1 黄 2 恐 3 政 4 不正当言论
}

type Message struct {
	Type             int               `json:"type"`             // 消息类型 1 文本 2 图片 3 视频 4 文件 5 语音 6 语音通话 7 视频通话 8 撤回消息 9 回复消息 10 引用消息
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
	AtMessage        *AtMessage        `json:"AtMessage"`        // 为11的时候使用,群聊才有
}

type ImgMessage struct {
	Title string `json:"title"`
	Src   string `json:"src"`
}

type VideoMessage struct {
	Title string `json:"title"`
	Src   string `json:"src"`
	Time  int    `json:"time"` //时长 单位秒

}

type FileMessage struct {
	Title string `json:"title"`
	Src   string `json:"src"`  //文件地址
	Size  int    `json:"size"` //文件大小 单位字节
	Type  string `json:"type"` //文件类型
}

type VoiceMessage struct {
	Src  string `json:"src"`
	Time int    `json:"time"` //时长 单位秒
}

type VoiceCallMessage struct {
	StartTime time.Time `json:"startTime"` // 通话开始时间
	EndTime   time.Time `json:"endTime"`   // 通话结束时间
	EndReason int       `json:"endReason"` // 通话结束原因  0 发起方挂断 1 接收方挂断 2 网络原因挂断 3 未打通
}

type VideoCallMessage struct {
	StartTime time.Time `json:"startTime"` // 通话开始时间
	EndTime   time.Time `json:"endTime"`   // 通话结束时间
	EndReason int       `json:"endReason"` // 通话结束原因  0 发起方挂断 1 接收方挂断 2 网络原因挂断 3 未打通
}

type WithdrawMessage struct {
	Content       string   `json:"content"`       // 撤回的提示词
	OriginMessage *Message `json:"originMessage"` //原消息
}

type ReplyMessage struct {
	MessageID uint     `json:"messageID"` // 消息ID
	Content   string   `json:"content"`   // 回复的内容
	Message   *Message `json:"message"`   // 回复的消息
}

type QuoteMessage struct {
	MessageID uint     `json:"messageID"`
	Content   string   `json:"content"`
	Message   *Message `json:"message"`
}

//@消息
type AtMessage struct {
	UserID  uint     `json:"userID"`
	Content string   `json:"content"`
	Message *Message `json:"message"`
}
