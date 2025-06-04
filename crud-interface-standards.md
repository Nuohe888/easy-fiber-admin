# CRUD接口开发规范

## ⚠️ 重要注意事项

### 🚨 ListAll接口使用限制
**ListAll接口仅在以下情况下才能使用：**
1. **数据量确定很小**（如角色表、字典表等配置类数据）
2. **与其他模型有关联关系**（如下拉选择框数据源）

**❌ 禁止在以下模型中使用ListAll接口：**
- User（用户表）- 数据量大
- UserCenter（用户中心表）- 数据量大  
- 日志表、订单表等业务数据表

**✅ 可以使用ListAll接口的模型：**
- Role（角色表）- 配置类数据，数量有限
- 字典表、分类表等配置数据

### 🚨 Status字段强制要求
**如果模型包含Status字段，必须额外提供以下接口：**
- `GET /api/模块名/状态接口` - 获取状态值说明

## 标准CRUD接口规范

### 基础接口（必须实现）

```go
// 标准CRUD接口 - 所有模型都必须实现
auth.Put("模型名/:id", controller.模型Ctl.Put)      // 更新
auth.Post("模型名", controller.模型Ctl.Add)         // 新增  
auth.Delete("模型名/:id", controller.模型Ctl.Del)   // 删除
auth.Get("模型名", controller.模型Ctl.Get)          // 获取单个
auth.Get("模型名/list", controller.模型Ctl.List)    // 获取列表（分页）
```

### 可选接口（按需实现）

```go
// ListAll接口 - 仅在符合条件时实现
auth.Get("模型名/list/all", controller.模型Ctl.ListAll) // 获取所有（不分页）

// Status接口 - 有Status字段时必须实现
auth.Get("模型名/status", controller.模型Ctl.GetStatus) // 获取状态说明
```

## Service层实现规范

### 必须实现的方法

```go
type 模型Srv struct {
    db  *gorm.DB
    log logger.ILog
}

// 基础CRUD方法
func (i *模型Srv) Add(model *system.模型) error
func (i *模型Srv) Del(id any) error  
func (i *模型Srv) Put(id any, model *system.模型) error
func (i *模型Srv) Get(id any) system.模型
func (i *模型Srv) List(page, limit int) *vo.List
```

### 可选实现的方法

```go
// 仅在符合条件时实现
func (i *模型Srv) ListAll() []system.模型

// 有Status字段时必须实现
func (i *模型Srv) GetStatus() map[string]interface{}
```

## Controller层实现规范

### 必须实现的方法

```go
type 模型Ctl struct {
    srv *service.模型Srv
}

// 基础CRUD接口
func (i *模型Ctl) Add(c *fiber.Ctx) error
func (i *模型Ctl) Del(c *fiber.Ctx) error
func (i *模型Ctl) Put(c *fiber.Ctx) error  
func (i *模型Ctl) Get(c *fiber.Ctx) error
func (i *模型Ctl) List(c *fiber.Ctx) error
```

### 可选实现的方法

```go
// 仅在符合条件时实现
func (i *模型Ctl) ListAll(c *fiber.Ctx) error

// 有Status字段时必须实现  
func (i *模型Ctl) GetStatus(c *fiber.Ctx) error
```

## 路由参数规范

### GET方法参数规范
- **单个查询**：`GET /api/模型名?id=1`
- **列表查询**：`GET /api/模型名/list?page=1&limit=20`
- **状态查询**：`GET /api/模型名/status`

### POST/PUT/DELETE方法参数规范
- **新增**：`POST /api/模型名` + Body参数
- **更新**：`PUT /api/模型名/:id` + Body参数
- **删除**：`DELETE /api/模型名/:id`

## Status字段处理规范

### Status字段定义要求
```go
type 模型 struct {
    Model
    // 其他字段...
    Status *int `json:"status"` // 0=禁用 1=启用
}
```

### GetStatus接口实现示例

**Service层：**
```go
func (i *模型Srv) GetStatus() map[string]interface{} {
    return map[string]interface{}{
        "0": "禁用",
        "1": "启用",
    }
}
```

**Controller层：**
```go
func (i *模型Ctl) GetStatus(c *fiber.Ctx) error {
    return vo.ResultOK(i.srv.GetStatus(), c)
}
```

**路由注册：**
```go
auth.Get("模型名/status", controller.模型Ctl.GetStatus)
```

## 实际应用示例

### ✅ 正确示例：Role模型（可以有ListAll）

```go
// 路由注册
auth.Put("role/:id", controller.RoleCtl.Put)
auth.Post("role", controller.RoleCtl.Add)
auth.Delete("role/:id", controller.RoleCtl.Del)
auth.Get("role", controller.RoleCtl.Get)
auth.Get("role/list", controller.RoleCtl.List)
auth.Get("role/list/all", controller.RoleCtl.ListAll) // ✅ 角色数据量小，可以使用
auth.Get("role/status", controller.RoleCtl.GetStatus) // ✅ 有Status字段，必须提供
```

### ❌ 错误示例：User模型（不应该有ListAll）

```go
// 路由注册
auth.Put("user/:id", controller.UserCtl.Put)
auth.Post("user", controller.UserCtl.Add)
auth.Delete("user/:id", controller.UserCtl.Del)
auth.Get("user", controller.UserCtl.Get)
auth.Get("user/list", controller.UserCtl.List)
// ❌ 用户数据量大，不应该提供ListAll接口
// auth.Get("user/list/all", controller.UserCtl.ListAll) 
auth.Get("user/status", controller.UserCtl.GetStatus) // ✅ 有Status字段，必须提供
```

## 开发检查清单

### 开发新CRUD接口时必须检查：

- [ ] 是否实现了5个基础CRUD接口
- [ ] 是否正确使用了路由参数规范
- [ ] 如果有Status字段，是否提供了GetStatus接口
- [ ] 是否评估了是否需要ListAll接口
- [ ] 如果提供ListAll接口，是否符合使用条件
- [ ] 是否在service和controller的init.go中注册
- [ ] 是否在router.go中添加路由

### 代码审查时必须检查：

- [ ] ListAll接口是否被滥用
- [ ] Status字段是否有对应的GetStatus接口
- [ ] 路由参数是否符合规范
- [ ] 错误处理是否使用errors.New()
- [ ] 是否使用了正确的vo包返回格式
