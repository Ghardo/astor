# astor

**astor** is a vuejs 3.x plugin for communicate with a go-astilectron app

**NOTE** vuejs 2 is only supported on astor v1.x

## Install

copy astor.js to PROJECTROOT/plugins/astor.js

## Usage

### Plugin registration

```:js
import { createApp } from 'vue'
import App from './App.vue'
import astor from './plugins/astor';

const app = createApp(App)

app.use(astor, { debug: true })

app.mount('#app')
```

### Sending and receiving with Options API

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

### Sending and receiving with Composition API

```vue
<template>
  <div>Hello Word</div>
</template>

<script setup>
import useAstor from './plugins/astor'
import onBeforeMount from 'vue'

const astor = useAstor()

onBeforeMount(() => {
  astor.trigger(
    'app-ready',
    {}, 
    (payload) => {
      window.console.log("ready", payload)
    }
  )
})
</script>

<style>
</style>
```

### Receiving only with Options API

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

### Receiving only with Composition API

```vue
<template>
  <div>Hello Word</div>
</template>

<script setup>
  import useAstor from './plugins/astor'
  import onBeforeMount from 'vue'

  const astor = useAstor()

  const cbMyCustomMessage = (payload) => {
    window.console.log('cbMyCustomMessage', payload)
  }

  onBeforeMount(() => {
    astor.listen(
      'my-custom-message',
      cbMyCustomMessage
    )
  })
</script>

<style>
</style>
```

```:go
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

```:shell
cd example/vue
npm install
npm run dev
cd ..
go run main.go
```
