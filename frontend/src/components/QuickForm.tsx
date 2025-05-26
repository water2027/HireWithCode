import type { PropType } from 'vue'
import type { CustomFormData } from '@/composables/FormExam'
import { createApp, defineComponent } from 'vue'

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
  },
  setup({ formName, formData, onSubmit, onClick }) {
    const { correct } = useFormExam(formData)
    return () => (
      <div class="h-full w-full flex flex-col bg-dark bg-op-40">
        <FomContainer class="m-a flex flex-col bg-white bg-op-80" formName={formName} formData={formData} disabled={!correct.value} onSubmitForm={onSubmit}>
          <button class="m-a mt-3 h-5 flex items-center justify-center rounded-xl p-3" onClick={onClick}>关闭</button>
        </FomContainer>
      </div>
    )
  },
})

export function useQuickForm(formName: string, formData: CustomFormData[], onSubmit: (...args: any[]) => any) {
  const div = document.createElement('div')
  document.body.appendChild(div)
  div.className = 'fixed top-0 left-0 h-full w-full z-50 flex items-center justify-center'
  const app = createApp(QuickForm, {
    formName,
    formData,
    onSubmit: (...args: any[]) => {
      onSubmit(...args)
      app.unmount()
      document.body.removeChild(div)
    },
    onClick: () => {
      app.unmount()
      document.body.removeChild(div)
    },
  })

  app.mount(div)
}
