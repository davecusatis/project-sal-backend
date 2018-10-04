package aggregator

import (
	"log"
	"time"

	"github.com/davecusatis/project-sal-backend/project-sal-backend/messages"
	"github.com/davecusatis/project-sal-backend/project-sal-backend/models"
)

// TODO: make these self cleaning
// the solution to these dying is like having a one minute ticker that when it recieves a message we return,
// the ticker should reset when a message is received i think it's doable

// GenericMessageSender is a message GenericMessageSender as to not eat through pubsub rate limits
type GenericMessageSender struct {
	MessageChan     chan *models.PubsubMessage
	ChatMessageChan chan *models.ChatMessage
	PubsubTicker    *time.Ticker
	ChatTicker      *time.Ticker
}

// Start begins the loop that aggregates and sends messages
func (gms *GenericMessageSender) Start(messageClient *messages.PubsubClient) {
	go func() {
		for {
			select {
			case <-gms.PubsubTicker.C:
				msg := <-gms.MessageChan
				log.Printf("Sending message: %v", msg)
				messageClient.SendPubsubBroadcastMessage(msg)
			}
		}
	}()

	go func() {
		for {
			select {
			case <-gms.ChatTicker.C:
				msg := <-gms.ChatMessageChan
				log.Printf("Sending chat message: %v", msg)
				messageClient.SendExtensionChatMessage(msg)
			}
		}
	}()
}
