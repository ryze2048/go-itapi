package news

import (
	"github.com/ryze2048/go-itapi/internal"
	"github.com/ryze2048/go-itapi/model"
)

// All 综合热点榜
func (n *NewsAPI) All(data *model.AllRequest) (allResponse *model.AllResponse, err error) {
	var result = make(map[string]string, 0)
	if result, err = internal.StructToMap(data.Common); err != nil {
		return allResponse, err
	}
	var bodyByte = make([]byte, 0)
	if bodyByte, err = internal.HttpGet(result, model.AllUrl); err != nil {
		return allResponse, err
	}
	if err = n.Unmarshal(bodyByte, &allResponse); err != nil {
		return allResponse, err
	}
	return allResponse, err
}
