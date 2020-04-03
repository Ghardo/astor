<template>
  <div>
    <div>Hello Astor</div>
    <div>{{ text  }}</div>
    <ul>
      <li v-for="(message, index) in messages" :key="index">
        {{ message }}
    </ul>
  </div>
</template>

<script>
  export default {
    data: () => ({
      text: '',
      messages: []
    }),

    created () {
      this.$astor.trigger('app-ready', {}, this.ready)
    },
    mounted () {
      this.$astor.listen('go-custom-message', this.customMessage)
    }, 
    methods: {
      ready (payload) {
        window.console.log("ready", payload)
        this.text = payload
      },
      customMessage (payload) {
        window.console.log("custom message", payload)
        this.messages.push(payload)
      }
    },
  }
</script>

<style>
</style>