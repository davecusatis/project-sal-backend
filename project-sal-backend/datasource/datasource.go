package datasource

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/davecusatis/project-sal-backend/project-sal-backend/models"

	// import pq driver
	_ "github.com/lib/pq"
)

const (
	prod = true
)

// Datasource is the datasource structs
type Datasource struct {
	db         *sql.DB
	scoreTable string
}

// NewDatasource returns a new datasource instance
func NewDatasource() *Datasource {
	log.Printf("In new datasource1")
	dbHost := "localhost"
	if prod {
		dbHost = "project-sal-db.cm9smw3zpm24.us-west-2.rds.amazonaws.com"
	}

	dbUser := "dave"
	dbPassword := "Pooppy1992"
	dbPort := "5432"
	dbName := "sal"
	connectionStr := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s sslmode=disable", dbHost, dbUser, dbName, dbPassword, dbPort)
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		log.Fatalf("Error creating database connection: %s", err)
	}
	log.Printf("In new datasource2")
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
	var scores []models.Score
	last24Hrs := time.Now().AddDate(0, 0, -1).UTC().Format(time.RFC3339)
	query := fmt.Sprintf(`
	SELECT *
	FROM ChannelScores
	WHERE channelId = '%s' AND recordedAt > '%s'
	LIMIT 100;`,
		channelID, last24Hrs)
	log.Printf("QUERY: %s", query)
	rows, err := d.db.Query(query)
	if err != sql.ErrNoRows && err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var score models.Score
		err = rows.Scan(
			&score.ID,
			&score.Score,
			&score.RecordedAt,
			&score.UserID,
			&score.ChannelID,
			&score.BitsUsed)
		if err != nil {
			return nil, err
		}
		scores = append(scores, score)
	}
	return scores, nil
}

func (d *Datasource) RecordScore(newScore models.Score) (*models.Score, error) {
	query := fmt.Sprintf(`
	INSERT INTO
	ChannelScores(channelId, userId, score, bitsUsed)
	VALUES ('%s', '%s', '%d', '%d')
	RETURNING id, recordedAt`,
		newScore.ChannelID,
		newScore.UserID,
		newScore.Score,
		newScore.BitsUsed)

	var ID string
	var recordedAt time.Time
	err := d.db.QueryRow(query).Scan(&ID, &recordedAt)
	if err != sql.ErrNoRows && err != nil {
		return nil, err
	}
	return &models.Score{
		ID:         ID,
		RecordedAt: recordedAt.String(),
		ChannelID:  newScore.ChannelID,
		UserID:     newScore.UserID,
		Score:      newScore.Score,
		BitsUsed:   newScore.BitsUsed,
	}, nil
}
