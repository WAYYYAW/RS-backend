export interface DeviceAddress {
  deviceId: string
  ip: string
  port: string
}

export interface DeviceParams {
  oilDensity: number
  pumpInsertionDepth: number
  inclination: number
  rodDensity: number
  area: number
  transmissionRatio: number
  distance: number
  strokesNumber: number
  motorSpeed: number
}

export interface DeviceParamGroup {
  title: string
  items: Array<{
    label: string
    value: string
    unit?: string
  }>
}
