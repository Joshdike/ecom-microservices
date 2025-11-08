package main

import (
	"product-service/internal/config"
	"product-service/internal/server/rest"
)

func main() {

	// Instantiate the server object
	cnf := config.Config{
		Port: ":8090",
	}
	s := rest.NewServer(cnf)
	s.Start()

}
