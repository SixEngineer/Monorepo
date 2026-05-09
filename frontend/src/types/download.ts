export type DownloadStatus = 
  | 'pending'
  | 'resolving'
  | 'resolved'
  | 'submitting'
  | 'downloading'
  | 'completed'
  | 'failed'
  | 'cancelled'

export interface DownloadTask {
  id: number
  task_id: string
  source_url: string
  source_type?: string
  file_name?: string
  file_size?: number
  direct_link?: string
  provider_account_id?: number
  provider_type?: string
  aria2_gid?: string
  status: DownloadStatus
  progress: number
  error_message?: string
  retry_count: number
  started_at?: string
  finished_at?: string
  created_at: string
  updated_at: string
}