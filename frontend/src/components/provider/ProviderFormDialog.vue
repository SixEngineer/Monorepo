<script setup lang="ts">
import { ref, watch } from 'vue'
import type { ProviderRecord } from '@/types/provider'

const props = defineProps<{
  visible: boolean
  provider?: ProviderRecord | null  // 编辑时传入
}>()

const emit = defineEmits<{
  (e: 'update:visible', value: boolean): void
  (e: 'submit', data: Partial<ProviderRecord>): void
}>()

const formData = ref<Partial<ProviderRecord>>({
  name: '',
  provider_type: 'mock',
  net_disk: 'mock',
  account_id: '',
  status: 'active',
  total_quota: 10737418240,  // 10GB
  used_quota: 0,
  available_quota: 10737418240
})

// 编辑时填充数据
watch(() => props.provider, (val) => {
  if (val) {
    formData.value = { ...val }
  } else {
    formData.value = {
      name: '',
      provider_type: 'mock',
      net_disk: 'mock',
      account_id: '',
      status: 'active',
      total_quota: 10737418240,
      used_quota: 0,
      available_quota: 10737418240
    }
  }
}, { immediate: true })

const providerTypes = [
  { value: 'mock', label: 'Mock' }
  // 后续可添加更多
]

const netDisks = [
  { value: 'mock', label: 'Mock' }
  // 后续可添加更多
]

const statusOptions = [
  { value: 'active', label: 'Active' },
  { value: 'disabled', label: 'Disabled' },
  { value: 'expired', label: 'Expired' },
  { value: 'error', label: 'Error' }
]

function handleClose() {
  emit('update:visible', false)
}

function handleSubmit() {
  emit('submit', formData.value)
  handleClose()
}

// 自动计算可用配额
function updateAvailableQuota() {
  const total = formData.value.total_quota || 0
  const used = formData.value.used_quota || 0
  formData.value.available_quota = total - used
}
</script>

<template>
  <div v-if="visible" class="dialog-overlay" @click.self="handleClose">
    <div class="dialog">
      <div class="dialog__header">
        <h3>{{ provider ? '编辑 Provider' : '注册 Provider' }}</h3>
        <button class="dialog__close" @click="handleClose">&times;</button>
      </div>
      
      <div class="dialog__body">
        <div class="form-group">
          <label>名称 *</label>
          <input v-model="formData.name" type="text" placeholder="如：阿里云盘测试" />
        </div>
        
        <div class="form-row">
          <div class="form-group">
            <label>Provider 类型</label>
            <select v-model="formData.provider_type">
              <option v-for="opt in providerTypes" :key="opt.value" :value="opt.value">
                {{ opt.label }}
              </option>
            </select>
          </div>
          
          <div class="form-group">
            <label>网盘类型</label>
            <select v-model="formData.net_disk">
              <option v-for="opt in netDisks" :key="opt.value" :value="opt.value">
                {{ opt.label }}
              </option>
            </select>
          </div>
        </div>
        
        <div class="form-group">
          <label>账户 ID</label>
          <input v-model="formData.account_id" type="text" placeholder="网盘账户标识" />
        </div>
        
        <div class="form-group">
          <label>状态</label>
          <select v-model="formData.status">
            <option v-for="opt in statusOptions" :key="opt.value" :value="opt.value">
              {{ opt.label }}
            </option>
          </select>
        </div>
        
        <div class="form-row">
          <div class="form-group">
            <label>总配额 (字节)</label>
            <input v-model.number="formData.total_quota" type="number" @input="updateAvailableQuota" />
          </div>
          
          <div class="form-group">
            <label>已用配额 (字节)</label>
            <input v-model.number="formData.used_quota" type="number" @input="updateAvailableQuota" />
          </div>
        </div>
        
        <div class="form-group">
          <label>可用配额</label>
          <input :value="formData.available_quota" type="number" disabled />
          <small>自动计算：总配额 - 已用配额</small>
        </div>
      </div>
      
      <div class="dialog__footer">
        <button class="btn btn--secondary" @click="handleClose">取消</button>
        <button class="btn btn--primary" @click="handleSubmit">
          {{ provider ? '保存' : '注册' }}
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.dialog-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.dialog {
  background: white;
  border-radius: 12px;
  width: 500px;
  max-width: 90vw;
  max-height: 80vh;
  overflow: hidden;
  display: flex;
  flex-direction: column;
  box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.1);
}

.dialog__header {
  padding: 20px 24px;
  border-bottom: 1px solid #e5e7eb;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.dialog__header h3 {
  margin: 0;
  font-size: 18px;
  font-weight: 600;
}

.dialog__close {
  background: none;
  border: none;
  font-size: 24px;
  cursor: pointer;
  color: #6b7280;
  padding: 0;
  line-height: 1;
}

.dialog__close:hover {
  color: #374151;
}

.dialog__body {
  padding: 24px;
  overflow-y: auto;
  flex: 1;
}

.form-group {
  margin-bottom: 16px;
}

.form-group label {
  display: block;
  font-size: 14px;
  font-weight: 500;
  margin-bottom: 6px;
  color: #374151;
}

.form-group input,
.form-group select {
  width: 100%;
  padding: 10px 12px;
  border: 1px solid #d1d5db;
  border-radius: 8px;
  font-size: 14px;
  box-sizing: border-box;
}

.form-group input:focus,
.form-group select:focus {
  outline: none;
  border-color: #3b82f6;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1);
}

.form-group input:disabled {
  background: #f3f4f6;
  color: #6b7280;
}

.form-group small {
  display: block;
  margin-top: 4px;
  font-size: 12px;
  color: #6b7280;
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
}

.dialog__footer {
  padding: 16px 24px;
  border-top: 1px solid #e5e7eb;
  display: flex;
  justify-content: flex-end;
  gap: 12px;
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

.btn--primary {
  background: #3b82f6;
  color: white;
}

.btn--primary:hover {
  background: #2563eb;
}

.btn--secondary {
  background: white;
  color: #374151;
  border: 1px solid #d1d5db;
}

.btn--secondary:hover {
  background: #f9fafb;
}
</style>