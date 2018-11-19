// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// matchmaking client
//
// Command:
// $ goa gen github.com/proepkes/speeddate/src/mmsvc/design

package matchmaking

import (
	"context"

	goa "goa.design/goa"
)

// Client is the "matchmaking" service client.
type Client struct {
	InsertEndpoint goa.Endpoint
}

// NewClient initializes a "matchmaking" service client given the endpoints.
func NewClient(insert goa.Endpoint) *Client {
	return &Client{
		InsertEndpoint: insert,
	}
}

// Insert calls the "insert" endpoint of the "matchmaking" service.
func (c *Client) Insert(ctx context.Context) (res string, err error) {
	var ires interface{}
	ires, err = c.InsertEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(string), nil
}
