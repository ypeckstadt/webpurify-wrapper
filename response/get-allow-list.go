package response

type WebPurifyGetAllowListResponseWrapper struct {
	Response WebPurifyGetAllowListResponse `json:"rsp"`
}

type WebPurifyGetAllowListResponse struct {
	Method string `json:"method"`
	Format string `json:"format"`
	Word []string `json:"word"`
	ApiKey string `json:"api_key"`
}