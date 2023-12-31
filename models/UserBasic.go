package models

import (
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	HostIP        string
	Port          int32
	Salt          string //随机数密码种子
	Email         string `valid:"email"`
	PhoneNumber   string `valid:"matches(^1[3-9]{1}\\d{9})"`
	Password      string
	LoginTime     time.Time
	HeartbeatTime uint64
	LoginOutTime  time.Time
}

// func GetUserLists()*[]UserBasic{
// 	data:=make([]UserBasic,10)

// }
