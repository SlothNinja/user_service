import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/components/home/Home'
import New from '@/components/user/New'
import Edit from '@/components/user/Edit'
import Show from '@/components/user/Show'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/new',
      name: 'newuser',
      component: New
    },
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/show/:id',
      name: 'show',
      component: Show
    },
    {
      path: '/edit/:id',
      name: 'edit',
      component: Edit
    },
    {
      path: '/logout/:redirect',
      name: 'logout',
      beforeEnter: (to, from, next) => {
        window.location = 'http://localhost:8080/' + to.params.redirect
        next()
      }
    }
  ]
})
