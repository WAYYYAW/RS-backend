<script setup lang="ts">
import { computed } from 'vue'
import { Activity, Droplets, Gauge, Zap } from '@lucide/vue'
import { useRealtimeStore } from '../../stores/realtime'
import MetricCard from './MetricCard.vue'

const realtime = useRealtimeStore()

function formatNumber(value: number) {
  if (!Number.isFinite(value)) return '0'
  if (Math.abs(value) >= 10000) return value.toExponential(2)
  return value.toFixed(value % 1 === 0 ? 0 : 2)
}

const metrics = computed(() => [
  {
    label: '载荷',
    value: formatNumber(realtime.current.load),
    unit: 'N',
    tone: 'green' as const,
    icon: Gauge,
  },
  {
    label: '冲程',
    value: formatNumber(realtime.current.position),
    unit: 'm',
    tone: 'cyan' as const,
    icon: Activity,
  },
  {
    label: '光杆功率',
    value: formatNumber(realtime.power),
    unit: 'W',
    tone: 'amber' as const,
    icon: Zap,
  },
  {
    label: '产液量',
    value: formatNumber(realtime.production),
    unit: 'L',
    tone: 'blue' as const,
    icon: Droplets,
  },
])
</script>

<template>
  <section class="metrics-overview">
    <MetricCard
      v-for="metric in metrics"
      :key="metric.label"
      :icon="metric.icon"
      :label="metric.label"
      :tone="metric.tone"
      :unit="metric.unit"
      :value="metric.value"
    />
  </section>
</template>

<style scoped lang="scss">
.metrics-overview {
  display: grid;
  grid-template-columns: repeat(4, minmax(0, 1fr));
  gap: 10px;
}
</style>
