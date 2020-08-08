package request

import (
	"encoding/json"
	"github.com/ypeckstadt/webpurify-wrapper/wrapper"
	"github.com/ypeckstadt/webpurify-wrapper/wrapper/response"
	"io/ioutil"
	"net/http"
)

// GetBlockList is a profanity word management method. Returns the custom block list of the associated license key.
func (w *WebPurifyRequest) GetBlockList(cdata bool, useDeepSearch bool) (response.WebPurifyGetBlockListResponse, error) {

	// create request parameter collection
	requestParameters := []wrapper.WebPurifyRequestParameter{}

	if cdata {
		requestParameters = append(requestParameters, wrapper.WebPurifyRequestParameter{Type: wrapper.CDATA, Value: "1"})
	}

	if useDeepSearch {
		requestParameters = append(requestParameters, wrapper.WebPurifyRequestParameter{Type: wrapper.DeepSearch, Value: "1"})
	}

	// build http request
	request, err := w.createRequest(GetBlockList, requestParameters...)
	if err != nil {
		return response.WebPurifyGetBlockListResponse{}, err
	}

	// create http client
	client := w.createHttpClient()

	// Execute http request
	httpResponse, err := client.Do(request)
	if err != nil {
		return response.WebPurifyGetBlockListResponse{}, err
	}

	// convert response to struct
	webPurifyResponse, err := w.convertToGetBlockListResponse(*httpResponse)
	if err != nil {
		return response.WebPurifyGetBlockListResponse{}, err
	}

	return webPurifyResponse, nil
}

// TODO: convert to generic function when Generics are finally available
func (w *WebPurifyRequest) convertToGetBlockListResponse(resp http.Response) (response.WebPurifyGetBlockListResponse, error) {
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return response.WebPurifyGetBlockListResponse{}, err
	}

	responseWrapper := response.WebPurifyGetBlockListResponseWrapper{}

	err = json.Unmarshal(body, &responseWrapper)
	if err != nil {
		return response.WebPurifyGetBlockListResponse{}, err
	}

	return responseWrapper.Response, nil
}