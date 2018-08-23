package datasource

import (
	"database/sql"
	"fmt"
	"log"

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
	dbHost := "project-sal-db.cm9smw3zpm24.us-west-2.rds.amazonaws.com"
	dbUser := "dave"
	dbPassword := ""
	dbPort := "5432"
	dbName := "sal"
	connectionStr := fmt.Sprintf("host= %s user=%s dbname=%s password=%s port=%s sslmode=disable", dbHost, dbUser, dbName, dbPassword, dbPort)
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		log.Fatalf("Error creating database connection: %s", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error: Could not establish a connection with the database: %s", err)
	}
	log.Printf("Connected to Database")
	return &Datasource{
		db:         db,
		scoreTable: "scores",
	}
}

func (d *Datasource) LeaderboardForChannelID(channelID string) ([]models.Score, error) {
	query := fmt.Sprintf(`
	SELECT *
	FROM ChannelScores
	WHERE channelId = '%s'`,
		channelID)

	var scores []models.Score
	rows, err := d.db.Query(query)
	if err != sql.ErrNoRows && err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		log.Printf("ROW")
		var id int
		var score models.Score
		err = rows.Scan(&id, &score)
		if err != nil {
			return nil, err
		}
		scores = append(scores, score)
	}
	return scores, nil
}

func (d *Datasource) RecordScore(newScore models.Score) error {
	query := fmt.Sprintf(`
	INSERT INTO
	ChannelScores(channelId, userId, score, bitsUsed)
	VALUES ('%s', '%s', '%d', '%d')
	RETURNING id`,
		newScore.ChannelID,
		newScore.UserID,
		newScore.Score,
		newScore.BitsUsed)

	var ID string
	err := d.db.QueryRow(query).Scan(&ID)
	if err != sql.ErrNoRows && err != nil {
		return err
	}
	return nil
}
