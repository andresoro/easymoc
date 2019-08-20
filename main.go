package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var cache = make(map[string]*Response)

// Response is the data structure we associate with an id
type Response struct {
	Code    int               `json:"code"`
	Headers map[string]string `json:"headers"`
	Body    []byte            `json:"body"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/r/{id}", reponseHandler).Methods("GET")
	r.HandleFunc("/r/", newResponse).Methods("PUT")

	http.ListenAndServe(":8080", r)
}

func reponseHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if resp, ok := cache[id]; ok {
		for k, v := range resp.Headers {
			w.Header().Set(k, v)
		}
		w.WriteHeader(resp.Code)
		w.Write(resp.Body)
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
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// generate uuid and cache response to handle later
	id := uuid.New().String()
	cache[id] = &resp

	// write back generated id
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(id))
	return
}
