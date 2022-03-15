package runner

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"github.com/thehunter365/gorunner/utils"
)

//Server type
type Server struct {
	Port   string
	Router *http.ServeMux

	Handlers map[string]HandlerFunc
}

//HandlerFunc type
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

//NewServer function
func NewServer(port string) *Server {
	r :=  http.NewServeMux()
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
	s.AddHandlerFunc("/", rootHandler)
	s.AddHandlerFunc("/go", handleGoRunner)

	for r, h := range s.Handlers {
		s.Router.HandleFunc(r, h)
	}

	log.Fatal(http.ListenAndServe(s.Port, s.Router))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from GoRunner !"))
}

func handleGoRunner(w http.ResponseWriter, r *http.Request) {
	defer utils.TimeTrack(time.Now(), "Http Handler")
	w.Header().Set("content-type", "application/json")
	log.Println("Handling client !!")
	var code RawCode
	_ = json.NewDecoder(r.Body).Decode(&code)
	if len(code.CodeLines) != 0 {
		log.Println("starting the runner")
		run := NewRunner(GO, code)
		run.ParseCode()
		log.Println(run.CodeLines)
		out := run.StartRunner()
		res := out
		json.NewEncoder(w).Encode(res)
	} else {
		res := Response{[]string{"Unable to read request body"}}
		log.Println(code)
		json.NewEncoder(w).Encode(res)
	}
}
