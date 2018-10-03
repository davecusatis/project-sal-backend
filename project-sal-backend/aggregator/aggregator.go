package aggregator

import (
	"log"
	"net/http"
	"time"

	"github.com/davecusatis/project-sal-backend/project-sal-backend/messages"
	"github.com/davecusatis/project-sal-backend/project-sal-backend/models"
)

// Aggregator is a message aggregator as to not eat through pubsub rate limits
type Aggregator struct {
	MessageChan     chan *models.PubsubMessage
	ChatMessageChan chan *models.ChatMessage
	PubsubTicker    *time.Ticker
	ChatTicker      *time.Ticker
	Pubsub          *messages.PubsubClient
}

// NewAggregator returns an instance of aggregator
func NewAggregator() *Aggregator {
	ps := messages.NewPubsubClient(&http.Client{})

	return &Aggregator{
		MessageChan:     make(chan *models.PubsubMessage),
		ChatMessageChan: make(chan *models.ChatMessage),
		PubsubTicker:    time.NewTicker(1 * time.Second),
		ChatTicker:      time.NewTicker(15 * time.Second),
		Pubsub:          ps,
	}
}

// Start begins the loop that aggregates and sends messages
func (a *Aggregator) Start() {
	go func() {
		for {
			select {
			case <-a.PubsubTicker.C:
				msg := <-a.MessageChan
				log.Printf("Sending message: %v", msg)
				a.Pubsub.SendPubsubBroadcastMessage(msg)
			}
		}
	}()

	go func() {
		for {
			select {
			case <-a.ChatTicker.C:
				msg := <-a.ChatMessageChan
				log.Printf("Sending chat message: %v", msg)
				a.Pubsub.SendExtensionChatMessage(msg)
			}
		}
	}()
}
