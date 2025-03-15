package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func ParseJSON(r *http.Request, v interface{}) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	return json.NewDecoder(r.Body).Decode(v)
}

func WriteJSON(w http.ResponseWriter, status int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	err = WriteJSON(w, status, map[string]string{"error": err.Error()})
	if err != nil {
		return
	}
}

func ExtractYear(dateStr string) (int, error) {
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return 0, err
	}
	return t.Year(), nil
}
