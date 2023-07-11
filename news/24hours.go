package news

import (
	"encoding/json"
	"github.com/ryze2048/go-itapi/internal"
	"github.com/ryze2048/go-itapi/model"
)

func (n *NewsAPI) Hours24(data *model.Common) (hours24Response *model.Hours24Response, err error) {
	var result = make(map[string]string, 0)
	if result, err = internal.StructToMap(data); err != nil {
		return nil, err
	}
	var bodyByte = make([]byte, 0)
	if bodyByte, err = internal.HttpGet(result, model.Hours24Url); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(bodyByte, &hours24Response); err != nil {
		return nil, err
	}
	return hours24Response, nil
}
