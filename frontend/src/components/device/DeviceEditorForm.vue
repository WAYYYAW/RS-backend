<script setup lang="ts">
import { reactive } from 'vue'
import { useDeviceStore } from '../../stores/device'

const device = useDeviceStore()
const draft = reactive({ ...device.params })

function apply() {
  device.updateParams({ ...draft })
}
</script>

<template>
  <form class="editor-form" @submit.prevent="apply">
    <div class="field">
      <label>原油密度</label>
      <input v-model.number="draft.oilDensity" type="number" />
    </div>
    <div class="field">
      <label>泵下入深度</label>
      <input v-model.number="draft.pumpInsertionDepth" type="number" />
    </div>
    <div class="field">
      <label>安装倾角</label>
      <input v-model.number="draft.inclination" type="number" />
    </div>
    <div class="field">
      <label>杆密度</label>
      <input v-model.number="draft.rodDensity" type="number" />
    </div>
    <div class="field">
      <label>截面积</label>
      <input v-model.number="draft.area" type="number" />
    </div>
    <div class="field">
      <label>传动比</label>
      <input v-model.number="draft.transmissionRatio" type="number" />
    </div>
    <div class="field">
      <label>冲程</label>
      <input v-model.number="draft.distance" type="number" />
    </div>
    <div class="field">
      <label>冲次</label>
      <input v-model.number="draft.strokesNumber" type="number" />
    </div>
    <div class="field">
      <label>电机转速</label>
      <input v-model.number="draft.motorSpeed" type="number" />
    </div>
    <button class="primary-button" type="submit">应用到本地</button>
  </form>
</template>

<style scoped lang="scss">
.editor-form {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 8px;
}

.primary-button {
  grid-column: 1 / -1;
}
</style>
