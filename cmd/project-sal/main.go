package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	sal "github.com/davecusatis/project-sal-backend/project-sal-backend"
	"github.com/davecusatis/project-sal-backend/project-sal-backend/api"
)

func main() {
	log.Printf("Hello")
	s3 := s3.New(session.New())

	log.Printf("after s3")

	api, err := api.NewAPI(s3)
	log.Printf("after api")
	server, err := sal.NewServer(api)
	if err != nil {
		log.Fatalf("Could not start server: %s", err)
	}
	log.Printf("after new server")
	server.Start()
}
