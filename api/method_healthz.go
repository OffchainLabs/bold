package api

import "net/http"

// healthzHandler returns OK if ready to serve requests.
func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
