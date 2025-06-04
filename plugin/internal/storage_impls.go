package internal

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

// StorageMinioImpl Minio存储实现
type StorageMinioImpl struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	IsSSL           bool
}

func (s *StorageMinioImpl) Init(endpoint, accessKeyId,
	secretAccessKey, filePath string, isSSL bool) error {
	s.Endpoint = endpoint
	s.AccessKeyID = accessKeyId
	s.SecretAccessKey = secretAccessKey
	s.IsSSL = isSSL
	return nil
}

func (s *StorageMinioImpl) UploadFile(file *multipart.FileHeader) (filePathres, key string, uploadErr error) {
	return "", "", uploadErr
}

func (s *StorageMinioImpl) DeleteFile(key string) error {
	return nil
}

// StorageLocalImpl 本地存储实现
type StorageLocalImpl struct {
	FilePath string
}

func (s *StorageLocalImpl) Init(endpoint, accessKeyId,
	secretAccessKey, filePath string, isSSL bool) error {
	s.FilePath = filePath
	return nil
}

func (s *StorageLocalImpl) UploadFile(file *multipart.FileHeader) (filePathres, key string, uploadErr error) {
	// 打开上传的文件
	src, err := file.Open()
	if err != nil {
		return "", "", err
	}
	defer src.Close()

	// 按日期创建子目录
	dateDir := time.Now().Format("2006/01/02")
	fullPath := filepath.Join(s.FilePath, dateDir)

	// 确保目录存在
	if err := os.MkdirAll(fullPath, 0755); err != nil {
		return "", "", err
	}

	// 获取原始文件名和扩展名
	originalName := file.Filename
	ext := filepath.Ext(originalName)

	// 生成随机文件名，不保留原文件名
	randomName := fmt.Sprintf("%s%s", uuid.New().String(), ext)

	// 完整的文件路径
	filePath := filepath.Join(fullPath, randomName)

	// 相对路径作为key (日期/随机文件名)
	key = filepath.Join(dateDir, randomName)

	// 创建目标文件
	dst, err := os.Create(filePath)
	if err != nil {
		return "", "", err
	}
	defer dst.Close()

	// 复制文件内容
	if _, err = io.Copy(dst, src); err != nil {
		return "", "", err
	}

	return filePath, key, nil
}

func (s *StorageLocalImpl) DeleteFile(key string) error {
	filePath := filepath.Join(s.FilePath, key)
	// 删除文件
	err := os.Remove(filePath)
	if err != nil {
		// 如果文件不存在，不视为错误
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	return nil
}

// 生成唯一文件名
func generateUniqueFileName(originalName string) string {
	timestamp := time.Now().UnixNano()
	ext := filepath.Ext(originalName)
	return filepath.Base(originalName[:len(originalName)-len(ext)]) +
		"_" + time.Now().Format("20060102150405") +
		"_" + fmt.Sprintf("%d", timestamp)[:8] + ext
}


