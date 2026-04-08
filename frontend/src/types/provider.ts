import type { HealthState } from './common'

export interface ProviderRecord {
  id: string
  name: string
  status: HealthState
  capability: string
  authMode: string
}
