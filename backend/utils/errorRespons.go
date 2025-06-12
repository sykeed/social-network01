package utils

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, statusCode int, message string,st interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(map[string]interface{}{
		"message": message,
		"status":  st,
	})
}
