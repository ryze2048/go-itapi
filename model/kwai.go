package model

const KwaiUrl string = `https://api.itapi.cn/api/hotnews/kuaishou`

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
