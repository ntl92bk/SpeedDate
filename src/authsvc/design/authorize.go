package design

import (
	. "goa.design/goa/dsl"
	cors "goa.design/plugins/cors/dsl" // Use CORS plugin
)

var _ = Service("authorize", func() {
	Description("The service makes it possible to generate login-tokens for valid authentification")

	// Sets CORS response headers for requests with any Origin header
	cors.Origin("*", func() {
		cors.Headers("Authorization")
		cors.Methods("OPTIONS", "POST")
		cors.Expose("Access-token")
		cors.MaxAge(600)
	})

	HTTP(func() {
		Path("/auth")
	})

	Method("login", func() {
		Description("Creates a valid JWT")

		// The signin endpoint is secured via basic auth
		Security(BasicAuth)

		Payload(func() {
			Description("Credentials used to authenticate to retrieve JWT token")
			Username("username", String, "Username used to perform signin", func() {
				Example("user")
			})
			Password("password", String, "Password used to perform signin", func() {
				Example("password")
			})
			Required("username", "password")
		})
		Result(func() {
			Description("Result defines a JWT Token.")
			Attribute("token", String, "Resulting token")
			Required("token")
		})
		HTTP(func() {
			POST("/login")
			// Use Authorization header to provide basic auth value.
			Response(StatusNoContent, func() {
				Header("token:Access-token", String, "Generated JWT")
			})
		})
	})
})
