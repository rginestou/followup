import Vue from 'vue'
import Router from 'vue-router'

import Main from '@/components/Main'
import Login from '@/components/Login'

Vue.use(Router)

let router = new Router({
  routes: [{
    path: '/login',
    name: 'login',
    component: Login,
  }, {
    path: '/',
    name: 'main',
    component: Main,
    meta: {
      requiresAuth: true,
    },
  }],
  mode: 'history',
})

export default router
