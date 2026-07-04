<script setup lang="ts">
import { useDeviceStore } from '../../stores/device'

const device = useDeviceStore()
</script>

<template>
  <section class="base-info">
    <div class="field">
      <label>设备标识</label>
      <input v-model="device.address.deviceId" />
    </div>
    <div class="field">
      <label>网络地址</label>
      <input v-model="device.address.ip" />
    </div>
    <div class="field">
      <label>端口号</label>
      <input v-model="device.address.port" />
    </div>
  </section>
</template>

<style scoped lang="scss">
.base-info {
  display: grid;
  grid-template-columns: 1fr 1fr 86px;
  gap: 8px;
}
</style>
