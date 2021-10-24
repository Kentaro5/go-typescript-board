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
                        <option v-for="(prefecture) in prefectures"
                                :key="prefecture.id"
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
                      <select v-if="isCityDataSet" name="city_code" @change="executeChangeWard($event)">
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
                      <select v-if="isWardsDataSet" name="ward_code">
                        <option value=""></option>
                        <option v-for="(wardList) in wardLists"
                                :key="wardList.id"
                                :value="wardList.code">
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
import {defineComponent, ref} from 'vue'
import {useUser} from "../../composables/user/user";
import {useSex} from "../../composables/sex/sex";
import {usePrefectures} from "../../composables/areas/prefecture";
import {useCities} from "../../composables/areas/city";
import {useWards} from "../../composables/areas/ward";

export default defineComponent({
  setup: () => {
    const err: boolean = false
    const isCityDataSet = ref<boolean>(false)
    const isWardsDataSet = ref<boolean>(false)

    const cityLists = ref<[] | null>(null)
    const wardLists = ref<[] | null>(null)

    const { user, useUserResult } = useUser()
    const { sexes, useSexResult } = useSex()
    const { prefectures, usePrefecturesResult } = usePrefectures()

    const resetAreaData = () => {
      cityLists.value = []
      wardLists.value = []
      isCityDataSet.value = false
      isWardsDataSet.value = false
    }

    const executeChangeCity = async (event: Event) => {
      resetAreaData()
      const prefCode = parseInt(event.target.value)
      const {cities, changeCityResult} = await useCities(prefCode)
      isCityDataSet.value = changeCityResult
      cityLists.value = cities
    }

    const executeChangeWard = async (event: Event) => {
      const cityCode = parseInt(event.target.value)
      const {wards, changeWardResult} = await useWards(cityCode)
      isWardsDataSet.value = changeWardResult
      wardLists.value = wards
    }

    return {
      err,
      user,
      sexes,
      prefectures,
      cityLists,
      wardLists,
      useUserResult,
      useSexResult,
      usePrefecturesResult,
      executeChangeCity,
      executeChangeWard,
      isCityDataSet,
      isWardsDataSet,
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
