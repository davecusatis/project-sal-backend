package datasource

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/davecusatis/project-sal-backend/project-sal-backend/models"
)

// Datasource is the datasource structs
type Datasource struct {
	dynamo     *dynamodb.DynamoDB
	scoreTable string
}

// NewDatasource returns a new datasource instance
func NewDatasource() *Datasource {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	}))

	return &Datasource{
		dynamo:     dynamodb.New(sess),
		scoreTable: "scores",
	}
}

func (d *Datasource) LeaderboardForChannelID(channelID string) ([]models.Score, error) {
	out, err := d.dynamo.Query(&dynamodb.QueryInput{
		TableName: aws.String(d.scoreTable),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":channelID": {
				S: aws.String(channelID),
			},
		},
		KeyConditionExpression: aws.String("ChannelID = :channelID"),
	})
	if err != nil {
		// alert
		log.Printf("Error getting leaderboard for channel %s: %s", channelID, err)
		return nil, err
	}

	var scores []models.Score
	err = dynamodbattribute.UnmarshalListOfMaps(out.Items, &scores)
	if err != nil {
		// alert
		log.Printf("Error unmarshalling leaderboard %#v: %s", out.Items, err)
		return nil, err
	}

	return scores, nil
}

func (d *Datasource) RecordScore(newScore models.Score) error {
	_, err := d.dynamo.PutItem(&dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"ChannelID": {
				S: aws.String(newScore.ChannelID),
			},
			"UserID": {
				S: aws.String(newScore.UserName),
			},
			"Score": {
				N: aws.String(string(newScore.Score)),
			},
			"ScoreValue": {
				N: aws.String(string(newScore.ScoreValue)),
			},
		},
		TableName: aws.String(d.scoreTable),
	})

	if err != nil {
		// alert
		log.Printf("Error recording score %#v: %s", newScore, err)
		return err
	}

	return nil
}
