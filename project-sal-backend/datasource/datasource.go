package datasource

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/davecusatis/project-sal-backend/project-sal-backend/models"
)

// Datasource is the datasource structs
type Datasource struct {
	dynamo *dynamodb.DynamoDB
}

// NewDatasource returns a new datasource instance
func NewDatasource() *Datasource {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	}))

	return &Datasource{
		dynamo: dynamodb.New(sess),
	}
}

func (d *Datasource) LeaderboardForChannelID(channelID string) ([]models.Score, error) {
	out, err := d.dynamo.GetItem(&dynamodb.GetItemInput{})
	if err != nil {
		// alert
		log.Printf("Error getting leaderboard for channel %s: %s", channelID, err)
	}
	return nil, nil
}

func (d *Datasource) RecordScore(newScore models.Score) error {
	_, err := d.dynamo.PutItem(&dynamodb.PutItemInput{})
	if err != nil {
		// alert
		log.Printf("Error recording score %#v: %s", newScore, err)
	}
	return nil
}
