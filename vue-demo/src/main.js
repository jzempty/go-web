import Vue from 'vue'
import App from './App.vue'

import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
Vue.use(ElementUI);

import axios from 'axios'
Vue.prototype.$axios = axios
Vue.config.productionTip = false;

//引入路由
// eslint-disable-next-line no-unused-vars
import router from './router/router'

Vue.config.productionTip = false

new Vue({
  el:'#app',
  render: h => h(App),
  router,
}).$mount('#app')
