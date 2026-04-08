import type { DownloadTaskRecord } from '@/types/task'

export const downloadTaskRecords: DownloadTaskRecord[] = [
  { id: 'DL-1024', name: 'Semester archive bundle', provider: 'Baidu Netdisk', target: '/downloads/archive.zip', speed: '22.4 MB/s', progress: 78, status: 'running' },
  { id: 'DL-1018', name: 'Research dataset mirror', provider: 'OneDrive', target: '/datasets/mirror.tar', speed: 'Queued', progress: 42, status: 'queued' },
  { id: 'DL-1012', name: 'Presentation media pack', provider: 'Google Drive', target: '/media/slides-assets.zip', speed: '0 MB/s', progress: 16, status: 'error' },
]
