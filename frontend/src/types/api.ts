import type { RealtimePayload } from './realtime'

export interface ApiResponse<T> {
  code: number
  msg: string
  data: T
}

export interface HistoryRecord {
  id?: number
  timestamp: number
  time: string
  position: number
  load: number
  distance?: number
}

export interface HistoryQueryResponse extends HistoryRecord {}

export interface HistoryRangeMeta {
  count: number
  limit: number
  start: number
  end: number
}

export interface HistoryRangeResponse {
  code: number
  msg: string
  data: HistoryRecord[]
  meta: HistoryRangeMeta
}

export type RealtimeResponse = ApiResponse<RealtimePayload>
