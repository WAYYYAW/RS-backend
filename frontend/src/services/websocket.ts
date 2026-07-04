type MessageHandler = (payload: unknown) => void
type StatusHandler = (connected: boolean) => void

export function createRealtimeSocket(onMessage: MessageHandler, onStatus: StatusHandler) {
  const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:'
  const url = `${protocol}//${window.location.host}/ws`
  let socket: WebSocket | null = null
  let retryTimer: number | undefined
  let closedByClient = false

  const connect = () => {
    socket = new WebSocket(url)

    socket.onopen = () => {
      onStatus(true)
    }

    socket.onmessage = (event) => {
      try {
        onMessage(JSON.parse(event.data))
      } catch {
        // Ignore malformed realtime frames. The next valid frame will recover the UI.
      }
    }

    socket.onerror = () => {
      onStatus(false)
    }

    socket.onclose = () => {
      onStatus(false)
      if (!closedByClient) {
        retryTimer = window.setTimeout(connect, 2000)
      }
    }
  }

  connect()

  return () => {
    closedByClient = true
    if (retryTimer) window.clearTimeout(retryTimer)
    socket?.close()
  }
}
