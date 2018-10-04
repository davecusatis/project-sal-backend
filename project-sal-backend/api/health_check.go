package api

import (
	"net/http"
)

// HealthCheck is the health check endpoint
func (a *API) HealthCheck(w http.ResponseWriter, req *http.Request) {

	w.Write([]byte("OK"))
}
