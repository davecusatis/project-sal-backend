package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/aws/aws-sdk-go/service/s3"

	sal "github.com/davecusatis/project-sal-backend/project-sal-backend"
	"github.com/davecusatis/project-sal-backend/project-sal-backend/api"
)

func main() {
	log.Printf("Hello")
	sess := session.New()
	s3 := s3.New(sess)
	cloudfront := cloudfront.New(sess)

	log.Printf("after s3")

	api, err := api.NewAPI(s3, cloudfront)
	log.Printf("after api")
	server, err := sal.NewServer(api)
	if err != nil {
		log.Fatalf("Could not start server: %s", err)
	}
	log.Printf("after new server")
	server.Start()
}
