import request from '@/utils/request'
import type { ApiResponse } from '@/types/api'
import type { DownloadTask } from '@/types/download'

// 获取下载任务列表
export function getTaskList(): Promise<ApiResponse<DownloadTask[]>> {
  return request.get('/tasks')
}

// 获取单个任务详情
export function getTaskDetail(taskId: string): Promise<ApiResponse<DownloadTask>> {
  return request.get('/task', { params: { task_id: taskId } })
}

// 创建下载任务
export function createTask(data: {
  source_url: string
  provider?: string
  file_name?: string
}): Promise<ApiResponse<{ task_id: string }>> {
  return request.post('/task', data)
}

// 取消任务
export function cancelTask(taskId: string): Promise<ApiResponse<null>> {
  return request.post('/task/cancel', { task_id: taskId })
}

// 重试任务
export function retryTask(taskId: string): Promise<ApiResponse<null>> {
  return request.post('/task/retry', { task_id: taskId })
}