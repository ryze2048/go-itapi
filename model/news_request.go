package model

type UnifyNews struct {
	Common
	RequestUrl string `json:"request_url"`
	Response   any    `json:"response"`
}

// GeneralNewsRequest 综合新闻请求参数
type GeneralNewsRequest struct {
	Common
	Type string `json:"type"`
	Num  string `json:"num"`
	Page string `json:"page"`
}

type AllRequest struct {
	Common
	Type string `json:"type"`
}
