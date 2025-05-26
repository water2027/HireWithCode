<script lang="ts" setup>
import { defineAsyncComponent } from 'vue'
import desc from '@/assets/desc.md?raw'

const MarkdownContainer = defineAsyncComponent(() => import('@/components/MarkdownContainer.vue'))
</script>

<template>
  <div class="root mx-a mt-1 max-h-full flex flex-col overflow-auto">
    <MarkdownContainer class="p-2 shadow-md" :content="desc" />
    <div class="mt-3 flex">
      <button class="m-a border border-2 border-b-blue rounded-xl bg-transparent p-2 transition-all duration-200 hover:bg-blueGray">
        接受挑战
      </button>
      <button class="m-a border border-2 border-b-blue rounded-xl bg-transparent p-2 transition-all duration-200 hover:bg-blueGray">
        提交作品
      </button>
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
