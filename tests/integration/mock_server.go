package integration

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/orbit-center/sdk/models"
)

func newMockServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/users/profile":
			json.NewEncoder(w).Encode(models.Response{
				Code: 0,
				Data: models.User{ID: 1, Account: "admin"},
			})
		case "/roles/list":
			json.NewEncoder(w).Encode(models.Response{
				Code: 0,
				Data: []models.Role{{ID: 1, Name: "admin"}},
			})
		case "/menus/list":
			json.NewEncoder(w).Encode(models.Response{
				Code: 0,
				Data: []models.Menu{{ID: 1, Name: "系统管理"}},
			})
		case "/dicts/list":
			json.NewEncoder(w).Encode(models.Response{
				Code: 0,
				Data: []models.Dict{{ID: 1, Name: "系统状态"}},
			})
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
}
