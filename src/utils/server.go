package utils

import (
	"go-graphql-test/src/handlers"

	"github.com/gin-gonic/gin"
)

func CreateServer() *gin.Engine {
	server := gin.Default()

	server.GET("/hello", handlers.HelloHandler)
	server.GET("/health", handlers.Health)
	server.GET("/graphql", handlers.GraphiQL)
	server.POST("/graphql", handlers.GraphqlHandler)

	return server
}
