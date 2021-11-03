<template>
  <section class="hero is-primary is-fullheight">
    <div class="hero-body">
      <div class="container">
        <div class="notification is-primary is-light" v-if="signUpResult">
          <button class="delete" @click="changeUpdatedStatus(false)"></button>
          <p>ユーザーが登録されました。以下のリンクよりログインください。</p>
          <router-link to="/login" class="button is-primary">ログイン画面へ遷移する</router-link>
        </div>
        <div class="columns is-centered">
          <div class="column is-5-tablet is-4-desktop is-3-widescreen">
            <div class="box">
              <form class="form-horizontal" method="POST" @submit.prevent="signUp">
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
                           v-model="name"
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
                      <select name="sex_code" v-if="useSexResult" v-model="sexCode">
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
                  <label for="password" class="label">パスワード</label>
                  <div class="control">
                    <input type="password" name="password" placeholder="*******" class="input" v-model="password">
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
                           v-model="email"
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
                      <select name="pref_code" @change="executeChangeCity($event)" v-model="prefCode">
                        <option value="">都道府県を選択してください</option>
                        <option v-for="(prefecture) in prefectures"
                                :key="prefecture.id"
                                :value="prefecture.code"
                        >
                          {{prefecture.name}}
                        </option>
                      </select>
                    </div>
                  </div>
                  <div class="control">
                    <div class="notification is-danger" v-if="err">
                      <strong></strong>
                    </div>
                    <label for="city_code" class="label">市</label>
                    <div class="select">
                      <select v-if="isCityDataSet" name="city_code" @change="executeChangeWard($event)" v-model="cityCode">
                        <option value=""></option>
                        <option v-for="(cityList) in cityLists"
                                :key="cityList.id"
                                :value="cityList.code"
                        >
                          {{cityList.name}}
                        </option>
                      </select>
                      <select v-else name="city_code" class="select-disabled">
                        <option value=""></option>
                      </select>
                    </div>

                  </div>
                  <div class="control">
                    <label for="ward_code" class="label">区</label>
                    <div class="select">
                      <select v-if="isWardsDataSet" name="ward_code" v-model="wardCode">
                        <option value=""></option>
                        <option v-for="(wardList) in wardLists"
                                :key="wardList.id"
                                :value="wardList.code"
                        >
                          {{wardList.name}}
                        </option>
                      </select>
                      <select v-else name="ward_code" class="select-disabled">
                        <option value=""></option>
                      </select>
                    </div>
                  </div>
                </div>
                <div class="field is-grouped">
                  <div class="control">
                    <button class="button is-link">更新
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
import {defineComponent, readonly, ref} from 'vue'
import {useSingUp, getCities, getWards} from "../../composables/signUp/signUp";

export default defineComponent({
  setup: () => {
    const err: boolean = false
    const isCityDataSet = ref<boolean>(false)
    const isWardsDataSet = ref<boolean>(false)

    const cityLists = ref<[] | null>(null)
    const wardLists = ref<[] | null>(null)

    const {
      signUpResult,
      signUp,
      name,
      sexCode,
      password,
      email,
      prefCode,
      cityCode,
      wardCode,
      prefectures,
      usePrefecturesResult,
      useSexResult,
      sexes,
      changeUpdatedStatus,
    } = useSingUp()

    const resetAreaData = () => {
      cityLists.value = []
      wardLists.value = []
      isCityDataSet.value = false
      isWardsDataSet.value = false
    }

    const executeChangeCity = async (event: Event) => {
      resetAreaData()
      const prefCode = parseInt(event.target.value)
      const {cities, changeCityResult} = await getCities(prefCode)
      isCityDataSet.value = changeCityResult
      cityLists.value = cities
    }

    const executeChangeWard = async (event: Event) => {
      const cityCode = parseInt(event.target.value)
      const {wards, changeWardResult} = await getWards(cityCode)
      isWardsDataSet.value = changeWardResult
      wardLists.value = wards
    }

    return {
      signUpResult,
      err,
      name,
      password,
      sexCode,
      email,
      prefCode,
      cityCode,
      wardCode,
      sexes,
      prefectures,
      cityLists,
      wardLists,
      useSexResult,
      usePrefecturesResult,
      executeChangeCity,
      executeChangeWard,
      changeUpdatedStatus,
      isCityDataSet,
      isWardsDataSet,
      signUp,
    }
  },
})
</script>

<style scoped lang="css">
.select-disabled {
  pointer-events: none;
  background-color: #f7f7f7;
}
</style>
