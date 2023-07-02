import {createApp} from 'vue'
import App from './App.vue'
import './tailwind.css'
import router from '@/router/index'
import store from '@/store/index'

const app = createApp(App)
    .use(router)
    .use(store)
    .mount('#app')
