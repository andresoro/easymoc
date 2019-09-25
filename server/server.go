package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Server struct {
	db     *Env
	router *mux.Router
	srv    *http.Server
}

// New server instance
func New(static string, db string) (*Server, error) {
	r := mux.NewRouter()
	e, err := NewEnv(db)
	if err != nil {
		return nil, err
	}

	r.HandleFunc("/r/{id}", reponseHandler(e)).Methods("GET")
	r.HandleFunc("/gimme", newResponse(e)).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(static)))

	srv := &http.Server{Addr: ":8080", Handler: r}

	return &Server{db: e, router: r, srv: srv}, nil

}

func (s *Server) Run() {
	//start server
	log.Printf("starting server")
	s.srv.ListenAndServe()
}

func (s *Server) ShutDown() {
	s.db.Close()
	s.srv.Shutdown(nil)
}

func reponseHandler(e *Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		resp, err := e.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("handling response with id: %s", id)

		w.Header().Set("Content-Type", resp.ContentType)
		w.WriteHeader(resp.Code)
		w.Write([]byte(resp.Body))
		return
	}
}

func newResponse(e *Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// read target response from request body
		var resp Response
		err := json.NewDecoder(r.Body).Decode(&resp)
		if err != nil {
			log.Printf("error decoding body: %e", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// generate uuid and cache response to handle later
		id := uuid.New().String()
		e.Set(id, &resp)

		// write back generated id
		log.Printf("generated response with id: %s", id)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(id))
		return
	}
}
