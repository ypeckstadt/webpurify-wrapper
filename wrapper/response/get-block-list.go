package response

type WebPurifyGetBlockListResponseWrapper struct {
	Response WebPurifyGetBlockListResponse `json:"rsp"`
}

type WebPurifyGetBlockListResponse struct {
	Method string `json:"method"`
	Format string `json:"format"`
	Word []string `json:"word"`
	ApiKey string `json:"api_key"`
}