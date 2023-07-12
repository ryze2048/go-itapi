package news

import (
	"encoding/json"
	"github.com/ryze2048/go-itapi/internal"
	"github.com/ryze2048/go-itapi/model"
)

// Kwai 快手热榜
// https://api.itapi.cn/user/doc?id=61
func (n *NewsAPI) Kwai(data *model.Common) (kwaiResponse *model.KwaiResponse, err error) {
	var result = make(map[string]string, 0)
	if result, err = internal.StructToMap(data); err != nil {
		return nil, err
	}
	var bodyByte = make([]byte, 0)
	if bodyByte, err = internal.HttpGet(result, model.KwaiUrl); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(bodyByte, &kwaiResponse); err != nil {
		return nil, err
	}
	return kwaiResponse, nil
}
