package request

import (
	"encoding/json"
	"github.com/ypeckstadt/webpurify-wrapper/wrapper"
	"github.com/ypeckstadt/webpurify-wrapper/wrapper/response"
	"io/ioutil"
	"net/http"
)

// CheckCount is a profanity checking method. Returns the number of profane words found in the submitted text. If the text is clean 0 (zero) is returned.
func (w *WebPurifyRequest) CheckCount(language wrapper.WebPurifyLanguage,text string, filterEmail bool, filterPhone bool, filterLink bool) (response.WebPurifyCheckCountResponse, error) {

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
	request, err := w.createRequest(CheckCount, requestParameters...)
	if err != nil {
		return response.WebPurifyCheckCountResponse{}, err
	}

	// create http client
	client := w.createHttpClient()

	// Execute http request
	httpResponse, err := client.Do(request)
	if err != nil {
		return response.WebPurifyCheckCountResponse{}, err
	}

	// convert response to struct
	webPurifyResponse, err := w.convertToCheckCountResponse(*httpResponse)
	if err != nil {
		return response.WebPurifyCheckCountResponse{}, err
	}

	return webPurifyResponse, nil
}

// TODO: convert to generic function when Generics are finally available
func (w *WebPurifyRequest) convertToCheckCountResponse(resp http.Response) (response.WebPurifyCheckCountResponse, error) {
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return response.WebPurifyCheckCountResponse{}, err
	}

	responseWrapper := response.WebPurifyCheckCountResponseWrapper{}

	err = json.Unmarshal(body, &responseWrapper)
	if err != nil {
		return response.WebPurifyCheckCountResponse{}, err
	}

	return responseWrapper.Response, nil
}
