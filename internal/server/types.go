package server

// запрос
type URLmessage struct {
	URL string `json:"url"`
}

// ответ
type URLanswer struct {
	Result string `json:"result"`
}
