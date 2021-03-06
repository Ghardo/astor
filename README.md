# astor
**astor** is a vuejs plugin for communicate with a go-astilectron app

**NOTE** vue3 support curently in development

## Install
copy astor.js to PROJECTROOT/plugins/astor.js

## Usage

### Plugin registration

```js
import Vue from 'vue'
import App from './App.vue'
import astor from './plugins/astor';

Vue.config.productionTip = false

Vue.use(
  astor,  {
    debug: true
  }
)


new Vue({
  render: h => h(App)
}).$mount('#app')
```

### Sending and recieving

```vue
<template>
  <div>Hello Word</div>
</template>

<script>
export default {
     created () {
      this.$astor.trigger(
        'app-ready', 
        {}, 
        (payload) => {
          window.console.log("ready", payload)
        }
      )
    },
}
</script>

<style>
</style>
```

### Recieving only

```vue
<template>
  <div>Hello Word</div>
</template>

<script>
export default {
     created () {
      this.$astor.listen(
        'my-custom-message', 
        this.cbMyCustomMessage
      )
    },
    methods: {
        cbMyCustomMessage(payload) {
            window.console.log('cbMyCustomMessage', payload)
        }
    }
}
</script>

<style>
</style>
```

```go

type AstorEvent struct {
	Name    string      `json:"name"`
	Payload interface{} `json:"payload"`
}
...
w.OnMessage(func(m *astilectron.EventMessage) interface{} {
    // Unmarshal
    var e AstorEvent
    m.Unmarshal(&e)

    l.Printf("%v", e)

    r := AstorEvent{Name: e.Name, Payload: "Ok"}
    return r
})
```


See example for more info

```
cd example/vue
yarn install
yarn build
cd ..
go run main.go
```

