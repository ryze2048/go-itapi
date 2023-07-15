package lifestyle

import (
	"fmt"
	"github.com/ryze2048/go-itapi/model"
	"testing"
	"time"
)

func TestLife_GetShippingFee(t *testing.T) {
	var LifeStyle = new(Life)
	req := &model.ShippingFeeRequest{
		Key:    "",
		From:   "",
		To:     "",
		Weight: 5,
		Time:   time.Now().Format("2006-01-02 15:04:05"),
	}
	if resp, err := LifeStyle.GetShippingFee(req); err != nil {
		t.Error("Err --> ", err)
	} else {
		fmt.Println(resp)
	}
}

func TestLife_Any(t *testing.T) {
	var LifeStyle = new(Life)

	request := &model.IDCardLocationRequest{
		Key: "",
		Id:  "",
	}
	var response model.IDCardLocationResponse
	if err := LifeStyle.Any(request, &response, model.IDCardLocationUrl); err != nil {
		t.Error("err --> ", err)
	}
	fmt.Println(response)
}

func TestGetOilPrice(t *testing.T) {
	var LifeStyle = new(Life)
	request := &model.OilPriceRequest{
		Key:  "",
		City: "广西",
	}
	var response model.OilPriceResponse
	if err := LifeStyle.Any(request, &response, model.OilPriceUrl); err != nil {
		t.Error("err --> ", err)
	}

	fmt.Println(response)
}

func TestGetServiceSuspendedAreas(t *testing.T) {
	var LifeStyle = new(Life)
	request := &model.ServiceSuspendedAreasRequest{
		Key:     "",
		Com:     "jd",
		Fromadd: "",
		Toadd:   "",
	}
	var response model.ServiceSuspendedAreasResponse
	if err := LifeStyle.Any(request, &response, model.ServiceSuspendedAreasUrl); err != nil {
		t.Error("err --> ", err)
	}
	fmt.Println(response)
}

func TestGetDeliveryRestrictions(t *testing.T) {
	var LifeStyle = new(Life)
	request := &model.DeliveryRestrictionsRequest{
		Key:     "",
		Address: "",
	}
	var response model.DeliveryRestrictionsResponse
	if err := LifeStyle.Any(request, &response, model.DeliveryRestrictionsUrl); err != nil {
		t.Error("err --> ", err)
	}
	fmt.Println(response)
}

func TestGetPhoneNumberLocation(t *testing.T) {
	var LifeStyle = new(Life)
	request := &model.PhoneNumberLocationRequest{
		Key: "",
		Tel: "",
	}
	var response model.PhoneNumberLocationResponse
	if err := LifeStyle.Any(request, &response, model.PhoneNumberLocationUrl); err != nil {
		t.Error("err --> ", err)
	}
	fmt.Println(response)
}

func TestGetWeatherInfo(t *testing.T) {
	var LifeStyle = new(Life)
	request := &model.WeatherInfoRequest{
		Key:  "",
		City: "六盘水",
		Ip:   "",
	}
	var response model.WeatherInfoResponse
	if err := LifeStyle.Any(request, &response, model.WeatherInfoUrl); err != nil {
		t.Error("err --> ", err)
	}
	fmt.Println(response)
}

func TestGetExpressInfo(t *testing.T) {
	var LifeStyle = new(Life)
	request := &model.ExpressInfoRequest{
		Key:    "",
		Number: "",
		Com:    "jd",
	}
	var response model.ExpressInfoResponse
	if err := LifeStyle.Any(request, &response, model.ExpressInfoUrl); err != nil {
		t.Error("err --> ", err)
	}
	fmt.Println(response)
}
