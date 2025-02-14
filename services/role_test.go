package services

import (
	"fmt"
	"testing"
)

func TestRoleService_GetRoleList(t *testing.T) {
	c := NewRoleService(&mockClient{
		response: []byte(`{
			"code": 0,
			"data": {
				"total": 1,
				"list": [
					{"id": 1, "name": "admin"}
				]
			}
		}`),
	})

	roles, err := c.GetRoleList()
	if err != nil {
		t.Fatalf("GetRoleList failed: %v", err)
	}
	if len(roles) != 1 {
		t.Errorf("expected 1 role, got %d", len(roles))
	}
	if roles[0].Name != "admin" {
		t.Errorf("expected role name 'admin', got %s", roles[0].Name)
	}
}

func TestRoleService_GetRolePermissions(t *testing.T) {
	c := NewRoleService(&mockClient{
		response: []byte(`{"code":0,"data":[{"id":1,"code":"system:user:view"}]}`),
	})

	perms, err := c.GetRolePermissions(1)
	if err != nil {
		t.Fatalf("GetRolePermissions failed: %v", err)
	}
	if len(perms) != 1 {
		t.Errorf("expected 1 permission, got %d", len(perms))
	}
}

func TestRoleService_Error(t *testing.T) {
	c := NewRoleService(&mockClient{
		err: fmt.Errorf("mock error"),
	})

	// 测试获取角色列表失败
	_, err := c.GetRoleList()
	if err == nil {
		t.Error("expected error, got nil")
	}

	// 测试分配角色失败
	err = c.AssignRole(1, 1)
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestRoleService_AssignRole(t *testing.T) {
	c := NewRoleService(&mockClient{
		response: []byte(`{"code":0}`),
	})

	err := c.AssignRole(1, 1)
	if err != nil {
		t.Fatalf("AssignRole failed: %v", err)
	}
}
