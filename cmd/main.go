package main

import (
	"log"

	sal "github.com/davecusatis/project-sal-backend/project-sal-backend"
	"github.com/davecusatis/project-sal-backend/project-sal-backend/api"
)

func main() {
	api, err := api.NewAPI()
	server, err := sal.NewServer(api)
	if err != nil {
		log.Fatalf("Could not start server: %s", err)
	}

	server.Start()
}
