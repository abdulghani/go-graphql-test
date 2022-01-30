package handlers

import (
	"go-graphql-test/src/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gin-gonic/gin"
)

const mb int64 = 1 << 20

func GraphqlHandler(c *gin.Context) {
	handler := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	// add multipart capability to handler
	handler.AddTransport(transport.POST{})
	handler.AddTransport(transport.MultipartForm{
		MaxMemory:     32 * mb,
		MaxUploadSize: 100 * mb,
	})

	// attach context to request
	request := c.Request.WithContext(c)

	handler.ServeHTTP(c.Writer, request)
}
