package api

import (
	"fmt"
	"net/http"

	"github.com/davecusatis/project-sal-backend/project-sal-backend/token"
)

// Ping is the health check endpoint
func (a *API) HealthCheck(w http.ResponseWriter, req *http.Request) {
	// validate token
	_, err := token.ExtractAndValidateTokenFromHeader(req.Header)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err)))
		return
	}
	w.Write([]byte("OK"))
}