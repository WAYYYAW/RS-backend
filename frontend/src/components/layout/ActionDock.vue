<script setup lang="ts">
import { Home, History, Gauge } from '@lucide/vue'
import { RouterLink, useRoute } from 'vue-router'

const route = useRoute()
</script>

<template>
  <aside class="action-dock">
    <RouterLink class="dock-button" :class="{ active: route.name === 'dashboard' }" title="主监控" to="/">
      <Home :size="17" />
    </RouterLink>
    <RouterLink class="dock-button" :class="{ active: route.name === 'history' }" title="历史" to="/history">
      <History :size="17" />
    </RouterLink>
    <RouterLink
      class="dock-button"
      :class="{ active: route.name === 'diagnostics' }"
      title="诊断"
      to="/diagnostics"
    >
      <Gauge :size="17" />
    </RouterLink>
  </aside>
</template>

<style scoped lang="scss">
.action-dock {
  position: absolute;
  right: 12px;
  bottom: 12px;
  display: flex;
  gap: 6px;
  padding: 6px;
  border: 1px solid var(--line-soft);
  border-radius: 8px;
  background: rgba(15, 21, 18, 0.72);
}

.dock-button {
  display: grid;
  width: 32px;
  height: 32px;
  place-items: center;
  border-radius: 5px;
  color: var(--muted);
  text-decoration: none;
}

.dock-button.active,
.dock-button:hover {
  color: var(--text);
  background: rgba(255, 255, 255, 0.1);
}
</style>
