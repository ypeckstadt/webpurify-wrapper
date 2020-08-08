package response

type WebPurifyCheckResponseWrapper struct {
	Response WebPurifyCheckResponse `json:"rsp"`
}

type WebPurifyCheckResponse struct {
	Method string `json:"method"`
	Format string `json:"format"`
	Found string `json:"found"`
	ApiKey string `json:"api_key"`
}
