package aggregator

import (
	"net/http"
	"time"

	"github.com/davecusatis/project-sal-backend/project-sal-backend/messages"
	"github.com/davecusatis/project-sal-backend/project-sal-backend/models"
)

// Aggregator is the struct that contains a map of message senders to rate limit messages on a per channel basis
type Aggregator struct {
	AggregatorMap map[string]*GenericMessageSender
	MessageClient *messages.PubsubClient
}

// NewAggregator returns an instance of aggregator
func NewAggregator() *Aggregator {
	messageClient := messages.NewPubsubClient(&http.Client{})

	return &Aggregator{
		AggregatorMap: make(map[string]*GenericMessageSender),
		MessageClient: messageClient,
	}
}

// NewGenericMessageSender creates a new GMS and attaches it to the aggregator
func (a *Aggregator) NewGenericMessageSender(channelID string) {
	a.AggregatorMap[channelID] = &GenericMessageSender{
		MessageChan:     make(chan *models.PubsubMessage),
		ChatMessageChan: make(chan *models.ChatMessage),
		PubsubTicker:    time.NewTicker(1 * time.Second),
		ChatTicker:      time.NewTicker(15 * time.Second),
	}
	a.AggregatorMap[channelID].Start(a.MessageClient)
}

// QueuePubsubMessage queues up the pubsub message to be send in a specific channel
func (a *Aggregator) QueuePubsubMessage(channelID string, msg *models.PubsubMessage) {
	if gms, ok := a.AggregatorMap[channelID]; ok {
		gms.MessageChan <- msg
		return
	}
	a.NewGenericMessageSender(channelID)
	a.AggregatorMap[channelID].MessageChan <- msg
}

// QueueChatMessage queues up the chat message to be send in a specific channel
func (a *Aggregator) QueueChatMessage(channelID string, msg *models.ChatMessage) {
	if gms, ok := a.AggregatorMap[channelID]; ok {
		gms.ChatMessageChan <- msg
		return
	}
	a.NewGenericMessageSender(channelID)
	a.AggregatorMap[channelID].ChatMessageChan <- msg
}
