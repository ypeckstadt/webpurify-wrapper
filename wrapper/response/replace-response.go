package response

type WebPurifyReplaceResponseWrapper struct {
	Response WebPurifyReplaceResponse `json:"rsp"`
}

type WebPurifyReplaceResponse struct {
	Method string `json:"method"`
	Format string `json:"format"`
	Found string `json:"found"`
	ApiKey string `json:"api_key"`
	Text string `json:"text"`
}