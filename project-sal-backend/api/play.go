package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/davecusatis/project-sal-backend/project-sal-backend/models"
	"github.com/davecusatis/project-sal-backend/project-sal-backend/slotmachine"
	"github.com/davecusatis/project-sal-backend/project-sal-backend/token"
)

// Play is the endpoint that recieves a request to play the slot game
func (a *API) Play(w http.ResponseWriter, req *http.Request) {
	// validate token
	tok, err := token.ExtractAndValidateTokenFromHeader(req.Header)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err)))
		return
	}
	log.Printf("Got play request")
	a.Aggregator.MessageChan <- &models.PubsubMessage{
		MessageType: "scoreUpdated",
		Data: models.MessageData{
			Score: slotmachine.GenerateRandomScore(),
		},
		Token: token.CreateServerToken(tok),
	}
	w.Write([]byte("OK"))
}
