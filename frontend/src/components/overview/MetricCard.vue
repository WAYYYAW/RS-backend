<script setup lang="ts">
import type { Component } from 'vue'

defineProps<{
  label: string
  value: string
  unit: string
  tone: 'green' | 'amber' | 'cyan' | 'blue'
  icon: Component
}>()
</script>

<template>
  <article class="metric-card" :class="tone">
    <div class="metric-head">
      <component :is="icon" :size="18" />
      <span>{{ label }}</span>
    </div>
    <div class="metric-value">
      <strong>{{ value }}</strong>
      <small>{{ unit }}</small>
    </div>
  </article>
</template>

<style scoped lang="scss">
.metric-card {
  display: grid;
  grid-template-rows: auto 1fr;
  min-width: 0;
  min-height: 102px;
  padding: 12px;
  border: 1px solid var(--line-soft);
  border-radius: 8px;
  background: linear-gradient(180deg, rgba(255, 255, 255, 0.08), rgba(255, 255, 255, 0.02));
}

.metric-card.green {
  border-left: 4px solid var(--green);
}

.metric-card.amber {
  border-left: 4px solid var(--amber);
}

.metric-card.cyan {
  border-left: 4px solid var(--cyan);
}

.metric-card.blue {
  border-left: 4px solid var(--blue);
}

.metric-head {
  display: flex;
  align-items: center;
  gap: 7px;
  color: var(--muted);
  font-size: 13px;
}

.metric-value {
  display: flex;
  align-items: baseline;
  justify-content: space-between;
  gap: 8px;
}

.metric-value strong {
  overflow: hidden;
  color: var(--text);
  font-size: 32px;
  line-height: 1;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.metric-value small {
  color: var(--muted);
  font-size: 15px;
}
</style>
