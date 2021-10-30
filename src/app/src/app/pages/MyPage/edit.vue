<template>
  <section class="hero is-primary is-fullheight">
    <div class="hero-body">
      <div class="container">
        <div class="notification is-primary is-light" v-if="updated.status === 'updated'">
          <button class="delete" @click="changeUpdatedStatus('unupdated')"></button>
          <p>データを更新しました。</p>
        </div>
        <div class="columns is-centered">
          <div class="column is-5-tablet is-4-desktop is-3-widescreen">
            <div class="box" v-if="useUserResult">
              <form class="form-horizontal" @submit.prevent="updateUserInfo">
                <input type="hidden" name="_token">
                <div class="field">
                  <div class="notification is-danger" v-if="err">
                    <strong></strong>
                  </div>
                  <label for="name" class="label">ユーザー名（半角英数字で入力してください。）</label>
                  <div class="control">
                    <input id="name" name="name"
                           class="input" type="text"
                           placeholder="ユーザー名を半角英数字で入力してください。"
                           v-model="user.name"
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
                      <select name="sex_code" v-if="useSexResult" v-model="userSexCode">
                        <option value="">性別を選択してください。</option>
                        <option :value="sex.code" v-for="(sex, index) in sexes" :key="index" :selected="sex.code === user.sexCode">{{sex.name}}</option>
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
                           v-model="user.email"
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
                      <select name="pref_code" @change="executeChangeCity($event)" v-model="userPrefectureCode">
                        <option value="">都道府県を選択してください</option>
                        <option v-for="(prefecture) in prefectures"
                                :key="prefecture.id"
                                :value="prefecture.code"
                                :selected="prefecture.code === user.prefectureCode"
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
                      <select v-if="isCityDataSet" name="city_code" @change="executeChangeWard($event)" v-model="userCityCode">
                        <option value=""></option>
                        <option v-for="(cityList) in cityLists"
                                :key="cityList.id"
                                :value="cityList.code"
                                :selected="cityList.code === user.cityCode"
                        >
                          {{cityList.name}}
                        </option>
                      </select>
                      <select v-else-if="user.cityLists.length > 0" name="city_code" v-model="userCityCode">
                        <option value=""></option>
                        <option v-for="(cityList) in user.cityLists"
                                :key="cityList.id"
                                :value="cityList.code"
                                :selected="cityList.code === user.cityCode"
                        >
                          {{cityList.name}}
                        </option>
                      </select>
                      <select v-else name="ward_code" class="select-disabled">
                        <option value=""></option>
                      </select>
                    </div>
                  </div>
                  <div class="control">
                    <label for="ward_code" class="label">区</label>
                    <div class="select">
                      <select v-if="isWardsDataSet" name="ward_code" v-model="userWardCode">
                        <option value=""></option>
                        <option v-for="(wardList) in wardLists"
                                :key="wardList.id"
                                :value="wardList.code"
                                :selected="wardList.code === user.wardCode"
                        >
                          {{wardList.name}}
                        </option>
                      </select>
                      <select v-else-if="user.wardLists.length > 0" name="ward_code" v-model="userWardCode">
                        <option value=""></option>
                        <option v-for="(wardList) in user.wardLists"
                                :key="wardList.id"
                                :value="wardList.code"
                                :selected="wardList.code === user.wardCode"
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
import {defineComponent, ref} from 'vue'
import {useEditUser} from "../../composables/user/user";
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

    const prefectureValue = ref<Number | null>(null)
    const cityValue = ref<Number | null>(null)
    const wardValue = ref<Number | null>(null)

    const {
      user,
      useUserResult,
      updateUserInfo,
      changeUpdatedStatus,
      userSexCode,
      userPrefectureCode,
      userCityCode,
      userWardCode,
      updated,
    } = useEditUser()
    const { sexes, useSexResult } = useSex()
    const { prefectures, usePrefecturesResult } = usePrefectures()
    cityValue.value = user.cityCode
    wardValue.value = user.wardCode
    prefectureValue.value = user.prefectureCode

    const resetAreaData = () => {
      user.value.cityLists = []
      user.value.wardLists = []
      cityLists.value = []
      wardLists.value = []
      isCityDataSet.value = false
      isWardsDataSet.value = false
      userCityCode.value = ''
      userWardCode.value = ''
    }

    const executeChangeCity = async (event: Event) => {
      resetAreaData()
      const prefCodeText = event.target.value
      if (prefCodeText !== '') {
        const prefCode = parseInt(prefCodeText)
        const {cities, changeCityResult} = await useCities(prefCode)
        isCityDataSet.value = changeCityResult
        console.log(cities);
        cityLists.value = cities
      }
    }

    const executeChangeWard = async (event: Event) => {
      console.log(event.target.value);
      const cityCodeText = event.target.value
      if (cityCodeText !== '') {
        const cityCode = parseInt(cityCodeText)
        const {wards, changeWardResult} = await useWards(cityCode)
        isWardsDataSet.value = changeWardResult
        wardLists.value = wards
      }
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
      updateUserInfo,
      changeUpdatedStatus,
      isCityDataSet,
      isWardsDataSet,
      userSexCode,
      userPrefectureCode,
      userCityCode,
      userWardCode,
      updated,
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
