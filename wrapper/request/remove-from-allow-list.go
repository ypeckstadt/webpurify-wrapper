package request

import (
	"encoding/json"
	"github.com/ypeckstadt/webpurify-wrapper/wrapper/response"
	"io/ioutil"
	"net/http"
)

// RemoveFromBlockList is a profanity word management method. Removes submitted word from the allow list of the associated license key.
func (w *WebPurifyRequest) RemoveFromAllowList(word string) (response.WebPurifyRemoveFromAllowListResponse, error) {

	// create request parameter collection
	requestParameters := []WebPurifyRequestParameter{
		{Type: Word, Value: word},
	}

	// build http request
	request, err := w.createRequest(RemoveFromAllowList, requestParameters...)
	if err != nil {
		return response.WebPurifyRemoveFromAllowListResponse{}, err
	}

	// create http client
	client := w.createHttpClient()

	// Execute http request
	httpResponse, err := client.Do(request)
	if err != nil {
		return response.WebPurifyRemoveFromAllowListResponse{}, err
	}

	// convert response to struct
	webPurifyResponse, err := w.convertToRemoveFromAllowListResponse(*httpResponse)
	if err != nil {
		return response.WebPurifyRemoveFromAllowListResponse{}, err
	}

	return webPurifyResponse, nil
}

// TODO: convert to generic function when Generics are finally available
func (w *WebPurifyRequest) convertToRemoveFromAllowListResponse(resp http.Response) (response.WebPurifyRemoveFromAllowListResponse, error) {
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return response.WebPurifyRemoveFromAllowListResponse{}, err
	}

	responseWrapper := response.WebPurifyRemoveFromAllowListResponseWrapper{}

	err = json.Unmarshal(body, &responseWrapper)
	if err != nil {
		return response.WebPurifyRemoveFromAllowListResponse{}, err
	}

	return responseWrapper.Response, nil
}
