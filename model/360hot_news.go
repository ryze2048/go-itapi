package model

const HotNews360Url string = `https://api.itapi.cn/api/hotnews/360`

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
