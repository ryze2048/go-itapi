package model

type ComRequest struct {
	Key  string `json:"key"`
	Type string `json:"type"`
}

type LogisticsRequest struct {
	Key    string `json:"key"`
	Com    string `json:"com"`
	Number string `json:"number"`
	Phone  string `json:"phone"`
}

type ShippingFeeRequest struct {
	Key    string `json:"key"`
	From   string `json:"from"`   // 始发详细地址
	To     string `json:"to"`     // 目的地详细地址
	Weight int    `json:"weight"` // 重量（单位：kg）
	Time   string `json:"time"`   // 寄件时间，格式：yyyy-MM-dd HH:mm:ss
}

type IDCardLocationRequest struct {
	Key string `json:"key"`
	Id  string `json:"id"`
}

type OilPriceRequest struct {
	Key  string `json:"key"`
	City string `json:"city"` // 各省级行政区名，比如湖北、上海。注意末尾不要带上省或市。
}

type ServiceSuspendedAreasRequest struct {
	Key     string `json:"key"`
	Com     string `json:"com"`     // 快递公司简称：jd、ems
	Fromadd string `json:"fromadd"` // 寄件人地址，如：张三，13800138000，广东省深圳市南山区 ！可选~
	Toadd   string `json:"toadd"`   // 智能识别收件人地址，如：张三，13800138000，广东省深圳市南山区 ！ 必填
}

type DeliveryRestrictionsRequest struct {
	Key     string `json:"key"`
	Address string `json:"address"` // 湖南省长沙市开福区
}

type PhoneNumberLocationRequest struct {
	Key string `json:"key"`
	Tel string `json:"tel"`
}

type WeatherInfoRequest struct {
	Key  string `json:"key"`
	City string `json:"city"`
	Ip   string `json:"ip"`
}

type ExpressInfoRequest struct {
	Key    string `json:"key"`
	Number string `json:"number"`
	Com    string `json:"com"`
}
