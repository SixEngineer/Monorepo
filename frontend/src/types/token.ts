export type NetDiskType = 'baidu' | 'aliyun' | 'quark' | 'mock'

export interface Token {
  id: number
  netDisk: NetDiskType
  accessToken: string
  refreshToken?: string
  expiresAt?: string
  createdAt: string
  updatedAt: string
}

// 网盘选项（用于下拉框）
export const netDiskOptions: { value: NetDiskType; label: string }[] = [
  { value: 'baidu', label: '百度网盘' },
  { value: 'aliyun', label: '阿里云盘' },
  { value: 'quark', label: '夸克网盘' },
  { value: 'mock', label: 'Mock 测试' },
]