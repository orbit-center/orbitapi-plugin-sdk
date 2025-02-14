// MenuService 实现了菜单相关的所有操作
package client

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/orbit-center/sdk/models"
)

type MenuService struct {
	Client *Client
}

// GetMenuList 获取菜单列表
func (s *MenuService) GetMenuList() ([]models.Menu, error) {
	data, err := s.Client.doRequest("POST", "/menus/list", nil)
	if err != nil {
		return nil, err
	}

	var resp models.Response
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal response failed: %w", err)
	}

	var menus []models.Menu
	if err := mapstructure.Decode(resp.Data, &menus); err != nil {
		return nil, fmt.Errorf("decode menus failed: %w", err)
	}

	return menus, nil
}

// AddMenu 添加菜单
func (s *MenuService) AddMenu(menu *models.Menu) error {
	_, err := s.Client.doRequest("POST", "/menus/create", menu)
	return err
}

// UpdateMenu 更新菜单
func (s *MenuService) UpdateMenu(menu *models.Menu) error {
	_, err := s.Client.doRequest("POST", "/menus/update", menu)
	return err
}

// DeleteMenu 删除菜单
func (s *MenuService) DeleteMenu(menuID int64) error {
	req := map[string]interface{}{
		"id": menuID,
	}
	_, err := s.Client.doRequest("POST", "/menus/delete", req)
	return err
}

// CheckMenuPermission 检查菜单权限
func (s *MenuService) CheckMenuPermission(menuID int64) bool {
	req := map[string]interface{}{
		"menu_id": menuID,
	}
	_, err := s.Client.doRequest("POST", "/menus/check-permission", req)
	return err == nil
}
