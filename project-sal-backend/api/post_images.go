package api

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	salImage "github.com/davecusatis/project-sal-backend/project-sal-backend/image"
	"github.com/davecusatis/project-sal-backend/project-sal-backend/token"
	"goji.io/pat"
)

// PostImages is the endpoint that uploads user images for use in the slot machine
func (a *API) PostImages(w http.ResponseWriter, req *http.Request) {
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

	err = req.ParseMultipartForm(200000) // grab the multipart form
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	for filename, files := range req.MultipartForm.File {
		for _, file := range files {
			reader, err := file.Open()
			if err != nil {
				log.Printf("Error opening file")
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(fmt.Sprintf("invalid file: %s", file.Filename)))
				return
			}
			defer reader.Close()

			img, format, err := image.Decode(reader)
			if err != nil {
				log.Printf("Error decoding %s file: %s", file.Filename, err)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(fmt.Sprintf("invalid file: %s", file.Filename)))
			}

			valid := salImage.ValidateImageFromFile(img, format)
			if !valid {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(fmt.Sprintf("invalid file: %s", file.Filename)))
				return
			}
			buffer := make([]byte, file.Size)
			reader.Read(buffer)
			fileBytes := bytes.NewReader(buffer)
			err = a.UploadImageToS3(fileBytes, tok.ChannelID, filename, format)
			if err != nil {
				log.Printf("Error uploading image: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf("unable to upload file: %s", file.Filename)))
				return
			}
		}
	}

	w.Write([]byte("OK"))
}

func (a *API) UploadImageToS3(fileBytes *bytes.Reader, userID string, filename string, format string) error {
	_, err := a.S3.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String("project-sal-distro"),
		Key:         aws.String(fmt.Sprintf("%s/%s", userID, filename)),
		ContentType: aws.String(fmt.Sprintf("image/%s", format)),
		ACL:         aws.String("public-read"),
		Body:        fileBytes,
	})

	if err != nil {
		return err
	}
	return nil
}
