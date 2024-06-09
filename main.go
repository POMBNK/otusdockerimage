package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func healthCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		statusOK := map[string]interface{}{
			"status": "OK",
		}
		statusJSON, err := json.Marshal(statusOK)
		if err != nil {
			slog.Error("Decoding error", err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(statusJSON)
	}

}

func main() {
	slog.Info("Starting server...")
	m := http.NewServeMux()
	m.Handle("GET /health/", healthCheck())
	slog.Info("Ready to handle requests")
	http.ListenAndServe("0.0.0.0:8000", m)
}
