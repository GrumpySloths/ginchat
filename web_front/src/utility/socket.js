import { GetMessageHistory,GetOnlineLists,GetGroupLists } from "./chat";

export class ClientSocket{
    /**
     * 
     * @param {string} username -客户端用户名称
     * @param {ref("")} messageReceiver -消息接收者，ref格式，其对象由UI界面用户自行确定
     * @param {ref("")} messageType -消息接受类型,ref格式，用于选择消息传输的对象是好友还是群聊
     * @param {ref("")} message -消息载体，用于在UI界面传输消息
     * @param {ref([])} OnlineLists -保存从服务器查询得到的用户在线名单
     * @param {ref([])} GroupLists -保存从服务器查询得到的用户群组关系列表
     * @param {ref([])} messageHistory -保存从服务器查询得到的用户聊天历史记录
     */
    constructor(username, messageReceiver,messageType, message,OnlineLists,GroupLists,messageHistory) {
        this.message = message
        this.messageHistory = messageHistory
        this.messageType=messageType
        this.messageReceiver = messageReceiver
        this.username = username
        this.OnlineLists = OnlineLists
        this.GroupLists=GroupLists
        this.socket = new WebSocket(`ws://localhost:8080/ws?name=${username}`);
    }


    //用于启动socket相关的服务，处理客户端和服务器之间的消息流
    socketStart() {
        // Connection opened
        //这里用(event)=>{}而不是function(event){}是因为使用=>的形式
        //this指向的是定义时绑定的,即当前所在的类，如果用function形式就变成运行时绑定
        //的了，即指向了this.socket,要注意区分
        this.socket.addEventListener("open", (event) =>{
            this.sendText("hello server")
        });
        //接收信道
        this.socket.onmessage = (event) => {
            const msg = JSON.parse(event.data);
            console.log("Message type from server ", msg);
            switch (msg.Type) {
                case "0":
                    //用户成功登录
                    console.log("有新用户成功登录")
                    GetOnlineLists(this.OnlineLists);//获取在线列表并更新
                    GetGroupLists(this.GroupLists, this.username);//获取当前用户群聊关系表
                    break;
                case "1":
                    //私聊功能处理
                    GetMessageHistory(this.messageHistory,this.username,this.messageReceiver,this.messageType);
                    break;
                case "2":
                    //群聊功能处理
                    GetMessageHistory(this.messageHistory, this.username, this.messageReceiver, this.messageType);
                    break;
                case "3":
                    //对机器人进行处理
                    GetMessageHistory(this.messageHistory, this.username, this.messageReceiver, this.messageType);
                    break;
                default:
                    break;
            }
        }
        //在socket成功打开后发送测试信息
        this.socket.onopen = (event) => {
            console.log("开始发送json信息");
            this.sendText("json 消息测试");
        };
    }
    /**
     * Send text to all users through the server
     * @param {string} msg_send 要发送的消息
     */
    sendText(msg_send) {
        //当消息类型为3即对象是chatgpt机器人时额外做出处理
        if (this.messageType.value == "3") {
            this.ChatGPT(msg_send)
        }
        // Construct a msg object containing the data the server needs to process the message from the chat client.
        const msg = {
            Type: this.messageType.value,
            Text: msg_send,
            Receiver: this.messageReceiver.value,
            Sender: this.username,
            Date: Date.now(),
        };
    
        // Send the msg object as a JSON-formatted string.
        this.socket.send(JSON.stringify(msg));
        this.message.value = ""
    }

    /**
     * 调用openai api key与chatgpt进行通话
     * @param {string} msg_send 要向chatgpt发送的消息
     */
    ChatGPT(msg_send) {
        let msg_receive=""
        const OPENAI_API_KEY = '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';
        const headers = {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${OPENAI_API_KEY}`
        };
    
        const data = {
        model: 'gpt-3.5-turbo',
        messages: [
            {
            role: 'user',
            content: msg_send
            }
        ]
        };
    
        fetch('https://api.openai.com/v1/chat/completions', {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(data)
        })
        .then(response => response.json())
            .then(result => {
            // 处理返回的结果
            msg_receive = result.choices[0].message.content
            // Construct a msg object containing the data the server needs to process the message from the chat client.
            const msg = {
                Type: "3",
                Text: msg_receive,
                Receiver: this.username,
                Sender: this.username+"'s Robot",
                Date: Date.now(),
            };
            this.socket.send(JSON.stringify(msg)); 
            console.log(result.choices[0].message.content)
            return msg_receive
        })
        .catch(error => {
            console.error(error);
            // 处理错误
        });
    }
    //文件上传功能
    inputFilter(newFile, oldFile, prevent) {
        if (newFile && !oldFile) {
            // 检查文件类型和大小
            if (['image/jpeg', 'image/png', 'image/gif'].indexOf(newFile.type) === -1) {
                this.$message.error('只能上传jpg、png或gif格式的图片')
                prevent()
            } else if (newFile.size > 1024 * 1024) {
                this.$message.error('文件大小不能超过1MB')
                prevent()
            }
        }
    }
    inputFile(event) {
        // 上传文件到assets文件夹下
        const files = event.target.files;
        if (files.length > 0) {
            const fileName = files[0].name;
            console.log(fileName);

            const reader = new FileReader()
            reader.readAsDataURL(files[0])
            // let imgUrl=URL.createObjectURL(files[0])
            // var imgElement = document.getElementById("test")
            // imgElement.src = dataUrl

            reader.addEventListener("load", (event) => {
            const dataUrl = reader.result
            var imgElement = document.getElementById("test")
            imgElement.src = dataUrl
            const msg = {
                Type: this.messageType.value,
                MsgType:"1",
                Text: dataUrl,
                Receiver: this.messageReceiver.value,
                Sender: this.username,
                Date: Date.now(),
            };
        
            // Send the msg object as a JSON-formatted string.
            this.socket.send(JSON.stringify(msg));
            this.message.value = ""
        })

        }
        // const reader = new FileReader()
        // console.log(event)
        // reader.readAsDataURL(event.file)
        // console.log("user name:",this.username)
        // reader.addEventListener("load", (event) => {
        //     const dataUrl = reader.result
        //     // console.log("dataurl:", dataUrl)

        //     var imgElement = document.getElementById("test")
        //     imgElement.src = dataUrl
        //     const msg = {
        //         Type: this.messageType.value,
        //         MsgType:"1",
        //         Text: dataUrl,
        //         Receiver: this.messageReceiver.value,
        //         Sender: this.username,
        //         Date: Date.now(),
        //     };
        
        //     // Send the msg object as a JSON-formatted string.
        //     this.socket.send(JSON.stringify(msg));
        //     this.message.value = ""
        // })

    }
}