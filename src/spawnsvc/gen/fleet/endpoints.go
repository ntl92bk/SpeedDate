// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// fleet endpoints
//
// Command:
// $ goa gen github.com/proepkes/speeddate/src/spawnsvc/design

package fleet

import (
	"context"

	goa "goa.design/goa"
)

// Endpoints wraps the "fleet" service endpoints.
type Endpoints struct {
	Add       goa.Endpoint
	Clear     goa.Endpoint
	Configure goa.Endpoint
}

// NewEndpoints wraps the methods of the "fleet" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Add:       NewAddEndpoint(s),
		Clear:     NewClearEndpoint(s),
		Configure: NewConfigureEndpoint(s),
	}
}

// Use applies the given middleware to all the "fleet" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Add = m(e.Add)
	e.Clear = m(e.Clear)
	e.Configure = m(e.Configure)
}

// NewAddEndpoint returns an endpoint function that calls the method "add" of
// service "fleet".
func NewAddEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Add(ctx)
	}
}

// NewClearEndpoint returns an endpoint function that calls the method "clear"
// of service "fleet".
func NewClearEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Clear(ctx)
	}
}

// NewConfigureEndpoint returns an endpoint function that calls the method
// "configure" of service "fleet".
func NewConfigureEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Configure(ctx)
	}
}