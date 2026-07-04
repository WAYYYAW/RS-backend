<script setup lang="ts">
import AiAnalyticsPanel from '../components/overview/AiAnalyticsPanel.vue'
import DynamometerMainChart from '../components/charts/DynamometerMainChart.vue'
import HistorySummaryCard from '../components/history/HistorySummaryCard.vue'
import MetricsOverview from '../components/overview/MetricsOverview.vue'
import SecondaryChartsPanel from '../components/charts/SecondaryChartsPanel.vue'
</script>

<template>
  <div class="dashboard-page">
    <section class="top-grid">
      <MetricsOverview />
      <AiAnalyticsPanel />
    </section>

    <section class="work-area">
      <DynamometerMainChart />
      <aside class="side-column">
        <SecondaryChartsPanel />
        <HistorySummaryCard />
      </aside>
    </section>
  </div>
</template>

<style scoped lang="scss">
.dashboard-page {
  display: grid;
  grid-template-rows: 116px 1fr;
  gap: 10px;
  height: 546px;
  padding: 10px 12px 12px;
}

.top-grid {
  display: grid;
  grid-template-columns: 1fr 244px;
  gap: 10px;
  min-width: 0;
}

.work-area {
  display: grid;
  grid-template-columns: 1fr 304px;
  gap: 10px;
  min-height: 0;
}

.side-column {
  display: grid;
  grid-template-rows: 1fr 74px;
  gap: 10px;
  min-height: 0;
}
</style>
