package devtoclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"os"
	"time"
)

// Client struct for defaults
type Client struct {
	BaseURL    string
	Key        string
	HTTPClient *http.Client
	RetryCount int
}

// ClientConfig struct for settings
type ClientConfig struct {
	BaseURL string
	Key     string
}

// NewClient method create a Client object from Client struct
func NewClient() *Client {
	baseURL := os.Getenv("DEVTO_BASE_URL")
	if baseURL == "" {
		baseURL = "https://dev.to/api/"
	}

	client := Client{
		BaseURL: baseURL,
		Key:     os.Getenv("DEVTO_API_KEY"),
		HTTPClient: &http.Client{
			Timeout: 15 * time.Second,
		},
		RetryCount: 0,
	}

	return &client
}

// UpdateConfig method create a ClientConfig object if not exists
func (c *Client) UpdateConfig(config *ClientConfig) {
	baseURL := config.BaseURL
	key := config.Key
	if baseURL != "" {
		c.BaseURL = baseURL
	}
	if key != "" {
		c.Key = key
	}
}

// Request and request methods create httpRequest
func (c *Client) Request(method string, url string, params, result interface{}) (res *http.Response, err error) {
	for i := 0; i <= c.RetryCount+1; i++ {
		retryDuration := time.Duration((math.Pow(2, float64(i))-1)/2*1000) * time.Millisecond
		time.Sleep(retryDuration)

		res, err = c.request(method, url, params, result)
		if res != nil && res.StatusCode == 429 {
			continue
		} else {
			break
		}
	}
	return res, err
}

func (c *Client) request(method string, url string, params, result interface{}) (res *http.Response, err error) {
	body := bytes.NewReader(make([]byte, 0))

	if params != nil {
		payloadBytes, err := json.Marshal(params)
		if err != nil {
			fmt.Println("test", err)
		}
		body = bytes.NewReader(payloadBytes)
	}
	fullURL := fmt.Sprintf("%s%s", c.BaseURL, url)
	req, err := http.NewRequest(method, fullURL, body)

	if err != nil {
		return res, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Api-Key", c.Key)
	req.Header.Add("User-Agent", "Go Dev.to API Client")
	res, err = c.HTTPClient.Do(req)
	if err != nil {
		return res, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		defer res.Body.Close()
		var devToError error
		decoder := json.NewDecoder(res.Body)
		if err := decoder.Decode(&devToError); err != nil {
			return res, err
		}

		return res, error(devToError)
	}

	if result != nil {
		decoder := json.NewDecoder(res.Body)
		if err = decoder.Decode(result); err != nil {
			return res, err
		}
	}
	return res, nil
}
