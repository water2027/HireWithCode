<script lang="ts" setup>
import type { CustomFormData } from '@/composables/FormExam'
import { defineAsyncComponent, reactive } from 'vue'
import desc from '@/assets/desc.md?raw'
import { useQuickForm } from '@/components/QuickForm'
import WaterButton from '@/components/WaterButton.vue'

const MarkdownContainer = defineAsyncComponent(() => import('@/components/MarkdownContainer.vue'))

const challengeForm = reactive<CustomFormData[]>([
  {
    id: 'Github ID',
    type: 'text',
    label: 'GitHub ID',
    value: '',
  },
  {
    id: 'Email',
    type: 'email',
    label: 'email',
    value: '',
    reg: /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\])|(([a-z\-0-9]+\.)+[a-z]{2,}))$/i,
  },
])
const submitForm = reactive<CustomFormData[]>([
  {
    id: 'repo url',
    type: 'text',
    value: '',
    label: 'repo url',
  },
  {
    id: '在线体验地址',
    type: 'text',
    value: '',
    label: '在线体验地址',
  },
])
async function acceptChallenge() {
  useQuickForm('来吧!', challengeForm, () => {
  })
}
async function submitWork() {
  useQuickForm('提交!', submitForm, () => {})
}
</script>

<template>
  <div class="root mx-a mt-1 max-h-full flex flex-col overflow-auto">
    <MarkdownContainer class="p-2 shadow-md" :content="desc" />
    <div class="mt-3 flex">
      <WaterButton @click="acceptChallenge">
        接受挑战
      </WaterButton>
      <WaterButton @click="submitWork">
        提交作品
      </WaterButton>
    </div>
  </div>
</template>

<style scoped lang="scss">
@import '@/assets/mixin.scss';

.root {
  @include ResponseTo('desktop') {
    width: 60%;
  }

  @include ResponseTo('phone') {
    width: 90%;
  }
}
</style>
