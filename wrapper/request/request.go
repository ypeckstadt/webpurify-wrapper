package request

import (
	"github.com/ypeckstadt/webpurify-wrapper/wrapper"
	"github.com/ypeckstadt/webpurify-wrapper/wrapper/response"
	"net/http"
)


type WebPurifyRequest struct {
	URL    string
	APIKey string
}

func (w *WebPurifyRequest) createRequest(check WebPurifyRequestMethod, parameters ...wrapper.WebPurifyRequestParameter) (*http.Request, error) {
	request, err := http.NewRequest("POST", w.URL, nil)
	if err != nil {
		return nil, err
	}

	// build new query
	q := request.URL.Query()

	// add default URL parameters
	q.Add(string(wrapper.Format), string(response.JSON))
	q.Add(string(wrapper.APIKey), w.APIKey)
	q.Add(string(wrapper.Method), string(check))

	// add request based parameters
	for _, s := range parameters {
		q.Add(string(s.Type), s.Value)
	}

	// encode to raw query
	request.URL.RawQuery = q.Encode()

	return request, nil
}

func (w *WebPurifyRequest) createHttpClient() http.Client {
	return http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}
}
