import Vue from 'vue'
import axios from 'axios'

import router from './router'
import App from './App.vue'
import SDK from './sdk/sdk'

Vue.config.productionTip = false

Vue.use(SDK)

Vue.http = axios
Vue.prototype.$http = axios

new Vue({
	router,
	render: h => h(App),
}).$mount('#app')
