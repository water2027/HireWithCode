<script lang="ts" setup>
import { onMounted, ref } from 'vue'
import WelcomeView from '@/views/WelcomeView.vue'

const show = ref(true)

onMounted(() => {
  setTimeout(() => {
    show.value = false
  }, 2000)
})
</script>

<template>
  <template v-if="show">
    <Transition name="fade" appear>
      <WelcomeView />
    </Transition>
  </template>
  <template v-else>
    <RouterView v-slot="{ Component }">
      <Transition name="fade" mode="out-in">
        <component :is="Component" />
      </Transition>
    </RouterView>
  </template>
</template>

<style scoped lang="scss">
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
