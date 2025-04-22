package services

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/lyb88999/pic2word/config"
)

// ConversionService 提供文档格式转换功能
type ConversionService struct {
	tempDir string
}

// NewConversionService 创建新的转换服务
func NewConversionService() *ConversionService {
	return &ConversionService{
		tempDir: config.AppConfig.TempDir,
	}
}

// LatexToDocx 将LaTeX转换为Word文档
func (s *ConversionService) LatexToDocx(latexContent string) (string, error) {
	log.Printf("开始LaTeX转Word处理，LaTeX内容长度: %d字符", len(latexContent))

	// 生成唯一的文件名前缀
	fileID := fmt.Sprintf("%d", os.Getpid())

	// 创建临时文件路径
	texFilePath := filepath.Join(s.tempDir, fileID+".tex")
	docxFilePath := filepath.Join(s.tempDir, fileID+".docx")

	log.Printf("创建临时文件路径: LaTeX=%s, Word=%s", texFilePath, docxFilePath)

	// 将LaTeX内容写入临时文件
	err := ioutil.WriteFile(texFilePath, []byte(latexContent), 0644)
	if err != nil {
		return "", fmt.Errorf("写入LaTeX文件失败: %v", err)
	}

	log.Printf("LaTeX内容已写入临时文件")

	// 使用pandoc将LaTeX转换为Word
	cmd := exec.Command("pandoc",
		"-f", "latex",
		"-t", "docx",
		"-o", docxFilePath,
		texFilePath)

	log.Printf("开始调用pandoc转换...")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Pandoc转换失败: %v, 输出: %s", err, output)
	}

	log.Printf("Pandoc转换成功，生成Word文档: %s", docxFilePath)

	return docxFilePath, nil
}

// CleanupFiles 清理临时文件
func (s *ConversionService) CleanupFiles(filePaths ...string) {
	for _, path := range filePaths {
		os.Remove(path)
	}
}
