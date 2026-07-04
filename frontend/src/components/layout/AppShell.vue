<script setup lang="ts">
import { computed } from 'vue'
import { useFixedScreenScale } from '../../composables/useFixedScreenScale'
import { useRealtimeSocket } from '../../composables/useRealtimeSocket'
import TopStatusBar from './TopStatusBar.vue'
import ActionDock from './ActionDock.vue'
import DeviceDrawer from '../device/DeviceDrawer.vue'

const { scale } = useFixedScreenScale()
useRealtimeSocket()

const shellStyle = computed(() => ({
  transform: `translate(-50%, -50%) scale(${scale.value})`,
}))
</script>

<template>
  <div class="viewport">
    <main class="app-shell" :style="shellStyle">
      <TopStatusBar />
      <section class="screen-page">
        <slot />
      </section>
      <ActionDock />
      <DeviceDrawer />
    </main>
  </div>
</template>

<style scoped lang="scss">
.viewport {
  position: fixed;
  inset: 0;
  overflow: hidden;
}

.app-shell {
  position: absolute;
  top: 50%;
  left: 50%;
  display: grid;
  grid-template-rows: 54px 1fr;
  width: 1024px;
  height: 600px;
  transform-origin: center;
  border: 1px solid rgba(217, 226, 220, 0.22);
  background:
    linear-gradient(135deg, rgba(255, 255, 255, 0.06), transparent 28%),
    #17211c;
  box-shadow: 0 24px 80px rgba(0, 0, 0, 0.35);
}

.screen-page {
  min-height: 0;
}
</style>
