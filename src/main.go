package main

import (
	"go-graphql-test/src/utils"
	"os"

	_ "github.com/99designs/gqlgen/cmd"
)

func main() {
	utils.LoadEnv()
	server := utils.CreateServer()

	server.Run(":" + os.Getenv("PORT"))
}
