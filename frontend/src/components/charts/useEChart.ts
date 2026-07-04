import type { ECharts, EChartsOption } from 'echarts'
import * as echarts from 'echarts/core'
import { onMounted, onUnmounted, shallowRef, watch, type Ref } from 'vue'

export function useEChart(el: Ref<HTMLElement | undefined>, option: () => EChartsOption) {
  const chart = shallowRef<ECharts>()
  let resizeObserver: ResizeObserver | undefined

  onMounted(() => {
    if (!el.value) return
    chart.value = echarts.init(el.value)
    chart.value.setOption(option())

    resizeObserver = new ResizeObserver(() => chart.value?.resize())
    resizeObserver.observe(el.value)
  })

  onUnmounted(() => {
    resizeObserver?.disconnect()
    chart.value?.dispose()
  })

  watch(option, (next) => chart.value?.setOption(next, true), { deep: true })

  return chart
}
