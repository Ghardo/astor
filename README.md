# astor
**astor** is a vuejs plugin for communicate with a go-astilectron app

## Install
copy astor.js to PROJECTROOT/plugins/astor.js

## Usage

### Plugin registration

```vue
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
```

The listener "astilectron-ready" around new Vue is required to prevent timing errors.


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
...


See example for more info

