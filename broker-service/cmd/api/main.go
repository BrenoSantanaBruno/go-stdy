package main

import (
	"fmt"
	"net/http"

	"../broker-service/internal/config"
)

const webPort = "80"

func main() {
	cfg := config.New()

	cfg.InfoLog.Printf("Starting broker service on port %s", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: api.Routes(cfg),
	}

	err := srv.ListenAndServe()
	if err != nil {
		cfg.ErrorLog.Panic(err)
	}
}
