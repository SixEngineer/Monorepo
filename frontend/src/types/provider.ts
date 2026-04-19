import type { HealthState } from './common'

export interface ProviderRecord {
  id: number                      // 数字类型
  name: string
  provider_type: string           // 服务商类型（如 aliyun、baidu）
  net_disk: string                // 网盘类型
  account_id: string              // 账户ID
  status: HealthState             // 状态
  access_token?: string
  refresh_token?: string
  token_expires_at?: string
  total_quota: number             // 总配额（字节）
  used_quota: number              // 已用配额
  available_quota: number         // 可用配额
  last_quota_sync_at?: string
  last_error?: string
  created_at: string
  updated_at: string
}