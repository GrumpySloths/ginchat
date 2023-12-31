<template>
    <div class="login">
        <h1>Login</h1>
        <form @submit.prevent="login">
            <div>
                <label for="user_name">Username:</label>
                <input type="text" id="user_name" v-model="username" required>
            </div>
            <div>
                <label for="password">Password:</label>
                <input type="password" id="password" v-model="password" required>
            </div>
            <button @click.once="login">Login</button>
        </form>
    </div>
</template>
  
  
<script >

export default {
    data() {
        return {
            username: "",
            password: "",
        }
    },

    methods: {
        login(event) {
            const formData = new FormData();
            formData.append("name", this.username)
            formData.append("passwd", this.password)
            formData.append("rePasswd", this.password)
            const request = new Request("http://localhost:8080/user/UserLogin", {
                method: "POST",
                body: formData,
            });
            console.log("login is trigged")
            fetch(request)
                .then((response) => {
                    if (response.status === 200) {
                        console.log(response.json())
                        this.$router.push({ path: `/mainchat/${this.username}` })
                        return true
                    } else {
                        throw new Error("Something went wrong on API server!");
                    }
                })
                .then((response) => {
                    console.log("logging is dubugging")
                    console.debug(response);
                    // …
                })
                .catch((error) => {
                    console.error(error);
                });
            if (event) {
                event.preventDefault()
            }

            return true
        }
        //socket连接验证
    }
}

</script>
  