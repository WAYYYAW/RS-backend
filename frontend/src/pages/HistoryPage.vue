<script setup lang="ts">
import HistoryQueryPanel from '../components/history/HistoryQueryPanel.vue'
import HistorySummaryCard from '../components/history/HistorySummaryCard.vue'
import HistoryTable from '../components/history/HistoryTable.vue'
</script>

<template>
  <div class="history-page">
    <section class="history-head">
      <HistorySummaryCard />
      <HistoryQueryPanel />
    </section>
    <HistoryTable />
  </div>
</template>

<style scoped lang="scss">
.history-page {
  display: grid;
  grid-template-rows: 112px 1fr;
  gap: 10px;
  height: 546px;
  padding: 10px 12px 12px;
}

.history-head {
  display: grid;
  grid-template-columns: 250px 1fr;
  gap: 10px;
  min-height: 0;
}
</style>
