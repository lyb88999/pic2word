<template>
  <div class="home-container">
    <div class="header">
      <h1>Pic2Word 转换器</h1>
      <p>轻松将图片转换为可编辑的Word文档</p>
    </div>
    
    <div class="upload-container">
      <UploadComponent @file-uploaded="handleFileUploaded" />
    </div>
    
    <div v-if="imagePreview" class="preview-container">
      <div class="image-preview">
        <img :src="imagePreview" alt="上传的图片" />
        <div class="actions">
          <el-button type="danger" size="small" @click="removeImage">删除图片</el-button>
        </div>
      </div>
      
      <div class="convert-settings">
        <h3>转换设置</h3>
        <el-form :model="convertSettings" label-width="120px">
          <el-form-item label="文档格式">
            <el-select v-model="convertSettings.format" placeholder="选择格式">
              <el-option label="DOCX" value="docx" />
            </el-select>
          </el-form-item>
          
          <el-form-item label="识别语言">
            <el-select v-model="convertSettings.language" placeholder="选择语言">
              <el-option label="中文" value="zh" />
              <el-option label="英文" value="en" />
            </el-select>
          </el-form-item>
        </el-form>
        
        <el-button type="primary" @click="startConversion" :loading="converting">
          {{ converting ? '正在转换...' : '开始转换' }}
        </el-button>
      </div>
    </div>
    
    <div v-if="conversionComplete" class="result-container">
      <el-result
        icon="success"
        title="转换完成!"
        sub-title="您的文档已准备好下载"
      >
        <template #extra>
          <el-button type="primary" @click="downloadDocument">下载Word文档</el-button>
          <el-button @click="resetConversion">再次转换</el-button>
        </template>
      </el-result>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'
import { saveAs } from 'file-saver'
import UploadComponent from '../components/UploadComponent.vue'
import api from '../api'

const imagePreview = ref<string | null>(null)
const uploadedFile = ref<File | null>(null)
const converting = ref(false)
const conversionComplete = ref(false)
const documentBlob = ref<Blob | null>(null)
const convertSettings = reactive({
  format: 'docx',
  language: 'zh'
})

const handleFileUploaded = (file: File, preview: string) => {
  uploadedFile.value = file
  imagePreview.value = preview
  conversionComplete.value = false
  documentBlob.value = null
}

const removeImage = () => {
  uploadedFile.value = null
  imagePreview.value = null
  documentBlob.value = null
}

const startConversion = async () => {
  if (!uploadedFile.value) {
    ElMessage.warning('请先上传图片')
    return
  }
  
  converting.value = true
  
  try {
    // 调用实际API
    documentBlob.value = await api.convertImageToWord(uploadedFile.value, {
      format: convertSettings.format,
      language: convertSettings.language
    })
    
    conversionComplete.value = true
    ElMessage.success('转换成功')
  } catch (error) {
    ElMessage.error('转换失败，请重试')
    console.error(error)
  } finally {
    converting.value = false
  }
}

const downloadDocument = () => {
  if (!documentBlob.value) {
    ElMessage.warning('文档不存在或已过期，请重新转换')
    return
  }
  
  const timeStamp = new Date().toISOString().replace(/[:.]/g, '-')
  saveAs(documentBlob.value, `pic2word_${timeStamp}.docx`)
}

const resetConversion = () => {
  uploadedFile.value = null
  imagePreview.value = null
  conversionComplete.value = false
  documentBlob.value = null
}
</script>

<style scoped>
.home-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

.header {
  text-align: center;
  margin-bottom: 40px;
}

.upload-container {
  margin-bottom: 40px;
}

.preview-container {
  display: flex;
  gap: 30px;
  margin-bottom: 40px;
}

.image-preview {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.image-preview img {
  max-width: 100%;
  max-height: 400px;
  border: 1px solid #ddd;
  border-radius: 4px;
  margin-bottom: 10px;
}

.convert-settings {
  flex: 1;
  padding: 20px;
  border: 1px solid #eee;
  border-radius: 8px;
}

.result-container {
  margin-top: 40px;
}

@media (max-width: 768px) {
  .preview-container {
    flex-direction: column;
  }
}
</style> 