<script setup lang="ts">
import { X } from '@lucide/vue'
import { ref } from 'vue'
import { useUiStore } from '../../stores/ui'
import DeviceBaseInfo from './DeviceBaseInfo.vue'
import DeviceEditorForm from './DeviceEditorForm.vue'
import DeviceParamGroups from './DeviceParamGroups.vue'

const ui = useUiStore()
const mode = ref<'view' | 'edit'>('view')
</script>

<template>
  <Transition name="drawer">
    <aside v-if="ui.deviceDrawerOpen" class="drawer">
      <header>
        <div>
          <strong>设备配置</strong>
          <span>本地显示参数</span>
        </div>
        <button class="icon-button" title="关闭" type="button" @click="ui.closeDeviceDrawer">
          <X :size="18" />
        </button>
      </header>

      <DeviceBaseInfo />

      <div class="mode-switch">
        <button :class="{ active: mode === 'view' }" type="button" @click="mode = 'view'">查看</button>
        <button :class="{ active: mode === 'edit' }" type="button" @click="mode = 'edit'">编辑</button>
      </div>

      <DeviceParamGroups v-if="mode === 'view'" />
      <DeviceEditorForm v-else />
    </aside>
  </Transition>
</template>

<style scoped lang="scss">
.drawer {
  position: absolute;
  top: 54px;
  right: 0;
  z-index: 20;
  display: grid;
  grid-template-rows: auto auto auto 1fr;
  gap: 12px;
  width: 360px;
  height: calc(100% - 54px);
  padding: 14px;
  border-left: 1px solid var(--line);
  background: rgba(24, 33, 29, 0.96);
  box-shadow: -18px 0 42px rgba(0, 0, 0, 0.24);
}

header {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

header strong {
  display: block;
  font-size: 18px;
}

header span {
  color: var(--muted);
  font-size: 12px;
}

.mode-switch {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 6px;
}

.mode-switch button {
  height: 32px;
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

.drawer-enter-active,
.drawer-leave-active {
  transition:
    transform 0.18s ease,
    opacity 0.18s ease;
}

.drawer-enter-from,
.drawer-leave-to {
  opacity: 0;
  transform: translateX(24px);
}
</style>
