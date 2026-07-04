import { computed, reactive } from 'vue'
import { defineStore } from 'pinia'
import type { DeviceAddress, DeviceParams } from '../types/device'

export const useDeviceStore = defineStore('device', () => {
  const address = reactive<DeviceAddress>({
    deviceId: 'RS-PLC-001',
    ip: '127.0.0.1',
    port: '5020',
  })

  const params = reactive<DeviceParams>({
    oilDensity: 0,
    pumpInsertionDepth: 0,
    inclination: 0,
    rodDensity: 0,
    area: 0,
    transmissionRatio: 0,
    distance: 0,
    strokesNumber: 0,
    motorSpeed: 0,
  })

  const groups = computed(() => [
    {
      title: '结构参数',
      items: [
        { label: '泵下入深度', value: params.pumpInsertionDepth.toString(), unit: 'm' },
        { label: '安装倾角', value: params.inclination.toString(), unit: 'deg' },
        { label: '截面积', value: params.area.toString(), unit: 'm2' },
      ],
    },
    {
      title: '流体参数',
      items: [
        { label: '原油密度', value: params.oilDensity.toString(), unit: 'kg/m3' },
        { label: '杆密度', value: params.rodDensity.toString(), unit: 'kg/m3' },
      ],
    },
    {
      title: '运行参数',
      items: [
        { label: '传动比', value: params.transmissionRatio.toString() },
        { label: '冲程', value: params.distance.toString(), unit: 'm' },
        { label: '冲次', value: params.strokesNumber.toString(), unit: 'min-1' },
        { label: '电机转速', value: params.motorSpeed.toString(), unit: 'rpm' },
      ],
    },
  ])

  function updateParams(next: Partial<DeviceParams>) {
    Object.assign(params, next)
  }

  return { address, params, groups, updateParams }
})
