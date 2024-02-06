package external

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type SSLLabsClient struct {
	httpClient *http.Client
}

func NewSSLLabsClient() *SSLLabsClient {
	return &SSLLabsClient{
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (c *SSLLabsClient) GetSSLInfo(url string) (map[string]interface{}, error) {
	baseSslLabsURL := os.Getenv("SSL_LABS_BASE_URL")
	resp, err := c.httpClient.Get(fmt.Sprintf("%s%s", baseSslLabsURL, url))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch SSL Labs info for URL %s: %v", url, err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode SSL Labs info for URL %s: %v", url, err)
	}

	return result, nil
}
