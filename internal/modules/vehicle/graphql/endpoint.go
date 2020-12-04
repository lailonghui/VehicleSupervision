package vehicle

import (
	"VehicleSupervision/internal/modules/vehicle/graphql/graph/generated"
	"VehicleSupervision/internal/modules/vehicle/graphql/graph/resolvers"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
)

func endpoint() *handler.Server {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}}))
	return srv

}

func GinEndpoint() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	srv := endpoint()

	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}
