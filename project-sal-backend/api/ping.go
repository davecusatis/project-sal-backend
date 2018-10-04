package api

import (
	"fmt"
	"net/http"

	"github.com/davecusatis/project-sal-backend/project-sal-backend/models"
	"github.com/davecusatis/project-sal-backend/project-sal-backend/token"
)

// Ping is the health check endpoint
func (a *API) Ping(w http.ResponseWriter, req *http.Request) {
	// validate token
	tok, err := token.ExtractAndValidateTokenFromHeader(req.Header)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err)))
		return
	}
	a.Aggregator.QueuePubsubMessage(tok.ChannelID, &models.PubsubMessage{
		MessageType: "ping",
		Data:        models.MessageData{},
		Token:       token.CreateServerToken(tok),
	})
	w.Write([]byte("OK"))
}
