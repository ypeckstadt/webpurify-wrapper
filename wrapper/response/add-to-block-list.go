package response

type WebPurifyAddToBlockListResponseWrapper struct {
	Response WebPurifyAddToBlockListResponse `json:"rsp"`
}

type WebPurifyAddToBlockListResponse struct {
	Method string `json:"method"`
	Format string `json:"format"`
	Success string `json:"success"`
	ApiKey string `json:"api_key"`
}