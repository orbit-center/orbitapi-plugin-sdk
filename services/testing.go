package services

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/orbit-center/sdk/interfaces"
	"github.com/orbit-center/sdk/models"
)

// mockClient 实现了 interfaces.Client 接口
type mockClient struct {
	response []byte
	err      error
}

func (m *mockClient) DoRequest(method, path string, body interface{}) (interfaces.Response, error) {
	if m.err != nil {
		return nil, m.err
	}
	return &mockResponse{data: m.response}, nil
}

type mockResponse struct {
	data []byte
	err  error
}

func (r *mockResponse) DecodeData(v interface{}) error {
	if r.err != nil {
		return r.err
	}

	var resp models.Response
	if err := json.Unmarshal(r.data, &resp); err != nil {
		return fmt.Errorf("unmarshal response failed: %w", err)
	}

	config := &mapstructure.DecoderConfig{
		Result:           v,
		TagName:          "json",
		WeaklyTypedInput: true,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return fmt.Errorf("create decoder failed: %w", err)
	}

	return decoder.Decode(resp.Data)
}
