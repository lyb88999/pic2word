# Pic2Word 后端服务

图片转Word应用的后端服务，基于Golang实现，通过调用Google Gemini API进行图片内容识别并转换为LaTeX格式，然后使用pandoc将LaTeX转换为Word文档。

## 技术栈

- **编程语言**: Golang
- **Web框架**: Gin
- **多模态识别**: Google Gemini 2.0 Flash API
- **文档转换**: Pandoc
- **依赖管理**: Go Modules

## 系统要求

- Go 1.18+
- Pandoc 2.0+
- Google Gemini API密钥

## 安装依赖

### 1. 安装Go

从[Go官网](https://golang.org/dl/)下载并安装最新版本的Go。

### 2. 安装Pandoc

#### macOS
```bash
brew install pandoc
```

#### Linux (Ubuntu/Debian)
```bash
sudo apt-get install pandoc
```

#### Windows
从[Pandoc官网](https://pandoc.org/installing.html)下载并安装。

### 3. 获取Google Gemini API密钥

1. 访问[Google AI Studio](https://aistudio.google.com/)
2. 创建API密钥
3. 复制API密钥并保存

## 配置

1. 复制`.env.example`文件为`.env`:
```bash
cp .env.example .env
```

2. 编辑`.env`文件，填入必要的配置信息：
```
GOOGLE_API_KEY=your_google_api_key_here
```

## 构建和运行

### 开发模式

```bash
go run main.go
```

### 构建并运行

```bash
go build -o pic2word .
./pic2word
```

## API接口

### 1. 图片转换接口

**请求**:
```
POST /api/convert
```

**参数**:
- `image`: 图片文件（multipart/form-data）
- `format`: 文档格式（目前仅支持'docx'）
- `language`: 识别语言（如'zh', 'en'）

**返回**:
- Word文档（application/vnd.openxmlformats-officedocument.wordprocessingml.document）

### 2. 获取支持的格式

**请求**:
```
GET /api/formats
```

**返回**:
```json
["docx"]
```

### 3. 获取支持的语言

**请求**:
```
GET /api/languages
```

**返回**:
```json
[
  {"code": "zh", "name": "中文"},
  {"code": "en", "name": "英文"}
]
```

## 项目结构

```
backend/
├── config/             # 配置管理
├── handlers/           # HTTP请求处理
├── middleware/         # HTTP中间件
├── services/           # 业务逻辑服务
├── utils/              # 工具函数
├── .env.example        # 环境变量样例
├── .gitignore          # Git忽略文件
├── go.mod              # Go模块定义
├── go.sum              # Go依赖校验和
├── main.go             # 主程序入口
└── README.md           # 项目说明
```

## 工作流程

1. 用户上传图片
2. 后端保存上传的图片并调用Gemini API识别图片内容
3. Gemini API返回LaTeX格式的文本
4. 后端将LaTeX格式转换为Word文档（使用pandoc）
5. 将生成的Word文档返回给用户

## 技术实现细节

### Gemini API调用

- 使用Google的Gemini 2.0 Flash进行多模态推理
- 提示词设计专注于生成LaTeX格式的输出
- 配置较低的temperature参数以获得确定性输出

### LaTeX转Word

- 使用pandoc作为转换工具
- 支持复杂数学公式和表格的转换
- 保留文本结构和布局

## 贡献与开发

欢迎提交问题和改进建议！

## 许可证

MIT 