<script setup lang="ts">
import { computed, ref } from 'vue'
import { GridComponent, TooltipComponent } from 'echarts/components'
import { LineChart } from 'echarts/charts'
import { CanvasRenderer } from 'echarts/renderers'
import * as echarts from 'echarts/core'
import type { EChartsOption } from 'echarts'
import { useEChart } from './useEChart'

echarts.use([GridComponent, TooltipComponent, LineChart, CanvasRenderer])

const props = defineProps<{
  title: string
  unit: string
  color: string
  values: Array<{ time: string; value: number }>
}>()

const chartEl = ref<HTMLElement>()

const option = computed<EChartsOption>(() => ({
  backgroundColor: 'transparent',
  grid: { left: 38, right: 12, top: 18, bottom: 24 },
  tooltip: { trigger: 'axis' },
  xAxis: {
    type: 'category',
    data: props.values.map((point) => point.time),
    axisLabel: { show: false },
    axisTick: { show: false },
    axisLine: { lineStyle: { color: '#54665b' } },
  },
  yAxis: {
    type: 'value',
    axisLabel: { color: '#9eada4', fontSize: 10 },
    splitLine: { lineStyle: { color: 'rgba(217,226,220,0.08)' } },
  },
  series: [
    {
      type: 'line',
      data: props.values.map((point) => point.value),
      showSymbol: false,
      smooth: true,
      lineStyle: { color: props.color, width: 2 },
      areaStyle: { color: `${props.color}22` },
    },
  ],
}))

useEChart(chartEl, () => option.value)
</script>

<template>
  <section class="trend-chart">
    <header>
      <strong>{{ title }}</strong>
      <span>{{ unit }}</span>
    </header>
    <div ref="chartEl" class="chart-canvas" />
  </section>
</template>

<style scoped lang="scss">
.trend-chart {
  min-height: 0;
  border: 1px solid var(--line-soft);
  border-radius: 8px;
  background: rgba(255, 255, 255, 0.04);
}

header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 30px;
  padding: 0 10px;
  border-bottom: 1px solid var(--line-soft);
}

strong {
  font-size: 13px;
}

span {
  color: var(--muted);
  font-size: 11px;
}

.chart-canvas {
  width: 100%;
  height: calc(100% - 30px);
}
</style>
