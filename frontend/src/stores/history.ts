import { ref } from 'vue'
import { defineStore } from 'pinia'
import { fetchHistoryPoint, fetchHistoryRange } from '../services/history'
import type { HistoryRangeMeta, HistoryRecord } from '../types/api'

export type HistoryQueryMode = 'point' | 'range'

export const useHistoryStore = defineStore('history', () => {
  const records = ref<HistoryRecord[]>([])
  const loading = ref(false)
  const error = ref('')
  const refreshedAt = ref('')
  const mode = ref<HistoryQueryMode>('point')
  const meta = ref<HistoryRangeMeta | null>(null)

  async function queryPoint(timestamp: number) {
    loading.value = true
    error.value = ''
    mode.value = 'point'
    meta.value = null

    try {
      const response = await fetchHistoryPoint(timestamp)
      if (response.code === 0) {
        records.value = [response.data]
        refreshedAt.value = new Date().toLocaleString()
      } else {
        records.value = []
        error.value = response.msg
      }
    } catch (err) {
      records.value = []
      error.value = err instanceof Error ? err.message : '历史数据查询失败'
    } finally {
      loading.value = false
    }
  }

  async function queryRange(start: number, end: number, limit = 1000) {
    loading.value = true
    error.value = ''
    mode.value = 'range'

    try {
      const response = await fetchHistoryRange(start, end, limit)
      if (response.code === 0) {
        records.value = response.data
        meta.value = response.meta
        refreshedAt.value = new Date().toLocaleString()
      } else {
        records.value = []
        meta.value = null
        error.value = response.msg
      }
    } catch (err) {
      records.value = []
      meta.value = null
      error.value = err instanceof Error ? err.message : '历史数据查询失败'
    } finally {
      loading.value = false
    }
  }

  return { records, loading, error, refreshedAt, mode, meta, queryPoint, queryRange }
})
