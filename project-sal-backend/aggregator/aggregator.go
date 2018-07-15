package aggregator

import (
	"log"
	"net/http"
	"time"

	"github.com/davecusatis/project-sal-backend/project-sal-backend/models"
	"github.com/davecusatis/project-sal-backend/project-sal-backend/pubsub"
)

// Aggregator is a message aggregator as to not eat through pubsub rate limits
type Aggregator struct {
	MessageChan chan *models.SongRequestMessage
	Ticker      *time.Ticker
	Pubsub      *pubsub.PubsubClient
}

// NewAggregator returns an instance of aggregator
func NewAggregator() *Aggregator {
	ps := pubsub.NewPubsubClient(&http.Client{})

	return &Aggregator{
		MessageChan: make(chan *models.SongRequestMessage),
		Ticker:      time.NewTicker(1 * time.Second),
		Pubsub:      ps,
	}
}

// Start begins the loop that aggregates and sends messages
func (a *Aggregator) Start() {
	go func() {
		for {
			select {
			case <-a.Ticker.C:
				msg := <-a.MessageChan
				log.Printf("Got message: %v", msg)
				a.Pubsub.SendPubsubBroadcastMessage(msg)
			}
		}
	}()
}
