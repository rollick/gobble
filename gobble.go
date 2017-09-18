// Package gobble is for Bubble Post API access (partial) using token authentication
package gobble

import "github.com/rollick/gobble/services"

//
// Client to wrap services
//

// Client is a tiny Bubble Post API client
type Client struct {
	DeliveryService *services.DeliveryService
	// TODO: Other service endpoints to be added
}

// NewClient returns a new Client
func NewClient(accessToken string) *Client {
	return &Client{
		DeliveryService: services.NewDeliveryService(accessToken),
		// TODO: Other service endpoints to be added
	}
}
