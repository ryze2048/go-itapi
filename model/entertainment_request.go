package model

type RandomQuoteRequest struct {
	Key string `json:"key"`
}

type JokeRequest struct {
	Key  string `json:"key"`
	Num  string `json:"num"`
	Page string `json:"page"`
}

type EditEmojiRequest struct {
	Key  string `json:"key"`
	Url  string `json:"url"`  // 图片链接（支持jpg，png，jpeg）
	Type string `json:"type"` // 0 代表张嘴笑(默认); 1 代表嘟嘟嘴; 2 代表不高兴; 3 代表闭嘴笑;（可选）
}
