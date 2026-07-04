import { computed, ref } from 'vue'
import { defineStore } from 'pinia'

export type AnalyticsStatus = 'idle' | 'pending' | 'ready' | 'error'

export const useAnalyticsStore = defineStore('analytics', () => {
  const status = ref<AnalyticsStatus>('idle')
  const label = ref('待接入本地分析模块')
  const confidence = ref(0)
  const riskLevel = ref('未分析')
  const suggestion = ref('预留示功图分析结果位置')
  const updatedAt = ref('--')

  const isReady = computed(() => status.value === 'ready')

  return {
    status,
    label,
    confidence,
    riskLevel,
    suggestion,
    updatedAt,
    isReady,
  }
})
