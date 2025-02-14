package client

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/orbit-center/sdk/models"
)

// UserService 实现了用户相关的所有操作
type UserService struct {
	Client *Client
}

// GetUserInfo 获取当前用户信息
func (s *UserService) GetUserInfo() (*models.User, error) {
	data, err := s.Client.doRequest("POST", "/users/profile", nil)
	if err != nil {
		return nil, err
	}

	var resp models.Response
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal response failed: %w", err)
	}

	var user models.User
	if err := mapstructure.Decode(resp.Data, &user); err != nil {
		return nil, fmt.Errorf("decode user failed: %w", err)
	}

	return &user, nil
}

// GetUserList 获取用户列表
func (s *UserService) GetUserList(page, pageSize int, keyword string) (*models.UserListResponse, error) {
	req := map[string]interface{}{
		"page":      page,
		"page_size": pageSize,
		"keyword":   keyword,
	}

	data, err := s.Client.doRequest("POST", "/users/list", req)
	if err != nil {
		return nil, err
	}

	var resp models.Response
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal response failed: %w", err)
	}

	var userList models.UserListResponse
	if err := mapstructure.Decode(resp.Data, &userList); err != nil {
		return nil, fmt.Errorf("decode user list failed: %w", err)
	}

	return &userList, nil
}

// UpdateUser 更新用户信息
func (s *UserService) UpdateUser(req *models.UpdateUserRequest) error {
	_, err := s.Client.doRequest("POST", "/users/profile/update", req)
	return err
}

// CheckPermission 检查用户权限
func (s *UserService) CheckPermission(permission string) bool {
	req := map[string]interface{}{
		"permission": permission,
	}
	_, err := s.Client.doRequest("POST", "/users/check-permission", req)
	return err == nil
}

// 实现方法...
