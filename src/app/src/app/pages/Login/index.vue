<template>
  <div class="container">
    Login
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  setup: () => {
    const err = false
    return {
      err,
    }
  },
})
</script>

<style scoped lang="css">
.container{
  width: 360px;
  margin: 0 auto;
  padding-top: 64px;
  padding-bottom: 64px;
}
</style>
