import type { PropType } from 'vue'
import type { CustomFormData } from '@/composables/FormExam'
import { computed, createApp, defineComponent } from 'vue'

import { useFormExam } from '@/composables/FormExam'
import FomContainer from './FomContainer.vue'

const QuickForm = defineComponent({
  props: {
    formName: {
      type: String,
      required: true,
    },
    formData: {
      type: Array<CustomFormData>,
      required: true,
    },
    onSubmit: {
      type: Function as PropType<(...args: any[]) => any>,
      required: true,
    },
    onClick: {
      type: Function as PropType<(payload: MouseEvent) => void>,
      required: true,
    },
    sendCode: {
      type: Function as PropType<() => void>,
      required: false,
      default: () => () => {},
    },
  },
  setup({ formName, formData, onSubmit, onClick, sendCode }) {
    const { correct } = useFormExam(formData)
    const email = formData.find(item => item.id === 'email')
    const hasEmail = email && email.type === 'email'
    const emailIsValid = computed(() => email?.reg?.test(email.value))
    return () => (
      <div class="h-full w-full flex flex-col bg-dark bg-op-40">
        <FomContainer class="m-a flex flex-col bg-white bg-op-80" formName={formName} formData={formData} disabled={!correct.value} onSubmitForm={onSubmit}>
          {hasEmail && (
            <button
              disabled={!emailIsValid.value}
              onClick={sendCode}
              class="mt-5 h-10 w-full flex cursor-pointer items-center justify-center border-0 rounded-[20px] bg-[#eb6b26] text-lg text-white disabled:bg-zinc-600 hover:bg-[#ff7e3b]"
              type="button"
            >
              { !emailIsValid.value ? '请填写邮箱' : '发送验证码' }
            </button>
          )}
          <button class="during-300 m-a mt-3 h-5 flex items-center justify-center rounded-xl p-3 transition-all hover:bg-gray" onClick={onClick}>关闭</button>
        </FomContainer>
      </div>
    )
  },
})

export function useQuickForm(formName: string, formData: CustomFormData[], onSubmit: (...args: any[]) => any, sendCode: () => void) {
  const div = document.createElement('div')
  document.body.appendChild(div)
  div.className = 'fixed top-0 left-0 h-full w-full z-50 flex items-center justify-center'
  const app = createApp(QuickForm, {
    formName,
    formData,
    onSubmit: async (...args: any[]) => {
      await onSubmit(...args)
      app.unmount()
      document.body.removeChild(div)
    },
    onClick: () => {
      app.unmount()
      document.body.removeChild(div)
    },
    sendCode,
  })

  app.mount(div)
}
