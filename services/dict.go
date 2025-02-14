// Package services 提供各种服务实现
package services

import (
	"fmt"

	"github.com/orbit-center/sdk/interfaces"
	"github.com/orbit-center/sdk/models"
)

// DictServiceImpl 实现了 interfaces.DictService 接口
type DictServiceImpl struct {
	client interfaces.Client
}

// NewDictService 创建字典服务实例
func NewDictService(client interfaces.Client) interfaces.DictService {
	return &DictServiceImpl{client: client}
}

// GetDictList 获取字典列表
func (s *DictServiceImpl) GetDictList() ([]models.Dict, error) {
	resp, err := s.client.DoRequest("GET", "/dict/list", nil)
	if err != nil {
		return nil, fmt.Errorf("get dict list failed: %w", err)
	}

	var listResp models.DictListResponse
	if err := resp.DecodeData(&listResp); err != nil {
		return nil, fmt.Errorf("decode dict list failed: %w", err)
	}

	return listResp.List, nil
}

// GetDictTypes 获取字典类型
func (s *DictServiceImpl) GetDictTypes(dictID int64) ([]models.DictType, error) {
	req := map[string]interface{}{
		"dict_id": dictID,
	}
	resp, err := s.client.DoRequest("POST", "/dict/type/list", req)
	if err != nil {
		return nil, fmt.Errorf("get dict types failed: %w", err)
	}

	var types []models.DictType
	if err := resp.DecodeData(&types); err != nil {
		return nil, fmt.Errorf("decode dict types failed: %w", err)
	}

	return types, nil
}

// GetDictContents 获取字典内容
func (s *DictServiceImpl) GetDictContents(typeID int64) ([]models.DictContent, error) {
	req := map[string]interface{}{
		"type_id": typeID,
	}
	resp, err := s.client.DoRequest("POST", "/dict/content/list", req)
	if err != nil {
		return nil, fmt.Errorf("get dict contents failed: %w", err)
	}

	var contents []models.DictContent
	if err := resp.DecodeData(&contents); err != nil {
		return nil, fmt.Errorf("decode dict contents failed: %w", err)
	}

	return contents, nil
}
