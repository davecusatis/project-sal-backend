package api

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/davecusatis/project-sal-backend/project-sal-backend/token"
	"goji.io/pat"
)

// PostTitle is the endpoint that uploads user images for use in the slot machine
func (a *API) PostTitle(w http.ResponseWriter, req *http.Request) {
	title := "test"
	userID := pat.Param(req, "userID")
	// validate token
	tok, err := token.ExtractAndValidateTokenFromHeader(req.Header)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err)))
		return
	}

	//TODO: fix this logic lmfao
	if tok.Role != "broadcaster" || tok.ChannelID != userID {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Not allowed"))
		return
	}

	a.UploadTextFileToS3(title, userID, "title.txt")

	// TODO: invalidate cache
	w.Write([]byte("OK"))
}

// UploadTextFileToS3 is the helper function that uploads a users slot machine title to s3 bucket
func (a *API) UploadTextFileToS3(title string, userID string, filename string) error {
	filepath := fmt.Sprintf("%s/%s", userID, filename)
	_, err := a.S3.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String("project-sal-distro"),
		Key:         aws.String(filepath),
		ContentType: aws.String("text/plain"),
		ACL:         aws.String("public-read"),
		Body:        aws.ReadSeekCloser(strings.NewReader(title)),
	})

	if err != nil {
		return err
	}

	_, err = a.CloudFront.CreateInvalidation(&cloudfront.CreateInvalidationInput{
		DistributionId: aws.String("E12APQ9IM9QY5"),
		InvalidationBatch: &cloudfront.InvalidationBatch{
			Paths: &cloudfront.Paths{
				Items: []*string{
					aws.String(filepath),
				},
				Quantity: aws.Int64(1),
			},
		},
	})

	if err != nil {
		return err
	}
	return nil
}
