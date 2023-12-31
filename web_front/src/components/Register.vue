<template>
  <div class="register">
    <h1>Register</h1>
    <form @submit.prevent="login">
      <div>
        <label for="user_name">Username:</label>
        <input type="text" id="user_name" v-model="username" required>
        <!-- <input placeholder="Username" className="username" v-model="username" required> -->
      </div>
      <div>
        <label for="password">Password:</label>
        <input type="password" id="password" v-model="password" required>
      </div>
      <div>
        <label for="repassword">Password Again:</label>
        <input type="password" id="repassword" v-model="rePassword" required>
      </div>
      <button @click="register">Register</button>
    </form>
  </div>
</template>
  
<script >
export default {
  data() {
    return {
      username: "",
      password: "",
      rePassword: ""
    }
  },
  methods: {
    register(event) {
      const formData = new FormData();
      formData.append("name", this.username)
      formData.append("passwd", this.password)
      formData.append("rePasswd", this.rePassword)
      const request = new Request("http://localhost:8080/user/CreateUser", {
        method: "POST",
        body: formData,
      });
      fetch(request)
        .then((response) => {
          if (response.status === 200) {
            console.log(response.json())
            return
          } else {
            throw new Error("Something went wrong on API server!");
          }
        })
        .then((response) => {
          console.debug(response);
          // …
        })
        .catch((error) => {
          console.error(error);
        });
      if (event) {
        event.preventDefault()
      }
    }
  }
}



</script>
  