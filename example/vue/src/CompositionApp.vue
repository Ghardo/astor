<script setup>
  import useAstor from './plugins/astor'
  import onBeforeMount from 'vue'

  const astor = useAstor()
  const text1 = ref("Astor example")
  const text2 = ref(null)
  const text3 = ref("Click me")

  const testEvent = (payload) => {
    text2.value = payload.text
  }

  const setButtonText = (payload) => {
    text3.value = payload.button + ' ' + payload.text;
  }

  const astorIsReady = () => {
    text1.value = "Astor is ready"
    astor.trigger('test.event', {}, testEvent)
    astor.listen('butten.text', setButtonText)
  }

  onBeforeMount(() => {
    astor.onIsReady(astorIsReady)
  })

  const btnClick = () => {
    astor.trigger('btn.click', {"buttonId": 1337})
  }
</script>

<template>
  <div id="app">
    <p>{{text1}}</p>
    <p>{{text2}}</p>
    <button @click="btnClick()">{{text3}}</button>
  </div>
</template>

<style>
#app {
  color: white;
  font-weight: bold;
  font-size: 24px;
}
</style>
