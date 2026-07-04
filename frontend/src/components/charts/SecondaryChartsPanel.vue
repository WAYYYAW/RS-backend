<script setup lang="ts">
import { computed } from 'vue'
import { useRealtimeStore } from '../../stores/realtime'
import TrendMiniChart from './TrendMiniChart.vue'

const realtime = useRealtimeStore()

const loadValues = computed(() => realtime.points.map((point) => ({ time: point.time, value: point.load })))
const strokeValues = computed(() => realtime.points.map((point) => ({ time: point.time, value: point.position })))
</script>

<template>
  <div class="secondary-charts">
    <TrendMiniChart color="#54c7d6" title="载荷趋势" unit="N" :values="loadValues" />
    <TrendMiniChart color="#e6bd55" title="冲程趋势" unit="m" :values="strokeValues" />
  </div>
</template>

<style scoped lang="scss">
.secondary-charts {
  display: grid;
  grid-template-rows: 1fr 1fr;
  gap: 10px;
  min-height: 0;
}
</style>
