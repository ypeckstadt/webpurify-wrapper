package request

import (
	"encoding/json"
	"github.com/ypeckstadt/webpurify-wrapper/wrapper/response"
	"io/ioutil"
	"net/http"
)

// AddToAllowList is a profanity word management method. Adds submitted word to the allow list of the associated license key.
func (w *WebPurifyRequest) AddToAllowList(word string, useDeepSearch bool) (response.WebPurifyAddToAllowListResponse, error) {

	// create request parameter collection
	requestParameters := []WebPurifyRequestParameter{
		{Type: Word, Value: word},
	}

	if useDeepSearch {
		requestParameters = append(requestParameters, WebPurifyRequestParameter{Type: DeepSearch, Value: "1"})
	}

	// build http request
	request, err := w.createRequest(AddToAllowList, requestParameters...)
	if err != nil {
		return response.WebPurifyAddToAllowListResponse{}, err
	}

	// create http client
	client := w.createHttpClient()

	// Execute http request
	httpResponse, err := client.Do(request)
	if err != nil {
		return response.WebPurifyAddToAllowListResponse{}, err
	}

	// convert response to struct
	webPurifyResponse, err := w.convertToAddToAllowListResponse(*httpResponse)
	if err != nil {
		return response.WebPurifyAddToAllowListResponse{}, err
	}

	return webPurifyResponse, nil
}

// TODO: convert to generic function when Generics are finally available
func (w *WebPurifyRequest) convertToAddToAllowListResponse(resp http.Response) (response.WebPurifyAddToAllowListResponse, error) {
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return response.WebPurifyAddToAllowListResponse{}, err
	}

	responseWrapper := response.WebPurifyAddToAllowListResponseWrapper{}

	err = json.Unmarshal(body, &responseWrapper)
	if err != nil {
		return response.WebPurifyAddToAllowListResponse{}, err
	}

	return responseWrapper.Response, nil
}
