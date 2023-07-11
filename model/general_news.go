package model

const GeneralNewsUrl string = `https://api.itapi.cn/api/news/news`

type GeneralNewsRequest struct {
	Common
	Type string `json:"type"`
	Num  string `json:"num"`
	Page string `json:"page"`
}

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
