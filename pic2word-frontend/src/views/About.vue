<template>
  <div class="about-container">
    <div class="header">
      <h1>关于 Pic2Word</h1>
    </div>
    
    <div class="content">
      <h2>项目介绍</h2>
      <p>Pic2Word 是一个简单易用的图片转Word文档工具。我们的目标是让用户能够轻松地将图片中的文字内容提取并转化为可编辑的Word文档。</p>
      
      <h2>技术栈</h2>
      <p>前端技术：Vue.js、Element Plus、TypeScript</p>
      <p>后端技术：Golang/Java（根据实际选择）</p>
      
      <h2>核心功能</h2>
      <ul>
        <li>图片上传与预览</li>
        <li>OCR文字识别</li>
        <li>文档格式转换</li>
        <li>多语言支持</li>
      </ul>
      
      <h2>联系我们</h2>
      <p>如果您有任何问题或建议，请联系我们：</p>
      <p>邮箱：example@pic2word.com</p>
    </div>
  </div>
</template>

<style scoped>
.about-container {
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

h2 {
  margin-top: 30px;
  margin-bottom: 15px;
  color: #333;
  border-bottom: 1px solid #eee;
  padding-bottom: 10px;
}

ul {
  padding-left: 20px;
}

li {
  margin-bottom: 8px;
}
</style> 