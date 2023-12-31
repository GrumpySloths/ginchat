//获取与好友或者群组的聊天记录
export function GetMessageHistory(messageHistory,username,messageReceiver,messageType) {
    console.log("GetMessageHistory is trigged")
    messageHistory.value = []
    const request = new Request(`http://localhost:8080/user/GetMessageHistory?Sender=${username}&Receiver=${messageReceiver.value}&Type=${messageType.value}`);
    fetch(request)
        .then((res) => res.json())
        .then((data) => {
            console.log(data)
            for (const msg in data.MessageHistorys) {
                messageHistory.value.push(data.MessageHistorys[msg])
            }
            console.log("msgHistory:", messageHistory)
        })
}

//获取用户列表信息
export function GetOnlineLists(OnlineLists) {
    console.log("getonlineLists is trigged")
    OnlineLists.value = []
    const request = new Request(`http://localhost:8080/user/GetOnlineUsers`);
    fetch(request)
        .then((res) => res.json())
        .then((data) => {
            console.log(data)
            for (const user in data.OnlineUsers) {
                console.log(data.OnlineUsers[user])
                //注意这里ref变量一定要用到.value方法才能真正意义上赋值
                OnlineLists.value.push(data.OnlineUsers[user])
            }
            console.log("onlineLists:", OnlineLists)
        })
}
//获取用户所在群组列表
export function GetGroupLists(GroupLists,username) {
    console.log("getonlineLists is trigged")
    GroupLists.value = []
    const request = new Request(`http://localhost:8080/user/GetGroups?ClientName=${username}`);
    fetch(request)
        .then((res) => res.json())
        .then((data) => {
            console.log(data)
            for (const group in data.GroupList) {
                console.log(data.GroupList[group])
                //注意这里ref变量一定要用到.value方法才能真正意义上赋值
                GroupLists.value.push(data.GroupList[group])
            }
            console.log("GroupLists:", GroupLists)
        })
}
//始终保持聊天界面在最底部
/**
 * 
 * @param {string} chat_compent 聊天css组件
 */
export function ChatScroll(chat_component) {
    const container = document.querySelector(chat_component);
    console.log("tick test:", container.scrollTop);
    container.scrollTop = container.scrollHeight;
    console.log("tick test:", container.scrollTop);
}
//创建聊天群
export function GroupCreate(groupInfo) {

    const request = new Request("http://localhost:8080/user/GroupCreate", {
        method: "POST",
        body: JSON.stringify(groupInfo),
    });
    fetch(request)
    .then((res) => res.json())
    .then((data) => {
        console.log(data)
    })

}
/**
 * 
 * @param {string} username 用户名称
 * @returns 用户头像在本地的地址路径
 */
export function GetUserAvatar(username) {
    
    return new URL(`../assets/${username}.png`,import.meta.url).href
} 

export function handleFileChange(event) {
    console.log(event)
    const files = event.target.files;
    if (files.length > 0) {
      const fileName = files[0].name;
      console.log(fileName);
    }
}