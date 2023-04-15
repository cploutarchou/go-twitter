package stream

import (
	"io"
	"net/http"
)

type Stream interface {
	StreamSample() ([]byte, error)
}

type streamImpl struct {
	Url    string
	Method string
	Auth   string
	Cookie string
	Client *http.Client
}

func NewStream(ulr, method, bearenToken, Cookie string) Stream {
	return &streamImpl{
		Url:    ulr,
		Method: method,
		Auth:   bearenToken,
		Cookie: Cookie,
	}
}

func (s *streamImpl) StreamSample() ([]byte, error) {
	req, err := http.NewRequest(s.Method, s.Url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", s.Auth)
	req.Header.Add("Cookie", s.Cookie)
	
	res, err := s.Client.Do(req)
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
