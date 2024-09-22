package webhook_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	URL     string
	AuthKey string
}

type Request struct {
	To      string `json:"to"`
	Content string `json:"content"`
}

type Response struct {
	Message   string `json:"message"`
	MessageID string `json:"messageId"`
}

func NewClient(url, authKey string) *Client {
	return &Client{
		URL:     url,
		AuthKey: authKey,
	}
}

func (c *Client) Send(to, content string) (*Response, error) {
	req := Request{
		To:      to,
		Content: content,
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequest("POST", c.URL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("x-ins-auth-key", c.AuthKey) // This may not be real auth key, not sure.

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()

	/* if resp.StatusCode != http.StatusAccepted {
		return nil, fmt.Errorf("unexpected status code: %d, was expecting %d", resp.StatusCode, http.StatusAccepted)
	} */

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("unexpected status code: %d, was expecting a value between 200 and 300", resp.StatusCode)
	}

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &response, nil
}
