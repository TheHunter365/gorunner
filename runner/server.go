package runner

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Server type
type Server struct {
	Port   string
	Router *mux.Router

	Handlers []HandlerFunc
}

//HandlerFunc type
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

//NewServer function
func NewServer(port string) *Server {
	r := mux.NewRouter()
	return &Server{
		Port:     port,
		Router:   r,
		Handlers: make([]HandlerFunc, 128),
	}
}

//AddHandlerFunc function
func (s *Server) AddHandlerFunc(handler HandlerFunc) {
	s.Handlers = append(s.Handlers, handler)
}

//Start func
func (s *Server) Start() {

	log.Fatal(http.ListenAndServe(s.Port, s.Router))
}
