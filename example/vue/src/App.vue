<template>
  <div id="app">
    <p>{{text1}}</p>
    <p>{{text2}}</p>
    <button @click="btnClick()">{{text3}}</button>
  </div>
</template>

<script>
export default {
  name: 'App',
  data: () => ({
    text1: "Astor Expample",
    text2: null,
    text3: "Click me"
  }),
  created () {
    this.$astor.onIsReady(this.astorIsReady);
  },
  methods: {
    astorIsReady () {
      this.text = "Astor is ready"
      this.$astor.trigger('test.event', {}, this.testEvent)
      this.$astor.listen('butten.text', this.setButtonText)
    },
    testEvent (payload) {
      this.text2 = payload.text
    },
    btnClick () {
      this.$astor.trigger('btn.click', {"buttonId": 1337})
    },
    setButtonText (payload) {
      this.text3 = payload.button + ' ' + payload.text;
    }
  }
}
</script>

<style>
#app {
  color: white;
  font-weight: bold;
  font-size: 24px;
}
</style>
