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
                      <select name="pref_code" @change="executeChangeCity($event)">
                        <option value="">都道府県を選択してください</option>
                        <option v-for="(prefecture, prefectureIndex) in prefectures"
                                :key="prefectureIndex"
                                :value="prefecture.code">
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
                      <select v-if="isCityDataSet" name="city_code">
                        <option value=""></option>
                        <option v-for="(cityList) in cityLists"
                                :key="cityList.id"
                                :value="cityList.code">
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
                      <select name="ward_code">
                        <option value=""></option>
                        <option value="11011">札幌市</option>
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
import {defineComponent, ref} from 'vue'
import {useUser} from "../../composables/user/user";
import {useSex} from "../../composables/sex/sex";
import {usePrefectures} from "../../composables/areas/prefecture";
import {useCities} from "../../composables/areas/city";

export default defineComponent({
  setup: () => {
    const err: boolean = false
    const isCityDataSet = ref<boolean>(false)
    const cityLists = ref<[] | null>(null)

    const { user, useUserResult } = useUser()
    const { sexes, useSexResult } = useSex()
    const { prefectures, usePrefecturesResult, prefectureIndex } = usePrefectures()
    const executeChangeCity = async (event: Event) => {
      const prefCode = parseInt(event.target.value)
      const {cities, changeCityResult} = await useCities(prefCode)
      isCityDataSet.value = changeCityResult
      cityLists.value = cities
    }

    return {
      err,
      user,
      sexes,
      prefectures,
      cityLists,
      useUserResult,
      useSexResult,
      usePrefecturesResult,
      prefectureIndex,
      executeChangeCity,
      isCityDataSet,
    }
  },
})
</script>

<style scoped lang="css">
.select-disabled {
  pointer-events: none;
  width: 100px;
  background-color: #f7f7f7;
}
</style>
