package model

type ComResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Name string `json:"name"`
		Code string `json:"code"`
	} `json:"data"`
	Debug    string  `json:"debug"`
	ExecTime float64 `json:"exec_time"`
	UserIp   string  `json:"user_ip"`
}

type LogisticsResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		State int    `json:"state"`
		Nu    string `json:"nu"`
		Com   string `json:"com"`
		Tel   string `json:"tel"`
		Data  []struct {
			Time     string `json:"time"`
			Address  string `json:"address"`
			Status   int    `json:"status"`
			Context  string `json:"context"`
			Location string `json:"location"`
		} `json:"data"`
		ComName      string `json:"com_name"`
		DeliveryTime string `json:"delivery_time"`
		UpdateTime   string `json:"update_time"`
	} `json:"data"`
	Debug    string  `json:"debug"`
	ExecTime float64 `json:"exec_time"`
	UserIp   string  `json:"user_ip"`
	LogId    string  `json:"log_id"`
}

type ShippingFeeResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Com  string `json:"com"`
		List []struct {
			Time  string `json:"time"`
			Price *int   `json:"price"`
			Type  string `json:"type"`
		} `json:"list"`
	} `json:"data"`
	Debug    string  `json:"debug"`
	ExecTime float64 `json:"exec_time"`
	UserIp   string  `json:"user_ip"`
	LogId    string  `json:"log_id"`
}

type IDCardLocationResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Local    string `json:"local"`
		Gender   string `json:"gender"`
		Birthday string `json:"birthday"`
		Info     struct {
			Id         string `json:"id"`
			ProvinceNo string `json:"provinceNo"`
			Province   string `json:"province"`
			CityNo     string `json:"cityNo"`
			City       string `json:"city"`
			DistrictNo string `json:"districtNo"`
			District   string `json:"district"`
			Update     string `json:"update"`
		} `json:"info"`
		BirthdayArray struct {
			Year  string `json:"year"`
			Month string `json:"month"`
			Day   string `json:"day"`
		} `json:"birthday_array"`
	} `json:"data"`
	Debug    string  `json:"debug"`
	ExecTime float64 `json:"exec_time"`
	UserIp   string  `json:"user_ip"`
}

type OilPriceResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Prov string `json:"prov"`
		P0   string `json:"p0"`
		P89  string `json:"p89"`
		P92  string `json:"p92"`
		P95  string `json:"p95"`
		P98  string `json:"p98"`
		Time string `json:"time"`
	} `json:"data"`
	Debug    string  `json:"debug"`
	ExecTime float64 `json:"exec_time"`
	UserIp   string  `json:"user_ip"`
	LogId    string  `json:"log_id"`
}

type ServiceSuspendedAreasResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Toadd string `json:"toadd"`
		List  []struct {
			Express string `json:"express"`
			Com     string `json:"com"`
			Status  int    `json:"status"`
			Reason  string `json:"reason"`
		} `json:"list"`
	} `json:"data"`
	Debug    string  `json:"debug"`
	ExecTime float64 `json:"exec_time"`
	UserIp   string  `json:"user_ip"`
	LogId    string  `json:"log_id"`
}

type DeliveryRestrictionsResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Toadd string `json:"toadd"`
		List  []struct {
			Express string `json:"express"`
			Com     string `json:"com"`
			Status  int    `json:"status"`
			Reason  string `json:"reason"`
		} `json:"list"`
	} `json:"data"`
	Debug    string  `json:"debug"`
	ExecTime float64 `json:"exec_time"`
	UserIp   string  `json:"user_ip"`
	LogId    string  `json:"log_id"`
}

type PhoneNumberLocationResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Pref     string `json:"pref"`
		Phone    string `json:"phone"`
		Province string `json:"province"`
		City     string `json:"city"`
		Isp      string `json:"isp"`
		PostCode string `json:"post_code"`
		CityCode string `json:"city_code"`
		AreaCode string `json:"area_code"`
	} `json:"data"`
	Debug    string  `json:"debug"`
	ExecTime float64 `json:"exec_time"`
	UserIp   string  `json:"user_ip"`
	LogId    string  `json:"log_id"`
}

type WeatherInfoResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		City        string      `json:"city"`
		CityEn      string      `json:"city_en"`
		Province    string      `json:"province"`
		ProvinceEn  string      `json:"province_en"`
		CityId      string      `json:"city_id"`
		Date        string      `json:"date"`
		UpdateTime  string      `json:"update_time"`
		Weather     string      `json:"weather"`
		WeatherCode string      `json:"weather_code"`
		Temp        string      `json:"temp"`
		MinTemp     string      `json:"min_temp"`
		MaxTemp     string      `json:"max_temp"`
		Wind        string      `json:"wind"`
		WindSpeed   string      `json:"wind_speed"`
		WindScale   string      `json:"wind_scale"`
		Rain        string      `json:"rain"`
		Rain24H     string      `json:"rain_24h"`
		Humidity    string      `json:"humidity"`
		Visibility  string      `json:"visibility"`
		Pressure    string      `json:"pressure"`
		TailNumber  string      `json:"tail_number"`
		Air         interface{} `json:"air"`
		AirPm25     interface{} `json:"air_pm25"`
		Sunrise     string      `json:"sunrise"`
		Sunset      string      `json:"sunset"`
		Aqi         struct {
			Air      interface{} `json:"air"`
			AirLevel interface{} `json:"air_level"`
			AirTips  string      `json:"air_tips"`
			Pm25     interface{} `json:"pm25"`
			Pm10     interface{} `json:"pm10"`
			Co       interface{} `json:"co"`
			No2      interface{} `json:"no2"`
			So2      interface{} `json:"so2"`
			O3       interface{} `json:"o3"`
		} `json:"aqi"`
		Index struct {
			Xiche     interface{} `json:"xiche"`
			Ganmao    interface{} `json:"ganmao"`
			Yundong   interface{} `json:"yundong"`
			Huazhuang interface{} `json:"huazhuang"`
			Chuangyi  interface{} `json:"chuangyi"`
			Ziwaixian interface{} `json:"ziwaixian"`
		} `json:"index"`
		Alarm []interface{} `json:"alarm"`
		Hour  []struct {
			Time      string `json:"time"`
			Temp      string `json:"temp"`
			Wea       string `json:"wea"`
			WeaCode   string `json:"wea_code"`
			Wind      string `json:"wind"`
			WindLevel string `json:"wind_level"`
		} `json:"hour"`
	} `json:"data"`
	Debug    string  `json:"debug"`
	ExecTime float64 `json:"exec_time"`
	UserIp   string  `json:"user_ip"`
	LogId    string  `json:"log_id"`
}

type ExpressInfoResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Nu         string `json:"nu"`
		Com        string `json:"com"`
		State      int    `json:"state"`
		Status     string `json:"status"`
		StatusDesc string `json:"status_desc"`
		Info       []struct {
			Content    string `json:"content"`
			Time       string `json:"time"`
			Status     string `json:"status"`
			StatusDesc string `json:"status_desc"`
		} `json:"info"`
	} `json:"data"`
	Debug    string  `json:"debug"`
	ExecTime float64 `json:"exec_time"`
	UserIp   string  `json:"user_ip"`
	LogId    string  `json:"log_id"`
}
