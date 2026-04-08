import type { ProviderRecord } from '@/types/provider'

export const providerRecords: ProviderRecord[] = [
  { id: 'p-baidu', name: 'Baidu Netdisk', status: 'healthy', capability: 'Quota sync, token bridge, raw-link', authMode: 'Access + refresh token' },
  { id: 'p-aliyun', name: 'Aliyun Drive', status: 'healthy', capability: 'Direct-link resolve, storage introspection', authMode: 'Cookie relay' },
  { id: 'p-onedrive', name: 'OneDrive', status: 'warning', capability: 'Quota sync, debug traces', authMode: 'OAuth token' },
]
