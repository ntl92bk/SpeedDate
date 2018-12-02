// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// gameserver client HTTP transport
//
// Command:
// $ goa gen github.com/proepkes/speeddate/src/spawnsvc/design

package client

import (
	"context"
	"net/http"

	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
)

// Client lists the gameserver service endpoint HTTP clients.
type Client struct {
	// Configure Doer is the HTTP client used to make requests to the configure
	// endpoint.
	ConfigureDoer goahttp.Doer

	// CORS Doer is the HTTP client used to make requests to the  endpoint.
	CORSDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the gameserver service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		ConfigureDoer:       doer,
		CORSDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Configure returns an endpoint that makes HTTP requests to the gameserver
// service configure server.
func (c *Client) Configure() goa.Endpoint {
	var (
		decodeResponse = DecodeConfigureResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildConfigureRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ConfigureDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("gameserver", "configure", err)
		}
		return decodeResponse(resp)
	}
}
