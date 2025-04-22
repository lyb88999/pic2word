package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/lyb88999/pic2word/config"
	"github.com/lyb88999/pic2word/handlers"
	"github.com/lyb88999/pic2word/middleware"
)

func main() {
	// 加载配置
	config.LoadConfig()

	// 检查pandoc是否已安装
	if !isPandocInstalled() {
		log.Fatal("错误: pandoc未安装或无法访问。请安装pandoc后再运行应用")
	}

	// 设置Gin模式
	if config.AppConfig.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 创建Gin引擎
	r := gin.New()

	// 添加中间件
	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.CorsMiddleware())
	r.Use(gin.Recovery())

	// 创建转换处理器
	convertHandler, err := handlers.NewConvertHandler()
	if err != nil {
		log.Fatalf("创建转换处理器失败: %v", err)
	}
	defer convertHandler.Close()

	// API路由
	api := r.Group("/api")
	{
		// 图片转Word接口
		api.POST("/convert", convertHandler.ConvertImage)

		// 获取支持的格式
		api.GET("/formats", convertHandler.GetSupportedFormats)

		// 获取支持的语言
		api.GET("/languages", convertHandler.GetSupportedLanguages)
	}

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// 启动服务器
	port := config.AppConfig.Port
	log.Printf("服务启动在 :%s\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}

// 检查pandoc是否已安装
func isPandocInstalled() bool {
	cmd := exec.Command("pandoc", "--version")
	if err := cmd.Run(); err != nil {
		fmt.Println("pandoc未找到:", err)
		return false
	}
	return true
}
