import Vue from 'vue'
import App from './App.vue'
import astor from './plugins/astor';

Vue.config.productionTip = false

Vue.use(
  astor,  {
    debug: process.env.NODE_ENV !== 'production'
  }
)

new Vue({
  render: h => h(App),
}).$mount('#app')
