# astor
**astor** is a vuejs 3.x plugin for communicate with a go-astilectron app

**NOTE** vuejs 2 is only supported on astor v1.x

## Install
copy astor.js to PROJECTROOT/plugins/astor.js

## Usage

### Plugin registration

```js
import { createApp } from 'vue'
import App from './App.vue'
import astor from './plugins/astor';

const app = createApp(App)

app.use(astor, { debug: true })

app.mount('#app')
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

