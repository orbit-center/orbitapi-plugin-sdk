// Package interfaces 定义了所有服务接口
package interfaces

import "github.com/orbit-center/sdk/models"

// UserService 用户服务接口定义了与用户相关的所有操作
type UserService interface {
	GetUserInfo() (*models.User, error)
	GetUserList(page, pageSize int, keyword string) (*models.UserListResponse, error)
	UpdateUser(req *models.UpdateUserRequest) error
	CheckPermission(permission string) bool
}

// RoleService 角色服务接口定义了与角色相关的所有操作
type RoleService interface {
	GetRoleList() ([]models.Role, error)
	GetRolePermissions(roleID int64) ([]models.Permission, error)
	AssignRole(userID, roleID int64) error
}

// MenuService 菜单服务接口定义了与菜单相关的所有操作
type MenuService interface {
	GetMenuList() ([]models.Menu, error)
	CheckMenuPermission(menuID int64) bool
	AddMenu(menu *models.Menu) error
	UpdateMenu(menu *models.Menu) error
	DeleteMenu(menuID int64) error
}

// DictService 字典服务接口定义了与数据字典相关的所有操作
type DictService interface {
	GetDictList() ([]models.Dict, error)
	GetDictTypes(dictID int64) ([]models.DictType, error)
	GetDictContents(typeID int64) ([]models.DictContent, error)
}
