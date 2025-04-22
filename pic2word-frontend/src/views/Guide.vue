<template>
  <div class="guide-container">
    <div class="header">
      <h1>使用指南</h1>
    </div>
    
    <div class="content">
      <div class="step">
        <div class="step-number">1</div>
        <div class="step-content">
          <h3>选择并上传图片</h3>
          <p>点击上传区域或将图片拖拽至上传区域。支持JPG、PNG、GIF等常见图片格式，文件大小不超过10MB。</p>
          <div class="step-image">
            <el-image src="https://via.placeholder.com/800x400?text=上传图片示例" fit="contain" />
          </div>
        </div>
      </div>
      
      <div class="step">
        <div class="step-number">2</div>
        <div class="step-content">
          <h3>配置转换设置</h3>
          <p>上传图片后，您可以选择文档格式和识别语言等设置。默认为DOCX格式和中文识别。</p>
          <div class="step-image">
            <el-image src="https://via.placeholder.com/800x400?text=设置示例" fit="contain" />
          </div>
        </div>
      </div>
      
      <div class="step">
        <div class="step-number">3</div>
        <div class="step-content">
          <h3>开始转换</h3>
          <p>点击"开始转换"按钮，系统将开始处理您的图片。转换过程可能需要几秒钟，请耐心等待。</p>
        </div>
      </div>
      
      <div class="step">
        <div class="step-number">4</div>
        <div class="step-content">
          <h3>下载文档</h3>
          <p>转换完成后，点击"下载Word文档"按钮将文档保存到您的设备。</p>
        </div>
      </div>
      
      <div class="tips">
        <h3>使用技巧</h3>
        <ul>
          <li>为获得最佳效果，请确保图片清晰且文字可读。</li>
          <li>对于复杂的表格或图表，可能需要在下载后手动调整格式。</li>
          <li>如遇转换失败，请尝试调整图片亮度/对比度后重新上传。</li>
        </ul>
      </div>
    </div>
  </div>
</template>

<style scoped>
.guide-container {
  max-width: 1000px;
  margin: 0 auto;
  padding: 20px;
}

.header {
  text-align: center;
  margin-bottom: 40px;
}

.content {
  background-color: #f9f9f9;
  padding: 30px;
  border-radius: 8px;
}

.step {
  display: flex;
  margin-bottom: 40px;
  padding-bottom: 30px;
  border-bottom: 1px dashed #ddd;
}

.step:last-of-type {
  border-bottom: none;
}

.step-number {
  flex-shrink: 0;
  width: 50px;
  height: 50px;
  border-radius: 50%;
  background-color: #409eff;
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 24px;
  font-weight: bold;
  margin-right: 20px;
}

.step-content {
  flex-grow: 1;
}

.step-image {
  margin-top: 15px;
  border: 1px solid #eee;
  border-radius: 4px;
  overflow: hidden;
}

.tips {
  margin-top: 40px;
  padding: 20px;
  background-color: #ecf8ff;
  border-left: 4px solid #409eff;
  border-radius: 4px;
}

h3 {
  margin-top: 0;
  margin-bottom: 15px;
  color: #333;
}

ul {
  padding-left: 20px;
}

li {
  margin-bottom: 8px;
}
</style> 