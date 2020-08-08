package response

type WebPurifyRemoveFromBlockListResponseWrapper struct {
	Response WebPurifyRemoveFromBlockListResponse `json:"rsp"`
}

type WebPurifyRemoveFromBlockListResponse struct {
	Method string `json:"method"`
	Format string `json:"format"`
	Success string `json:"success"`
	ApiKey string `json:"api_key"`
}