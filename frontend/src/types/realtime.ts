export interface RealtimePayload {
  timestamp?: number
  realtime?: string
  time?: string
  position?: number
  Position?: number
  load?: number
  Load?: number
  motorSpeed?: number
  MotorSpeed?: number
  strokesNumber?: number
  StrokesNumber?: number
  distance?: number
  Distance?: number
  rodDensity?: number
  RodDensity?: number
  transmissionRatio?: number
  TransmissionRatio?: number
  area?: number
  Area?: number
  inclination?: number
  Inclination?: number
  pumpInsertionDepth?: number
  PumpInsertionDepth?: number
  oilDensity?: number
  OilDensity?: number
}

export interface RealtimeData {
  timestamp: number
  time: string
  position: number
  load: number
  motorSpeed: number
  strokesNumber: number
  distance: number
  rodDensity: number
  transmissionRatio: number
  area: number
  inclination: number
  pumpInsertionDepth: number
  oilDensity: number
}

export interface ChartPoint {
  time: string
  position: number
  load: number
  power: number
  production: number
}
