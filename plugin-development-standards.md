# 插件开发规范

## 插件架构设计原则

### 核心原则
- **接口与实现分离**：接口定义在外层，具体实现在internal包
- **工厂模式**：使用工厂函数创建实例，支持动态切换实现
- **统一注册**：所有插件实现在init.go中统一注册
- **线程安全**：使用互斥锁保证并发安全

## 标准文件结构

```
plugin/
├── init.go                    # 统一注册所有插件
├── 插件名.go                   # 接口定义和工厂管理
├── internal/
│   └── 插件名_impls.go         # 具体实现
└── 插件名_example.md           # 使用说明（可选）
```

## 文件职责分工

### `plugin/插件名.go`
**职责**：接口定义、常量、工厂函数、全局管理
**包含内容**：
- 接口定义（如`ICrypto`、`IStorage`）
- 类型常量（如`CryptoTypeMD5`、`StorageTypeLocal`）
- 工厂函数类型定义
- 工厂映射表和互斥锁
- 全局实例和互斥锁
- 注册函数（register开头，小写）
- 初始化函数（Init开头，大写）
- 获取实例函数（Get开头，大写）

### `plugin/internal/插件名_impls.go`
**职责**：具体实现类
**包含内容**：
- 各种算法/类型的具体实现结构体
- 实现接口的所有方法
- 实现相关的辅助函数

### `plugin/init.go`
**职责**：统一注册所有插件实现
**包含内容**：
- 导入internal包
- 在Init()函数中注册所有插件的所有实现

## 命名规范

### 接口命名
- 格式：`I + 功能名`
- 示例：`ICrypto`、`IStorage`、`ICache`

### 常量命名
- 格式：`功能名Type + 算法名`
- 示例：`CryptoTypeMD5`、`StorageTypeLocal`

### 实现类命名
- 格式：`功能名 + 算法名 + Impl`
- 示例：`CryptoMD5Impl`、`StorageLocalImpl`

### 函数命名
- 注册函数：`register + 功能名`（小写，内部使用）
- 初始化函数：`Init + 功能名`（大写，外部调用）
- 获取实例：`Get + 功能名`（大写，外部调用）

## 标准代码模板

### 接口文件模板（plugin/crypto.go）
```go
package plugin

import (
    "fmt"
    "sync"
)

// ICrypto 密码加密接口
type ICrypto interface {
    EncryptPassword(password string) (string, error)
    VerifyPassword(password, hashedPassword string) bool
    GetAlgorithm() string
}

// 加密算法类型常量
const (
    CryptoTypeMD5    = "md5"
    CryptoTypeSHA256 = "sha256"
)

// 工厂函数类型
type CryptoFactory func() ICrypto

// 工厂映射表
var cryptoFactories = make(map[string]CryptoFactory)
var cryptoFactoryMutex sync.RWMutex

// 全局实例
var globalCrypto ICrypto
var cryptoMutex sync.RWMutex

// 注册实现（内部使用）
func registerCrypto(cryptoType string, factory CryptoFactory) {
    cryptoFactoryMutex.Lock()
    defer cryptoFactoryMutex.Unlock()
    cryptoFactories[cryptoType] = factory
}

// 初始化插件（外部调用）
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

// 获取实例（外部调用）
func GetCrypto() ICrypto {
    cryptoMutex.RLock()
    defer cryptoMutex.RUnlock()

    if globalCrypto == nil {
        panic("加密插件未初始化")
    }

    return globalCrypto
}
```

### 实现文件模板（plugin/internal/crypto_impls.go）
```go
package internal

import (
    "crypto/md5"
    "encoding/hex"
    "fmt"
)

// CryptoMD5Impl MD5加密实现
type CryptoMD5Impl struct{}

func (c *CryptoMD5Impl) EncryptPassword(password string) (string, error) {
    if password == "" {
        return "", fmt.Errorf("密码不能为空")
    }
    
    hash := md5.Sum([]byte(password))
    return hex.EncodeToString(hash[:]), nil
}

func (c *CryptoMD5Impl) VerifyPassword(password, hashedPassword string) bool {
    encrypted, err := c.EncryptPassword(password)
    if err != nil {
        return false
    }
    return encrypted == hashedPassword
}

func (c *CryptoMD5Impl) GetAlgorithm() string {
    return "md5"
}
```

### 注册文件模板（plugin/init.go）
```go
package plugin

import "your-project/plugin/internal"

func Init() {
    // 注册存储插件
    registerStorage(StorageTypeLocal, func() IStorage {
        return &internal.StorageLocalImpl{}
    })

    // 注册加密插件
    registerCrypto(CryptoTypeMD5, func() ICrypto {
        return &internal.CryptoMD5Impl{}
    })
}
```

## 开发流程

1. **设计接口**：在`plugin/插件名.go`中定义接口和常量
2. **实现功能**：在`plugin/internal/插件名_impls.go`中实现具体功能
3. **注册插件**：在`plugin/init.go`中注册所有实现
4. **编写文档**：创建使用说明文档
5. **测试验证**：确保插件功能正常

## 注意事项

- **不要在接口文件中写具体实现**
- **所有实现必须在internal包中**
- **必须在init.go中注册实现**
- **遵循项目的错误处理规范**
- **使用项目规定的日志库**
- **保持线程安全**
