// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// gameserver endpoints
//
// Command:
// $ goa gen github.com/proepkes/speeddate/src/spawnsvc/design

package gameserver

import (
	"context"

	goa "goa.design/goa"
)

// Endpoints wraps the "gameserver" service endpoints.
type Endpoints struct {
	Configure goa.Endpoint
}

// NewEndpoints wraps the methods of the "gameserver" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Configure: NewConfigureEndpoint(s),
	}
}

// Use applies the given middleware to all the "gameserver" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Configure = m(e.Configure)
}

// NewConfigureEndpoint returns an endpoint function that calls the method
// "configure" of service "gameserver".
func NewConfigureEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Configure(ctx)
	}
}
