<script setup lang="ts">
import { ref, onMounted } from 'vue'
import PageHeader from '@/components/common/PageHeader.vue'
import { useConsoleStore } from '@/stores/console'

const store = useConsoleStore()

// 选中的 Provider
const selectedProvider = ref('mock')
const providers = ref(['mock'])  // 后续可以从 store.providers 获取

// 格式化字节
function formatBytes(bytes: number): string {
  if (bytes === 0) return '0 B'
  const k = 1024
  const sizes = ['B', 'KB', 'MB', 'GB', 'TB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

// 查询配额
async function handleQuery() {
  await store.fetchQuota(selectedProvider.value)
}

// 同步配额
async function handleSync() {
  if (!confirm('同步配额会调用远端接口，确定吗？')) return
  await store.syncProviderQuota(selectedProvider.value)
}

onMounted(() => {
  // 如果有 Provider 列表，可以从这里获取
  if (store.providers.length > 0) {
    providers.value = store.providers.map(p => p.net_disk)
    selectedProvider.value = providers.value[0]
  }
  handleQuery()
})
</script>

<template>
  <section class="page">
    <PageHeader 
      title="Quota Management" 
      description="查询和同步网盘配额信息"
    />

    <div class="quota-controls">
      <select v-model="selectedProvider" class="provider-select">
        <option v-for="p in providers" :key="p" :value="p">{{ p }}</option>
      </select>
      
      <div class="button-group">
        <button 
          class="btn btn--secondary" 
          @click="handleQuery"
          :disabled="store.quotaLoading"
        >
          查询配额
        </button>
        <button 
          class="btn btn--primary" 
          @click="handleSync"
          :disabled="store.quotaLoading"
        >
          {{ store.quotaLoading ? '同步中...' : '同步配额' }}
        </button>
      </div>
    </div>

    <div v-if="store.currentQuota" class="quota-card">
      <div class="quota-card__header">
        <h3>{{ store.currentQuota.provider }}</h3>
        <span class="quota-card__time">
          更新时间: {{ new Date(store.currentQuota.updated_at).toLocaleString() }}
        </span>
      </div>
      
      <div class="quota-stats">
        <div class="quota-stat">
          <span class="quota-stat__label">总配额</span>
          <span class="quota-stat__value">{{ formatBytes(store.currentQuota.total) }}</span>
        </div>
        <div class="quota-stat">
          <span class="quota-stat__label">已使用</span>
          <span class="quota-stat__value">{{ formatBytes(store.currentQuota.used) }}</span>
        </div>
        <div class="quota-stat">
          <span class="quota-stat__label">可用</span>
          <span class="quota-stat__value quota-stat__value--available">
            {{ formatBytes(store.currentQuota.available) }}
          </span>
        </div>
      </div>
      
      <div class="quota-progress">
        <div 
          class="quota-progress__bar" 
          :style="{ width: `${(store.currentQuota.used / store.currentQuota.total) * 100}%` }"
        ></div>
      </div>
    </div>

    <div v-else class="empty-state">
      <p>暂无配额数据，请点击"查询配额"或"同步配额"</p>
    </div>
  </section>
</template>

<style scoped>
.quota-controls {
  display: flex;
  align-items: center;
  gap: 20px;
  margin-bottom: 30px;
  padding: 20px;
  background: white;
  border-radius: 12px;
  border: 1px solid #e5e7eb;
}

.provider-select {
  padding: 10px 16px;
  border: 1px solid #d1d5db;
  border-radius: 8px;
  font-size: 14px;
  min-width: 200px;
}

.button-group {
  display: flex;
  gap: 12px;
  margin-left: auto;
}

.btn {
  padding: 10px 20px;
  border-radius: 8px;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  border: none;
  transition: all 0.2s;
}

.btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.btn--primary {
  background: #3b82f6;
  color: white;
}

.btn--primary:hover:not(:disabled) {
  background: #2563eb;
}

.btn--secondary {
  background: white;
  color: #374151;
  border: 1px solid #d1d5db;
}

.btn--secondary:hover:not(:disabled) {
  background: #f9fafb;
}

.quota-card {
  background: white;
  border-radius: 12px;
  padding: 24px;
  border: 1px solid #e5e7eb;
}

.quota-card__header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 24px;
}

.quota-card__header h3 {
  margin: 0;
  font-size: 20px;
  font-weight: 600;
  text-transform: capitalize;
}

.quota-card__time {
  font-size: 13px;
  color: #6b7280;
}

.quota-stats {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 24px;
  margin-bottom: 24px;
}

.quota-stat {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.quota-stat__label {
  font-size: 14px;
  color: #6b7280;
}

.quota-stat__value {
  font-size: 28px;
  font-weight: 600;
  color: #111827;
}

.quota-stat__value--available {
  color: #10b981;
}

.quota-progress {
  height: 8px;
  background: #e5e7eb;
  border-radius: 4px;
  overflow: hidden;
}

.quota-progress__bar {
  height: 100%;
  background: #3b82f6;
  border-radius: 4px;
  transition: width 0.3s;
}

.empty-state {
  text-align: center;
  padding: 60px 20px;
  color: #6b7280;
  background: #f9fafb;
  border-radius: 12px;
  border: 1px dashed #d1d5db;
}
</style>