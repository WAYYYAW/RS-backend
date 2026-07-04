<script setup lang="ts">
import { Activity, BarChart3, Clock3, Database, Settings } from '@lucide/vue'
import { RouterLink } from 'vue-router'
import { useClock } from '../../composables/useClock'
import { useRealtimeStore } from '../../stores/realtime'
import { useUiStore } from '../../stores/ui'

const { now } = useClock()
const realtime = useRealtimeStore()
const ui = useUiStore()
</script>

<template>
  <header class="top-bar">
    <div class="brand">
      <Activity :size="22" />
      <div>
        <strong>采油机智能管理系统</strong>
        <span>RS BACKEND LOCAL HMI</span>
      </div>
    </div>

    <div class="status" :class="{ offline: !realtime.connected }">
      <span class="status-dot" />
      <strong>{{ realtime.connected ? '连接正常' : '连接异常' }}</strong>
      <small>最后更新 {{ realtime.lastUpdated }}</small>
    </div>

    <div class="clock">
      <Clock3 :size="17" />
      <span>{{ now }}</span>
    </div>

    <nav class="nav-actions">
      <button class="icon-button" title="设备配置" type="button" @click="ui.openDeviceDrawer">
        <Settings :size="18" />
      </button>
      <RouterLink class="icon-button" title="历史数据" to="/history">
        <Database :size="18" />
      </RouterLink>
      <RouterLink class="icon-button" title="诊断页" to="/diagnostics">
        <BarChart3 :size="18" />
      </RouterLink>
    </nav>
  </header>
</template>

<style scoped lang="scss">
.top-bar {
  display: grid;
  grid-template-columns: 282px 244px 1fr 136px;
  align-items: center;
  gap: 12px;
  height: 54px;
  padding: 8px 12px;
  border-bottom: 1px solid var(--line-soft);
  background: rgba(15, 21, 18, 0.72);
}

.brand {
  display: flex;
  align-items: center;
  gap: 10px;
}

.brand strong {
  display: block;
  font-size: 17px;
  line-height: 18px;
}

.brand span {
  display: block;
  margin-top: 2px;
  color: var(--muted);
  font-size: 10px;
}

.status {
  display: grid;
  grid-template-columns: 12px auto;
  column-gap: 8px;
  align-items: center;
  padding: 6px 10px;
  border: 1px solid rgba(75, 212, 123, 0.35);
  border-radius: 6px;
  background: rgba(75, 212, 123, 0.09);
}

.status.offline {
  border-color: rgba(227, 101, 91, 0.45);
  background: rgba(227, 101, 91, 0.1);
}

.status-dot {
  width: 10px;
  height: 10px;
  border-radius: 50%;
  background: var(--green);
  box-shadow: 0 0 14px var(--green);
}

.offline .status-dot {
  background: var(--red);
  box-shadow: 0 0 14px var(--red);
}

.status strong {
  font-size: 13px;
}

.status small {
  grid-column: 2;
  color: var(--muted);
  font-size: 10px;
}

.clock {
  display: flex;
  align-items: center;
  justify-content: flex-end;
  gap: 7px;
  color: var(--muted);
  font-size: 14px;
}

.nav-actions {
  display: flex;
  justify-content: flex-end;
  gap: 7px;
}

.nav-actions a {
  text-decoration: none;
}
</style>
