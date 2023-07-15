package model

const (
	_                        string = ``
	ComUrl                   string = `https://api.itapi.cn/api/kuaidi/com`     // 快递公司查询
	LogisticsUrl             string = `https://api.itapi.cn/api/kuaidi/v2`      // 全球快递物流(含轨迹)查询
	ShippingFeeUrl           string = `https://api.itapi.cn/api/kuaidi/yunfei`  // 快递运费查询
	IDCardLocationUrl        string = `https://api.itapi.cn/api/idcard/query`   // 身份证归属地查询
	OilPriceUrl              string = `https://api.itapi.cn/api/oilprice`       // 实时油价
	ServiceSuspendedAreasUrl string = `https://api.itapi.cn/api/kuaidi/stopone` // 快递停发区域查询(需指定快递公司)
	DeliveryRestrictionsUrl  string = `https://api.itapi.cn/api/kuaidi/stop`    // 快递停发区域查询
	PhoneNumberLocationUrl   string = `https://api.itapi.cn/api/tel`            // 手机号码归属地查询
	WeatherInfoUrl           string = `https://api.itapi.cn/api/tianqi`         // 国内天气查询接口
	ExpressInfoUrl           string = `https://api.itapi.cn/api/kuaidi/index`   // 快递查询
)
