import { computed, ref } from 'vue'
import { defineStore } from 'pinia'
import { fetchRealtime } from '../services/realtime'
import type { ChartPoint, RealtimeData, RealtimePayload } from '../types/realtime'

const emptyRealtime: RealtimeData = {
  timestamp: 0,
  time: '--',
  position: 0,
  load: 0,
  motorSpeed: 0,
  strokesNumber: 0,
  distance: 0,
  rodDensity: 0,
  transmissionRatio: 0,
  area: 0,
  inclination: 0,
  pumpInsertionDepth: 0,
  oilDensity: 0,
}

function numeric(value: unknown) {
  const parsed = Number(value)
  return Number.isFinite(parsed) ? parsed : 0
}

function normalizePayload(payload: unknown): RealtimeData {
  const envelope = payload as { data?: RealtimePayload }
  const data = envelope.data ?? (payload as RealtimePayload)

  return {
    timestamp: numeric(data.timestamp) || Math.floor(Date.now() / 1000),
    time: data.time ?? data.realtime ?? new Date().toISOString(),
    position: numeric(data.position ?? data.Position),
    load: numeric(data.load ?? data.Load),
    motorSpeed: numeric(data.motorSpeed ?? data.MotorSpeed),
    strokesNumber: numeric(data.strokesNumber ?? data.StrokesNumber),
    distance: numeric(data.distance ?? data.Distance),
    rodDensity: numeric(data.rodDensity ?? data.RodDensity),
    transmissionRatio: numeric(data.transmissionRatio ?? data.TransmissionRatio),
    area: numeric(data.area ?? data.Area),
    inclination: numeric(data.inclination ?? data.Inclination),
    pumpInsertionDepth: numeric(data.pumpInsertionDepth ?? data.PumpInsertionDepth),
    oilDensity: numeric(data.oilDensity ?? data.OilDensity),
  }
}

export const useRealtimeStore = defineStore('realtime', () => {
  const connected = ref(false)
  const current = ref<RealtimeData>({ ...emptyRealtime })
  const points = ref<ChartPoint[]>([])
  const lastError = ref('')

  const power = computed(() => current.value.position * current.value.load)
  const production = computed(() => 0)
  const lastUpdated = computed(() => current.value.time)

  function setConnected(value: boolean) {
    connected.value = value
  }

  function applyPayload(payload: unknown) {
    const next = normalizePayload(payload)
    current.value = next
    lastError.value = ''

    points.value.push({
      time: next.time,
      position: next.position,
      load: next.load,
      power: next.position * next.load,
      production: production.value,
    })

    if (points.value.length > 160) {
      points.value.splice(0, points.value.length - 160)
    }
  }

  async function loadSnapshot() {
    try {
      const response = await fetchRealtime()
      if (response.code === 0) {
        applyPayload(response.data)
      }
    } catch (error) {
      lastError.value = error instanceof Error ? error.message : '实时数据拉取失败'
    }
  }

  return {
    connected,
    current,
    points,
    power,
    production,
    lastUpdated,
    lastError,
    setConnected,
    applyPayload,
    loadSnapshot,
  }
})
