package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lyb88999/pic2word/utils"
)

func main() {
	// 加载.env文件
	err := godotenv.Load()
	if err != nil {
		log.Println("未找到.env文件，尝试使用命令行参数")
	}

	// 获取API密钥
	apiKey := os.Getenv("GOOGLE_API_KEY")

	// 如果环境变量中没有，尝试从命令行参数获取
	if apiKey == "" && len(os.Args) > 1 {
		apiKey = os.Args[1]
	}

	if apiKey == "" {
		fmt.Println("错误: 未提供API密钥。请在.env文件中设置GOOGLE_API_KEY或作为命令行参数提供")
		fmt.Println("用法: apitest [API_KEY]")
		os.Exit(1)
	}

	fmt.Printf("正在测试API密钥 (前5个字符): %s***\n", apiKey[:5])

	// 测试API密钥
	err = utils.ValidateGeminiAPIKey(apiKey)
	if err != nil {
		fmt.Printf("❌ API测试失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ API测试成功！密钥有效且可以正常使用")
}
