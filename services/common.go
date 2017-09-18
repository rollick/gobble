package services

import (
	"fmt"

	"github.com/dghubble/sling"
)

const (
	baseURL    = "https://api.bubblepost.eu"
	apiVersion = "v1.0"
)

// BubblePostError represents a Bubble Post API error response
type BubblePostError struct {
	Message string `json:"msg"`
}

// TrackingParam is the param for a list / fetch request
// http://apidocs.bubblepost.eu/get_delivery/#definition-get
type TrackingParam struct {
	Trn string `url:"trn,omitempty"`
}

// NewClient returns a new Bubble Post client
func NewClient(accessToken string) *sling.Sling {
	// Create bubble post api client
	client := sling.New().Client(nil).Base(fmt.Sprintf("%s/%s/", baseURL, apiVersion))

	// Add request headers
	client.Set("Authorization", accessToken)
	client.Set("user-agent", "Gobble/1.0 Go/1.8 OpenSSL/1.0.2d")

	return client
}

// Error is a formatted Bubble Post error
func (e BubblePostError) Error() string {
	return fmt.Sprintf("Bubble Post error: %v", e.Message)
}
