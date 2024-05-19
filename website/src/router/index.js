import Vue from 'vue'
import Router from 'vue-router'
import IndexView from '@/components/IndexView'
import HomeView from '@/components/HomeView'
import Chat from '@/components/Chat'
import Agents from '@/components/Agents'
import Login from '@/components/login'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/', // 定义路由路径
      name: 'index', // 定义路由名称
      component: IndexView // 关联的组件
    },
    {
      path: '/home',
      name: 'home',
      component: HomeView
    },
    {
      path: '/chat',
      name: 'chat',
      component: Chat
    },
    {
      path: '/agents',
      name: 'agents',
      component: Agents
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    }
  ]
})
