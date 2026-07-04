import { onMounted, onUnmounted, ref } from 'vue'

export function useFixedScreenScale(width = 1024, height = 600) {
  const scale = ref(1)

  const updateScale = () => {
    scale.value = Math.min(window.innerWidth / width, window.innerHeight / height)
  }

  onMounted(() => {
    updateScale()
    window.addEventListener('resize', updateScale)
  })

  onUnmounted(() => {
    window.removeEventListener('resize', updateScale)
  })

  return { scale }
}
