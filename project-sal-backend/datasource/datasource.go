package datasource

import (
	"database/sql"
	"fmt"

	"github.com/davecusatis/project-sal-backend/project-sal-backend/models"

	// import pq driver
	_ "github.com/lib/pq"
)

// Datasource is the datasource structs
type Datasource struct {
	db         *sql.DB
	scoreTable string
}

// NewDatasource returns a new datasource instance
func NewDatasource() *Datasource {
	// dbHost := config.MustGetConfigValue("DB_HOST")
	// dbUser := config.MustGetConfigValue("DB_USER")
	// dbPassword := config.MustGetConfigValue("DB_PASSWORD")
	// dbPort := config.MustGetConfigValue("DB_PORT")
	// dbName := config.MustGetConfigValue("DB_NAME")

	// dbHost := "DB_HOST"
	// dbUser := "DB_USER"
	// dbPassword := "DB_PASSWORD"
	// dbPort := "DB_PORT"
	// dbName := "DB_NAME"
	// connectionStr := fmt.Sprintf("host= %s user=%s dbname=%s password=%s port=%s sslmode=disable", dbHost, dbUser, dbName, dbPassword, dbPort)
	// db, err := sql.Open("postgres", connectionStr)
	// if err != nil {
	// 	log.Fatalf("Error creating database connection: %s", err)
	// }

	// err = db.Ping()
	// if err != nil {
	// 	log.Fatal("Error: Could not establish a connection with the database")
	// }
	return &Datasource{
		db:         nil,
		scoreTable: "scores",
	}
}

func (d *Datasource) LeaderboardForChannelID(channelID string) ([]models.Score, error) {
	query := fmt.Sprintf(`SELECT score FROM scores WHERE channel_id = '%s' LIMIT 10`, channelID)

	var scores []models.Score
	err := d.db.QueryRow(query).Scan(&scores)
	if err != sql.ErrNoRows && err != nil {
		return nil, err
	}
	return scores, nil

	// TODO: delete comments out when this is verified to work
	// out, err := d.dynamo.Query(&dynamodb.QueryInput{
	// 	TableName: aws.String(d.scoreTable),
	// 	ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
	// 		":channelID": {
	// 			S: aws.String(channelID),
	// 		},
	// 	},
	// 	KeyConditionExpression: aws.String("ChannelID = :channelID"),
	// })
	// if err != nil {
	// 	// alert
	// 	log.Printf("Error getting leaderboard for channel %s: %s", channelID, err)
	// 	return nil, err
	// }

	// var scores []models.Score
	// err = dynamodbattribute.UnmarshalListOfMaps(out.Items, &scores)
	// if err != nil {
	// 	// alert
	// 	log.Printf("Error unmarshalling leaderboard %#v: %s", out.Items, err)
	// 	return nil, err
	// }
}

func (d *Datasource) RecordScore(newScore models.Score) error {
	query := fmt.Sprintf(`INSERT INTO scores(channel_id, user_id, score, score_value) VALUES ('%s', '%s', '%d', '%d') RETURNING id`,
		newScore.ChannelID, newScore.UserName, newScore.Score, newScore.ScoreValue)

	var ID string
	err := d.db.QueryRow(query).Scan(&ID)
	if err != sql.ErrNoRows && err != nil {
		return err
	}
	return nil

	// _, err := d.dynamo.PutItem(&dynamodb.PutItemInput{
	// 	Item: map[string]*dynamodb.AttributeValue{
	// 		"ChannelID": {
	// 			S: aws.String(newScore.ChannelID),
	// 		},
	// 		"UserID": {
	// 			S: aws.String(newScore.UserName),
	// 		},
	// 		"Score": {
	// 			N: aws.String(string(newScore.Score)),
	// 		},
	// 		"ScoreValue": {
	// 			N: aws.String(string(newScore.ScoreValue)),
	// 		},
	// 	},
	// 	TableName: aws.String(d.scoreTable),
	// })

	// if err != nil {
	// 	// alert
	// 	log.Printf("Error recording score %#v: %s", newScore, err)
	// 	return err
	// }
}
