package api

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/davecusatis/project-sal-backend/project-sal-backend/token"
// )

// // GetImages is the endpoint retrieves image URLs from s3 to load into the app
// func (a *API) GetImages(w http.ResponseWriter, req *http.Request) {
// 	// validate token
// 	tok, err := token.ExtractAndValidateTokenFromHeader(req.Header)
// 	if err != nil {
// 		w.Write([]byte(fmt.Sprintf("error %s", err)))
// 		return
// 	}

// 	w.WriteHeader(http.StatusNotFound)
// 	w.Write([]byte("404 - Images not uploaded"))
// }
