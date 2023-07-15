package entertainment

import (
	"encoding/json"
	"fmt"
	"github.com/ryze2048/go-itapi/internal"
)

type Entertainment struct{}

func (e *Entertainment) Any(request any, response any, requestUrl string) (err error) {
	var result = make(map[string]string, 0)
	if result, err = internal.StructToMap(request); err != nil {
		return err
	}
	var bodyByte = make([]byte, 0)
	if bodyByte, err = internal.HttpGet(result, requestUrl); err != nil {
		return err
	}
	fmt.Println("body byte --> ", string(bodyByte))
	if err = json.Unmarshal(bodyByte, &response); err != nil {
		return err
	}
	return
}
