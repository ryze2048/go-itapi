package news

import (
	"bytes"
	"encoding/json"
	"github.com/ryze2048/go-itapi/model"
	"io"
	"net/http"
)

type NewsAPI struct{}

func (n *NewsAPI) GeneralNews(data *model.GeneralNewsRequest) (generalNewsResponse *model.GeneralNewsResponse, err error) {
	var byteInfo = make([]byte, 0)
	if byteInfo, err = json.Marshal(data); err != nil {
		return nil, err
	}
	var response *http.Response
	if response, err = http.Post(model.GeneralNewsUrl, "application/json", bytes.NewBuffer(byteInfo)); err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()

	var bodyByte = make([]byte, 0)
	if bodyByte, err = io.ReadAll(response.Body); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(bodyByte, &generalNewsResponse); err != nil {
		return nil, err
	}
	return
}
