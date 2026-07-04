<script setup lang="ts">
import AiAnalyticsPanel from '../components/overview/AiAnalyticsPanel.vue'
import DynamometerMainChart from '../components/charts/DynamometerMainChart.vue'
</script>

<template>
  <div class="diagnostics-page">
    <DynamometerMainChart />
    <AiAnalyticsPanel />
  </div>
</template>

<style scoped lang="scss">
.diagnostics-page {
  display: grid;
  grid-template-columns: 1fr 280px;
  gap: 10px;
  height: 546px;
  padding: 10px 12px 12px;
}
</style>
