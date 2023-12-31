import { createRouter, createWebHashHistory } from 'vue-router'
import Login from '@/components/Login.vue'
import Register from '@/components/Register.vue'
import MainChat from '@/components/MainChat.vue'
import Avatar from '@/components/Avatar.vue'
// 2. Define some routes
// Each route should map to a component.
// We'll talk about nested routes later.
const routes = [
    { path: '/', name: "Login", component: Login },
    { path: '/register', name: "Register", component: Register },
    { path: '/mainchat/:name', name: "mainchat", component: MainChat },
    { path:'/avatar',name:"Avatar",component:Avatar}
]

// 3. Create the router instance and pass the `routes` option
// You can pass in additional options here, but let's
// keep it simple for now.
const router = createRouter({
    // 4. Provide the history implementation to use. We are using the hash history for simplicity here.
    history: createWebHashHistory(),
    routes, // short for `routes: routes`
})

export default router 