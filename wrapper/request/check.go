package request

import (
	"encoding/json"
	"github.com/ypeckstadt/webpurify-wrapper/wrapper"
	"github.com/ypeckstadt/webpurify-wrapper/wrapper/response"
	"io/ioutil"
	"net/http"
)

// Check is a profanity checking method. If profanity is found it returns 1. If the text is clean 0 (zero) is returned.
func (w *WebPurifyRequest) Check(language wrapper.WebPurifyLanguage, text string, filterEmail bool, filterPhone bool, filterLink bool) (response.WebPurifyCheckResponse, error) {

	//rsp (Optional)
	//To include our response time in the result. set = 1

	// create request parameter collection
	requestParameters := []wrapper.WebPurifyRequestParameter{
		{Type: wrapper.Text, Value: text},
		{Type: wrapper.Language, Value: string(language)},
	}

	if filterEmail {
		requestParameters = append(requestParameters, wrapper.WebPurifyRequestParameter{Type: wrapper.Email, Value: "1"})
	}

	if filterPhone {
		requestParameters = append(requestParameters, wrapper.WebPurifyRequestParameter{Type: wrapper.Phone, Value: "1"})
	}

	if filterLink {
		requestParameters = append(requestParameters, wrapper.WebPurifyRequestParameter{Type: wrapper.Link, Value: "1"})
	}

	// build http request
	request, err := w.createRequest(Check, requestParameters...)
	if err != nil {
		return response.WebPurifyCheckResponse{}, err
	}

	// create http client
	client := w.createHttpClient()

	// Execute http request
	httpResponse, err := client.Do(request)
	if err != nil {
		return response.WebPurifyCheckResponse{}, err
	}

	// convert response to struct
	webPurifyResponse, err := w.convertToCheckResponse(*httpResponse)
	if err != nil {
		return response.WebPurifyCheckResponse{}, err
	}

	return webPurifyResponse, nil
}

// TODO: convert to generic function when Generics are finally available
func (w *WebPurifyRequest) convertToCheckResponse(resp http.Response) (response.WebPurifyCheckResponse, error) {
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return response.WebPurifyCheckResponse{}, err
	}

	webPurifyCheckResponseWrapper := response.WebPurifyCheckResponseWrapper{}

	err = json.Unmarshal(body, &webPurifyCheckResponseWrapper)
	if err != nil {
		return response.WebPurifyCheckResponse{}, err
	}

	return webPurifyCheckResponseWrapper.Response, nil
}
