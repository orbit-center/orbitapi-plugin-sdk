// Package services 提供各种服务实现
package services

import (
	"fmt"

	"github.com/orbit-center/sdk/interfaces"
	"github.com/orbit-center/sdk/models"
)

// MenuServiceImpl 实现了 interfaces.MenuService 接口
type MenuServiceImpl struct {
	client interfaces.Client
}

// NewMenuService 创建菜单服务实例
func NewMenuService(client interfaces.Client) interfaces.MenuService {
	return &MenuServiceImpl{client: client}
}

// GetMenuList 获取菜单列表
func (s *MenuServiceImpl) GetMenuList() ([]models.Menu, error) {
	resp, err := s.client.DoRequest("GET", "/menus/list", nil)
	if err != nil {
		return nil, fmt.Errorf("get menus failed: %w", err)
	}

	var listResp models.MenuListResponse
	if err := resp.DecodeData(&listResp); err != nil {
		return nil, fmt.Errorf("decode menus failed: %w", err)
	}

	return listResp.List, nil
}

// AddMenu 添加菜单
func (s *MenuServiceImpl) AddMenu(menu *models.Menu) error {
	_, err := s.client.DoRequest("POST", "/menus/create", menu)
	return err
}

// UpdateMenu 更新菜单
func (s *MenuServiceImpl) UpdateMenu(menu *models.Menu) error {
	_, err := s.client.DoRequest("POST", "/menus/update", menu)
	return err
}

// DeleteMenu 删除菜单
func (s *MenuServiceImpl) DeleteMenu(menuID int64) error {
	req := map[string]interface{}{
		"id": menuID,
	}
	_, err := s.client.DoRequest("POST", "/menus/delete", req)
	return err
}

// CheckMenuPermission 检查菜单权限
func (s *MenuServiceImpl) CheckMenuPermission(menuID int64) bool {
	req := map[string]interface{}{
		"menu_id": menuID,
	}
	_, err := s.client.DoRequest("POST", "/menus/check-permission", req)
	return err == nil
}
