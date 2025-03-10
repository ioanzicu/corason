package health

import (
	"encoding/json"
	"log"
	"net/http"
)

const HealthPath = "/health"

type Health struct{}

func NewHealth() Health {
	return Health{}
}

func (h Health) Handle() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = "OK"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	})
}
