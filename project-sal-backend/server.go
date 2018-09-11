package sal

import (
	"log"
	"net/http"

	"github.com/davecusatis/project-sal-backend/project-sal-backend/api"
	cors "github.com/heppu/simple-cors"
	"goji.io"
	"goji.io/pat"
)

const (
	apiBase = "/api/v0"
)

// Server is the struct representing the doorman Server
type Server struct {
	Port string
	Mux  *goji.Mux
}

// NewServer returns a new instance of the doorman Server
func NewServer(api *api.API) (*Server, error) {
	mux := goji.NewMux()

	mux.HandleFunc(pat.Post(apiBase+"/ping"), api.Ping)
	mux.HandleFunc(pat.Post(apiBase+"/scores"), api.Play)
	mux.HandleFunc(pat.Get(apiBase+"/scores"), api.GetScores)
	mux.HandleFunc(pat.Post(apiBase+"/user/:userID/images"), api.PostImages)
	mux.HandleFunc(pat.Get("/debug/running"), api.HealthCheck)
	return &Server{
		Port: "3030",
		Mux:  mux,
	}, nil
}

// Start starts the webserver
func (s *Server) Start() {
	log.Printf("Starting server on port %s", s.Port)
	log.Fatal(http.ListenAndServe(":"+s.Port, cors.CORS(s.Mux)))
}
