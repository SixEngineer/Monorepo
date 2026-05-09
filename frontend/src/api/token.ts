import request from '@/utils/request'
import type { ApiResponse } from '@/types/api'
import type { Token } from '@/types/token'

// 上传/更新 Token
export function uploadToken(data: Partial<Token>): Promise<ApiResponse<null>> {
  return request.post('/token', data)
}

// 获取 Token 列表（如果后端有的话，目前没有，预留）
export function getTokenList(): Promise<ApiResponse<Token[]>> {
  return request.get('/token/list')
}

// 获取单个 Token
export function getToken(netDisk: string): Promise<ApiResponse<Token>> {
  return request.get('/token', { params: { net_disk: netDisk } })
}

// 删除 Token（如果后端有的话，预留）
export function deleteToken(id: number): Promise<ApiResponse<null>> {
  return request.delete('/token', { params: { id } })
}