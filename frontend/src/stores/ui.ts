import { ref } from 'vue'
import { defineStore } from 'pinia'

export const useUiStore = defineStore('ui', () => {
  const deviceDrawerOpen = ref(false)
  const activeChart = ref<'load' | 'stroke' | 'power' | null>(null)

  function openDeviceDrawer() {
    deviceDrawerOpen.value = true
  }

  function closeDeviceDrawer() {
    deviceDrawerOpen.value = false
  }

  function toggleChart(chart: 'load' | 'stroke' | 'power') {
    activeChart.value = activeChart.value === chart ? null : chart
  }

  return {
    deviceDrawerOpen,
    activeChart,
    openDeviceDrawer,
    closeDeviceDrawer,
    toggleChart,
  }
})
