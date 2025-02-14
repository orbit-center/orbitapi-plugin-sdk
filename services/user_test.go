package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/orbit-center/sdk/client"
	"github.com/orbit-center/sdk/models"
)

func TestGetUserInfo(t *testing.T) {
	// 创建测试服务器
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := models.Response{
			Code: 0,
			Data: models.User{
				ID:       1,
				Account:  "test",
				Nickname: "测试用户",
			},
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer ts.Close()

	// 创建客户端和服务
	c := client.NewClient(ts.URL, "test-token")
	userSrv := NewUserService(c)

	// 测试获取用户信息
	user, err := userSrv.GetUserInfo()
	if err != nil {
		t.Fatalf("GetUserInfo failed: %v", err)
	}
	if user.Account != "test" {
		t.Errorf("expected account 'test', got %s", user.Account)
	}
}

func TestUserService_GetUserList(t *testing.T) {
	c := NewUserService(&mockClient{
		response: []byte(`{"code":0,"data":{"total":1,"list":[{"id":1,"account":"test"}]}}`),
	})

	list, err := c.GetUserList(1, 10, "")
	if err != nil {
		t.Fatalf("GetUserList failed: %v", err)
	}
	if list.Total != 1 {
		t.Errorf("expected total 1, got %d", list.Total)
	}
}

func TestUserService_CheckPermission(t *testing.T) {
	c := NewUserService(&mockClient{
		response: []byte(`{"code":0,"data":true}`),
	})

	hasPermission := c.CheckPermission("system:user:view")
	if !hasPermission {
		t.Error("expected to have permission")
	}
}

func TestUserService_Error(t *testing.T) {
	c := NewUserService(&mockClient{
		err: fmt.Errorf("mock error"),
	})

	// 测试获取用户信息失败
	_, err := c.GetUserInfo()
	if err == nil {
		t.Error("expected error, got nil")
	}

	// 测试更新用户信息失败
	err = c.UpdateUser(&models.UpdateUserRequest{
		Nickname: "test",
	})
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestUserService_UpdateUser(t *testing.T) {
	c := NewUserService(&mockClient{
		response: []byte(`{"code":0}`),
	})

	err := c.UpdateUser(&models.UpdateUserRequest{
		Nickname: "新昵称",
		Email:    "test@example.com",
	})
	if err != nil {
		t.Fatalf("UpdateUser failed: %v", err)
	}
}
