package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	// Set headers first
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	payload := jsonResponse{
		Error:   false,
		Message: "Broker service reached successfully",
		Data: map[string]string{
			"status": "active",
			"time":   time.Now().Format(time.RFC3339),
		},
	}

	// Better error handling
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		app.ErrorLog.Println("Failed to encode JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.InfoLog.Println("Successfully processed broker request")
}
func (app *Config) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.InfoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Method, r.URL.Path, r.Proto)
		next.ServeHTTP(w, r)
	})
}
