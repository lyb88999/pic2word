package utils

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// ValidateGeminiAPIKey 测试Gemini API密钥是否有效
func ValidateGeminiAPIKey(apiKey string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// 创建客户端
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return fmt.Errorf("创建Gemini客户端失败: %v", err)
	}
	defer client.Close()

	// 创建简单模型
	model := client.GenerativeModel("gemini-2.0-flash")

	// 发送简单请求
	resp, err := model.GenerateContent(ctx, genai.Text("测试API连接，请回复'连接成功'"))
	if err != nil {
		return fmt.Errorf("API测试失败: %v", err)
	}

	// 检查响应
	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return fmt.Errorf("API返回空响应")
	}

	log.Printf("API测试成功，收到响应")
	return nil
}
