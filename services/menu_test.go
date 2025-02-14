package services

import (
	"fmt"
	"testing"

	"github.com/orbit-center/sdk/models"
)

func TestMenuService_GetMenuList(t *testing.T) {
	c := NewMenuService(&mockClient{
		response: []byte(`{"code":0,"data":[{"id":1,"name":"系统管理"}]}`),
	})

	menus, err := c.GetMenuList()
	if err != nil {
		t.Fatalf("GetMenuList failed: %v", err)
	}
	if len(menus) != 1 {
		t.Errorf("expected 1 menu, got %d", len(menus))
	}
}

func TestMenuService_AddMenu(t *testing.T) {
	c := NewMenuService(&mockClient{
		response: []byte(`{"code":0}`),
	})

	err := c.AddMenu(&models.Menu{
		Name:     "测试菜单",
		Path:     "/test",
		ParentID: 0,
		Sort:     1,
	})
	if err != nil {
		t.Fatalf("AddMenu failed: %v", err)
	}
}

func TestMenuService_UpdateMenu(t *testing.T) {
	c := NewMenuService(&mockClient{
		response: []byte(`{"code":0}`),
	})

	err := c.UpdateMenu(&models.Menu{
		ID:   1,
		Name: "更新菜单",
	})
	if err != nil {
		t.Fatalf("UpdateMenu failed: %v", err)
	}
}

func TestMenuService_DeleteMenu(t *testing.T) {
	c := NewMenuService(&mockClient{
		response: []byte(`{"code":0}`),
	})

	err := c.DeleteMenu(1)
	if err != nil {
		t.Fatalf("DeleteMenu failed: %v", err)
	}
}

func TestMenuService_Error(t *testing.T) {
	c := NewMenuService(&mockClient{
		err: fmt.Errorf("mock error"),
	})

	// 测试获取菜单列表失败
	_, err := c.GetMenuList()
	if err == nil {
		t.Error("expected error, got nil")
	}

	// 测试添加菜单失败
	err = c.AddMenu(&models.Menu{Name: "test"})
	if err == nil {
		t.Error("expected error, got nil")
	}
}
