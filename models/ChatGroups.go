package models

import (
	"gorm.io/gorm"
)

// 这里切记发给服务器的message信息中不要包含id字段，哪怕是clientid也不可以
// gorm会自动提取包含"id"的字段并读取其中的值导致id的值被错误改变
type ChatGroups struct {
	gorm.Model
	Name        string   //群组名称
	Description string   //群组描述
	Creator     string   //群组创建者用户id
	Members     []string //群组成员id
}

// 用于存储群组的基本信息
type GroupBasic struct {
	gorm.Model
	Name        string //群组名称
	Description string //群组描述
}

// 存储用户群组关系表
type User_Group struct {
	gorm.Model
	UserName  string
	GroupName string
	GroupId   uint
	Status    uint //0:群主 1:管理员 2:普通用户
}

// 群组内消息存储
type Group_Message struct {
	gorm.Model
	Sender    string
	GroupName string
	GroupId   uint
}
