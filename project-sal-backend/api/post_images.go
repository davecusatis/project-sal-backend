package api

import (
	"fmt"
	"net/http"

	_ "image/jpeg"
	_ "image/png"

	"github.com/davecusatis/project-sal-backend/project-sal-backend/image"
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

	for _, files := range req.MultipartForm.File {
		for _, file := range files {
			valid := image.ValidateImageFromFile(file)
			if !valid {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(fmt.Sprintf("invalid file: %s", file.Filename)))
				return
			}
		}
	}

	w.Write([]byte("OK"))
}
