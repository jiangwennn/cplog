import History from '@/views/History.vue'
import { createRouter, createWebHistory } from 'vue-router'


const routes = [
    {
        path: '/',
        name: 'History',
        component: History
    },
    {
        path: '/like',
        name: 'Like',
        component: () => import('@/views/Like.vue')
    },
    {
        path: '/config',
        name: 'Config',
        component: () => import('@/views/Config.vue')
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router