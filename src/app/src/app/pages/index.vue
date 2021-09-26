<template>
  <section class="hero is-primary is-fullheight">
    <p>Your are Logged IN !!!</p>
    <button @click="createNewAccessToken">Send Request</button>
  </section>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import {checkAuth} from "../composables/auth/auth";
import {createNewAccessToken} from "../composables/jwt/jwt";

export default defineComponent({
  setup: () => {
    const err = false
    checkAuth();

    return {
      createNewAccessToken,
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
