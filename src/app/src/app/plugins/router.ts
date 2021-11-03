import { createRouter, createWebHistory } from 'vue-router'
import SignUp from '~/pages/SigunUp/index.vue'
import Login from '~/pages/Login/index.vue'
import MyPage from '~/pages/MyPage/index.vue'
import ChangePassword from '~/pages/MyPage/changePassword.vue'
import EditMyPage from '~/pages/MyPage/edit.vue'

// 2. Define some routes
// Each route should map to a component.
// We'll talk about nested routes later.
const routes = [
    { path: '/signUp', component: SignUp },
    { path: '/login', component: Login },
    { path: '/', component: MyPage },
    { path: '/myPage', component: MyPage },
    { path: '/myPage/edit', component: EditMyPage },
    { path: '/myPage/change-password', component: ChangePassword },
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
