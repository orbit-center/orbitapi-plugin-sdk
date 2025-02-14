// Package services 提供各种服务实现
package services

import (
	"fmt"

	"github.com/orbit-center/sdk/interfaces"
	"github.com/orbit-center/sdk/models"
)

// UserServiceImpl 实现了 interfaces.UserService 接口
type UserServiceImpl struct {
	client interfaces.Client
}

// NewUserService 创建用户服务实例
func NewUserService(client interfaces.Client) interfaces.UserService {
	return &UserServiceImpl{client: client}
}

// GetUserInfo 获取用户信息
func (s *UserServiceImpl) GetUserInfo() (*models.User, error) {
	resp, err := s.client.DoRequest("GET", "/users/profile", nil)
	if err != nil {
		return nil, fmt.Errorf("get user info failed: %w", err)
	}

	var user models.User
	if err := resp.DecodeData(&user); err != nil {
		return nil, fmt.Errorf("decode user failed: %w", err)
	}

	return &user, nil
}

// GetUserList 获取用户列表
func (s *UserServiceImpl) GetUserList(page, pageSize int, keyword string) (*models.UserListResponse, error) {
	req := map[string]interface{}{
		"page":      page,
		"page_size": pageSize,
		"keyword":   keyword,
	}
	resp, err := s.client.DoRequest("GET", "/users/list", req)
	if err != nil {
		return nil, fmt.Errorf("get user list failed: %w", err)
	}

	var userList models.UserListResponse
	if err := resp.DecodeData(&userList); err != nil {
		return nil, fmt.Errorf("decode user list failed: %w", err)
	}

	return &userList, nil
}

// UpdateUser 更新用户信息
func (s *UserServiceImpl) UpdateUser(req *models.UpdateUserRequest) error {
	_, err := s.client.DoRequest("POST", "/users/update", req)
	return err
}

// CheckPermission 检查权限
func (s *UserServiceImpl) CheckPermission(permission string) bool {
	req := map[string]interface{}{
		"permission": permission,
	}
	_, err := s.client.DoRequest("POST", "/users/check-permission", req)
	return err == nil
}
