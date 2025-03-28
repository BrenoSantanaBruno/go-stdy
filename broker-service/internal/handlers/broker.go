package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/brenosantanabruno/go-stdy/config"
)

type JSONResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type BrokerHandler struct {
	cfg *config.Config
}

func NewBrokerHandler(cfg *config.Config) *BrokerHandler {
	return &BrokerHandler{cfg: cfg}
}

func (h *BrokerHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	payload := JSONResponse{
		Error:   false,
		Message: "Broker service reached successfully",
		Data: map[string]string{
			"status": "active",
			"time":   time.Now().Format(time.RFC3339),
		},
	}

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		h.cfg.ErrorLog.Printf("Failed to encode JSON: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	h.cfg.InfoLog.Println("Successfully processed broker request")
}
