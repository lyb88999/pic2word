package services

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"path/filepath"
	"strings"
	"time"

	"github.com/lyb88999/pic2word/config"
	"google.golang.org/api/option"

	"github.com/google/generative-ai-go/genai"
)

// GeminiService 提供与Gemini API交互的功能
type GeminiService struct {
	client *genai.Client
}

// NewGeminiService 创建一个新的Gemini服务实例
func NewGeminiService() (*GeminiService, error) {
	ctx := context.Background()

	// 输出API密钥前几个字符用于调试
	apiKey := config.AppConfig.GoogleAPIKey
	if len(apiKey) > 5 {
		log.Printf("使用Gemini API密钥: %s*** (密钥前5个字符)", apiKey[:5])
	} else {
		log.Printf("警告: API密钥长度不正确")
	}

	// 使用API密钥创建客户端
	client, err := genai.NewClient(ctx, option.WithAPIKey(config.AppConfig.GoogleAPIKey))
	if err != nil {
		return nil, fmt.Errorf("创建Gemini客户端失败: %v", err)
	}

	return &GeminiService{
		client: client,
	}, nil
}

// ImageToLatex 将图片转换为LaTeX格式
func (s *GeminiService) ImageToLatex(ctx context.Context, imagePath string) (string, error) {
	log.Printf("开始处理图片转LaTeX: %s", imagePath)

	// 读取图片文件
	imageData, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return "", fmt.Errorf("读取图片文件失败: %v", err)
	}

	log.Printf("成功读取图片文件，大小: %d 字节", len(imageData))

	// 获取MIME类型
	ext := filepath.Ext(imagePath)
	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		mimeType = "image/jpeg" // 默认MIME类型
	}

	// 创建生成模型
	model := s.client.GenerativeModel("gemini-2.0-flash")

	// 设置请求上下文超时
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// 设置生成参数
	model.SetTemperature(0.2) // 较低的温度使输出更确定性
	model.SetTopP(0.95)
	model.SetTopK(40)
	model.SetMaxOutputTokens(4096) // 较大的输出token以容纳完整LaTeX

	// 构建提示词
	prompt := "帮我把这个图片的内容转换为完整的overleaf格式（LaTeX格式），确保生成的LaTeX代码是完整且可编译的，包括必要的导言区内容。仅返回LaTeX代码，无需其他说明。"

	// 创建图片部分
	img := genai.Blob{
		MIMEType: mimeType,
		Data:     imageData,
	}

	log.Printf("开始调用Gemini API...")

	// 构建实际请求
	resp, err := model.GenerateContent(ctx, genai.Text(prompt), img)
	if err != nil {
		switch {
		case ctx.Err() == context.DeadlineExceeded:
			return "", fmt.Errorf("Gemini API调用超时（30秒）: %v", err)
		case strings.Contains(err.Error(), "authentication"):
			return "", fmt.Errorf("Gemini API认证失败，请检查API密钥: %v", err)
		case strings.Contains(err.Error(), "quota"):
			return "", fmt.Errorf("Gemini API配额已用尽: %v", err)
		default:
			return "", fmt.Errorf("生成内容失败: %v", err)
		}
	}

	log.Printf("Gemini API调用完成，正在解析响应...")

	// 解析响应
	if resp == nil || len(resp.Candidates) == 0 {
		return "", fmt.Errorf("未生成任何内容")
	}

	if len(resp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("生成的内容为空")
	}

	// 提取LaTeX文本
	for _, part := range resp.Candidates[0].Content.Parts {
		if textPart, ok := part.(genai.Text); ok {
			latexContent := string(textPart)
			log.Printf("成功生成LaTeX内容，长度：%d字符", len(latexContent))
			return latexContent, nil
		}
	}

	return "", fmt.Errorf("未找到文本内容")
}

// Close 关闭Gemini客户端
func (s *GeminiService) Close() {
	if s.client != nil {
		s.client.Close()
	}
}
