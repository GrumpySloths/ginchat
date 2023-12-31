package service

import (
	"fmt"
	"ginchat/models"
	"ginchat/pool"
	"ginchat/utility"
	"log"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// @BasePath /api/v1
// @Summary GetUserLists
// @Schemes
// @Description get user list from mysql
// @Tags service/UserServices
// @Accept json
// @Produce json
// @Success 200 {code} message
// @Router /user/getUserLists [get]
func GetUserLists(c *gin.Context) {
	UserLists := make([]models.UserBasic, 10)
	utility.DB.Find(&UserLists)

	c.JSON(200, gin.H{
		"message": UserLists,
	})
}

// @BasePath /api/v1
// @Summary user login
// @Schemes
// @Tags service/UserServices
// @Accept json
// @Produce json
// @param name formData string false "用户名"
// @param passwd formData string false "密码"
// @Success 200 {code} message
// @Router /user/UserLogin [post]
func UserLogin(c *gin.Context) {

	name := c.PostForm("name")
	passwd := c.PostForm("passwd")
	var user models.UserBasic
	utility.DB.Where("name = ?", name).First(&user)
	if user.Name == "" {
		c.JSON(-1, gin.H{"message": "用户不存在"})
		return
	}
	//密码校验
	flag := utility.ValidEncode(passwd, user.Salt, user.Password)
	if !flag {
		c.JSON(-1, gin.H{"message": "密码不正确"})
		return
	}

	c.JSON(200, gin.H{
		"message": user,
	})
}

// @BasePath /api/v1
// @Summary Create new user
// @Schemes
// @Tags service/UserServices
// @Accept json
// @Produce json
// @param name formData  string false "用户名"
// @param passwd formData string false "密码"
// @param rePasswd formData  string false "再次输入密码"
// @Success 200 {code} message
// @Router /user/CreateUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.PostForm("name")
	passwd := c.PostForm("passwd")
	user.Salt = fmt.Sprintf("%06d", rand.Int31())
	// 这里Query询问的内容和添加的变量要完全一致,不然读取不到
	rePasswd := c.PostForm("rePasswd")
	//用户注册名字重复验证
	query_data := models.UserBasic{}
	fmt.Println("___________________________")
	fmt.Println("user.name:" + user.Name + " user.passwd:" + user.Password)
	fmt.Println("___________________________")
	utility.DB.Where("name = ?", user.Name).First(&query_data)
	fmt.Println(query_data)
	fmt.Println("hello world _____________")
	if query_data.Name != "" {
		c.JSON(-1, gin.H{"message": "用户名已存在"})
		return
	}
	//用户密码注册验证
	if passwd != rePasswd {
		c.JSON(-1, gin.H{"message": "两次输入密码不一致"})
		return
	}

	user.Password = utility.MakeEncode(passwd, user.Salt)
	user.LoginTime = time.Now().UTC().Local()
	user.LoginOutTime = time.Now().UTC().Local()

	utility.DB.Create(&user)
	c.JSON(200, gin.H{
		"message": "创建新用户成功",
	})
}

// @BasePath /api/v1
// @Summary delete user by id
// @Schemes
// @Tags service/UserServices
// @Accept json
// @Produce json
// @param id query string false "要删除的用户id"
// @Success 200 {code} message
// @Router /user/DeleteUser [post]
func DeleteUser(c *gin.Context) {
	var user models.UserBasic
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	utility.DB.Delete(&user)
	c.JSON(200, gin.H{
		"message": "删除用户成功",
	})
}

// @BasePath /api/v1
// @Summary update user by id
// @Schemes
// @Tags service/UserServices
// @Accept json
// @Produce json
// @param id query string false "要修改的用户id"
// @param name query string false "用户名"
// @param passwd query string false "密码"
// @param email query string false "用户邮箱"
// @param phone query string false "用户手机号"
// @Success 200 {code} message
// @Router /user/UpdateUser [post]
func UpdateUser(c *gin.Context) {
	var user models.UserBasic
	id, _ := strconv.Atoi(c.Query("id"))
	fmt.Printf("id=%d", id)
	user.ID = uint(id)
	user.Name = c.Query("name")
	user.Password = c.Query("passwd")
	user.PhoneNumber = c.Query("phone")
	user.Email = c.Query("email")
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println("err:", err)
		c.JSON(-1, gin.H{
			"message": "修改用户失败,部分信息格式不正确",
		})
	} else {
		utility.DB.Model(&user).Updates(&models.UserBasic{Name: user.Name,
			Password: user.Password, Email: user.Email, PhoneNumber: user.PhoneNumber})
		c.JSON(200, gin.H{
			"message": "修改用户成功",
		})
	}
}

var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		fmt.Println("requset:", r)
		fmt.Println("url", r.URL)
		return true
	},
}

