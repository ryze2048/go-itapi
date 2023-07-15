package news

import (
	"encoding/json"
	"github.com/ryze2048/go-itapi/internal"
	"github.com/ryze2048/go-itapi/model"
)

func (n *NewsAPI) Unify(data *model.UnifyNews) (bodyByte []byte, err error) {
	var result = make(map[string]string, 0)
	if result, err = internal.StructToMap(data.Common); err != nil {
		return nil, err
	}
	return internal.HttpGet(result, data.RequestUrl)
}

func (n *NewsAPI) Unmarshal(byteInfo []byte, data any) (err error) {
	return json.Unmarshal(byteInfo, &data)
}
