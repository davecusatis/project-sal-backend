package api

import (
	"net/http"

	"github.com/davecusatis/project-sal-backend/project-sal-backend/aggregator"
	"github.com/davecusatis/project-sal-backend/project-sal-backend/datasource"
	"github.com/davecusatis/project-sal-backend/project-sal-backend/twitch"
)

// API struct
type API struct {
	Aggregator   *aggregator.Aggregator
	Datasource   *datasource.Datasource
	TwitchClient *twitch.TwitchClient
}

// NewAPI creates a new instance of an API
func NewAPI() (*API, error) {
	a := aggregator.NewAggregator()
	a.Start()
	return &API{
		Aggregator:   a,
		Datasource:   datasource.NewDatasource(),
		TwitchClient: twitch.NewTwitchClient(&http.Client{}),
	}, nil
}
