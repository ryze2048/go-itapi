package model

type RandomQuoteResponse struct {
	Code     int     `json:"code"`
	Msg      string  `json:"msg"`
	Data     string  `json:"data"`
	Debug    string  `json:"debug"`
	ExecTime float64 `json:"exec_time"`
	UserIp   string  `json:"user_ip"`
	LogId    string  `json:"log_id"`
}

type JokeResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Title   string `json:"title"`
		Content string `json:"content"`
		Time    string `json:"time"`
	} `json:"data"`
	Debug    string  `json:"debug"`
	ExecTime float64 `json:"exec_time"`
	UserIp   string  `json:"user_ip"`
	LogId    string  `json:"log_id"`
}

type EditEmojiResponse struct {
	Code     int     `json:"code"`
	Msg      string  `json:"msg"`
	Data     string  `json:"data"`
	Debug    string  `json:"debug"`
	ExecTime float64 `json:"exec_time"`
	UserIp   string  `json:"user_ip"`
	LogId    string  `json:"log_id"`
}
