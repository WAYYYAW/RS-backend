<script setup lang="ts">
import { Database } from '@lucide/vue'
import { useHistoryStore } from '../../stores/history'

const history = useHistoryStore()
</script>

<template>
  <section class="history-summary">
    <Database :size="18" />
    <div>
      <strong>历史数据</strong>
      <span>{{ history.refreshedAt ? `刷新 ${history.refreshedAt}` : '请选择时间查询' }}</span>
    </div>
    <small>{{ history.records.length }} 条</small>
  </section>
</template>

<style scoped lang="scss">
.history-summary {
  display: grid;
  grid-template-columns: auto 1fr auto;
  align-items: center;
  gap: 10px;
  padding: 12px;
  border: 1px solid var(--line-soft);
  border-radius: 8px;
  background: rgba(255, 255, 255, 0.04);
}

strong,
span {
  display: block;
}

span {
  color: var(--muted);
  font-size: 12px;
}

small {
  color: var(--cyan);
  font-size: 18px;
}
</style>
