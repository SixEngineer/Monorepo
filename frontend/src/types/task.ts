export interface DownloadTaskRecord {
  id: string
  name: string
  provider: string
  target: string
  speed: string
  progress: number
  status: 'running' | 'queued' | 'error' | 'complete'
}
