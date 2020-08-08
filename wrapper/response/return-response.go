package response

type WebPurifyReturnResponseWrapper struct {
	Response WebPurifyReturnResponse `json:"rsp"`
}

type WebPurifyReturnResponse struct {
	Method string `json:"method"`
	Format string `json:"format"`
	Found string `json:"found"`
	ApiKey string `json:"api_key"`
	Expletive []string `json:"expletive"`
}