// 设置app接听者
func ServerHandler(c *gin.Context) {
	// 升级成 websocket 连接
	ws, err := upgrade.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Fatalln(err)
	}
	clientName := c.Query("name")
	fmt.Println("client name:", clientName)
	//将当前线程添加至用户在线线程池,这里应该加锁但临界区应该越少越好尽量少写并发
	pool.OnlineMap[clientName] = ws
	//打印coon连接双方的地址
	fmt.Println("server.addr:" + ws.LocalAddr().String() + " remote addr:" + ws.RemoteAddr().String())
	// 完成时关闭连接释放资源
	defer func() {
		fmt.Println("socket has closed")
		delete(pool.OnlineMap, clientName)
		pool.ChatPublic(&models.Message{Type: "0", Text: fmt.Sprintf("%s 连接已断开", clientName)})
		ws.Close()
	}()
	go func() {
		// 监听连接“完成”事件，其实也可以说丢失事件
		<-c.Done()
		// 这里也可以做用户在线/下线功能
		fmt.Println("ws lost connection")
	}()
	pool.ChatPublic(&models.Message{Type: "0", Text: "连接成功"})

	for {
		// 读取客户端发送过来的消息，如果没发就会一直阻塞住
		//这里处理用户输入逻辑，包括群聊信息，私聊信息等
		message := models.Message{}
		ws.ReadJSON(&message)
		// fmt.Println("message.type:", message.Type)
		switch message.Type {
		case "1":
			//实现私聊功能
			fmt.Println("私聊功能实现,接收方:", message.Receiver)
			if _, ok := pool.OnlineMap[message.Receiver]; ok {
				fmt.Println("发送成功")
				utility.DB.Create(&message)
				//将数据库中成功更新的聊天记录发送给聊天双方
				ws.WriteJSON(&message)
				pool.ChatPrivate(&message)
			} else {
				fmt.Println("用户不存在，请选择一个存在的用户后重新尝试")
			}

		case "2":
			//实现群聊功能
			fmt.Println("群聊信息接受成功，群聊对象:", message.Receiver)
			utility.DB.Create(&message)
			//从数据库中得到群组中的所有人员
			var GroupMembers []models.User_Group
			utility.DB.Where(&models.User_Group{GroupName: message.Receiver}).Find(&GroupMembers)
			//这里暂时只考虑把消息发送回给自己就可以了，因为消息已经写到数据库里了
			//上面的注释说明考虑并不完善，要让群聊用户可以在线时获取最新消息应该要用到redis消息缓存队列
			ws.WriteJSON(&message)

		case "3":
			//接收到向机器人发送消息时的处理
			fmt.Println("与机器人聊天功能实现,接收方:", message.Receiver)
			utility.DB.Create(&message)
			ws.WriteJSON(&message)

		default:
			fmt.Println("未知选项，信息发送失败")
		}
	}
}

// 获取当前在线的用户列表清单
func GetOnlineUsers(c *gin.Context) {
	var Onlinelist []string
	for key := range pool.OnlineMap {
		Onlinelist = append(Onlinelist, key)
	}
	sort.Strings(Onlinelist)
	c.JSON(200, gin.H{"OnlineUsers": Onlinelist})
}

// 获取当前用户所在的群聊记录
func GetGroups(c *gin.Context) {
	ClientName := c.Query("ClientName") //发送用户的名字
	var Groups []models.User_Group
	utility.DB.Where(&models.User_Group{UserName: ClientName}).Find(&Groups)

	var GroupList []string
	for _, group := range Groups {
		GroupList = append(GroupList, group.GroupName)
	}

	c.JSON(200, gin.H{"GroupList": GroupList})
}

// 获取聊天信息记录
func GetMessageHistory(c *gin.Context) {
	Sender, Receiver, Type := c.Query("Sender"), c.Query("Receiver"), c.Query("Type")
	fmt.Println("sender:" + Sender + " receiver:" + Receiver)
	var messages []models.Message

	switch Type {
	case "1": //获取和好友的聊天记录
		utility.DB.Where(&models.Message{Sender: Sender,
			Receiver: Receiver, Type: Type}).Or(&models.Message{Sender: Receiver,
			Receiver: Sender, Type: Type}).Find(&messages)
	case "2": //获取相应的群聊记录
		utility.DB.Where(&models.Message{Receiver: Receiver, Type: Type}).Find(&messages)
	case "3": //获取和本地机器人的聊天记录
		utility.DB.Where(&models.Message{Sender: Sender,
			Receiver: Receiver, Type: Type}).Or(&models.Message{Sender: Receiver,
			Receiver: Sender, Type: Type}).Find(&messages)
	default:
		fmt.Println("未知选项，获取消息记录失败")
	}

	c.JSON(200, gin.H{"MessageHistorys": messages})
}

// 创建新的群组
func GroupCreate(c *gin.Context) {
	var group_msg models.ChatGroups
	c.BindJSON(&group_msg)
	//处理接收到的群组创建信息并将其写入数据库
	// GroupsInfo := models.Groups{Name: group_msg.Name, Description: group_msg.Description}
	// utility.DB.Create(&GroupsInfo)
	utility.DB.Create(&models.GroupBasic{Name: group_msg.Name, Description: group_msg.Description})
	//查询当前群组id
	query_data := models.GroupBasic{}
	utility.DB.Where("name=?", group_msg.Name).First(&query_data)
	//用户群组关系处理
	for _, name := range group_msg.Members {
		utility.DB.Create(&models.User_Group{UserName: name, GroupId: query_data.ID, GroupName: query_data.Name})
	}
	fmt.Println(group_msg)

	c.JSON(200, gin.H{"message": "success"})
}

// 获取客户端上传的文件
func FileUpload(c *gin.Context) {
	// 单文件
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	dst := "./" + file.Filename
	// 上传文件至指定的完整文件路径
	c.SaveUploadedFile(file, dst)
	c.JSON(200, gin.H{"message": "success"})

}
