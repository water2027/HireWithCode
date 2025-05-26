<script lang="ts" setup>
import { defineAsyncComponent } from 'vue'
import desc from '@/assets/desc.md?raw'
import WaterButton from '@/components/WaterButton.vue'

const MarkdownContainer = defineAsyncComponent(() => import('@/components/MarkdownContainer.vue'))

async function acceptChallenge() {}
async function submitWork() {}
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
