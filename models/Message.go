package models

import (
	"gorm.io/gorm"
)

// 这里切记发给服务器的message信息中不要包含id字段，哪怕是clientid也不可以
// gorm会自动提取包含"id"的字段并读取其中的值导致id的值被错误改变
type Message struct {
	gorm.Model
	Type     string //'1':私聊,'2':群聊，‘0’：用户登录标志,
	MsgType  string //'0':文本，‘1’:图片url格式
	Text     string
	Sender   string //发送方名字
	Receiver string //接收方名字
	Date     int64
}
