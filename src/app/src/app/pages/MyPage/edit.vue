<template>
  <section class="hero is-primary is-fullheight">
    <div class="hero-body">
      <div class="container">
        <div class="columns is-centered">
          <div class="column is-5-tablet is-4-desktop is-3-widescreen">
            <div class="box" v-if="useUserResult">
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
                           placeholder="ユーザー名を半角英数字で入力してください。"
                           :value="user.name"
                    >
                  </div>
                </div>
                <div class="field ">
                  <div class="notification is-danger" v-if="err">
                    <strong></strong>
                  </div>
                  <label for="sex_code" class="label">性別</label>
                  <div class="control">
                    <div class="select">
                      <select name="sex_code" v-if="useSexResult">
                        <option value="">性別を選択してください。</option>
                        <option :value="sex.code" v-for="(sex, index) in sexes" :key="index">{{sex.name}}</option>
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
                           placeholder="example@example.com"
                           :value="user.email"
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
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script lang="ts">
import {defineComponent} from 'vue'
import {useUser} from "../../composables/user/user";
import {useSex} from "../../composables/sex/sex";

export default defineComponent({
  setup: () => {
    const err = false
    const { user, useUserResult } = useUser()
    const { sexes, useSexResult } = useSex()

    return {
      err,
      user,
      sexes,
      useUserResult,
      useSexResult,
    }
  },
})
</script>

<style scoped lang="css">
</style>
