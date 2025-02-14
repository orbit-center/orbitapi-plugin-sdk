// Package models 定义所有服务共用的数据模型
package models

// Response 通用响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ListResponse 列表响应结构
type ListResponse[T any] struct {
	Total int64 `json:"total"`
	List  []T   `json:"list"`
}

// RoleListResponse 角色列表响应
type RoleListResponse = ListResponse[Role]

// MenuListResponse 菜单列表响应
type MenuListResponse = ListResponse[Menu]

// DictListResponse 字典列表响应
type DictListResponse = ListResponse[Dict]

// UserListResponse 用户列表响应
type UserListResponse = ListResponse[User]

// User 用户信息
type User struct {
	ID       int64  `json:"id"`
	Account  string `json:"account"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar"`
	Status   int    `json:"status"`
}

// Role 角色信息
type Role struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

// Menu 菜单信息
type Menu struct {
	ID         int64  `json:"id"`
	ParentID   int64  `json:"parent_id"`
	Name       string `json:"name"`
	Path       string `json:"path"`
	Icon       string `json:"icon"`
	Sort       int    `json:"sort"`
	Type       int    `json:"type"`
	Permission string `json:"permission"`
	Component  string `json:"component"`
	Hidden     bool   `json:"hidden"`
}

// Dict 字典信息
type Dict struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Desc   string `json:"desc"`
	Sort   int    `json:"sort"`
	Status int    `json:"status"`
	Type   string `json:"type"`
}

// DictType 字典类型
type DictType struct {
	ID     int64  `json:"id"`
	DictID int64  `json:"dict_id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Status int    `json:"status"`
}

// DictContent 字典内容
type DictContent struct {
	ID     int64  `json:"id"`
	TypeID int64  `json:"type_id"`
	Label  string `json:"label"`
	Value  string `json:"value"`
	Sort   int    `json:"sort"`
	Status int    `json:"status"`
}

// Permission 权限信息
type Permission struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

// UpdateUserRequest 更新用户请求
type UpdateUserRequest struct {
	ID       int64  `json:"id"`
	Nickname string `json:"nickname,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
	Status   int    `json:"status,omitempty"`
}
