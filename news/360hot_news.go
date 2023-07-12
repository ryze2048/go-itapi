package news

import (
	"encoding/json"
	"github.com/ryze2048/go-itapi/internal"
	"github.com/ryze2048/go-itapi/model"
)

// HotNews360 360热搜
// https://api.itapi.cn/user/doc?id=60
func (n *NewsAPI) HotNews360(data *model.Common) (hotNews360Response *model.HotNews360Response, err error) {
	var result = make(map[string]string, 0)
	if result, err = internal.StructToMap(data); err != nil {
		return nil, err
	}
	var bodyByte = make([]byte, 0)
	if bodyByte, err = internal.HttpGet(result, model.HotNews360Url); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(bodyByte, &hotNews360Response); err != nil {
		return nil, err
	}
	return
}
