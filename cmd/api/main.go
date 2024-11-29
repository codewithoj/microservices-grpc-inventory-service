package main

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct{}

const webPort = "8082"

func main() {
	app := &Config{}

	// define the server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	log.Println("starting inventory service on port: ", webPort)

	// listen the server

	err := srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
