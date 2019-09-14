package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func main() {
	// router
	r := mux.NewRouter()

	// init persistence mechanisms
	env, err := NewEnv()
	if err != nil {
		log.Fatal(err)
	}
	//handlers
	r.HandleFunc("/r/{id}", reponseHandler(env)).Methods("GET")
	r.HandleFunc("/gimme", newResponse(env)).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./client/build/")))

	// handle SIGINT
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			env.Close()
			os.Exit(0)
		}
	}()

	//start server
	log.Printf("starting server")
	http.ListenAndServe(":8080", r)
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

		w.WriteHeader(resp.Code)
		w.Header().Set("Content-Type", resp.ContentType)
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
