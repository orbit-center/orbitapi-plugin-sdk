package integration

import (
	"testing"

	"github.com/orbit-center/sdk/client"
)

func TestClientIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	// 创建 mock 服务器
	ts := newMockServer()
	defer ts.Close()

	c := client.NewClient(ts.URL, "test-token")
	c.Use(client.WithRetry(3), client.WithLogging())

	// 测试用户服务
	t.Run("UserService", func(t *testing.T) {
		user, err := c.User().GetUserInfo()
		if err != nil {
			t.Fatalf("GetUserInfo failed: %v", err)
		}
		if user.ID == 0 {
			t.Error("expected non-zero user ID")
		}
	})

	// 测试角色服务
	t.Run("RoleService", func(t *testing.T) {
		roles, err := c.Role().GetRoleList()
		if err != nil {
			t.Fatalf("GetRoleList failed: %v", err)
		}
		if len(roles) == 0 {
			t.Error("expected non-empty role list")
		}
	})
}
