<template>
    <div class="mainchat">
        <h1>main chat UI {{ username }}</h1>
        <!-- 用户聊天界面设置 -->
        <div class="chat-container">
            <div class="chat-messages">
                <div v-for="message in messageHistory" :key="message.id" class="message"
                    :class="{ 'sent': message.Sender === username, 'received': message.Sender !== username }">
                    <!-- 消息发送者是本人的处理 -->
                    <div v-if="message.Sender === username" :class="{ 'message-row': true }">
                        <div v-if="message.MsgType == '1'" class="img-display">
                            <img :src="message.Text" alt="blank" width="150" height="150">
                        </div>
                        <div v-else class="message-content">
                            {{ message.Text }}
                        </div>
                        <!-- <span>{{ message.Sender }}</span> -->
                        <img :src="GetUserAvatar(message.Sender)" alt="user avatar" width="50" height="50">
                    </div>
                    <!-- 消息接受者是本人的处理 -->
                    <div v-else :class="{ 'message-row': true }">
                        <img :src="GetUserAvatar(message.Sender)" alt="user avatar" width="50" height="50">
                        <!-- <span>{{ message.Sender }}</span> -->
                        <div v-if="message.MsgType == '1'" class="img-display">
                            <img :src="message.Text" alt="blank" width="150" height="150">
                        </div>
                        <div v-else class="message-content">
                            {{ message.Text }}
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <!-- 输入消息文本框设置 -->
        <div class="msg-container">
            请输入聊天内容: <input class="msg-box" v-model="message" @keydown.enter="socket.sendText(message)">
            <br>
            <button type="button" @click="socket.sendText(message)">发送</button>
        </div>
        <!-- 在线用户界面设置 -->
        <h2> 在线用户列表</h2>
        <div class="onlineUser-container">
            <div v-for="user in OnlineLists" :key="user.id" @click="SelectChater(user)" class="message">
                <div class="message-content">
                    {{ user }}
                </div>
            </div>
        </div>
        <!-- 群组界面设置 -->
        <h2>群组列表</h2>
        <div class="onlineUser-container">
            <div v-for="group in GroupLists" :key="group.id" @click="SelectGroup(group)" class="message">
                <div class="message-content">
                    {{ group }}
                </div>
            </div>
        </div>
        <!-- 机器人聊天界面设置 -->
        <h2>机器人列表</h2>
        <div class="onlineUser-container">
            <div @click="SelectRobot()" class="message">
                <div class="message-content">
                    robot
                </div>
            </div>
        </div>
        <!-- 表情界面设置 -->
        <button @click="emojiFlag = !emojiFlag">表情界面切换</button>
        <div v-if="emojiFlag">
            <div class="row">
                <Picker :data="emojiIndex" set="twitter" @select="showEmoji" />
            </div>

            <div class="row">
                <div>
                    {{ emojisOutput }}
                </div>
            </div>
        </div>
        <!-- 群组创建-->
        <button @click="groupFlag = !groupFlag"> 群聊创建界面</button>
        <div v-if="groupFlag">
            群聊名称:<input v-model="groupInfo.Name"> {{ groupInfo.Name }}
            <br>
            群聊描述:<input v-model="groupInfo.Description"> {{ groupInfo.Description }}
            <br>
            群聊群主:<input v-model="groupInfo.Creator"> {{ groupInfo.Creator }}
            <br>
            群聊成员:<input v-model="groupInfo.Members">
            <div>
                <label v-for="user in OnlineLists" :key="user.id">
                    <input type="checkbox" v-model="groupInfo.Members" :id="user.id" :value="user">
                    {{ user }}
                </label>
            </div>

            <br>
            <button @click="GroupCreate(groupInfo)">群聊创建信息发送</button>
        </div>
        <!-- 用户头像上传 -->
        <file-upload ref="fileUpload" name="avatar" @input-filter="socket.inputFilter" @input-file="inputFile">
            <button>发送图片</button>
        </file-upload>
        <img id="test" alt="image test" width="50" height="50">
        <input type="file" id="msg_img" @change="socket.inputFile">
    </div>
</template>

