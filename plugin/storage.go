package plugin

import (
	"fmt"
	"mime/multipart"
	"sync"
)

// IStorage 存储接口
type IStorage interface {
	Init(endpoint, accessKeyId,
		secretAccessKey, filePath string, isSSL bool) error
	UploadFile(file *multipart.FileHeader) (filePathres, key string, uploadErr error)
	DeleteFile(key string) error
}

// 存储类型常量
const (
	StorageTypeLocal = "local"
	StorageTypeMinio = "minio"
)

// 存储工厂函数类型
type StorageFactory func() IStorage

// 存储工厂映射表
var storageFactories = make(map[string]StorageFactory)
var factoryMutex sync.RWMutex

// 全局存储实例
var globalStorage IStorage
var storageMutex sync.RWMutex

// RegisterStorage 注册存储实现
func RegisterStorage(storageType string, factory StorageFactory) {
	factoryMutex.Lock()
	defer factoryMutex.Unlock()

	storageFactories[storageType] = factory
}

// InitStorage 初始化插件系统，指定存储类型
func InitStorage(storageType string) error {
	factoryMutex.RLock()
	factory, exists := storageFactories[storageType]
	factoryMutex.RUnlock()

	if !exists {
		return fmt.Errorf("未知的存储类型: %s", storageType)
	}

	storageMutex.Lock()
	globalStorage = factory()
	storageMutex.Unlock()

	return nil
}

// GetStorage 获取全局存储实例
func GetStorage() IStorage {
	storageMutex.RLock()
	defer storageMutex.RUnlock()

	if globalStorage == nil {
		panic("存储插件未初始化")
	}

	return globalStorage
}
