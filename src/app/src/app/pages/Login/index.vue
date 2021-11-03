<template>
  <section class="hero is-primary is-fullheight">
    <div class="hero-body">
      <div class="container">
        <div class="columns is-centered">
          <div class="column is-5-tablet is-4-desktop is-3-widescreen">
            <div class="notification is-danger is-light" v-if="loginStatus === 'failed'" @click="changeLoginStatus('')">
              <button class="delete"></button>
              <p>ログインに失敗しました。</p>
            </div>
            <form class="box" @submit.prevent="login">
              <div class="field">
                <label for="email" class="label">Email</label>
                <div class="control has-icons-left">
                  <input type="email" name="email"  placeholder="e.g. bobsmith@gmail.com" v-model="loginEmail" class="input" required>
                  <span class="icon is-small is-left">
                  <i class="fa fa-envelope"></i>
                </span>
                </div>
              </div>
              <div class="field">
                <label for="password" class="label">Password</label>
                <div class="control has-icons-left">
                  <input type="password" name="password" placeholder="*******" class="input" v-model="loginPassword" required>
                  <span class="icon is-small is-left">
                  <i class="fa fa-lock"></i>
                </span>
                </div>
              </div>
              <div class="field">
                <button class="button is-success">
                  Login
                </button>
              </div>
              <div class="field is-grouped">
                <div class="control">
                  <router-link to="/signUp" class="button is-link">登録</router-link>
                </div>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import {useLogin} from "../../composables/auth/login";

export default defineComponent({
  setup: () => {
    const {loginEmail, loginPassword, login, loginStatus, changeLoginStatus} = useLogin()

    return {
      loginEmail,
      loginPassword,
      login,
      loginStatus,
      changeLoginStatus,
    }
  },
})
</script>

<style scoped lang="css">
</style>
