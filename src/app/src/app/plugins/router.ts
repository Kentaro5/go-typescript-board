import { createRouter, createWebHistory } from 'vue-router'
import SignUp from '~/pages/sigunUp/index.vue'
import Login from '~/pages/Login/index.vue'
import Top from '~/pages/index.vue'

// 2. Define some routes
// Each route should map to a component.
// We'll talk about nested routes later.
const routes = [
    { path: '/signUp', component: SignUp },
    { path: '/login', component: Login },
    { path: '/', component: Top },
]

// 3. Create the router instance and pass the `routes` option
// You can pass in additional options here, but let's
// keep it simple for now.
export const router = createRouter({
    // 4. Provide the history implementation to use. We are using the hash history for simplicity here.
    history: createWebHistory(),
    routes, // short for `routes: routes`
})

// Now the app has started!
