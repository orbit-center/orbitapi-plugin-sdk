package services

import (
	"fmt"
	"testing"
)

func TestDictService_GetDictList(t *testing.T) {
	c := NewDictService(&mockClient{
		response: []byte(`{
			"code": 0,
			"data": {
				"total": 1,
				"list": [
					{"id": 1, "name": "系统状态", "type": "status"}
				]
			}
		}`),
	})

	dicts, err := c.GetDictList()
	if err != nil {
		t.Fatalf("GetDictList failed: %v", err)
	}
	if len(dicts) != 1 {
		t.Errorf("expected 1 dict, got %d", len(dicts))
	}
	if dicts[0].Type != "status" {
		t.Errorf("expected type 'status', got %s", dicts[0].Type)
	}
}

func TestDictService_GetDictTypes(t *testing.T) {
	c := NewDictService(&mockClient{
		response: []byte(`{"code":0,"data":[{"id":1,"name":"正常","value":"1"}]}`),
	})

	types, err := c.GetDictTypes(1)
	if err != nil {
		t.Fatalf("GetDictTypes failed: %v", err)
	}
	if len(types) != 1 {
		t.Errorf("expected 1 type, got %d", len(types))
	}
}

func TestDictService_GetDictContents(t *testing.T) {
	c := NewDictService(&mockClient{
		response: []byte(`{"code":0,"data":[{"id":1,"label":"启用","value":"1"}]}`),
	})

	contents, err := c.GetDictContents(1)
	if err != nil {
		t.Fatalf("GetDictContents failed: %v", err)
	}
	if len(contents) != 1 {
		t.Errorf("expected 1 content, got %d", len(contents))
	}
}

func TestDictService_Error(t *testing.T) {
	c := NewDictService(&mockClient{
		err: fmt.Errorf("mock error"),
	})

	_, err := c.GetDictList()
	if err == nil {
		t.Error("expected error, got nil")
	}
}
