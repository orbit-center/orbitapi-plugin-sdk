---
title: 基础版
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.28"

---

# 基础版

Base URLs:

* <a href="http://127.0.0.1:8081/api/v1">开发环境: http://127.0.0.1:8081/api/v1</a>

* <a href="http://orbit-son.api.dev.vm1776.top/api/v1">测试环境: http://orbit-son.api.dev.vm1776.top/api/v1</a>

# Authentication

# 1. 用户模块

## POST 1. 注册新用户

POST /users/register

> Body 请求参数

```json
"{\r\n    \"account\": \"admin\",        // 账号(必填)\r\n    \"password\": \"123456\",     // 密码(必填,最小6位)\r\n    \"nickname\": \"测试用户\",    // 昵称(选填)\r\n    \"email\": \"admin@example.com\", // 邮箱(选填)\r\n    \"phone\": \"13800138000\"    // 手机号(选填)\r\n}"
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 2. 用户登录

POST /users/login

> Body 请求参数

```json
{
  "account": "admin",
  "password": "123456"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 3. 获取用户列表

POST /users/list

> Body 请求参数

```json
{
  "page": 1,
  "pageSize": 10,
  "keyword": "",
  "deptId": -1
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 4. 更新用户信息

POST /users/update

> Body 请求参数

```json
"{\r\n    \"id\": 4, // 用户ID(必填)\r\n    \"nickname\": \"test3\", // 昵称(选填)\r\n    \"email\": \"test3@163.com\", // 邮箱(选填)\r\n    \"phone\": \"13800138000\", // 手机号(选填)\r\n    \"deptId\": 3, // 部门ID(选填)\r\n    \"status\": 1, // 状态(选填):0=禁用,1=启用\r\n    \"roles\": [\r\n        \"user\",\r\n        \"test\"\r\n    ] // 角色编码数组(选填)\r\n}"
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 4.1 新增用户

POST /users/add

> Body 请求参数

```json
{
  "account": "test3",
  "password": "123456",
  "nickname": "test3",
  "email": "test3@163.com",
  "deptId": 3,
  "roles": [
    "user",
    "test"
  ]
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 5. 删除用户

POST /users/delete

> Body 请求参数

```json
{
  "user_id": 12345
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 6. 修改密码

POST /users/password

> Body 请求参数

```json
"{\r\n    \"oldPassword\": \"123456\",  // 原密码(必填)\r\n    \"newPassword\": \"654321\"   // 新密码(必填,最小6位)\r\n}"
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 7. 获取用户角色列表

POST /users/roles

> Body 请求参数

```json
{
  "user_id": 1
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 8. 分配用户角色

POST /users/assign-roles

> Body 请求参数

```json
{
  "user_id": 1,
  "role_ids": [
    1,
    2
  ]
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 9. 获取个人信息

POST /users/profile

> Body 请求参数

```json
{}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 10. 更新个人信息

POST /users/profile/update

> Body 请求参数

```json
"{\r\n    \"nickname\": \"新昵称\",         // 选填，最大长度50\r\n    \"email\": \"new@example.com\",  // 选填，需要验证邮箱格式\r\n    \"phone\": \"13800138000\",      // 选填，11位手机号\r\n    \"avatar\": \"avatar_url\",      // 选填，头像URL\r\n    \"gender\": 1                  // 选填，0=未知,1=男,2=女\r\n}"
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 11. 修改密码

POST /users/password/change

> Body 请求参数

```json
"{\r\n    \"old_password\": \"123456\",      // 必填，原密码\r\n    \"new_password\": \"654321\",      // 必填，新密码\r\n    \"confirm_password\": \"654321\"   // 必填，确认密码\r\n}"
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 12. 上传头像

POST /users/avatar/upload

> Body 请求参数

```yaml
file: ""

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|
|» file|body|string(binary)| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

# 2. 角色模块

## POST 4. 获取角色列表

POST /roles/list

> Body 请求参数

```json
{}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 1. 创建角色

POST /roles/create

> Body 请求参数

```json
"{\r\n    \"name\": \"管理员\",    // 角色名称(必填)\r\n    \"desc\": \"系统管理员\"  // 角色描述(选填)\r\n}"
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 2. 更新角色

POST /roles/update

> Body 请求参数

```json
"{\r\n    \"id\": 1,           // 角色ID(必填)\r\n    \"name\": \"管理员\",    // 角色名称(必填)\r\n    \"desc\": \"系统管理员\"  // 角色描述(选填)\r\n}"
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 3. 删除角色

POST /roles/delete

> Body 请求参数

```json
"{\r\n    \"id\": 1  // 角色ID(必填)\r\n}"
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 5. 分配菜单权限

POST /roles/assign

> Body 请求参数

```json
"{\r\n    \"role_id\": 1,           // 角色ID(必填)\r\n    \"menu_ids\": [1, 2, 3]   // 菜单ID列表(必填)\r\n}"
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

# 3. 菜单模块

## POST 4. 获取菜单列表

POST /menus/list

> Body 请求参数

```json
{}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 5. 获取用户菜单

POST /menus/user

> Body 请求参数

```json
{}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 1. 新增菜单

POST /menus/create

> Body 请求参数

```json
"{\r\n    \"parent_id\": 0,        // 父菜单ID(必填)，0表示一级菜单\r\n    \"name\": \"系统管理\",     // 菜单名称(必填)\r\n    \"path\": \"/system\",     // 路由路径\r\n    \"icon\": \"setting\",     // 菜单图标\r\n    \"sort\": 1,            // 排序，数值越小越靠前\r\n    \"type\": 1,            // 菜单类型(必填) 1:主菜单 2:二级菜单 3:页面 4:按钮\r\n    \"permission\": \"system:view\",  // 权限标识\r\n    \"component\": \"system/index\",  // 组件路径\r\n    \"hidden\": false       // 是否隐藏\r\n}"
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 2. 更新菜单

POST /menus/update

> Body 请求参数

```json
"{\r\n    \"id\": 1,             // 菜单ID(必填)\r\n    \"parent_id\": 0,      // 父菜单ID\r\n    \"name\": \"系统管理\",   // 菜单名称\r\n    \"path\": \"/system\",   // 路由路径\r\n    \"icon\": \"setting\",   // 菜单图标\r\n    \"sort\": 1,          // 排序\r\n    \"type\": 1,          // 菜单类型\r\n    \"permission\": \"system:view\",  // 权限标识\r\n    \"component\": \"system/index\",  // 组件路径\r\n    \"hidden\": false     // 是否隐藏\r\n}"
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 3. 删除菜单

POST /menus/delete

> Body 请求参数

```json
"{\r\n    \"id\": 1  // 菜单ID(必填)\r\n}"
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

# 4. 部门管理

## POST 4. 获取部门树

POST /depts/tree

> Body 请求参数

```json
{}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 1. 创建部门

POST /depts/create

> Body 请求参数

```json
"{\r\n    \"parent_id\": null,     // 父级部门ID(选填)，null表示顶级部门\r\n    \"name\": \"技术部\",      // 部门名称(必填)\r\n    \"order\": 1,           // 排序(选填)，默认0\r\n    \"leader\": \"张三\",     // 部门负责人(选填)\r\n    \"phone\": \"13800138000\", // 联系电话(选填)\r\n    \"email\": \"tech@example.com\", // 邮箱(选填)\r\n    \"status\": 1           // 状态(选填) 1:正常 0:禁用\r\n}"
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 2. 更新部门

POST /depts/update

> Body 请求参数

```json
"{\r\n    \"id\": 1,             // 部门ID(必填)\r\n    \"parent_id\": null,   // 父级部门ID(选填)\r\n    \"name\": \"技术部\",    // 部门名称(必填)\r\n    \"order\": 1,         // 排序(选填)\r\n    \"leader\": \"张三\",   // 部门负责人(选填)\r\n    \"phone\": \"13800138000\", // 联系电话(选填)\r\n    \"email\": \"tech@example.com\", // 邮箱(选填)\r\n    \"status\": 1         // 状态(选填)\r\n}"
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 3. 删除部门

POST /dept/batch-delete

> Body 请求参数

```json
"{\r\n    \"id\": 1  // 部门ID(必填)\r\n}"
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

# 5. 数据字典

## POST 1. 创建字典

POST /dict/create

> Body 请求参数

```json
"{\r\n    \"name\": \"系统状态\",        // 字典名称\r\n    \"code\": \"system_status\",  // 字典编码\r\n    \"desc\": \"系统通用状态\",    // 描述\r\n    \"sort\": 1,               // 排序(可选)\r\n    \"status\": 1              // 状态(1=启用,0=禁用)\r\n}"
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 5. 创建字典类型

POST /dict/type/create

> Body 请求参数

```json
{
  "parent_id": "1ac32a63-265c-4823-a419-5365e103a8d4",
  "name": "技术测试",
  "code": "dev_test",
  "desc": "技术测试",
  "sort": 1,
  "status": 1
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 10. 更新字典内容

POST /dict/content/create

> Body 请求参数

```json
"{\r\n    \"id\": 1,                 // 内容ID\r\n    \"label\": \"正常\",         // 显示标签\r\n    \"value\": \"1\",           // 实际值\r\n    \"desc\": \"正常状态\",      // 描述\r\n    \"sort\": 1,              // 排序\r\n    \"status\": 1             // 状态\r\n}"
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 6. 更新字典类型

POST /dict/type/update

> Body 请求参数

```json
"{\r\n    \"id\": 1,                  // 类型ID\r\n    \"name\": \"用户状态\",        // 类型名称\r\n    \"code\": \"user_status\",    // 类型编码\r\n    \"desc\": \"用户状态定义\",    // 描述\r\n    \"sort\": 1                 // 排序\r\n}"
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 2. 更新字典

POST /dict/update

> Body 请求参数

```json
"{\r\n    \"id\": 1,                 // 字典ID\r\n    \"name\": \"系统状态\",        // 字典名称\r\n    \"code\": \"system_status\",  // 字典编码\r\n    \"desc\": \"系统通用状态\",    // 描述\r\n    \"sort\": 1,               // 排序\r\n    \"status\": 1              // 状态\r\n}"
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 3. 删除字典

POST /dict/delete

> Body 请求参数

```json
"{\r\n    \"id\": 1  // 部门ID(必填)\r\n}"
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 4. 获取字典表

POST /dict/list

> Body 请求参数

```json
"{\r\n    // \"query\": \"系统\"  // 可选，搜索关键字\r\n}"
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 4.1 获取字典类型表

POST /dict/type/list

> Body 请求参数

```json
"{\r\n    \"dict_id\": 1  // 必填，字典ID\r\n}"
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 4.2 获取字典内容表

POST /dict/content/list

> Body 请求参数

```json
"{\r\n    \"type_id\": 1  // 必填，字典类型ID\r\n}"
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

# 数据模型

