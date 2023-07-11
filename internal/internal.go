package internal

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

func HttpGet(data map[string]string, requestUrl string) (bodyByte []byte, err error) {
	client := http.Client{}
	values := url.Values{}
	var request *http.Request
	if request, err = http.NewRequest(http.MethodGet, requestUrl, nil); err != nil {
		return nil, err
	}
	if data != nil {
		for key, val := range data {
			values.Set(key, val)
		}
		request.URL.RawQuery = values.Encode()
	}
	var response *http.Response
	if response, err = client.Get(request.URL.String()); err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	if bodyByte, err = io.ReadAll(response.Body); err != nil {
		return nil, err
	}
	return bodyByte, nil
}

func StructToMap(data any) (result map[string]string, err error) {
	var byteInfo = make([]byte, 0)
	if byteInfo, err = json.Marshal(data); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(byteInfo, &result); err != nil {
		return nil, err
	}
	return result, nil
}
