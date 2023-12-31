package pool

import (
	"ginchat/models"

	"github.com/gorilla/websocket"
)

var OnlineMap map[string]*websocket.Conn

// 线程池初始化实现
func Pool_init() {
	OnlineMap = make(map[string]*websocket.Conn)
}

// 用户私聊信息发送
func ChatPrivate(msg *models.Message) {
	OnlineMap[msg.Receiver].WriteJSON(msg)
}

// 集体广播功能实现
func ChatPublic(msg *models.Message) {
	for key := range OnlineMap {
		OnlineMap[key].WriteJSON(msg)
	}
}
