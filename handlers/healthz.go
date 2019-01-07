package handlers

import (
	"fmt"
	"net/http"
)

// Healthz check
func Healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s OK", r.URL.Path)
}
