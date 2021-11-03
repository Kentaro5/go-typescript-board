<template>
  <section class="hero is-primary is-fullheight">
    <div class="hero-body">
      <div class="container">
        <div class="columns is-centered">
          <div class="column is-5-tablet is-4-desktop is-3-widescreen">
            <div class="box">
              <div class="notification is-primary is-light" v-if="updated.status === 'updated'">
                <button class="delete" @click="changeUpdatedStatus('unupdated')"></button>
                <p>データを更新しました。</p>
              </div>
              <form class="form-horizontal" @submit.prevent="updatePassword">
                <input type="hidden" name="_token">
                <OldPassword
                    v-model:oldPassword="oldPassword"
                    v-model:confirmOldPassword="oldConfirmPassword"
                />
                <NewPassword v-model:newPassword="newPassword"/>
                <div class="field is-grouped">
                  <div class="control">
                    <button class="button is-link">パスワード変更</button>
                  </div>
                </div>
                <div class="field is-grouped">
                  <div class="control">
                    <router-link to="/myPage/" class="button is-link">戻る</router-link>
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
import {useEditPassword} from "../../composables/user/password";
import OldPassword from "../../components/OldPassword.vue";
import NewPassword from "../../components/NewPassword.vue";

export default defineComponent({
  components: {OldPassword, NewPassword},
  setup: () => {
  const {oldPassword, oldConfirmPassword, newPassword, updated, updatePassword, changeUpdatedStatus} = useEditPassword()
    return {
      oldPassword,
      oldConfirmPassword,
      newPassword,
      updated,
      updatePassword,
      changeUpdatedStatus,
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
