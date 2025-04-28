package pkg

import (
	"encoding/json"
	"net/http"
)

func SetReponse(w http.ResponseWriter, statusCode int, headers map[string]string, writeMessage string, jsonData interface{}) {
	w.WriteHeader(statusCode)

	for key, value := range headers {
		w.Header().Set(key, value)
	}

	w.Write([]byte(writeMessage))

	json.NewEncoder(w).Encode(&jsonData)
}
