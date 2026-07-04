import { onMounted, onUnmounted, ref } from 'vue'

function formatClock(date: Date) {
  const pad = (value: number) => value.toString().padStart(2, '0')
  return `${date.getFullYear()}-${pad(date.getMonth() + 1)}-${pad(date.getDate())} ${pad(
    date.getHours(),
  )}:${pad(date.getMinutes())}:${pad(date.getSeconds())}`
}

export function useClock() {
  const now = ref(formatClock(new Date()))
  let timer: number | undefined

  onMounted(() => {
    timer = window.setInterval(() => {
      now.value = formatClock(new Date())
    }, 1000)
  })

  onUnmounted(() => {
    if (timer) window.clearInterval(timer)
  })

  return { now }
}
