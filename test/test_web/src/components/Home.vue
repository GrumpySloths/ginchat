<template>
    <input type="text" v-model="username" placeholder="请输入用户名称">
    <h1>{{ username }}</h1>
</template>

<script setup>
import ref from 'vue'
username = ref("")
</script>
