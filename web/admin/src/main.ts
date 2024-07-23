import './assets/main.css'

// 国际化
import ElementPlus from 'element-plus'
import zhCn from 'element-plus/es/locale/lang/zh-cn'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(ElementPlus, {
    locale: zhCn,
})

app.use(createPinia())

app.use(router)

app.mount('#app')
