package main

import (
	"fmt"
	"log"

	"ApuestaTotal/cmd/providers"
	"ApuestaTotal/config"
	"ApuestaTotal/internal/products/infrastructure/router"
	"github.com/labstack/echo/v4"
)

var (
	serverPort = config.Environments().ServerPort
	serverHost = config.Environments().ServerHost
)

func main() {
	container := providers.BuildContainer()
	err := container.Invoke(func(router *router.Router, server *echo.Echo) {

		router.Init()
		server.Logger.Fatal(server.Start(fmt.Sprintf("%s:%d", serverHost, serverPort)))
	})

	if err != nil {
		log.Panic(err)
	}
}
