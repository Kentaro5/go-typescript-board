<template>
  <div class="field">
    <label for="old_password" class="label">現状のパスワード</label>
    <div class="control">
      <input id="old_password" name="old_password"
             class="input" type="text"
             placeholder="半角英数字で入力してください"
             :value="oldPassword"
             @input="$emit('update:oldPassword', $event.target.value)"
      >
    </div>
  </div>
  <div class="field">
    <label for="old_confirm_password" class="label">確認用のパスワード（再入力）</label>
    <div class="control">
      <input id="old_confirm_password" name="new_password"
             class="input" type="text"
             placeholder="半角英数字で入力してください"
             :value="confirmOldPassword"
             @input="$emit('update:confirmOldPassword', $event.target.value)"
      >
    </div>
  </div>
</template>
<script lang="ts">
import {defineComponent, defineProps} from "vue";

export default defineComponent({
  props: {
    oldPassword: {
      type: String,
      required: true,
    },
    confirmOldPassword: {
      type: String,
      required: true,
    },
  },
  emits:['update:oldPassword', 'update:confirmOldPassword'],
})
</script>

