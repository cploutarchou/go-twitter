package stream

import (
	"encoding/json"
	"fmt"
	"github.com/cploutarchou/go-twitter/sdk"
	"io"
	"net/http"
	"net/url"
)

const sampleURL = "https://api.twitter.com/2/tweets/sample/stream?"

type Stream interface {
	StreamSample() ([]byte, error)
	Close() error
	Filter(strings []string) error
	Next() ([]sdk.Tweet, interface{})
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

func (s *streamImpl) Filter(strings []string) error {
	queryParams := make(url.Values)
	queryParams.Set("tweet.fields", "text")
	
	// Add query params for each string in the slice
	for _, str := range strings {
		queryParams.Add("expansions", "author_id")
		queryParams.Add("user.fields", "username")
		queryParams.Add("hashtags", str)
	}
	
	// Build the URL for the filtered stream
	filteredURL := sampleURL + queryParams.Encode()
	
	// Create a new HTTP request for the filtered stream
	req, err := http.NewRequest("GET", filteredURL, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", s.bearer)
	req.Header.Add("Cookie", fmt.Sprintf("%s=%s", s.cookie.Name, s.cookie.Value))
	
	// Send the HTTP request and return any errors
	res, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	
	// Return an error if the response status code is not OK
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected response code: %d", res.StatusCode)
	}
	
	// TODO: Implement reading and processing the filtered stream data
	
	return nil
}

func (s *streamImpl) Close() error {
	// Create a new HTTP request to close the stream connection
	req, err := http.NewRequest("DELETE", sampleURL, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", s.bearer)
	req.Header.Add("Cookie", fmt.Sprintf("%s=%s", s.cookie.Name, s.cookie.Value))
	
	// Send the HTTP request and return any errors
	res, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	
	// Return an error if the response status code is not OK
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected response code: %d", res.StatusCode)
	}
	
	return nil
}

func (s *streamImpl) Next() ([]sdk.Tweet, interface{}) {
	// Create a new HTTP request to get the next tweet from the stream
	req, err := http.NewRequest("GET", sampleURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", s.bearer)
	req.Header.Add("Cookie", fmt.Sprintf("%s=%s", s.cookie.Name, s.cookie.Value))
	
	// Send the HTTP request and return any errors
	res, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	// Read the next tweet from the stream
	var tweet []sdk.Tweet
	err = json.NewDecoder(res.Body).Decode(&tweet)
	if err != nil {
		return nil, err
	}
	
	// TODO: Implement reading and processing the next tweet from the stream data
	// Recommended to use the same approach as in the Filter() method above to read the stream data in chunks to avoid memory issues and improve performance also is good to send data to
	// a channel to be processed by another goroutine
	if err := res.Body.Close(); err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response code: %d", res.StatusCode)
	}
	
	if len(tweet) == 0 {
		return nil, nil
	}
	
	// send the tweet to a channel to be processed by another goroutine and return the channel  to the caller
	// return tweet, tweetChannel
	
	return tweet, nil
}
