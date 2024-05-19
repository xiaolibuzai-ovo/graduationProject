// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import VueParticles from 'vue-particles'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
Vue.use(VueParticles)

Vue.config.productionTip = false
Vue.use(ElementUI);
/* eslint-disable no-new */
Vue.prototype.$logined = false;
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
