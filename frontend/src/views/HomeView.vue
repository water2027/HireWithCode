<script lang="ts" setup>
import type { CustomFormData } from '@/composables/FormExam'
import { defineAsyncComponent, reactive } from 'vue'
import { acceptChallenge } from '@/api/work/accept'
import { sendCode } from '@/api/work/send_code'
import { submitWork } from '@/api/work/submit'
import desc from '@/assets/desc.md?raw'
import { showMsg } from '@/components/MessageBox'
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
    id: 'email',
    type: 'email',
    label: 'email',
    value: '',
    reg: /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\])|(([a-z\-0-9]+\.)+[a-z]{2,}))$/i,
  },
  {
    id: 'code',
    type: 'text',
    label: '验证码',
    value: '',
  },
])
async function acceptChallengeForm() {
  useQuickForm('来吧!', challengeForm, async () => {
    const githubId = challengeForm[0].value
    const email = challengeForm[1].value
    const code = challengeForm[2].value

    if (!githubId || !email || !code) {
      showMsg('请填写完整信息')
      return
    }

    try {
      const resp = await acceptChallenge({
        github_id: githubId,
        email,
        code,
      })
      resp && showMsg(resp)
    }
    catch (error) {
      console.error(error)
      showMsg(error as string)
    }
  }, () => {
    const email = challengeForm[1].value
    sendCode({ email }).then(() => {
      showMsg('验证码已发送，请注意查收')
    }).catch((error) => {
      console.error(error)
      showMsg(error as string)
    })
  })
}
const submitForm = reactive<CustomFormData[]>([
  {
    id: 'email',
    type: 'email',
    value: '',
    label: 'email',
    reg: /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\])|(([a-z\-0-9]+\.)+[a-z]{2,}))$/i,
  },
  {
    id: 'code',
    type: 'text',
    label: '验证码',
    value: '',
  },
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
async function submitWorkForm() {
  useQuickForm('提交!', submitForm, async () => {
    const email = submitForm[0].value
    const code = submitForm[1].value
    const githubUrl = submitForm[2].value
    const onlineUrl = submitForm[3].value

    if (!email || !code || !githubUrl || !onlineUrl) {
      showMsg('请填写完整信息')
      return
    }

    try {
      const resp = await submitWork({
        email,
        code,
        github_url: githubUrl,
        online_url: onlineUrl,
      })
      resp && showMsg(resp)
    }
    catch (error) {
      console.error(error)
      showMsg(error as string)
    }
  }, () => {
    const email = submitForm[0].value
    sendCode({ email }).then(() => {
      showMsg('验证码已发送，请注意查收')
    }).catch((error) => {
      console.error(error)
      showMsg('验证码发送失败，请稍后再试')
    })
  })
}
</script>

<template>
  <div class="root mx-a mt-1 max-h-full flex flex-col overflow-auto">
    <MarkdownContainer class="p-2 shadow-md" :content="desc" />
    <div class="mb-10 mt-3 flex">
      <WaterButton @click="acceptChallengeForm">
        接受挑战
      </WaterButton>
      <WaterButton @click="submitWorkForm">
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
