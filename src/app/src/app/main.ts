import { createApp } from 'vue'
import App from '~/App.vue'
import { router } from './plugins/router.ts'
import 'bulma'
// 5. Create and mount the root instance.

createApp(App)
    .use(router)
    .mount('#app')
