import { createRouter, createWebHistory } from 'vue-router'
import SignUrlView from '@/views/SignUrlView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'signurl',
      component: SignUrlView
    }
  ]
})

export default router
