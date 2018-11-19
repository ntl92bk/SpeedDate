package design

import . "goa.design/goa/http/dsl"

var _ = API("spawnsvc", func() {
	Title("Speeddate Spawnerservice")
	Description("HTTP service for managing game-instances")

	Server("spawner", func() {
		Description("Serves repository and swagger service")
		Services("spawn", "swagger")
		Host("localhost", func() {
			Description("Host on local machine")
			URI("http://:8001")
		})
	})
})

// JWTAuth defines a security scheme that uses JWT tokens.
var JWTAuth = JWTSecurity("jwt", func() {
	Description(`Secures endpoint by requiring a valid JWT token retrieved via the signin endpoint. Supports scopes "api:read" and "api:write".`)
	Scope("api:read", "Read-only access")
	Scope("api:write", "Read and write access")
})