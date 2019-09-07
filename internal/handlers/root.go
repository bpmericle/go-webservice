package handlers

import (
	"fmt"
	"net/http"
)

// Root returns a HTTP 200 status code and a hello message
func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello there friend, I hope you are doing well")
	w.WriteHeader(http.StatusOK)
}
