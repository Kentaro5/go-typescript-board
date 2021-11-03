<template>
  <div class="field">
    <label for="new_password" class="label">新しいパスワード</label>
    <div class="control">
      <input id="new_password" name="new_password"
             class="input" type="text"
             placeholder="半角英数字で入力してください"
             :value="newPassword"
             @input="$emit('update:newPassword', $event.target.value)"
      >
    </div>
  </div>
</template>

<script lang="ts">
import {defineComponent, defineProps} from "vue"

export default defineComponent({
  props: {
    newPassword: {
      type: String,
      required: true,
    },
  },
  emits:['update:newPassword']
})
</script>

