<template>
  <div class="upload-component">
    <el-upload
      class="upload-area"
      drag
      action="#"
      :auto-upload="false"
      :show-file-list="false"
      :on-change="handleChange"
      accept="image/*"
    >
      <el-icon class="upload-icon"><el-icon-upload /></el-icon>
      <div class="upload-text">
        <span class="main-text">点击上传图片或拖拽图片至此处</span>
        <span class="sub-text">支持JPG、PNG、GIF等格式</span>
      </div>
    </el-upload>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { Upload as ElIconUpload } from '@element-plus/icons-vue'

const emit = defineEmits(['file-uploaded'])

const handleChange = (file: any) => {
  const isImage = file.raw.type.indexOf('image/') !== -1
  if (!isImage) {
    ElMessage.error('请上传图片文件')
    return
  }

  // 限制文件大小 (10MB)
  const isLt10M = file.raw.size / 1024 / 1024 < 10
  if (!isLt10M) {
    ElMessage.error('图片大小不能超过10MB')
    return
  }

  // 创建图片预览
  const reader = new FileReader()
  reader.onload = (e) => {
    if (e.target?.result) {
      emit('file-uploaded', file.raw, e.target.result)
    }
  }
  reader.readAsDataURL(file.raw)
}
</script>

<style scoped>
.upload-component {
  width: 100%;
}

.upload-area {
  width: 100%;
  border: 2px dashed #dcdfe6;
  border-radius: 8px;
  padding: 60px 0;
  text-align: center;
  cursor: pointer;
  transition: border-color 0.3s;
}

.upload-area:hover {
  border-color: #409eff;
}

.upload-icon {
  font-size: 48px;
  color: #c0c4cc;
  margin-bottom: 16px;
}

.upload-text {
  display: flex;
  flex-direction: column;
}

.main-text {
  font-size: 16px;
  color: #606266;
  margin-bottom: 8px;
}

.sub-text {
  font-size: 14px;
  color: #909399;
}
</style> 