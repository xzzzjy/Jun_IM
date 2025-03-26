package models

import "IM_server/common/models"

// UserModel 用户表
type UserModel struct { // 用户模型，继承公共模型common/
	models.Model
	Password string `json:"password"` // 密码
	Avatar   string `json:"avatar"`   // 头像
	Nickname string `json:"nickname"` // 昵称
	Abstract string `json:"abstract"` // 简介
	IP       string `json:"ip"`       // IP地址
	Address  string `json:"address"`  // 地址
}
