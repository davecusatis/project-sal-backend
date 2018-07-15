package pubsub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/davecusatis/project-sal-backend/project-sal-backend/models"
)

type PubsubClient struct {
	Client *http.Client
}

// NewPubsubClient returns an instance of our pubsub client
func NewPubsubClient(client *http.Client) *PubsubClient {
	return &PubsubClient{
		Client: client,
	}
}

func newPubsubMessageRequest(token *models.TokenData, data []byte) *http.Request {
	r, _ := http.NewRequest("POST",
		fmt.Sprintf("https://api.twitch.tv/extensions/message/%s", token.ChannelID),
		bytes.NewReader(data))

	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.Token))
	r.Header.Add("Client-Id", "cm5nkhrq18yy02yy9tp108lx745vcx")
	r.Header.Add("Content-Type", "application/json")
	return r
}

// SendPubsubBroadcastMessage sends a pubsub message to twitch broadcast topic
func (p *PubsubClient) SendPubsubBroadcastMessage(message *models.PubsubMessage) error {
	srMessage, _ := json.Marshal(message)
	postData, _ := json.Marshal(&models.PostData{
		ContentType: "application/json",
		Message:     string(srMessage),
		Targets:     []string{"broadcast"},
	})

	req := newPubsubMessageRequest(message.Token, postData)
	resp, err := p.Client.Do(req)
	if err != nil {
		e := fmt.Errorf("Error sending pubsub message: %s", err)
		log.Println(e)
		return e
	}
	if resp.StatusCode != http.StatusNoContent {
		e := fmt.Errorf("Error from twitch API: expected 204 got %d, %s", resp.StatusCode, resp.Status)
		log.Println(e)
		return e
	}

	return nil
}
