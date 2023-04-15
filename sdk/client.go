package sdk

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

// Client defines the behavior of a Twitter client.
type Client interface {
	Stream() (Stream, error)
	getBearerToken() (string, error)
}

// clientImpl is a concrete implementation of the Client interface.
type clientImpl struct {
	client *http.Client
	config *Config
	once   sync.Once
	mu     sync.Mutex
}

// Config defines the configuration for a Twitter client.
type Config struct {
	Url            string
	ConsumerKey    string
	ConsumerSecret string
	Bearer         string
	Cookie         *http.Cookie
}

// NewClient creates a new Twitter client.
func NewClient(config *Config) (Client, error) {
	cl := &clientImpl{
		config: &Config{
			Url:            config.Url,
			ConsumerKey:    config.ConsumerKey,
			ConsumerSecret: config.ConsumerSecret,
			Cookie:         nil,
		},
	}
	bearer, err := cl.getBearerToken()
	if err != nil {
		return nil, fmt.Errorf("error getting bearer token! Error :  %w", err)
	}
	cl.config.Bearer = bearer
	return cl, nil
}

// Stream returns a new Twitter stream.
func (c *clientImpl) Stream() (Stream, error) {
	c.once.Do(func() {
		c.client = http.DefaultClient
	})
	c.mu.Lock()
	defer c.mu.Unlock()
	return NewStream(c.config.Bearer, c.config.Cookie, c.client)
}

// GetBearerToken returns a bearer token for Twitter API authentication.
func (c *clientImpl) getBearerToken() (string, error) {
	// Create a new HTTP request to get a bearer token.
	req, err := http.NewRequest(http.MethodPost, "https://api.twitter.com/oauth2/token", nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	
	// Encode the consumer key and secret as a base64-encoded string.
	credentials := url.QueryEscape(c.config.ConsumerKey) + ":" + url.QueryEscape(c.config.ConsumerSecret)
	credentialsEncoded := base64.StdEncoding.EncodeToString([]byte(credentials))
	
	// Set the HTTP request headers.
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", credentialsEncoded))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	
	// Set the request body.
	body := "grant_type=client_credentials"
	req.Body = nopCloser{Reader: strings.NewReader(body)}
	req.ContentLength = int64(len(body))
	
	// Send the HTTP request and parse the response.
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	fmt.Println(res.StatusCode)
	defer res.Body.Close()
	
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected response code: %d", res.StatusCode)
	}
	
	var response struct {
		TokenType   string `json:"token_type"`
		AccessToken string `json:"access_token"`
	}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}
	
	return fmt.Sprintf("%s %s", response.TokenType, response.AccessToken), nil
}

// nopCloser is a wrapper for an io.Reader that implements io.ReadCloser.
type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }

// Path: sdk/stream/stream.go
