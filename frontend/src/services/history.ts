import { http } from './http'
import type { ApiResponse, HistoryQueryResponse, HistoryRangeResponse } from '../types/api'

export async function fetchHistoryPoint(timestamp: number) {
  const response = await http.get<ApiResponse<HistoryQueryResponse>>('/history', {
    params: { timestamp },
  })
  return response.data
}

export async function fetchHistoryRange(start: number, end: number, limit = 1000) {
  const response = await http.get<HistoryRangeResponse>('/history/range', {
    params: { start, end, limit },
  })
  return response.data
}
