import { onMounted, onUnmounted } from 'vue'
import { useRealtimeStore } from '../stores/realtime'
import { createRealtimeSocket } from '../services/websocket'

export function useRealtimeSocket() {
  const realtime = useRealtimeStore()
  let dispose: (() => void) | undefined

  onMounted(() => {
    realtime.loadSnapshot()
    dispose = createRealtimeSocket(
      (payload) => realtime.applyPayload(payload),
      (connected) => realtime.setConnected(connected),
    )
  })

  onUnmounted(() => {
    dispose?.()
  })
}
