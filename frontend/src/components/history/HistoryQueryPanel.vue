<script setup lang="ts">
import { computed, ref } from 'vue'
import { Search } from '@lucide/vue'
import { useHistoryStore } from '../../stores/history'

const history = useHistoryStore()
const queryMode = ref<'point' | 'range'>('point')
const pointTime = ref(toDatetimeLocal(new Date()))
const startTime = ref(toDatetimeLocal(new Date(Date.now() - 60 * 60 * 1000)))
const endTime = ref(toDatetimeLocal(new Date()))
const limit = ref(500)
const validation = ref('')

const canQuery = computed(() => !history.loading)

function toDatetimeLocal(date: Date) {
  const pad = (value: number) => value.toString().padStart(2, '0')
  return `${date.getFullYear()}-${pad(date.getMonth() + 1)}-${pad(date.getDate())}T${pad(
    date.getHours(),
  )}:${pad(date.getMinutes())}:${pad(date.getSeconds())}`
}

function parseDatetimeLocal(value: string) {
  const date = new Date(value)
  return Number.isNaN(date.getTime()) ? null : Math.floor(date.getTime() / 1000)
}

function submitQuery() {
  validation.value = ''

  if (queryMode.value === 'point') {
    const timestamp = parseDatetimeLocal(pointTime.value)
    if (timestamp === null) {
      validation.value = '请选择有效的查询时间'
      return
    }
    history.queryPoint(timestamp)
    return
  }

  const start = parseDatetimeLocal(startTime.value)
  const end = parseDatetimeLocal(endTime.value)
  if (start === null || end === null) {
    validation.value = '请选择有效的开始和结束时间'
    return
  }
  if (end < start) {
    validation.value = '结束时间不能早于开始时间'
    return
  }
  history.queryRange(start, end, limit.value)
}
</script>

<template>
  <section class="query-panel">
    <div class="mode-switch">
      <button :class="{ active: queryMode === 'point' }" type="button" @click="queryMode = 'point'">
        时间点
      </button>
      <button :class="{ active: queryMode === 'range' }" type="button" @click="queryMode = 'range'">
        时间段
      </button>
    </div>

    <form :class="['query-form', queryMode]" @submit.prevent="submitQuery">
      <template v-if="queryMode === 'point'">
        <div class="field">
          <label>查询时间</label>
          <input v-model="pointTime" step="1" type="datetime-local" />
        </div>
      </template>

      <template v-else>
        <div class="field">
          <label>开始时间</label>
          <input v-model="startTime" step="1" type="datetime-local" />
        </div>
        <div class="field">
          <label>结束时间</label>
          <input v-model="endTime" step="1" type="datetime-local" />
        </div>
        <div class="field limit-field">
          <label>最大条数</label>
          <input v-model.number="limit" max="1000" min="1" type="number" />
        </div>
      </template>

      <button class="primary-button" :disabled="!canQuery" type="submit">
        <Search :size="16" />
        <span>{{ history.loading ? '查询中' : '查询' }}</span>
      </button>
    </form>

    <p v-if="validation" class="validation">{{ validation }}</p>
  </section>
</template>

<style scoped lang="scss">
.query-panel {
  display: grid;
  grid-template-columns: 148px 1fr;
  gap: 10px;
  min-height: 86px;
  padding: 12px;
  border: 1px solid var(--line-soft);
  border-radius: 8px;
  background: rgba(255, 255, 255, 0.04);
}

.mode-switch {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 6px;
  align-self: start;
}

.mode-switch button {
  height: 34px;
  border: 1px solid var(--line);
  border-radius: 5px;
  color: var(--muted);
  background: rgba(255, 255, 255, 0.04);
  cursor: pointer;
}

.mode-switch button.active {
  color: var(--text);
  border-color: rgba(84, 199, 214, 0.55);
  background: rgba(84, 199, 214, 0.14);
}

.query-form {
  display: grid;
  align-items: end;
  gap: 8px;
  min-width: 0;
}

.query-form.point {
  grid-template-columns: minmax(0, 1fr) 92px;
}

.query-form.range {
  grid-template-columns: minmax(0, 1fr) minmax(0, 1fr) 96px 92px;
}

.query-form :deep(.field) {
  min-width: 0;
}

.query-form :deep(.field input) {
  min-width: 0;
}

.limit-field {
  min-width: 0;
}

.primary-button {
  display: inline-flex;
  align-items: center;
  gap: 7px;
  justify-content: center;
  min-width: 0;
  min-width: 92px;
  white-space: nowrap;
}

.primary-button:disabled {
  cursor: wait;
  opacity: 0.62;
}

.validation {
  grid-column: 2;
  margin: -4px 0 0;
  color: var(--amber);
  font-size: 12px;
}
</style>
