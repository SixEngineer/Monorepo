export interface ApiResponse<T> {
  code: number
  message: string
  data: T
}

export interface NavItem {
  label: string
  path: string
  description: string
}

export type HealthState = 'healthy' | 'warning' | 'offline'
