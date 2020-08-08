package response

type WebPurifyRemoveFromAllowListResponseWrapper struct {
	Response WebPurifyRemoveFromAllowListResponse `json:"rsp"`
}

type WebPurifyRemoveFromAllowListResponse struct {
	Method string `json:"method"`
	Format string `json:"format"`
	Success string `json:"success"`
	ApiKey string `json:"api_key"`
}