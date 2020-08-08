package response

type WebPurifyAddToAllowListResponseWrapper struct {
	Response WebPurifyAddToAllowListResponse `json:"rsp"`
}

type WebPurifyAddToAllowListResponse struct {
	Method string `json:"method"`
	Format string `json:"format"`
	Success string `json:"success"`
	ApiKey string `json:"api_key"`
}