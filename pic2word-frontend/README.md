# Pic2Word 图片转Word应用

一个简单易用的图片转Word文档工具，支持OCR文字识别和文档格式转换。

## 功能特点

- 图片上传与预览
- OCR文字识别（支持多语言）
- 转换为可编辑的Word文档
- 简洁直观的用户界面

## 技术栈

### 前端
- Vue.js 3
- TypeScript
- Element Plus UI框架
- Axios HTTP客户端
- Vue Router
- Vite 构建工具

### 后端
- Golang/Java（待确定）
- OCR引擎（待确定）
- Word文档生成库

## 开发环境设置

1. 安装依赖

```bash
# 使用yarn安装依赖
yarn
```

2. 运行开发服务器

```bash
# 启动开发服务器
yarn dev
```

3. 构建生产版本

```bash
# 构建生产版本
yarn build
```

## API接口

### 图片转Word接口

```
POST /api/convert
```

**参数**：
- `image`: 图片文件（支持JPG、PNG、GIF等格式）
- `format`: 文档格式（默认'docx'）
- `language`: 识别语言（如'zh', 'en'）

**返回**：
- Word文档（Blob）

## 项目结构

```
pic2word-frontend/
├── public/              # 静态资源
├── src/
│   ├── api/             # API接口
│   ├── assets/          # 图片等资源
│   ├── components/      # 可复用组件
│   ├── router/          # 路由配置
│   ├── views/           # 页面组件
│   ├── App.vue          # 根组件
│   └── main.ts          # 入口文件
├── package.json         # 依赖管理
└── vite.config.ts       # Vite配置
```

## 后续开发计划

- 添加更多文档格式支持
- 增强OCR识别精度
- 添加图片预处理功能
- 批量处理功能

## 许可证

MIT
