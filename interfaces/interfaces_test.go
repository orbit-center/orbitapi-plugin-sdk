package interfaces

import (
	"errors"
	"testing"

	"github.com/mitchellh/mapstructure"
)

type mockResponse struct {
	data interface{}
	err  error
}

func (r *mockResponse) DecodeData(v interface{}) error {
	if r.err != nil {
		return r.err
	}

	// 使用 mapstructure 进行类型转换
	return mapstructure.Decode(r.data, v)
}

func TestResponse_DecodeData(t *testing.T) {
	tests := []struct {
		name    string
		resp    Response
		wantErr bool
	}{
		{
			name:    "success",
			resp:    &mockResponse{data: map[string]interface{}{"id": 1}},
			wantErr: false,
		},
		{
			name:    "error",
			resp:    &mockResponse{err: errors.New("decode error")},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var data interface{}
			err := tt.resp.DecodeData(&data)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestResponse_DecodeDataTypes(t *testing.T) {
	tests := []struct {
		name    string
		data    interface{}
		target  interface{}
		wantErr bool
	}{
		{
			name:    "decode map",
			data:    map[string]interface{}{"id": 1},
			target:  &map[string]interface{}{},
			wantErr: false,
		},
		{
			name:    "decode slice",
			data:    []interface{}{1, 2, 3},
			target:  &[]interface{}{},
			wantErr: false,
		},
		{
			name:    "invalid target",
			data:    map[string]interface{}{"id": 1},
			target:  nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &mockResponse{data: tt.data}
			err := resp.DecodeData(tt.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
