import request from '@/utils/request'
import type { ApiResponse } from '@/types/api'
import type { QuotaInfo } from '@/types/quota'

// 查询配额
export function queryQuota(provider: string): Promise<ApiResponse<QuotaInfo>> {
  return request.post('/quota/query', { name: provider })  // ← 改成 name
}

// 同步配额
export function syncQuota(provider: string): Promise<ApiResponse<QuotaInfo>> {
  return request.post('/quota/sync', { name: provider })   // ← 改成 name
}