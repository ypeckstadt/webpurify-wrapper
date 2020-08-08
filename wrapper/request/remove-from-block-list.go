package request

import (
	"encoding/json"
	"github.com/ypeckstadt/webpurify-wrapper/wrapper"
	"github.com/ypeckstadt/webpurify-wrapper/wrapper/response"
	"io/ioutil"
	"net/http"
)

// RemoveFromBlockList is a profanity word management method. Removes submitted word from the blocklist of the associated license key.
func (w *WebPurifyRequest) RemoveFromBlockList(word string) (response.WebPurifyRemoveFromBlockListResponse, error) {

	// create request parameter collection
	requestParameters := []wrapper.WebPurifyRequestParameter{
		{Type: wrapper.Word, Value: word},
	}

	// build http request
	request, err := w.createRequest(RemoveFromBlockList, requestParameters...)
	if err != nil {
		return response.WebPurifyRemoveFromBlockListResponse{}, err
	}

	// create http client
	client := w.createHttpClient()

	// Execute http request
	httpResponse, err := client.Do(request)
	if err != nil {
		return response.WebPurifyRemoveFromBlockListResponse{}, err
	}

	// convert response to struct
	webPurifyResponse, err := w.convertToRemoveFromBlockListResponse(*httpResponse)
	if err != nil {
		return response.WebPurifyRemoveFromBlockListResponse{}, err
	}

	return webPurifyResponse, nil
}

// TODO: convert to generic function when Generics are finally available
func (w *WebPurifyRequest) convertToRemoveFromBlockListResponse(resp http.Response) (response.WebPurifyRemoveFromBlockListResponse, error) {
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return response.WebPurifyRemoveFromBlockListResponse{}, err
	}

	responseWrapper := response.WebPurifyRemoveFromBlockListResponseWrapper{}

	err = json.Unmarshal(body, &responseWrapper)
	if err != nil {
		return response.WebPurifyRemoveFromBlockListResponse{}, err
	}

	return responseWrapper.Response, nil
}
