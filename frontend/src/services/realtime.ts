import { http } from './http'
import type { RealtimeResponse } from '../types/api'

export async function fetchRealtime() {
  const response = await http.get<RealtimeResponse>('/realtime')
  return response.data
}
