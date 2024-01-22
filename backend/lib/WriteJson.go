package lib

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, status int, v any) interface{} {
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
