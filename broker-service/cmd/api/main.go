package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct {
}

func main() {
	app := Config{}
	log.Println("Starting broker service on port:", webPort)

	// define hhtp server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: routes(&app),
	}
	// start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
