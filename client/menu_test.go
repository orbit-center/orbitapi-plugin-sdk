package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/orbit-center/sdk/models"
)

func TestGetMenuList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := models.Response{
			Code:    0,
			Message: "success",
			Data: []models.Menu{
				{
					ID:   1,
					Name: "系统管理",
					Path: "/system",
				},
			},
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer ts.Close()

	c := NewClient(ts.URL, "test-token")
	menuSrv := &MenuService{Client: c}

	menus, err := menuSrv.GetMenuList()
	if err != nil {
		t.Fatalf("GetMenuList failed: %v", err)
	}
	if len(menus) != 1 {
		t.Errorf("expected 1 menu, got %d", len(menus))
	}
}
