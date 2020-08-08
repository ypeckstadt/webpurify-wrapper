package request

import (
	"encoding/json"
	"github.com/ypeckstadt/webpurify-wrapper/wrapper"
	"github.com/ypeckstadt/webpurify-wrapper/wrapper/response"
	"io/ioutil"
	"net/http"
)

// Replace is a profanity search and replace method. Returns the number of profane words found and the submitted text with profane words replaced with symbol provided. If the text is clean 0 (zero) is returned.
func (w *WebPurifyRequest) Replace(
	language wrapper.WebPurifyLanguage,
	text string,
	replaceSymbol string,
	filterEmail bool,
	filterPhone bool,
	filterLink bool,
	useCDATA bool,
	) (response.WebPurifyReplaceResponse, error) {

	// create request parameter collection
	requestParameters := []wrapper.WebPurifyRequestParameter{
		{Type: wrapper.Text, Value: text},
		{Type: wrapper.ReplaceSymbol, Value: replaceSymbol},
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
	request, err := w.createRequest(Replace, requestParameters...)
	if err != nil {
		return response.WebPurifyReplaceResponse{}, err
	}

	// create http client
	client := w.createHttpClient()

	// Execute http request
	httpResponse, err := client.Do(request)
	if err != nil {
		return response.WebPurifyReplaceResponse{}, err
	}

	// convert response to struct
	webPurifyResponse, err := w.convertToReplaceResponse(*httpResponse)
	if err != nil {
		return response.WebPurifyReplaceResponse{}, err
	}

	return webPurifyResponse, nil
}

// TODO: convert to generic function when Generics are finally available
func (w *WebPurifyRequest) convertToReplaceResponse(resp http.Response) (response.WebPurifyReplaceResponse, error) {
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return response.WebPurifyReplaceResponse{}, err
	}

	responseWrapper := response.WebPurifyReplaceResponseWrapper{}

	err = json.Unmarshal(body, &responseWrapper)
	if err != nil {
		return response.WebPurifyReplaceResponse{}, err
	}

	return responseWrapper.Response, nil
}
