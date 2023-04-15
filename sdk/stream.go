package sdk

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
	
	"io"
	"net/http"
	"net/url"
)

const sampleURL = "https://api.twitter.com/2/tweets/sample/stream?"

type FilterFields struct {
	TweetFields     []string   // Tweet fields to include in the response
	Expansions      []string   // Expandable fields to include in the response
	MediaFields     []string   // Media fields to include in the response
	PollFields      []string   // Poll fields to include in the response
	PlaceFields     []string   // Place fields to include in the response
	UserFields      []string   // User fields to include in the response
	BackfillMinutes *int       // Number of minutes to backfill results
	StartTime       *time.Time // Start time for filtering by time range
	EndTime         *time.Time // End time for filtering by time range
}

type Stream interface {
	StreamSample() ([]byte, error)
	Close() error
	Filter(strings FilterFields) error
	Next() ([]Tweet, interface{})
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

func (s *streamImpl) Filter(filter FilterFields) error {
	queryParams := make(url.Values)
	
	// Set the tweet fields to include in the response
	if len(filter.TweetFields) > 0 {
		queryParams.Set("tweet.fields", strings.Join(filter.TweetFields, ","))
	}
	
	// Set the expandable fields to include in the response
	if len(filter.Expansions) > 0 {
		queryParams.Set("expansions", strings.Join(filter.Expansions, ","))
	}
	
	// Set the media fields to include in the response
	if len(filter.MediaFields) > 0 {
		queryParams.Set("media.fields", strings.Join(filter.MediaFields, ","))
	}
	
	// Set the poll fields to include in the response
	if len(filter.PollFields) > 0 {
		queryParams.Set("poll.fields", strings.Join(filter.PollFields, ","))
	}
	
	// Set the place fields to include in the response
	if len(filter.PlaceFields) > 0 {
		queryParams.Set("place.fields", strings.Join(filter.PlaceFields, ","))
	}
	
	// Set the user fields to include in the response
	if len(filter.UserFields) > 0 {
		queryParams.Set("user.fields", strings.Join(filter.UserFields, ","))
	}
	
	// Set the number of minutes to backfill results
	if filter.BackfillMinutes != nil {
		queryParams.Set("backfill_minutes", strconv.Itoa(*filter.BackfillMinutes))
	}
	
	// Set the time range to filter by
	if filter.StartTime != nil {
		queryParams.Set("start_time", filter.StartTime.Format(time.RFC3339))
	}
	if filter.EndTime != nil {
		queryParams.Set("end_time", filter.EndTime.Format(time.RFC3339))
	}
	
	// Build the URL for the filtered stream
	filteredURL := sampleURL + queryParams.Encode()
	
	// Create a new HTTP request for the filtered stream
	req, err := http.NewRequest("GET", filteredURL, nil)
	if err != nil {
		return err
	}
	// check if authorization is not empty
	if s.bearer == "" {
		return fmt.Errorf("bearer token is empty")
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.bearer))
	if s.cookie == nil {
		// create a new cookie
		s.cookie = &http.Cookie{
			Name:  "personalization_id",
			Value: fmt.Sprintf("v1_%s", RandString(10)),
		}
	}
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

func (s *streamImpl) Next() ([]Tweet, interface{}) {
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
	var tweet []Tweet
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
