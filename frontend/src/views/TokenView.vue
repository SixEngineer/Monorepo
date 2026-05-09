<script setup lang="ts">
import { ref } from 'vue'
import PageHeader from '@/components/common/PageHeader.vue'
import { useConsoleStore } from '@/stores/console'
import { netDiskOptions, type NetDiskType } from '@/types/token'

const store = useConsoleStore()

const formData = ref({
  netDisk: 'mock' as NetDiskType,
  accessToken: '',
  refreshToken: '',
  expiresAt: ''
})

const loading = ref(false)
const resultMessage = ref<{ type: 'success' | 'error'; text: string } | null>(null)

async function handleSubmit() {
  if (!formData.value.accessToken) {
    resultMessage.value = { type: 'error', text: 'Access Token 不能为空' }
    return
  }

  loading.value = true
  resultMessage.value = null

  try {
    const data: any = {
      netDisk: formData.value.netDisk,
      accessToken: formData.value.accessToken
    }
    
    if (formData.value.refreshToken) {
      data.refreshToken = formData.value.refreshToken
    }
    
    if (formData.value.expiresAt) {
      data.expiresAt = new Date(formData.value.expiresAt).toISOString()
    }

    const result = await store.saveToken(data)
    
    if (result.success) {
      resultMessage.value = { type: 'success', text: 'Token 上传成功！' }
      // 清空表单（可选）
      formData.value.accessToken = ''
      formData.value.refreshToken = ''
      formData.value.expiresAt = ''
    } else {
      resultMessage.value = { type: 'error', text: result.message }
    }
  } catch (error) {
    resultMessage.value = { type: 'error', text: '上传失败，请稍后重试' }
  } finally {
    loading.value = false
  }
}

function handleReset() {
  formData.value = {
    netDisk: 'mock',
    accessToken: '',
    refreshToken: '',
    expiresAt: ''
  }
  resultMessage.value = null
}
</script>

<template>
  <section class="page">
    <PageHeader
      title="Token Management"
      description="上传和管理网盘的 Access Token"
    />

    <div class="token-form-container">
      <div class="token-card">
        <h3 class="token-card__title">上传 Token</h3>
        <p class="token-card__desc">
          目前仅支持 Mock 网盘，后续可扩展百度、阿里云等
        </p>

        <form @submit.prevent="handleSubmit" class="token-form">
          <div class="form-group">
            <label>网盘类型 *</label>
            <select v-model="formData.netDisk" class="form-control">
              <option v-for="opt in netDiskOptions" :key="opt.value" :value="opt.value">
                {{ opt.label }}
              </option>
            </select>
          </div>

          <div class="form-group">
            <label>Access Token *</label>
            <textarea
              v-model="formData.accessToken"
              class="form-control form-control--textarea"
              placeholder="粘贴 Access Token"
              rows="4"
            />
          </div>

          <div class="form-group">
            <label>Refresh Token（可选）</label>
            <textarea
              v-model="formData.refreshToken"
              class="form-control form-control--textarea"
              placeholder="粘贴 Refresh Token"
              rows="3"
            />
          </div>

          <div class="form-group">
            <label>过期时间（可选）</label>
            <input
              v-model="formData.expiresAt"
              type="datetime-local"
              class="form-control"
            />
          </div>

          <div v-if="resultMessage" class="message" :class="`message--${resultMessage.type}`">
            {{ resultMessage.text }}
          </div>

          <div class="form-actions">
            <button type="button" class="btn btn--secondary" @click="handleReset" :disabled="loading">
              重置
            </button>
            <button type="submit" class="btn btn--primary" :disabled="loading">
              {{ loading ? '上传中...' : '上传 Token' }}
            </button>
          </div>
        </form>
      </div>

      <div class="info-card">
        <h4>说明</h4>
        <ul>
          <li>Token 用于访问网盘 API，获取配额信息</li>
          <li>上传后会自动覆盖同网盘类型的旧 Token</li>
          <li>Access Token 必填，Refresh Token 可选</li>
          <li>过期时间用于提醒，非必填</li>
        </ul>
      </div>
    </div>
  </section>
</template>

<style scoped>
.token-form-container {
  display: grid;
  grid-template-columns: 1fr 320px;
  gap: 24px;
}

.token-card {
  background: white;
  border-radius: 12px;
  padding: 24px;
  border: 1px solid #e5e7eb;
}

.token-card__title {
  margin: 0 0 8px 0;
  font-size: 20px;
  font-weight: 600;
}

.token-card__desc {
  margin: 0 0 24px 0;
  color: #6b7280;
  font-size: 14px;
}

.token-form {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.form-group label {
  font-size: 14px;
  font-weight: 500;
  color: #374151;
}

.form-control {
  padding: 10px 12px;
  border: 1px solid #d1d5db;
  border-radius: 8px;
  font-size: 14px;
  font-family: inherit;
}

.form-control:focus {
  outline: none;
  border-color: #3b82f6;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1);
}

.form-control--textarea {
  resize: vertical;
  min-height: 80px;
}

.message {
  padding: 12px 16px;
  border-radius: 8px;
  font-size: 14px;
}

.message--success {
  background: #dcfce7;
  color: #16a34a;
  border: 1px solid #bbf7d0;
}

.message--error {
  background: #fee2e2;
  color: #dc2626;
  border: 1px solid #fecaca;
}

.form-actions {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  margin-top: 8px;
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

.info-card {
  background: #f9fafb;
  border-radius: 12px;
  padding: 20px;
  border: 1px solid #e5e7eb;
  height: fit-content;
}

.info-card h4 {
  margin: 0 0 12px 0;
  font-size: 16px;
  font-weight: 600;
}

.info-card ul {
  margin: 0;
  padding-left: 20px;
  color: #6b7280;
  font-size: 14px;
  line-height: 1.6;
}

.info-card li {
  margin-bottom: 8px;
}
</style>