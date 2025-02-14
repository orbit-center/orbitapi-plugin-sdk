package integration

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/orbit-center/sdk/client"
	"github.com/orbit-center/sdk/models"
)

func TestFullIntegration(t *testing.T) {
	// 创建模拟服务器
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/users/profile":
			json.NewEncoder(w).Encode(models.Response{
				Code: 0,
				Data: models.User{
					ID:      1,
					Account: "test",
				},
			})
		case "/roles/list":
			json.NewEncoder(w).Encode(models.Response{
				Code: 0,
				Data: models.RoleListResponse{
					Total: 1,
					List: []models.Role{
						{ID: 1, Name: "admin"},
					},
				},
			})
		case "/menus/list":
			json.NewEncoder(w).Encode(models.Response{
				Code: 0,
				Data: models.MenuListResponse{
					Total: 1,
					List: []models.Menu{
						{ID: 1, Name: "系统管理"},
					},
				},
			})
		case "/dict/list":
			json.NewEncoder(w).Encode(models.Response{
				Code: 0,
				Data: models.DictListResponse{
					Total: 1,
					List: []models.Dict{
						{ID: 1, Name: "系统状态"},
					},
				},
			})
		case "/not-found":
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(models.Response{
				Code:    404,
				Message: "not found",
			})
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	// 创建客户端
	c := client.NewClient(ts.URL, "test-token")

	// 测试用户服务
	t.Run("UserService", func(t *testing.T) {
		user, err := c.User().GetUserInfo()
		if err != nil {
			t.Fatalf("GetUserInfo failed: %v", err)
		}
		if user.ID != 1 || user.Account != "test" {
			t.Errorf("unexpected user: %+v", user)
		}
	})

	t.Run("RoleService", func(t *testing.T) {
		roles, err := c.Role().GetRoleList()
		if err != nil {
			t.Fatalf("GetRoleList failed: %v", err)
		}
		if len(roles) != 1 {
			t.Errorf("expected 1 role, got %d", len(roles))
		}
	})

	t.Run("MenuService", func(t *testing.T) {
		menus, err := c.Menu().GetMenuList()
		if err != nil {
			t.Fatalf("GetMenuList failed: %v", err)
		}
		if len(menus) != 1 {
			t.Errorf("expected 1 menu, got %d", len(menus))
		}
	})

	t.Run("DictService", func(t *testing.T) {
		dicts, err := c.Dict().GetDictList()
		if err != nil {
			t.Fatalf("GetDictList failed: %v", err)
		}
		if len(dicts) != 1 {
			t.Errorf("expected 1 dict, got %d", len(dicts))
		}
	})

	t.Run("ErrorHandling", func(t *testing.T) {
		_, err := c.DoRequest("GET", "/not-found", nil)
		if err == nil {
			t.Error("expected error for not found")
		}
	})

	// ... 添加更多集成测试场景
}
