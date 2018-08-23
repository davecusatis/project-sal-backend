package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/davecusatis/project-sal-backend/project-sal-backend/token"
)

// GetScores is the endpoint that gets scores from DB and returns them
func (a *API) GetScores(w http.ResponseWriter, req *http.Request) {
	// validate token
	tok, err := token.ExtractAndValidateTokenFromHeader(req.Header)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err)))
		return
	}
	// s := slotmachine.GenerateRandomScore()

	// 1. Record score
	// 2. Update clients
	// 3. message in chat
	scores, err := a.Datasource.LeaderboardForChannelID(tok.ChannelID)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err)))
		return
	}
	log.Printf("Got scores: %#v", scores)
	// a.Aggregator.MessageChan <- &models.PubsubMessage{
	// 	MessageType: "scoreUpdated",
	// 	Data: models.MessageData{
	// 		Score: s,
	// 	},
	// 	Token: token.CreateServerToken(tok),
	// }
	out, _ := json.Marshal(scores)
	w.Write([]byte(out))
}
