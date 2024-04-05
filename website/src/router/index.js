import Vue from 'vue'
import Router from 'vue-router'
import HomeView from '@/components/HomeView'
import IndexView from '@/components/IndexView'

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
    }
  ]
})
