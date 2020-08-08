package request

import (
	"encoding/json"
	"github.com/ypeckstadt/webpurify-wrapper/wrapper/response"
	"io/ioutil"
	"net/http"
)

// GetAllowList is a profanity word management method. Returns the custom allow list of the associated license key.
func (w *WebPurifyRequest) GetAllowList(cdata bool, useDeepSearch bool) (response.WebPurifyGetAllowListResponse, error) {

	// create request parameter collection
	requestParameters := []WebPurifyRequestParameter{}

	if cdata {
		requestParameters = append(requestParameters, WebPurifyRequestParameter{Type: CDATA, Value: "1"})
	}

	if useDeepSearch {
		requestParameters = append(requestParameters, WebPurifyRequestParameter{Type: DeepSearch, Value: "1"})
	}

	// build http request
	request, err := w.createRequest(GetAllowList, requestParameters...)
	if err != nil {
		return response.WebPurifyGetAllowListResponse{}, err
	}

	// create http client
	client := w.createHttpClient()

	// Execute http request
	httpResponse, err := client.Do(request)
	if err != nil {
		return response.WebPurifyGetAllowListResponse{}, err
	}

	// convert response to struct
	webPurifyResponse, err := w.convertToGetAllowListResponse(*httpResponse)
	if err != nil {
		return response.WebPurifyGetAllowListResponse{}, err
	}

	return webPurifyResponse, nil
}

// TODO: convert to generic function when Generics are finally available
func (w *WebPurifyRequest) convertToGetAllowListResponse(resp http.Response) (response.WebPurifyGetAllowListResponse, error) {
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return response.WebPurifyGetAllowListResponse{}, err
	}

	responseWrapper := response.WebPurifyGetAllowListResponseWrapper{}

	err = json.Unmarshal(body, &responseWrapper)
	if err != nil {
		return response.WebPurifyGetAllowListResponse{}, err
	}

	return responseWrapper.Response, nil
}
