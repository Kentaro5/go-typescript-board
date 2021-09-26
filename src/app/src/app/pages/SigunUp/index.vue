<template>
    <div class="container">
      <form class="form-horizontal" method="POST" action="http://localhost:8000/signUp">
        <input type="hidden" name="_token">
        <div class="field ">
          <div class="notification is-danger" v-if="err">
            <strong></strong>
          </div>
          <label for="name" class="label">ユーザー名（半角英数字で入力してください。）</label>
          <div class="control">
            <input id="name" name="name"
                   class="input" type="text"
                   placeholder="ユーザー名を半角英数字で入力してください。">
          </div>
        </div>
        <div class="field ">
          <div class="notification is-danger" v-if="err">
            <strong></strong>
          </div>
          <label for="sex_code" class="label">性別</label>
          <div class="control">
            <div class="select">
              <select name="sex_code">
                <option value="">性別を選択してください。</option>
                <option value="0">その他</option>
              </select>
            </div>

          </div>
        </div>
        <div class="field ">
          <div class="notification is-danger" v-if="err">
            <strong></strong>
          </div>
          <label for="email" class="label">メールアドレス</label>
          <div class="control">
            <input id="email" name="email"
                   class="input" type="text"
                   placeholder="example@example.com">
          </div>
        </div>
        <div class="field ">
          <div class="notification is-danger" v-if="err">
            <strong></strong>
          </div>
          <label for="password" class="label">パスワード</label>
          <div class="control">
            <input id="password" name="password"
                   class="input"
                   type="password"
                   >
          </div>
        </div>
        <div class="field ">
          <div class="control">
            <div class="notification is-danger" v-if="err">
              <strong></strong>
            </div>
            <label for="pref_code" class="label">都道府県</label>
            <div class="select">
              <select name="pref_code">
                <option value="">都道府県を選択してください</option>
                <option value="10006">北海道</option>
              </select>
            </div>
          </div>
          <div class="control">
            <div class="notification is-danger" v-if="err">
              <strong></strong>
            </div>
            <label for="city_code" class="label">市</label>
            <div class="select">
              <select name="city_code">
                <option value=""></option>
                <option value="11002">札幌市</option>
              </select>
            </div>
          </div>
          <div class="control">
            <label for="ward_code" class="label">区</label>
            <div class="select">
              <select name="ward_code">
                <option value=""></option>
                <option value="11011">札幌市</option>
              </select>
            </div>

          </div>
        </div>
        <div class="field is-grouped">
          <div class="control">
            <button class="button is-link">登録
            </button>
          </div>
        </div>
      </form>
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
</style>
