package main

import (
	"fmt"

	"github.com/CylonSam/go-rest-api/routes"
)

func main() {
	fmt.Println("Server running with Go!")
	routes.HandleRequest()
}
