<script setup lang="ts">
import { BrainCircuit } from '@lucide/vue'
import { useAnalyticsStore } from '../../stores/analytics'

const analytics = useAnalyticsStore()
</script>

<template>
  <article class="ai-panel">
    <div class="panel-title">
      <BrainCircuit :size="18" />
      <span>AI Analytics</span>
      <b>{{ analytics.riskLevel }}</b>
    </div>

    <div class="ai-body">
      <strong>{{ analytics.label }}</strong>
      <p>{{ analytics.suggestion }}</p>
    </div>

    <div class="ai-footer">
      <span>置信度 {{ analytics.confidence }}%</span>
      <span>更新 {{ analytics.updatedAt }}</span>
    </div>
  </article>
</template>

<style scoped lang="scss">
.ai-panel {
  display: grid;
  grid-template-rows: auto 1fr auto;
  min-height: 102px;
  padding: 12px;
  border: 1px solid rgba(230, 189, 85, 0.38);
  border-radius: 8px;
  background:
    linear-gradient(135deg, rgba(230, 189, 85, 0.13), transparent 54%),
    rgba(255, 255, 255, 0.04);
}

.panel-title {
  display: flex;
  align-items: center;
  gap: 7px;
  color: var(--amber);
  font-size: 13px;
}

.panel-title b {
  margin-left: auto;
  color: var(--muted);
  font-weight: 500;
}

.ai-body {
  align-self: center;
}

.ai-body strong {
  display: block;
  color: var(--text);
  font-size: 18px;
}

.ai-body p,
.ai-footer {
  color: var(--muted);
  font-size: 12px;
}

.ai-body p {
  margin: 6px 0 0;
}

.ai-footer {
  display: flex;
  justify-content: space-between;
}
</style>
