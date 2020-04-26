package runner

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Server type
type Server struct {
	Port   string
	Router *mux.Router

	Handlers map[string]HandlerFunc
}

//HandlerFunc type
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

//NewServer function
func NewServer(port string) *Server {
	r := mux.NewRouter()
	return &Server{
		Port:     port,
		Router:   r,
		Handlers: make(map[string]HandlerFunc, 128),
	}
}

//AddHandlerFunc function
func (s *Server) AddHandlerFunc(route string, handler HandlerFunc) {
	s.Handlers[route] = handler
}

//Start func
func (s *Server) Start() {
	s.AddHandlerFunc("/go", handleGoRunner)

	for r, h := range s.Handlers {
		s.Router.HandleFunc(r, h)
	}

	log.Fatal(http.ListenAndServe(s.Port, s.Router))
}

func handleGoRunner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var runner Runner

	err := json.NewDecoder(r.Body).Decode(&runner)
	if err != nil {
		log.Println(err)
		w.Write([]byte("Malformed body exception"))
	}
	runner.ParseCode()
	out := runner.StartRunner()
	w.Write([]byte(out))
}
