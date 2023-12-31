console.log("开始发送json信息")
const msg_group = {
    Name: "mygo乐队聊天群",
    Description: "感受重力的旋涡，一辈子都要组乐队!!!(*^▽^*)",
    Creator_id: 1,
    Member_ids:[1,3,5]
};
// const formData = new FormData();
// formData.append("Name", "mygo乐队聊天群")
// formData.append("Description", "感受重力的旋涡，一辈子都要组乐队!!!(*^▽^*)")
// formData.append("Creator_id", 1)
// formData.append("Member_ids", [1,3,5])
const request = new Request("http://localhost:8080/user/GroupCreate", {
    method: "POST",
    body: JSON.stringify(msg_group),
});
console.log("debug point")
fetch(request)
    .then((res) => res.json())
    .then((data) => {
        console.log(data)
    })