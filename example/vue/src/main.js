import { createApp } from 'vue'
import App from './App.vue'
import astor from './plugins/astor';

const app = createApp(App)

app.use(astor, {debug: true})

app.mount('#app')
