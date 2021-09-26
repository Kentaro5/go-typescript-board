<template>
  <section class="hero is-primary is-fullheight">
    <div class="hero-body">
      <div class="container">
        <div class="columns is-centered">
          <div class="column is-5-tablet is-4-desktop is-3-widescreen">
            <form class="box" @submit.prevent="sendRequest">
              <div class="field">
                <label for="" class="label">Email</label>
                <div class="control has-icons-left">
                  <input type="email" name="email"  placeholder="e.g. bobsmith@gmail.com" class="input" required>
                  <span class="icon is-small is-left">
                  <i class="fa fa-envelope"></i>
                </span>
                </div>
              </div>
              <div class="field">
                <label for="" class="label">Password</label>
                <div class="control has-icons-left">
                  <input type="password" name="password" placeholder="*******" class="input" required>
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
            </form>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import axios from "axios"

export default defineComponent({
  setup: () => {
    const err = false

    const sendRequest = () => {
      // TODO: あとで、フォームの値をバインドする。
      const data = {
        email: 'test@example.com',
        password: 'test',
      }
      axios.post('http://localhost:8000/login', data).then(function (response) {
        const result = response.data
        console.log(result.data);
        if (result.status === 200) {
          localStorage.setItem('accessToken', result.data.access_token)
          localStorage.setItem('refreshToken', result.data.refresh_token)
          localStorage.setItem('user', result.data.user)
          location.href = '/'
        }
      })
    }
    return {
      err,
      sendRequest,
    }
  },
})
</script>

<style scoped lang="css">
</style>
