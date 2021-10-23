<template>
  <section class="hero is-primary is-fullheight">
    <div class="hero-body">
      <div class="container">
        <div class="columns is-centered">
          <div class="column is-5-tablet is-4-desktop is-3-widescreen">
            <div class="box" v-if="useUserResult">
              <div class="field">
                <label class="label">名前</label>
                <div class="control">
                  <p>{{user.name}}</p>
                </div>
              </div>
              <div class="field">
                <label class="label">メールアドレス</label>
                <div class="control">
                  <p>{{user.email}}</p>
                </div>
              </div>
              <router-link to="/myPage/edit" class="button is-primary">編集</router-link>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script lang="ts">
import {defineComponent} from 'vue'
import {useUser} from "../../composables/user/user";

export default defineComponent({
  setup: () => {
    const err = false
    const { user, useUserResult } = useUser()

    return {
      err,
      user,
      useUserResult,
    }
  },
})
</script>

<style scoped lang="css">
</style>
