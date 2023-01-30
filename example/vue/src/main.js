import { createApp } from 'vue'
// use Options API
//import App from './App.vue'
// or Composition API
import CompositionApp from './App.vue'
import astor from './plugins/astor';

// use Options API
//const app = createApp(App)
// or Composition API
const app = createApp(CompositionApp)

app.use(astor, {debug: true})

app.mount('#app')
