package rest

import (
	"encoding/json"
	"net/http"
)

func Health() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		response := map[string]string{"status": "ok"}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			return
		}
		_, err = w.Write(jsonResponse)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
