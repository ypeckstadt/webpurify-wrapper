package request

import (
	"encoding/json"
	"github.com/ypeckstadt/webpurify-wrapper/response"
	"io/ioutil"
	"net/http"
)

// AddToBlockList is a profanity word management method. Adds submitted word to the block list of the associated license key. The word will be filtered along with words WebPurify filters by default.
func (w *WebPurifyRequest) AddToBlockList(word string, useDeepSearch bool) (response.WebPurifyAddToBlockListResponse, error) {

	// create request parameter collection
	requestParameters := []WebPurifyRequestParameter{
		{Type: Word, Value: word},
	}

	if useDeepSearch {
		requestParameters = append(requestParameters, WebPurifyRequestParameter{Type: DeepSearch, Value: "1"})
	}

	// build http request
	request, err := w.createRequest(AddToBlockList, requestParameters...)
	if err != nil {
		return response.WebPurifyAddToBlockListResponse{}, err
	}

	// create http client
	client := w.createHttpClient()

	// Execute http request
	httpResponse, err := client.Do(request)
	if err != nil {
		return response.WebPurifyAddToBlockListResponse{}, err
	}

	// convert response to struct
	webPurifyResponse, err := w.convertToAddToBlockListResponse(*httpResponse)
	if err != nil {
		return response.WebPurifyAddToBlockListResponse{}, err
	}

	return webPurifyResponse, nil
}

// TODO: convert to generic function when Generics are finally available
func (w *WebPurifyRequest) convertToAddToBlockListResponse(resp http.Response) (response.WebPurifyAddToBlockListResponse, error) {
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return response.WebPurifyAddToBlockListResponse{}, err
	}

	responseWrapper := response.WebPurifyAddToBlockListResponseWrapper{}

	err = json.Unmarshal(body, &responseWrapper)
	if err != nil {
		return response.WebPurifyAddToBlockListResponse{}, err
	}

	return responseWrapper.Response, nil
}
