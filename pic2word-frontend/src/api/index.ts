import axios from 'axios'

// 创建axios实例
const apiClient = axios.create({
  baseURL: import.meta.env.VITE_API_URL || '',
  timeout: 60000, // 因为图片处理可能需要较长时间，所以设置较长的超时时间
  headers: {
    'Content-Type': 'multipart/form-data',
  }
})

// API接口
export const api = {
  /**
   * 上传图片并转换为Word文档
   * @param imageFile 图片文件
   * @param options 转换选项
   * @returns Blob Word文档数据
   */
  convertImageToWord: async (imageFile: File, options: {
    format: string, // 文档格式, 如 'docx'
    language: string // 识别语言, 如 'zh', 'en'
  }) => {
    const formData = new FormData()
    formData.append('image', imageFile)
    formData.append('format', options.format)
    formData.append('language', options.language)
    
    const response = await apiClient.post('/convert', formData, {
      responseType: 'blob'
    })
    
    return response.data
  },
  
  /**
   * 获取支持的文档格式列表
   * @returns 格式列表
   */
  getSupportedFormats: async () => {
    const response = await apiClient.get('/formats')
    return response.data
  },
  
  /**
   * 获取支持的识别语言列表
   * @returns 语言列表
   */
  getSupportedLanguages: async () => {
    const response = await apiClient.get('/languages')
    return response.data
  }
}

export default api 