<script>
import { useRoute } from 'vue-router'
import { ref } from 'vue'
import { GetMessageHistory, GetOnlineLists, ChatScroll, GroupCreate, GetUserAvatar } from "../utility/chat"
import { ClientSocket } from "../utility/socket"
// Import data/twitter.json to reduce size, all.json contains data for
// all emoji sets.
import data from "emoji-mart-vue-fast/data/all.json";
// Import default CSS
import "emoji-mart-vue-fast/css/emoji-mart.css";
// Vue 3, import components from `/src`:
import { Picker, EmojiIndex } from "emoji-mart-vue-fast/src";
import VueUploadComponent from 'vue-upload-component';

export default {

    setup() {
        const route = useRoute()
        const username = route.params.name
        const message = ref("")
        const OnlineLists = ref([]) //保存在线用户列表
        const GroupLists = ref([])//保存用户所在群列表
        const messageReceiver = ref("") //客户端消息接收者
        const messageHistory = ref([])  //保存用户聊天历史记录
        const messageType = ref("") //消息类型 1:好友  2:群组
        const groupInfo = ref({ Members: [] }) //群组创建信息
        const emojiFlag = ref(false);
        const groupFlag = ref(false);
        const emojiIndex = new EmojiIndex(data)
        let emojisOutput = ref("")
        //建立websocket信道连接过程
        const socket = new ClientSocket(username, messageReceiver, messageType, message, OnlineLists, GroupLists, messageHistory)
        socket.socketStart()
        //显示表情函数
        function showEmoji(emoji) {
            emojisOutput = emojisOutput + emoji.native;
            message.value = message.value + emoji.native;
        }
        // 选择要聊天的好友
        function SelectChater(user) {
            messageReceiver.value = user;
            messageType.value = "1";
            console.log("messagetype:", messageType.value)
            GetMessageHistory(messageHistory, username, messageReceiver, messageType);
        }
        //选择要聊天的群组
        function SelectGroup(user) {
            messageReceiver.value = user;
            messageType.value = "2";
            console.log("messagetype:", messageType.value)
            GetMessageHistory(messageHistory, username, messageReceiver, messageType);
        }
        //选择和机器人进行聊天
        function SelectRobot() {
            messageReceiver.value = username + "'s Robot"
            messageType.value = "3";
            console.log("正在和机器人聊天", messageReceiver.value)
            GetMessageHistory(messageHistory, username, messageReceiver, messageType);
        }

        return {
            username, message, OnlineLists, GroupLists, messageReceiver, messageHistory, emojiFlag, groupFlag,
            emojiIndex, emojisOutput, socket, groupInfo, messageType,
            GetOnlineLists, GetMessageHistory, SelectChater, SelectGroup, SelectRobot, showEmoji, GroupCreate,
            GetUserAvatar
        }
    },
    components: {
        Picker,
        'file-upload': VueUploadComponent
    },
    mounted() {
    },
    updated() {
        ChatScroll(".chat-messages")
    }

}
</script> 

<style>
/* 文本款布局测试 */
.msg-container {
    position: relative;
    top: 500px;
    left: 400px;
    width: 20%;
    height: 20%;
}

.msg-box {
    padding: 15px;
    border-radius: 5px;
    background-color: #f0f0f0a8;
    width: 980px;
    height: 100px;
}

/* 用户在线列表测试*/
.onlineUser-container {
    width: 10%;
    height: 30%;
    display: flex;
    flex-direction: column;
    border: 5px solid #ebedef;
}

/* 聊天框布局设置 */
.chat-container {
    width: 60%;
    height: 60%;
    display: flex;
    flex-direction: column;
    border: 5px solid #ebedef;

    position: absolute;
    top: 30px;
    left: 400px;
}

.chat-messages {
    flex: 1;
    overflow-y: auto;
}

.message {
    display: flex;
    justify-content: flex-start;
    align-items: center;
    margin-bottom: 10px;
}

.message.sent {
    justify-content: flex-end;
}

.message-content {
    padding: 15px;
    border-radius: 5px;
    background-color: #f0f0f0a8;
}

.message.sent .message-content {
    background-color: #dcf8c6;
    margin-right: 10px;

}

.message.sent .img-display {
    margin-right: 10px;

}

.message.received .message-content {
    background-color: rgba(93, 91, 91, 0.238);
    margin-left: 10px;
}

.message.received .img-display {
    margin-left: 10px;

}

.message-row {
    display: flex;
    align-items: center;
}
</style>