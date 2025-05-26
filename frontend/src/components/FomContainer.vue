<script lang="ts" setup>
import type { CustomFormData } from '@/composables/FormExam'

import FormInput from '@/components/FormInput.vue'

defineProps({
  formName: {
    type: String,
    required: false,
    default: '提交',
  },
  disabled: {
    type: Boolean,
    required: false,
    default: false,
  },
  formData: {
    // 要传递数组
    type: Array<CustomFormData>,
    required: true,
  },
})
defineEmits(['submitForm'])
</script>

<template>
  <form
    class="rounded-2rem"
    @submit.prevent="$emit('submitForm')"
  >
    <FormInput
      v-for="item in formData"
      :id="item.id"
      :key="item.id"
      v-model="item.value"
      :label="item.label"
      :type="item.type || ''"
      :autocomplete="item.autocomplete || 'off'"
    />
    <button
      :disabled="disabled"
      class="mt-5 h-10 w-full flex cursor-pointer items-center justify-center border-0 rounded-[20px] bg-[#eb6b26] text-lg text-white disabled:bg-zinc-600 hover:bg-[#ff7e3b]"
      type="submit"
    >
      {{ disabled ? '请填写完整信息' : formName }}
    </button>
    <slot />
  </form>
</template>

<style scoped lang="scss">
@import '@/assets/mixin.scss';

form {
  @include ResponseTo('phone') {
    width: 80%;
    padding: 2rem;
  }

  @include ResponseTo('desktop') {
    width: 50%;
    padding: 4rem;
    padding-top: 8rem;
  }
}
</style>
