import type { QuotaRecord } from '@/types/quota'

export const quotaRecords: QuotaRecord[] = [
  { provider: 'Baidu Netdisk', used: 880, total: 1000, updatedAt: '5 minutes ago' },
  { provider: 'Aliyun Drive', used: 420, total: 1024, updatedAt: '2 minutes ago' },
  { provider: 'OneDrive', used: 690, total: 1000, updatedAt: '18 minutes ago' },
]
