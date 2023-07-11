package model

const Hours24Url string = `https://api.itapi.cn/api/news/24hours`

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
