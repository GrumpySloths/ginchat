<template>
    <div class="about">
        <h1>this is about page</h1>
        <input type="text" v-model="username" placeholder="请输入用户名称">
        <h1>{{ username }}</h1>
    </div>
</template>

<script>

export default {
    data() {
        return {
            username: ""
        }
    }
}
</script> 