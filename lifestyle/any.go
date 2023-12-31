package lifestyle

import (
	"encoding/json"
	"github.com/ryze2048/go-itapi/internal"
)

func (l *Life) Any(request any, response any, requestUrl string) (err error) {
	var result = make(map[string]string, 0)
	if result, err = internal.StructToMap(request); err != nil {
		return err
	}
	var bodyByte = make([]byte, 0)
	if bodyByte, err = internal.HttpGet(result, requestUrl); err != nil {
		return err
	}
	if err = json.Unmarshal(bodyByte, &response); err != nil {
		return err
	}
	return
}
