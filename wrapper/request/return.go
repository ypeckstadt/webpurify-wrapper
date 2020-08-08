package request

import (
	"encoding/json"
	"github.com/ypeckstadt/webpurify-wrapper/wrapper"
	"github.com/ypeckstadt/webpurify-wrapper/wrapper/response"
	"io/ioutil"
	"net/http"
)

// Return is a profanity search method. Returns the number of profane words found and a list of the profane words. If the text is clean 0 (zero) is returned.
func (w *WebPurifyRequest) Return(
	language wrapper.WebPurifyLanguage,
	text string,
	filterEmail bool,
	filterPhone bool,
	filterLink bool,
	useCDATA bool,
	) (response.WebPurifyReturnResponse, error) {

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

	if useCDATA {
		requestParameters = append(requestParameters, wrapper.WebPurifyRequestParameter{Type: wrapper.CDATA, Value: "1"})
	}

	// build http request
	request, err := w.createRequest(Return, requestParameters...)
	if err != nil {
		return response.WebPurifyReturnResponse{}, err
	}

	// create http client
	client := w.createHttpClient()

	// Execute http request
	httpResponse, err := client.Do(request)
	if err != nil {
		return response.WebPurifyReturnResponse{}, err
	}

	// convert response to struct
	webPurifyResponse, err := w.convertToReturnResponse(*httpResponse)
	if err != nil {
		return response.WebPurifyReturnResponse{}, err
	}

	return webPurifyResponse, nil
}

func (w *WebPurifyRequest) convertToReturnResponse(resp http.Response) (response.WebPurifyReturnResponse, error) {
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return response.WebPurifyReturnResponse{}, err
	}

	responseWrapper := response.WebPurifyReturnResponseWrapper{}

	err = json.Unmarshal(body, &responseWrapper)
	if err != nil {
		return response.WebPurifyReturnResponse{}, err
	}

	return responseWrapper.Response, nil
}
