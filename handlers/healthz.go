package handlers

import (
	"fmt"
	"net/http"
)

// Healthz check
func Healthz(w http.ResponseWriter, r *http.Request) {
	responseBody := fmt.Sprintf("%s OK", r.URL.Path)
	_, err := w.Write([]byte(responseBody))

	if err != nil {
		panic(err)
	}
}
