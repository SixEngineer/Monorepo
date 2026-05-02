import { ref } from 'vue'
import { defineStore } from 'pinia'
import { getProviderList, deleteProvider } from '@/api/provider'
import type { ProviderRecord } from '@/types/provider'

import { queryQuota, syncQuota } from '@/api/quota'
import type { QuotaInfo } from '@/types/quota'

import { alertItems, metricCards, systemStatuses, taskDigests } from '@/mock/dashboard'
import { quotaRecords } from '@/mock/quota'
import { downloadTaskRecords } from '@/mock/tasks'

import { getTaskList, createTask, cancelTask, retryTask } from '@/api/task'
import type { DownloadTask } from '@/types/download'

import { uploadToken } from '@/api/token'
import type { Token } from '@/types/token'

export const useConsoleStore = defineStore('console', () => {
  const metrics = ref(metricCards)
  const statuses = ref(systemStatuses)
  const tasks = ref(taskDigests)
  const alerts = ref(alertItems)
  const taskTable = ref(downloadTaskRecords)
  const quotas = ref(quotaRecords)

  const providers = ref<ProviderRecord[]>([])

  const currentQuota = ref<QuotaInfo | null>(null)
  const quotaLoading = ref(false)

  const downloadTasks = ref<DownloadTask[]>([])
  const tasksLoading = ref(false)

  const tokens = ref<Token[]>([])
  const tokenLoading = ref(false)

  // 获取 Provider 列表
  async function fetchProviders() {
    try {
      const res = await getProviderList()
      if (res.code === 1000) {
        providers.value = res.data
      }
    } catch (error) {
      console.error('获取 Provider 列表失败', error)
      providers.value = []
    }
  }

  // 删除 Provider
  async function removeProvider(id: number) {
    try {
      const res = await deleteProvider(id)
      if (res.code === 1000) {
        await fetchProviders()
        return true
      }
      return false
    } catch (error) {
      console.error('删除失败', error)
      return false
    }
  }

  // 查询配额
  async function fetchQuota(provider: string) {
    quotaLoading.value = true
    try {
      const res = await queryQuota(provider)
      // 改成兼容 code: 0 和 code: 1000
      if (res.code === 0 || res.code === 1000) {
        currentQuota.value = res.data
      }
      return res
    } catch (error) {
      console.error('查询配额失败', error)
      return null
    } finally {
      quotaLoading.value = false
    }
  }

  // 同步配额
  async function syncProviderQuota(provider: string) {
    quotaLoading.value = true
    try {
      const res = await syncQuota(provider)
      if (res.code === 0 || res.code === 1000) {
        currentQuota.value = res.data
      }
      return res
    } catch (error) {
      console.error('同步配额失败', error)
      return null
    } finally {
      quotaLoading.value = false
    }
  }
  // 获取任务列表
  async function fetchTasks() {
    tasksLoading.value = true
    try {
      const res = await getTaskList()
      if (res.code === 0 || res.code === 1000) {
        downloadTasks.value = res.data
      }
      return res
    } catch (error) {
      console.error('获取任务列表失败', error)
      return null
    } finally {
      tasksLoading.value = false
    }
  }

  // 创建任务
  async function addTask(data: { source_url: string; provider?: string; file_name?: string }) {
    try {
      const res = await createTask(data)
      if (res.code === 0 || res.code === 1000) {
        await fetchTasks()
      }
      return res
    } catch (error) {
      console.error('创建任务失败', error)
      return null
    }
  }

  // 取消任务
  async function cancelDownloadTask(taskId: string) {
    try {
      const res = await cancelTask(taskId)
      if (res.code === 0 || res.code === 1000) {
        await fetchTasks()
      }
      return res
    } catch (error) {
      console.error('取消任务失败', error)
      return null
    }
  }

  // 重试任务
  async function retryDownloadTask(taskId: string) {
    try {
      const res = await retryTask(taskId)
      if (res.code === 0 || res.code === 1000) {
        await fetchTasks()
      }
      return res
    } catch (error) {
      console.error('重试任务失败', error)
      return null
    }
  }

  async function saveToken(data: Partial<Token>) {
    tokenLoading.value = true
    try {
      const res = await uploadToken(data)
      if (res.code === 0 || res.code === 1000) {
        // 成功后可以刷新列表（如果有列表接口的话）
        return { success: true, message: '上传成功' }
      }
      return { success: false, message: res.msg || res.msg }
    } catch (error) {
      console.error('上传 Token 失败', error)
      return { success: false, message: '上传失败' }
    } finally {
      tokenLoading.value = false
    }
  }


  // 只有一个 return，放在最后
  return {
    metrics,
    statuses,
    tasks,
    alerts,
    providers,
    quotas,
    taskTable,
    fetchProviders,
    removeProvider,
    currentQuota,
    quotaLoading,
    fetchQuota,
    syncProviderQuota,
    downloadTasks,
    tasksLoading,
    fetchTasks,
    addTask,
    cancelDownloadTask,
    retryDownloadTask,
    tokens,
    tokenLoading,
    saveToken,
  }
})