package model

// GeneralNewsResponse 综合新闻响应参数
type GeneralNewsResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Title   string `json:"title"`
		Source  string `json:"source"`
		Content string `json:"content"`
		Link    string `json:"link"`
		Time    string `json:"time"`
	} `json:"data"`
	Debug    string  `json:"debug"`
	ExecTime float64 `json:"exec_time"`
	UserIp   string  `json:"user_ip"`
	LogId    string  `json:"log_id"`
}

// Hours24Response 7*24小时热点
type Hours24Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Title       string   `json:"title"`
		Summary     string   `json:"summary"`
		Content     string   `json:"content"`
		Tags        []string `json:"tags"`
		Category    []string `json:"category"`
		Author      string   `json:"author"`
		PublishTime string   `json:"publish_time"`
	} `json:"data"`
	Debug    string  `json:"debug"`
	ExecTime float64 `json:"exec_time"`
	UserIp   string  `json:"user_ip"`
	LogId    string  `json:"log_id"`
}

// KwaiResponse 快手热榜
type KwaiResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Rank    int    `json:"rank"`
		Name    string `json:"name"`
		Img     string `json:"img"`
		Date    string `json:"date"`
		Viewnum string `json:"viewnum"`
		Url     string `json:"url"`
	} `json:"data"`
	Debug    string  `json:"debug"`
	ExecTime float64 `json:"exec_time"`
	UserIp   string  `json:"user_ip"`
	LogId    string  `json:"log_id"`
}

// TiebaResponse 贴吧热点
type TiebaResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Name    string `json:"name"`
		Brief   string `json:"brief"`
		Date    string `json:"date"`
		Viewnum int    `json:"viewnum"`
		Url     string `json:"url"`
	} `json:"data"`
	Debug    string  `json:"debug"`
	ExecTime float64 `json:"exec_time"`
	UserIp   string  `json:"user_ip"`
	LogId    string  `json:"log_id"`
}

// HotNews360Response 360热搜
type HotNews360Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Name    string `json:"name"`
		Date    string `json:"date"`
		Viewnum string `json:"viewnum"`
		Url     string `json:"url"`
	} `json:"data"`
	Debug    string  `json:"debug"`
	ExecTime float64 `json:"exec_time"`
	UserIp   string  `json:"user_ip"`
	LogId    string  `json:"log_id"`
}

// BaiduResponse 百度热点
type BaiduResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Name    string `json:"name"`
		Brief   string `json:"brief"`
		Date    string `json:"date"`
		Viewnum string `json:"viewnum"`
		Url     string `json:"url"`
	} `json:"data"`
	Debug    string  `json:"debug"`
	ExecTime float64 `json:"exec_time"`
	UserIp   string  `json:"user_ip"`
	LogId    string  `json:"log_id"`
}

// ToutiaoResponse 头条热点
type ToutiaoResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Name    string `json:"name"`
		Date    string `json:"date"`
		Viewnum string `json:"viewnum"`
		Url     string `json:"url"`
	} `json:"data"`
	Debug    string  `json:"debug"`
	ExecTime float64 `json:"exec_time"`
	UserIp   string  `json:"user_ip"`
	LogId    string  `json:"log_id"`
}

// WeixinResponse 微信热点
type WeixinResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Name   string `json:"name"`
		Brief  string `json:"brief"`
		Date   string `json:"date"`
		Writer string `json:"writer"`
		Url    string `json:"url"`
	} `json:"data"`
	Debug    string  `json:"debug"`
	ExecTime float64 `json:"exec_time"`
	UserIp   string  `json:"user_ip"`
	LogId    string  `json:"log_id"`
}

// WeiboResponse 微博热点榜
type WeiboResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Name    string `json:"name"`
		Date    string `json:"date"`
		Viewnum int    `json:"viewnum"`
		Url     string `json:"url"`
	} `json:"data"`
	Debug    string  `json:"debug"`
	ExecTime float64 `json:"exec_time"`
	UserIp   string  `json:"user_ip"`
	LogId    string  `json:"log_id"`
}

// DouyinResponse 抖音热点榜
type DouyinResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Name      string `json:"name"`
		Date      string `json:"date"`
		Viewnum   int    `json:"viewnum"`
		WordCover struct {
			UrlList []string `json:"url_list"`
			Uri     string   `json:"uri"`
		} `json:"word_cover"`
		Url string `json:"url"`
	} `json:"data"`
	Debug    string  `json:"debug"`
	ExecTime float64 `json:"exec_time"`
	UserIp   string  `json:"user_ip"`
	LogId    string  `json:"log_id"`
}

// BilibiliResponse B站热点
type BilibiliResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Name    string `json:"name"`
		Date    string `json:"date"`
		Viewnum int    `json:"viewnum"`
		Url     string `json:"url"`
	} `json:"data"`
	Debug    string  `json:"debug"`
	ExecTime float64 `json:"exec_time"`
	UserIp   string  `json:"user_ip"`
	LogId    string  `json:"log_id"`
}

// ZaobaoResponse 早报热点
type ZaobaoResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Date      string   `json:"date"`
		News      []string `json:"news"`
		Weiyu     string   `json:"weiyu"`
		Image     string   `json:"image"`
		HeadImage string   `json:"head_image"`
	} `json:"data"`
	Debug    string  `json:"debug"`
	ExecTime float64 `json:"exec_time"`
	UserIp   string  `json:"user_ip"`
	LogId    string  `json:"log_id"`
}

// ZhihuResponse 知乎热点榜
type ZhihuResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Name    string `json:"name"`
		Date    string `json:"date"`
		Viewnum string `json:"viewnum"`
		Url     string `json:"url"`
	} `json:"data"`
	Debug    string  `json:"debug"`
	ExecTime float64 `json:"exec_time"`
	UserIp   string  `json:"user_ip"`
	LogId    string  `json:"log_id"`
}

// AllResponse 综合热榜
type AllResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Rank    int    `json:"rank"`
		Name    string `json:"name"`
		Viewnum string `json:"viewnum"`
		Url     string `json:"url"`
	} `json:"data"`
	Debug    string  `json:"debug"`
	ExecTime float64 `json:"exec_time"`
	UserIp   string  `json:"user_ip"`
	LogId    string  `json:"log_id"`
}
