package plugin

import (
	"fmt"
	"sync"
)

// ICrypto 密码加密接口
type ICrypto interface {
	// EncryptPassword 加密密码
	EncryptPassword(password string) (string, error)
	// VerifyPassword 验证密码
	VerifyPassword(password, hashedPassword string) bool
	// GetAlgorithm 获取当前算法名称
	GetAlgorithm() string
}

// 加密算法类型常量
const (
	CryptoTypeMD5    = "md5"
	CryptoTypeSHA1   = "sha1"
	CryptoTypeSHA256 = "sha256"
	CryptoTypeSHA512 = "sha512"
)

// 加密工厂函数类型
type CryptoFactory func() ICrypto

// 加密工厂映射表
var cryptoFactories = make(map[string]CryptoFactory)
var cryptoFactoryMutex sync.RWMutex

// 全局加密实例
var globalCrypto ICrypto
var cryptoMutex sync.RWMutex

// registerCrypto 注册加密实现
func registerCrypto(cryptoType string, factory CryptoFactory) {
	cryptoFactoryMutex.Lock()
	defer cryptoFactoryMutex.Unlock()

	cryptoFactories[cryptoType] = factory
}

// InitCrypto 初始化加密插件系统，指定加密类型
func InitCrypto(cryptoType string) error {
	cryptoFactoryMutex.RLock()
	factory, exists := cryptoFactories[cryptoType]
	cryptoFactoryMutex.RUnlock()

	if !exists {
		return fmt.Errorf("未知的加密类型: %s", cryptoType)
	}

	cryptoMutex.Lock()
	globalCrypto = factory()
	cryptoMutex.Unlock()

	return nil
}

// GetCrypto 获取全局加密实例
func GetCrypto() ICrypto {
	cryptoMutex.RLock()
	defer cryptoMutex.RUnlock()

	if globalCrypto == nil {
		panic("加密插件未初始化")
	}

	return globalCrypto
}
