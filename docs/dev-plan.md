# 子站点 SDK 开发方案

## 一、目标说明

子站点 SDK 的主要目标是 **为插件提供基础功能的接口**，专注于与子站点的基础数据交互，包括用户、角色、菜单、数据字典等基本功能。生命周期管理等复杂功能由 Go Micro 处理。

## 二、功能清单

### 1. 用户管理
- 获取当前用户信息
- 用户权限验证
- 获取用户列表
- 更新用户信息

### 2. 角色管理
- 获取角色列表
- 获取角色权限
- 角色分配与修改

### 3. 菜单管理
- 获取菜单列表
- 菜单权限控制
- 菜单的增删改

### 4. 数据字典管理
- 获取数据字典列表
- 数据字典的增删改
- 获取字典类型
- 获取字典内容

## 三、技术方案

### 1. 基础客户端
```go
type Client struct {
    baseURL    string
    httpClient *http.Client
    token      string 
}
```

### 2. 核心服务
```go
type UserService interface {
    GetUserInfo() (*User, error)
    GetUserList(page, pageSize int, keyword string) (*UserListResponse, error)
    UpdateUser(req UpdateUserRequest) error
    CheckPermission(permission string) bool
}

type RoleService interface {
    GetRoleList() ([]Role, error)
    GetRolePermissions(roleID int64) ([]Permission, error)
    AssignRole(userID, roleID int64) error
}

type MenuService interface {
    GetMenuList() ([]Menu, error)
    CheckMenuPermission(menuID int64) bool
    AddMenu(menu *Menu) error
    UpdateMenu(menu *Menu) error
    DeleteMenu(menuID int64) error
}

type DictService interface {
    GetDictList() ([]Dict, error)
    GetDictTypes(dictID int64) ([]DictType, error)
    GetDictContents(typeID int64) ([]DictContent, error)
}
```

## 四、开发计划

### 阶段一：基础功能开发（3-4天）
1. 实现基础 HTTP 客户端
2. 实现用户服务
3. 实现角色服务

### 阶段二：扩展功能开发（3-4天）
1. 实现菜单服务
2. 实现数据字典服务
3. 完善错误处理

### 阶段三：测试与文档（2-3天）
1. 编写单元测试
2. 编写使用文档
3. 提供示例代码

## 五、使用示例

```go
// 创建客户端
client := NewClient("http://api.example.com", "your-token")

// 获取用户信息
userSrv := client.User()
userInfo, err := userSrv.GetUserInfo()
if err != nil {
    log.Fatal(err)
}

// 检查权限
if !userSrv.CheckPermission("menu:view") {
    log.Fatal("permission denied")
}

// 获取菜单列表
menuSrv := client.Menu()
menus, err := menuSrv.GetMenuList()
if err != nil {
    log.Fatal(err)
}
```

## 六、注意事项

1. SDK 不处理认证逻辑，token 由外部传入
2. 错误处理采用标准 error 接口
3. 所有接口都是同步调用
4. 不包含数据库直接操作，统一通过 HTTP API 交互
