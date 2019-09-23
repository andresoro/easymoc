package server

import (
	"bytes"
	"encoding/json"
	"log"
	"time"

	"github.com/allegro/bigcache"
	"github.com/dgraph-io/badger"
)

// Response is the data structure we associate with an id
type Response struct {
	Code        int    `json:"code"`
	ContentType string `json:"content"`
	Body        string `json:"body"`
}

// Env holds persistence mechanisms
type Env struct {
	db    *badger.DB
	cache *bigcache.BigCache
}

func NewEnv(path string) (*Env, error) {
	e := &Env{}

	// persistent db
	db, err := badger.Open(badger.DefaultOptions(path))
	if err != nil {
		return nil, err
	}

	// cache config
	config := bigcache.DefaultConfig(time.Minute * 5)

	// cache
	cache, err := bigcache.NewBigCache(config)
	if err != nil {
		return nil, err
	}

	e.db = db
	e.cache = cache

	return e, nil
}

// Get a resp from cache or db
func (e *Env) Get(id string) (*Response, error) {

	data, err := e.cache.Get(id)
	if err != nil {
		// if entry is not in cache check db
		if err == bigcache.ErrEntryNotFound {
			return e.GetDB(id)
		}
		return nil, err
	}
	log.Println("cache hit")
	var resp Response
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil

}

// Set adds a resp to cache
func (e *Env) Set(id string, resp *Response) error {
	data := new(bytes.Buffer)
	json.NewEncoder(data).Encode(resp)
	err := e.cache.Set(id, data.Bytes())
	if err != nil {
		return err
	}

	go func() {
		e.WriteDB(id, data.Bytes())
	}()

	return nil
}

// WriteDB writes to filesystem
func (e *Env) WriteDB(key string, resp []byte) error {
	return e.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), resp)
	})
}

func (e *Env) GetDB(key string) (*Response, error) {

	var resp Response
	err := e.db.View(func(txn *badger.Txn) error {

		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}

		err = item.Value(func(val []byte) error {
			log.Println("db hit")
			log.Println("writing to cache")

			// write to cache
			err := e.cache.Set(key, val)
			if err != nil {
				return err
			}

			err = json.Unmarshal(val, &resp)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
		return nil
	})

	return &resp, err
}

func (e *Env) Close() {
	e.db.Close()
	e.cache.Close()
}
