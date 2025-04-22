package handlers

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lyb88999/pic2word/config"
	"github.com/lyb88999/pic2word/services"
)

// ConvertHandler 处理图片转换请求
type ConvertHandler struct {
	geminiService     *services.GeminiService
	conversionService *services.ConversionService
}

// NewConvertHandler 创建新的转换处理器
func NewConvertHandler() (*ConvertHandler, error) {
	geminiService, err := services.NewGeminiService()
	if err != nil {
		return nil, err
	}

	conversionService := services.NewConversionService()

	return &ConvertHandler{
		geminiService:     geminiService,
		conversionService: conversionService,
	}, nil
}

// ConvertImage 处理图片转Word的请求
func (h *ConvertHandler) ConvertImage(c *gin.Context) {
	// 获取上传的图片文件
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无法获取上传的图片: " + err.Error(),
		})
		return
	}

	// 获取转换设置
	format := c.DefaultPostForm("format", "docx")
	language := c.DefaultPostForm("language", "zh")

	// 记录转换请求信息
	log.Printf("收到转换请求: 文件=%s, 格式=%s, 语言=%s", file.Filename, format, language)

	// 目前我们只支持docx格式
	if format != "docx" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "目前仅支持docx格式",
		})
		return
	}

	// 创建临时图片文件
	timePrefix := time.Now().Format("20060102150405")
	imageTempPath := filepath.Join(config.AppConfig.TempDir, fmt.Sprintf("%s_%s", timePrefix, file.Filename))

	if err := c.SaveUploadedFile(file, imageTempPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "保存上传的图片失败: " + err.Error(),
		})
		return
	}
	defer os.Remove(imageTempPath) // 最后清理临时图片文件

	// 设置请求上下文，带有超时（增加到120秒以适应可能较长的API处理时间）
	ctx, cancel := context.WithTimeout(c.Request.Context(), 120*time.Second)
	defer cancel()

	// 使用Gemini进行图片识别并转换为LaTeX
	latex, err := h.geminiService.ImageToLatex(ctx, imageTempPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "图片识别失败: " + err.Error(),
		})
		return
	}

	// 将LaTeX转换为Word文档
	docxFilePath, err := h.conversionService.LatexToDocx(latex)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "转换为Word文档失败: " + err.Error(),
		})
		return
	}
	defer h.conversionService.CleanupFiles(docxFilePath) // 清理临时文件

	// 读取生成的Word文档
	docxData, err := ioutil.ReadFile(docxFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "读取生成的Word文档失败: " + err.Error(),
		})
		return
	}

	// 设置响应头并发送文件
	filename := fmt.Sprintf("pic2word_%s.docx", timePrefix)
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filename))
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Expires", "0")
	c.Header("Cache-Control", "must-revalidate")
	c.Header("Pragma", "public")
	c.Header("Content-Length", fmt.Sprintf("%d", len(docxData)))

	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.wordprocessingml.document", docxData)
}

// GetSupportedFormats 返回支持的文档格式
func (h *ConvertHandler) GetSupportedFormats(c *gin.Context) {
	// 目前只支持docx
	formats := []string{"docx"}
	c.JSON(http.StatusOK, formats)
}

// GetSupportedLanguages 返回支持的语言
func (h *ConvertHandler) GetSupportedLanguages(c *gin.Context) {
	languages := []map[string]string{
		{"code": "zh", "name": "中文"},
		{"code": "en", "name": "英文"},
	}
	c.JSON(http.StatusOK, languages)
}

// Close 关闭资源
func (h *ConvertHandler) Close() {
	if h.geminiService != nil {
		h.geminiService.Close()
	}
}
