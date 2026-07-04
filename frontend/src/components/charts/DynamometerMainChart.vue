<script setup lang="ts">
import { computed, ref } from 'vue'
import { GridComponent, TooltipComponent } from 'echarts/components'
import { ScatterChart } from 'echarts/charts'
import { CanvasRenderer } from 'echarts/renderers'
import * as echarts from 'echarts/core'
import type { EChartsOption } from 'echarts'
import { useRealtimeStore } from '../../stores/realtime'
import { useEChart } from './useEChart'

echarts.use([GridComponent, TooltipComponent, ScatterChart, CanvasRenderer])

const chartEl = ref<HTMLElement>()
const realtime = useRealtimeStore()

const option = computed<EChartsOption>(() => ({
  backgroundColor: 'transparent',
  grid: { left: 50, right: 18, top: 28, bottom: 38 },
  tooltip: { trigger: 'item' },
  xAxis: {
    type: 'value',
    name: '冲程',
    nameTextStyle: { color: '#9eada4' },
    axisLine: { lineStyle: { color: '#54665b' } },
    axisLabel: { color: '#9eada4' },
    splitLine: { lineStyle: { color: 'rgba(217,226,220,0.08)' } },
  },
  yAxis: {
    type: 'value',
    name: '载荷',
    nameTextStyle: { color: '#9eada4' },
    axisLine: { lineStyle: { color: '#54665b' } },
    axisLabel: { color: '#9eada4' },
    splitLine: { lineStyle: { color: 'rgba(217,226,220,0.08)' } },
  },
  series: [
    {
      type: 'scatter',
      data: realtime.points.map((point) => [point.position, point.load]),
      symbolSize: 5,
      itemStyle: { color: '#4bd47b' },
    },
  ],
}))

useEChart(chartEl, () => option.value)
</script>

<template>
  <section class="chart-panel main-chart">
    <header>
      <div>
        <strong>示功图</strong>
        <span>实时曲线采样</span>
      </div>
      <b>{{ realtime.points.length }} 点</b>
    </header>
    <div ref="chartEl" class="chart-canvas" />
  </section>
</template>

<style scoped lang="scss">
.chart-panel {
  min-height: 0;
  border: 1px solid var(--line-soft);
  border-radius: 8px;
  background: rgba(255, 255, 255, 0.04);
}

header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 44px;
  padding: 0 12px;
  border-bottom: 1px solid var(--line-soft);
}

strong {
  display: block;
  font-size: 17px;
}

span,
b {
  color: var(--muted);
  font-size: 12px;
  font-weight: 500;
}

.chart-canvas {
  width: 100%;
  height: calc(100% - 44px);
}
</style>
