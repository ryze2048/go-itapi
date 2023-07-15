package lifestyle

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ryze2048/go-itapi/model"
	"io"
	"net/http"
)

// GetShippingFee 快递运费查询
func (l *Life) GetShippingFee(data *model.ShippingFeeRequest) (shippingFeeResponse *model.ShippingFeeResponse, err error) {
	var byteInfo = make([]byte, 0)
	if byteInfo, err = json.Marshal(data); err != nil {
		return nil, err
	}

	var response *http.Response
	if response, err = http.Post(model.ShippingFeeUrl, "application/json", bytes.NewBuffer(byteInfo)); err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()

	var bodyByte = make([]byte, 0)
	if bodyByte, err = io.ReadAll(response.Body); err != nil {
		return nil, err
	}

	fmt.Println(string(bodyByte))
	if err = json.Unmarshal(bodyByte, &shippingFeeResponse); err != nil {
		return nil, err
	}
	return
}
