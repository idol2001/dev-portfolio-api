package handlers

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"dev-portfolio-api/models"
	"dev-portfolio-api/pkg/global"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
)

// allowedTypes 允许的文件类型（扩展名）
var allowedTypes = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".svg":  true,
	".webp": true,
	".gif":  true,
}

// UploadFile 文件上传（阿里云 OSS 优先，未启用时降级为本地存储）
func UploadFile(c *gin.Context) {
	fileType := c.Param("type")
	if fileType == "" {
		models.FailWithMessage("请指定上传类型", c)
		return
	}

	// 获取上传的文件
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		models.FailWithMessage("请上传文件", c)
		return
	}
	defer file.Close()

	// 校验文件类型
	ext := filepath.Ext(header.Filename)
	if !allowedTypes[ext] {
		models.FailWithMessage("不支持的文件类型，仅支持: jpg, png, svg, webp, gif", c)
		return
	}

	// 校验文件大小（10MB）
	const maxFileSize = 10 << 20
	if header.Size > maxFileSize {
		models.FailWithMessage("文件大小不能超过 10MB", c)
		return
	}

	// 生成 OSS key
	dirPrefix := global.Conf.AliyunOss.DirPrefix
	if dirPrefix == "" {
		dirPrefix = "dev-portfolio"
	}
	now := time.Now()
	objectKey := fmt.Sprintf("%s/%s/%d/%02d/%d_%d%s",
		dirPrefix,
		fileType,
		now.Year(),
		now.Month(),
		now.UnixMilli(),
		now.Nanosecond(),
		ext,
	)

	// 如果启用了 OSS，上传到阿里云 OSS
	if global.Conf.AliyunOss.Enable {
		fileURL, err := uploadToOSS(file, objectKey, header.Size)
		if err != nil {
			models.FailWithMessage(fmt.Sprintf("OSS 上传失败: %v", err), c)
			return
		}
		models.OkWithData(gin.H{
			"url":      fileURL,
			"filename": header.Filename,
			"size":     header.Size,
		}, c)
		return
	}

	// 降级为本地存储（开发环境）
	fileURL, err := saveToLocal(file, fileType, ext)
	if err != nil {
		models.FailWithMessage(fmt.Sprintf("本地保存失败: %v", err), c)
		return
	}

	models.OkWithData(gin.H{
		"url":      fileURL,
		"filename": header.Filename,
		"size":     header.Size,
	}, c)
}

// uploadToOSS 上传文件到阿里云 OSS
func uploadToOSS(file multipart.File, objectKey string, size int64) (string, error) {
	cfg := global.Conf.AliyunOss

	// 创建 OSS 客户端
	client, err := oss.New(cfg.Endpoint, cfg.AccessKeyId, cfg.AccessKeySecret)
	if err != nil {
		return "", fmt.Errorf("创建 OSS 客户端失败: %w", err)
	}

	// 获取存储空间
	bucket, err := client.Bucket(cfg.BucketName)
	if err != nil {
		return "", fmt.Errorf("获取 Bucket 失败: %w", err)
	}

	// 读取文件内容到内存（小文件适用，10MB 以内安全）
	data, err := io.ReadAll(io.LimitReader(file, 10<<20))
	if err != nil {
		return "", fmt.Errorf("读取文件失败: %w", err)
	}

	// 上传
	err = bucket.PutObject(objectKey, bytes.NewReader(data))
	if err != nil {
		return "", fmt.Errorf("上传到 OSS 失败: %w", err)
	}

	// 构建可访问 URL
	if cfg.BucketDomain != "" {
		return fmt.Sprintf("%s/%s", cfg.BucketDomain, objectKey), nil
	}
	return fmt.Sprintf("https://%s.%s/%s", cfg.BucketName, cfg.Endpoint, objectKey), nil
}

// saveToLocal 降级保存为本地文件
func saveToLocal(file multipart.File, fileType string, ext string) (string, error) {
	now := time.Now()
	dir := fmt.Sprintf("uploads/%s/%d/%02d", fileType, now.Year(), now.Month())
	filename := fmt.Sprintf("%d_%d%s", now.UnixMilli(), now.Nanosecond(), ext)
	fullPath := dir + "/" + filename

	// 确保目录存在
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", err
	}

	// 保存文件
	dst, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, file); err != nil {
		return "", err
	}

	return "/" + fullPath, nil
}
