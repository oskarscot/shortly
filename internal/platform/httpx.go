package httpx

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorEnvelope struct {
	Message string `json:"message"`
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Printf("failed to encode %v: %v", data, err)
	}
}

func Error(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, ErrorEnvelope{Message: message})
}

func JSON(w http.ResponseWriter, status int, data any) {
	writeJSON(w, status, data)
}
