package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	sal "github.com/davecusatis/project-sal-backend/project-sal-backend"
	"github.com/davecusatis/project-sal-backend/project-sal-backend/api"
	"github.com/davecusatis/project-sal-backend/project-sal-backend/image"
)

func main() {
	s3 := s3.New(session.New())

	api, err := api.NewAPI(s3)
	server, err := sal.NewServer(api)
	if err != nil {
		log.Fatalf("Could not start server: %s", err)
	}
	image.GenerateImageFromURLS()
	server.Start()
}
