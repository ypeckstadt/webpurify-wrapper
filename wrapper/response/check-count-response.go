package response

type WebPurifyCheckCountResponseWrapper struct {
	Response WebPurifyCheckCountResponse `json:"rsp"`
}

type WebPurifyCheckCountResponse struct {
	Method string `json:"method"`
	Format string `json:"format"`
	Found string `json:"found"`
	ApiKey string `json:"api_key"`
}
