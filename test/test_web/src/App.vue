<template>
    <h1>Hello App!</h1>
    <p>
        <!-- use the router-link component for navigation. -->
        <!-- specify the link by passing the `to` prop. -->
        <!-- `<router-link>` will render an `<a>` tag with the correct `href` attribute -->
        <router-link to="/">Go to Home</router-link> |
        <router-link to="/about">Go to About</router-link>|
        <router-link to="/chat">Go to chat</router-link>|
        <router-link to="/emoji">Go to emoji</router-link>

    </p>
    <!-- route outlet -->
    <!-- component matched by the route will render here -->
    <router-view></router-view>
</template>
<script setup>

// import Home from './components/Home.vue'


// 1. Define route components.
// These can be imported from other files
// const route = VueRouter.useRoute()

// Now the app has started!
</script>
