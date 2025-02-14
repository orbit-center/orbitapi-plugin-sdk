package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/orbit-center/sdk/models"
)

func TestRoleService_GetRoleList(t *testing.T) {
	tests := []struct {
		name     string
		response interface{}
		wantErr  bool
	}{
		{
			name: "success",
			response: models.Response{
				Code:    0,
				Message: "success",
				Data: []models.Role{
					{ID: 1, Name: "admin"},
				},
			},
			wantErr: false,
		},
		{
			name: "error_response",
			response: models.Response{
				Code:    500,
				Message: "internal error",
				Data:    nil,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if tt.wantErr {
					w.WriteHeader(http.StatusInternalServerError)
				}
				json.NewEncoder(w).Encode(tt.response)
			}))
			defer ts.Close()

			c := NewClient(ts.URL, "test-token")
			roles, err := c.Role().GetRoleList()

			if (err != nil) != tt.wantErr {
				t.Errorf("GetRoleList() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && len(roles) != 1 {
				t.Errorf("GetRoleList() got %v roles, want 1", len(roles))
			}
		})
	}
}
