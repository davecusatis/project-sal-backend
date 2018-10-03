package messages

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/davecusatis/project-sal-backend/project-sal-backend/models"
)

// TODO: make an interface in which these implement
func newExtensionsChatMessage(token *models.TokenData, data []byte) *http.Request {
	r, _ := http.NewRequest("POST",
		fmt.Sprintf("https://api.twitch.tv/extensions/6uwsgp1z9ymm816pyd0a7ga8zdir1n/0.0.1/channels/%s/chat", token.ChannelID),
		bytes.NewReader(data))

	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.Token))
	r.Header.Add("Client-Id", "6uwsgp1z9ymm816pyd0a7ga8zdir1n")
	r.Header.Add("Content-Type", "application/json")
	return r
}

type extChatMessage struct {
	Text string `json:"text"`
}

func (p *PubsubClient) SendExtensionChatMessage(message *models.ChatMessage) error {
	postData, _ := json.Marshal(&extChatMessage{
		Text: message.Message,
	})
	req := newExtensionsChatMessage(message.Token, postData)

	resp, err := p.Client.Do(req)
	if err != nil {
		e := fmt.Errorf("Error sending ext chat message: %s", err)
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
