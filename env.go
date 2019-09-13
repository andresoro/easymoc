package main

import (
	"time"

	"github.com/allegro/bigcache"
	"github.com/dgraph-io/badger"
)

// Env holds persistence mechanisms
type Env struct {
	db    *badger.DB
	cache *bigcache.BigCache
}

func NewEnv() (*Env, error) {
	e := &Env{}

	db, err := badger.Open(badger.DefaultOptions("./db"))
	if err != nil {
		return nil, err
	}

	config := bigcache.DefaultConfig(time.Minute * 5)
	config.OnRemove = func(key string, entry []byte) {
		e.WriteDB(key, entry)
	}
	cache, err := bigcache.NewBigCache(config)
	if err != nil {
		return nil, err
	}

	e.db = db
	e.cache = cache

	return e, nil
}

func (e *Env) WriteDB(key string, entry []byte) {

}
