package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var cache = make(map[string]*Response)

// Response is the data structure we associate with an id
type Response struct {
	Code        int    `json:"code"`
	ContentType string `json:"content"`
	Body        string `json:"body"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/r/{id}", reponseHandler).Methods("GET")
	r.HandleFunc("/gimme", newResponse).Methods("POST")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./client/build/")))

	log.Printf("starting server")
	http.ListenAndServe(":8080", r)
}

func reponseHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if resp, ok := cache[id]; ok {
		log.Printf("handling response with id: %s", id)

		w.WriteHeader(resp.Code)
		w.Header().Set("Content-Type", resp.ContentType)
		w.Write([]byte(resp.Body))
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	return
}

func newResponse(w http.ResponseWriter, r *http.Request) {
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
	cache[id] = &resp

	// write back generated id
	log.Printf("generated response with id: %s", id)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
	return
}
