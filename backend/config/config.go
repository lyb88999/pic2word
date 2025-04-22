package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// Config 应用配置结构
type Config struct {
	Port         string
	Env          string
	GoogleAPIKey string
	ProjectID    string
	Location     string
	TempDir      string
	AllowOrigins []string
}

// AppConfig 全局配置变量
var AppConfig Config

// LoadConfig 从环境变量加载配置
func LoadConfig() {
	// 加载.env文件
	err := godotenv.Load()
	if err != nil {
		log.Println("未找到.env文件，尝试从环境变量加载")
	}

	// 初始化配置
	AppConfig = Config{
		Port:         getEnv("PORT", "8080"),
		Env:          getEnv("ENV", "development"),
		GoogleAPIKey: getEnv("GOOGLE_API_KEY", ""),
		ProjectID:    getEnv("GOOGLE_PROJECT_ID", ""),
		Location:     getEnv("GOOGLE_LOCATION", "us-central1"),
		TempDir:      getEnv("TEMP_DIR", "./tmp"),
		AllowOrigins: strings.Split(getEnv("ALLOW_ORIGINS", "http://localhost:5173"), ","),
	}

	// 确保临时目录存在
	ensureTempDirExists()

	log.Printf("配置加载完成: 端口=%s, 环境=%s\n", AppConfig.Port, AppConfig.Env)
}

// 获取环境变量，若不存在则返回默认值
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// 确保临时目录存在
func ensureTempDirExists() {
	err := os.MkdirAll(AppConfig.TempDir, 0755)
	if err != nil {
		log.Fatalf("无法创建临时目录: %v", err)
	}
}
