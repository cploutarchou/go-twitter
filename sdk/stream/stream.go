package stream

import (
	"fmt"
	"io"
	"net/http"
)

const sampleURL = "https://api.twitter.com/2/tweets/sample/stream?"

type Stream interface {
	StreamSample() ([]byte, error)
}

type streamImpl struct {
	client *http.Client
	bearer string
	cookie *http.Cookie
}

func NewStream(bearerToken string, cookie *http.Cookie, client *http.Client) (Stream, error) {
	if client == nil {
		return nil, fmt.Errorf("client is not initialized")
	}
	return &streamImpl{
		bearer: bearerToken,
		cookie: cookie,
		client: client,
	}, nil
}

func (s *streamImpl) StreamSample() ([]byte, error) {
	req, err := http.NewRequest("GET", sampleURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", s.bearer)
	req.Header.Add("Cookie", fmt.Sprintf("%s=%s", s.cookie.Name, s.cookie.Value))
	
	res, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	
	return body, nil
}
