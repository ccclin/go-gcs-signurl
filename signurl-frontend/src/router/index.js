import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

function view (name) {
  return function (resolve) {
    import(`../components/${name}.vue`).then(mod => {
      resolve(mod)
    })
  }
}

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'SignUrl',
      component: view('SignUrl')
    }
  ],
  scrollBehavior (to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    } else {
      return { y: 0 }
    }
  }
})
