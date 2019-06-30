import Vue from 'vue'
import axios from 'axios'

import router from './router'
import App from './App.vue'
import SDK from './sdk/sdk'
import auth from './sdk/auth'

Vue.config.productionTip = false

Vue.use(SDK)

Vue.http = axios
Vue.prototype.$http = axios

Vue.config.errorHandler = (err, vm, info) => {
  if (err.response && err.response.status === 401)
    vm.$router.push('/login')
}

new Vue({
  router,
  render: h => h(App),
  data: {
    user: auth().getUser(),
  },
}).$mount('#app')
