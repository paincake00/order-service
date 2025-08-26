package jsonutil

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

func WriteJSONError(w http.ResponseWriter, status int, message string) error {
	type wrapped struct {
		Error string `json:"error"`
	}

	return WriteJSON(w, status, &wrapped{Error: message})
}

func WriteJSONResponse(w http.ResponseWriter, status int, data any) error {
	type wrapped struct {
		Data any `json:"data"`
	}

	return WriteJSON(w, status, &wrapped{Data: data})
}
