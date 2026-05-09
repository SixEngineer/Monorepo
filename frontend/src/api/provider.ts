import request from '@/utils/request'
import type { ApiResponse } from '@/types/api'
import type { ProviderRecord } from '@/types/provider'

// 获取 Provider 列表
export function getProviderList(): Promise<ApiResponse<ProviderRecord[]>> {
  return request.get('/provider/list')
}

// 注册 Provider
export function registerProvider(data: any): Promise<ApiResponse<any>> {
  return request.post('/provider', data)
}

// 删除 Provider
export function deleteProvider(id: number): Promise<ApiResponse<null>> {
  return request.delete('/provider', { params: { id } })
}

// 更新 Provider
export function updateProvider(data: Partial<ProviderRecord> & { id: number }): Promise<ApiResponse<null>> {
  return request.put('/provider/', data) 
}

// 获取单个 Provider 信息
export function getProviderInfo(id: string): Promise<ApiResponse<ProviderRecord>> {
  return request.get('/provider/info', { params: { id } })
}