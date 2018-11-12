package design

import . "goa.design/goa/http/design"
import . "goa.design/goa/http/dsl"

var _ = Service("authorize", func() {
	Description("The service makes it possible to ...")

	HTTP(func() {
		Path("/")
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
			Description("Result defines a single field which is the sum.")
			Attribute("auth", String, "Resulting sum")
			Required("auth")
		})
		HTTP(func() {
			POST("/login")
			// Use Authorization header to provide basic auth value.
			Response(StatusNoContent, func() {
				Header("auth:Authorization", String, "Generated JWT")
			})
		})
	})
})