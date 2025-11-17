package picsumclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Image struct {
	ID          string `json:"id"`
	Author      string `json:"author"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	URL         string `json:"url"`
	DownloadUrl string `json:"download_url"`
}

type Client struct {
	http *http.Client
	baseUrl string
}

func NewClient() *Client {
	return &Client{
		http: &http.Client{Timeout: 10 * time.Second},
		baseUrl: "https://picsum.photos",
	}
}


func (c *Client) ListImages(page, Limit int) ([]Image, error) {

	url := fmt.Sprintf("%s/v2/list?page=%d&limit=%d", c.baseUrl, page, Limit)

	resp, err := c.http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("kunde inte nå picsum: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("picsum svarade %d", resp.StatusCode)
	}


	var images []Image

	if err := json.NewDecoder(resp.Body).Decode(&images); err != nil {
		return nil, fmt.Errorf("kunde inte avkoda json: %w", err)
	}

	return images, nil


}
