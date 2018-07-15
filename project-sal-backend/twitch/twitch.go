package twitch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/davecusatis/project-sal-backend/project-sal-backend/models"
)

type TwitchClient struct {
	ClientID string
	Client   *http.Client
}

// NewTwitchClient returns an instance of our Twitch client
func NewTwitchClient(client *http.Client) *TwitchClient {
	return &TwitchClient{
		ClientID: "cm5nkhrq18yy02yy9tp108lx745vcx",
		Client:   client,
	}
}

func (c *TwitchClient) GetLogin(userID string) string {
	r, _ := http.NewRequest("GET", fmt.Sprintf("https://api.twitch.tv/helix/users?id=%s", userID), nil)
	r.Header.Add("Client-Id", c.ClientID)
	r.Header.Add("Content-Type", "application/json")
	resp, err := c.Client.Do(r)
	if err != nil {
		log.Printf("Error getting user data for %s", userID)
	}
	defer resp.Body.Close()

	respBody := new(models.TwitchUserResponse)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body")
	}
	err = json.Unmarshal(body, &respBody)
	if err != nil {
		log.Printf("Error parsing data from twitch")
	}

	return respBody.Data[0].DisplayName
}
