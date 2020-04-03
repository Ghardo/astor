import Vue from 'vue'
import App from './App.vue'
import astor from './plugins/astor';

Vue.config.productionTip = false

Vue.use(
  astor,  {
    debug: true
  }
)

document.addEventListener('astilectron-ready', function() {
  new Vue({
    render: h => h(App)
  }).$mount('#app')
})

