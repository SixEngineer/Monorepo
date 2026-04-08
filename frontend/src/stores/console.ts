import { computed, ref } from 'vue'
import { defineStore } from 'pinia'
import { alertItems, metricCards, systemStatuses, taskDigests } from '@/mock/dashboard'
import { providerRecords } from '@/mock/provider'
import { quotaRecords } from '@/mock/quota'
import { downloadTaskRecords } from '@/mock/tasks'

export const useConsoleStore = defineStore('console', () => {
  const metrics = ref(metricCards)
  const statuses = ref(systemStatuses)
  const tasks = ref(taskDigests)
  const alerts = ref(alertItems)
  const providers = ref(providerRecords)
  const quotas = ref(quotaRecords)
  const taskTable = ref(downloadTaskRecords)

  const healthyServices = computed(() => statuses.value.filter((item) => item.state === 'healthy').length)

  return { metrics, statuses, tasks, alerts, providers, quotas, taskTable, healthyServices }
})
