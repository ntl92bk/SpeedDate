// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// gameserver client
//
// Command:
// $ goa gen github.com/proepkes/speeddate/src/spawnsvc/design

package gameserver

import (
	"context"

	goa "goa.design/goa"
)

// Client is the "gameserver" service client.
type Client struct {
	ConfigureEndpoint goa.Endpoint
}

// NewClient initializes a "gameserver" service client given the endpoints.
func NewClient(configure goa.Endpoint) *Client {
	return &Client{
		ConfigureEndpoint: configure,
	}
}

// Configure calls the "configure" endpoint of the "gameserver" service.
func (c *Client) Configure(ctx context.Context) (res string, err error) {
	var ires interface{}
	ires, err = c.ConfigureEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(string), nil
}
