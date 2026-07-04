<script setup lang="ts">
import { useDeviceStore } from '../../stores/device'

const device = useDeviceStore()
</script>

<template>
  <section class="param-groups">
    <article v-for="group in device.groups" :key="group.title" class="param-group">
      <h3>{{ group.title }}</h3>
      <dl>
        <template v-for="item in group.items" :key="item.label">
          <dt>{{ item.label }}</dt>
          <dd>{{ item.value || '0' }}<span v-if="item.unit">{{ item.unit }}</span></dd>
        </template>
      </dl>
    </article>
  </section>
</template>

<style scoped lang="scss">
.param-groups {
  display: grid;
  gap: 10px;
}

.param-group {
  padding: 10px;
  border: 1px solid var(--line-soft);
  border-radius: 8px;
  background: rgba(255, 255, 255, 0.04);
}

h3 {
  margin: 0 0 8px;
  color: var(--text);
  font-size: 14px;
}

dl {
  display: grid;
  grid-template-columns: 1fr auto;
  gap: 7px 12px;
  margin: 0;
  font-size: 13px;
}

dt {
  color: var(--muted);
}

dd {
  margin: 0;
}

dd span {
  margin-left: 4px;
  color: var(--muted);
  font-size: 11px;
}
</style>
