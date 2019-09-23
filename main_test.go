package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/andresoro/easymock/server"
)

func TestServer(t *testing.T) {
	t.Log("Starting server test...")
	s, err := server.New("./client/build", "./db_test")
	if err != nil {
		t.Fatal("unable to init server")
	}
	defer func() {
		s.ShutDown()
		os.RemoveAll("./db_test")
	}()

	t.Run("testing code 200 creation", func(t *testing.T) {
		// create our request
		d := &server.Response{Code: 200, ContentType: "application/json", Body: "test"}
		body, err := json.Marshal(d)
		if err != nil {
			t.Error(err)
		}
		req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/gimme", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != 200 {
			t.Errorf("status code on create should be 200 got: %d", resp.StatusCode)
		}

		// read id from response body
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal("unable to read creation response body")
		}
		id := string(bodyBytes)

		// test newly created endpoint
		url := fmt.Sprintf("http://localhost:8080/r/%s", id)

		req, _ = http.NewRequest(http.MethodGet, url, nil)

		resp, err = client.Do(req)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != 200 {
			t.Error("not getting proper status code")
		}

		bodyBytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		if string(bodyBytes) != d.Body {
			t.Log(string(bodyBytes))
			t.Error("body should be the same as the original object")
		}
	})

}
