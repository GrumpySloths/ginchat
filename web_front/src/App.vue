<template>
  <p>
    <!-- use the router-link component for navigation. -->
    <!-- specify the link by passing the `to` prop. -->
    <!-- `<router-link>` will render an `<a>` tag with the correct `href` attribute -->
    <router-link to="/">Go to Login</router-link> |
    <router-link to="/register">Go to register</router-link>|
    <router-link to="/avatar">Go to Avatar</router-link>
  </p>
  <!-- route outlet -->
  <!-- component matched by the route will render here -->
  <router-view></router-view>
</template>
<script setup>

// import Home from './components/Home.vue'


</script>
