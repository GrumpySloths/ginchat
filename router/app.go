package router

import (
	docs "ginchat/docs"

	"ginchat/service"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/index", service.GetIndex)
	r.GET("/user/getUserLists", service.GetUserLists)
	r.POST("/user/CreateUser", service.CreateUser)
	r.POST("/user/DeleteUser", service.DeleteUser)
	r.POST("/user/UpdateUser", service.UpdateUser)
	r.POST("/user/UserLogin", service.UserLogin)
	r.GET("/ws", service.ServerHandler)
	r.GET("/user/GetOnlineUsers", service.GetOnlineUsers)
	r.GET("/user/GetGroups", service.GetGroups)
	r.GET("/user/GetMessageHistory", service.GetMessageHistory)
	r.POST("/user/GroupCreate", service.GroupCreate)
	r.POST("/user/FileUpload", service.FileUpload)

	return r
}
