// Package services 提供各种服务实现
package services

import (
	"fmt"

	"github.com/orbit-center/sdk/interfaces"
	"github.com/orbit-center/sdk/models"
)

// RoleServiceImpl 实现了 interfaces.RoleService 接口
type RoleServiceImpl struct {
	client interfaces.Client
}

// NewRoleService 创建角色服务实例
func NewRoleService(client interfaces.Client) interfaces.RoleService {
	return &RoleServiceImpl{client: client}
}

// GetRoleList 获取角色列表
func (s *RoleServiceImpl) GetRoleList() ([]models.Role, error) {
	resp, err := s.client.DoRequest("GET", "/roles/list", nil)
	if err != nil {
		return nil, fmt.Errorf("get roles failed: %w", err)
	}

	var listResp models.RoleListResponse
	if err := resp.DecodeData(&listResp); err != nil {
		return nil, fmt.Errorf("decode roles failed: %w", err)
	}

	return listResp.List, nil
}

// GetRolePermissions 获取角色权限
func (s *RoleServiceImpl) GetRolePermissions(roleID int64) ([]models.Permission, error) {
	req := map[string]interface{}{
		"role_id": roleID,
	}
	resp, err := s.client.DoRequest("GET", "/roles/permissions", req)
	if err != nil {
		return nil, fmt.Errorf("get role permissions failed: %w", err)
	}

	var permissions []models.Permission
	if err := resp.DecodeData(&permissions); err != nil {
		return nil, fmt.Errorf("decode permissions failed: %w", err)
	}

	return permissions, nil
}

// AssignRole 分配角色
func (s *RoleServiceImpl) AssignRole(userID, roleID int64) error {
	req := map[string]interface{}{
		"user_id": userID,
		"role_id": roleID,
	}
	_, err := s.client.DoRequest("POST", "/roles/assign", req)
	return err
}
