// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// spawn HTTP server
//
// Command:
// $ goa gen github.com/proepkes/speeddate/src/spawnsvc/design

package server

import (
	"context"
	"net/http"

	spawn "github.com/proepkes/speeddate/src/spawnsvc/gen/spawn"
	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
)

// Server lists the spawn service endpoint HTTP handlers.
type Server struct {
	Mounts []*MountPoint
	New    http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the spawn service endpoints.
func New(
	e *spawn.Endpoints,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"New", "POST", "/spawn"},
		},
		New: NewNewHandler(e.New, mux, dec, enc, eh),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "spawn" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.New = m(s.New)
}

// Mount configures the mux to serve the spawn endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountNewHandler(mux, h.New)
}

// MountNewHandler configures the mux to serve the "spawn" service "new"
// endpoint.
func MountNewHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/spawn", f)
}

// NewNewHandler creates a HTTP handler which loads the HTTP request and calls
// the "spawn" service "new" endpoint.
func NewNewHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		encodeResponse = EncodeNewResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "new")
		ctx = context.WithValue(ctx, goa.ServiceKey, "spawn")

		res, err := endpoint(ctx, nil)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}
