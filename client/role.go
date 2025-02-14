package client

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/orbit-center/sdk/models"
)

// RoleService 实现了角色相关的所有操作
type RoleService struct {
	Client *Client
}

// 实现 interfaces.RoleService 接口
func (s *RoleService) GetRoleList() ([]models.Role, error) {
	data, err := s.Client.doRequest("POST", "/roles/list", nil)
	if err != nil {
		return nil, err
	}

	var resp models.Response
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal response failed: %w", err)
	}

	var roles []models.Role
	if err := mapstructure.Decode(resp.Data, &roles); err != nil {
		return nil, fmt.Errorf("decode roles failed: %w", err)
	}

	return roles, nil
}

// AssignRole 为用户分配角色
func (s *RoleService) AssignRole(userID, roleID int64) error {
	req := map[string]interface{}{
		"user_id": userID,
		"role_id": roleID,
	}
	_, err := s.Client.doRequest("POST", "/roles/assign", req)
	return err
}

// GetRolePermissions 获取角色权限
func (s *RoleService) GetRolePermissions(roleID int64) ([]models.Permission, error) {
	req := map[string]interface{}{
		"role_id": roleID,
	}
	data, err := s.Client.doRequest("POST", "/roles/permissions", req)
	if err != nil {
		return nil, err
	}

	var resp models.Response
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal response failed: %w", err)
	}

	var permissions []models.Permission
	if err := mapstructure.Decode(resp.Data, &permissions); err != nil {
		return nil, fmt.Errorf("decode permissions failed: %w", err)
	}

	return permissions, nil
}
