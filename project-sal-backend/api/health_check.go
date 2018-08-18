package api

import (
	"net/http"
)

// Ping is the health check endpoint
func (a *API) HealthCheck(w http.ResponseWriter, req *http.Request) {
	// validate token
	// _, err := token.ExtractAndValidateTokenFromHeader(req.Header)
	// if err != nil {
	// 	w.Write([]byte(fmt.Sprintf("error %s", err)))
	// 	return
	// }
	w.Write([]byte("OK"))
}
