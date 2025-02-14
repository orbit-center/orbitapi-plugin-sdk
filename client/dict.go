// DictService 实现了字典相关的所有操作
package client

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/orbit-center/sdk/models"
)

type DictService struct {
	Client *Client
}

// 实现 interfaces.DictService 接口
func (s *DictService) GetDictList() ([]models.Dict, error) {
	data, err := s.Client.doRequest("POST", "/dict/list", nil)
	if err != nil {
		return nil, err
	}

	var resp models.Response
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal response failed: %w", err)
	}

	var dicts []models.Dict
	if err := mapstructure.Decode(resp.Data, &dicts); err != nil {
		return nil, fmt.Errorf("decode dicts failed: %w", err)
	}

	return dicts, nil
}

// GetDictTypes 获取字典类型
func (s *DictService) GetDictTypes(dictID int64) ([]models.DictType, error) {
	req := map[string]interface{}{
		"dict_id": dictID,
	}
	data, err := s.Client.doRequest("POST", "/dict/type/list", req)
	if err != nil {
		return nil, err
	}

	var resp models.Response
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal response failed: %w", err)
	}

	var types []models.DictType
	if err := mapstructure.Decode(resp.Data, &types); err != nil {
		return nil, fmt.Errorf("decode dict types failed: %w", err)
	}

	return types, nil
}

// GetDictContents 获取字典内容
func (s *DictService) GetDictContents(typeID int64) ([]models.DictContent, error) {
	req := map[string]interface{}{
		"type_id": typeID,
	}
	data, err := s.Client.doRequest("POST", "/dict/content/list", req)
	if err != nil {
		return nil, err
	}

	var resp models.Response
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal response failed: %w", err)
	}

	var contents []models.DictContent
	if err := mapstructure.Decode(resp.Data, &contents); err != nil {
		return nil, fmt.Errorf("decode dict contents failed: %w", err)
	}

	return contents, nil
}
