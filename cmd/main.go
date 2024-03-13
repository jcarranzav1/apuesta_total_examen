package main

import (
	"fmt"

	"ApuestaTotal/config"
)

var (
	serverPort = config.Environments().ServerPort
)

func main() {
	fmt.Println("Hola, mundo")
}
