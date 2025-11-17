package uploader

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	baseUrl string
	http    *http.Client
}


func NewClient(url string) *Client {
	return &Client{
		baseUrl: url,
		http: &http.Client{Timeout: 10 * time.Second},
	}
}


type uploadBody struct {
	Key string `json:"key"`
	Content string `json:"content"`
}



func (c *Client) Upload(key, content string) error {

	body := uploadBody{
		Key: key,
		Content: content,
	}

	data, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("marshal error: %w", err)
	}

	requestUrl := c.baseUrl + "/upload"

	resp, err := c.http.Post(requestUrl, "application/json",
	bytes.NewReader(data))

	if err != nil {
		return fmt.Errorf("endpoint error %w, ", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("uploader svarade inte: %d", resp.StatusCode)
	}

	return nil
	
}
