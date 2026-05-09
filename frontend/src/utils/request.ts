import axios from 'axios'

const request = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || '/api/v1',
  timeout: 10000,
})

request.interceptors.request.use(
  (config) => {
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

request.interceptors.response.use(
  (response) => {
    const res = response.data
    // 兼容 code: 0 和 code: 1000 都算成功
    if (res.code !== 0 && res.code !== 1000) {
      console.error('API Error:', res.message || res.msg)
      return Promise.reject(new Error(res.message || res.msg || 'Error'))
    }
    return res
  },
  (error) => {
    console.error('Request Error:', error)
    return Promise.reject(error)
  }
)

export default request