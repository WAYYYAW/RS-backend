<script setup lang="ts">
import { useHistoryStore } from '../../stores/history'

const history = useHistoryStore()

function formatPower(position: number, load: number) {
  const value = position * load
  if (!Number.isFinite(value)) return '0'
  return Math.round(value).toString()
}
</script>

<template>
  <section class="history-table">
    <div class="table-meta">
      <strong>{{ history.mode === 'point' ? '时间点查询' : '时间段查询' }}</strong>
      <span v-if="history.meta">
        {{ history.meta.start }} - {{ history.meta.end }}，最多 {{ history.meta.limit }} 条
      </span>
      <span v-else>查询结果</span>
    </div>

    <div class="table-head">
      <span>时间</span>
      <span>载荷</span>
      <span>冲程</span>
      <span>功率估算</span>
    </div>

    <div v-if="history.loading" class="empty">加载中</div>
    <div v-else-if="history.error" class="empty">{{ history.error }}</div>
    <div v-else-if="history.records.length === 0" class="empty">暂无历史数据</div>
    <div v-else class="rows">
      <div v-for="record in history.records" :key="`${record.timestamp}-${record.id ?? 0}`" class="row">
        <span>{{ record.time }}</span>
        <span>{{ record.load }}</span>
        <span>{{ record.position }}</span>
        <span>{{ formatPower(record.position, record.load) }}</span>
      </div>
    </div>
  </section>
</template>

<style scoped lang="scss">
.history-table {
  min-height: 0;
  border: 1px solid var(--line-soft);
  border-radius: 8px;
  background: rgba(255, 255, 255, 0.04);
  overflow: hidden;
}

.table-meta {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 34px;
  padding: 0 12px;
  border-bottom: 1px solid var(--line-soft);
}

.table-meta strong {
  font-size: 13px;
}

.table-meta span {
  color: var(--muted);
  font-size: 11px;
}

.table-head,
.row {
  display: grid;
  grid-template-columns: 1.4fr 0.8fr 0.8fr 0.9fr;
  gap: 12px;
  align-items: center;
  min-height: 38px;
  padding: 0 12px;
}

.table-head {
  color: var(--muted);
  border-bottom: 1px solid var(--line-soft);
  font-size: 12px;
}

.rows {
  overflow: auto;
  max-height: 332px;
}

.row {
  border-bottom: 1px solid rgba(217, 226, 220, 0.08);
  font-size: 13px;
}

.empty {
  display: grid;
  min-height: 220px;
  place-items: center;
  color: var(--muted);
}
</style>
