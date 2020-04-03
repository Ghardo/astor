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