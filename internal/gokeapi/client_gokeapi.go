package gokeapi

import (
	"net/http"
	"time"

	"github.com/tidwall/buntdb"
	"github.com/zyra426/gokedex/internal/db"
	"github.com/zyra426/gokedex/internal/gokecache"
)

type Client struct {
	cache      gokecache.Cache
	db         *buntdb.DB
	httpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: gokecache.NewCache(cacheInterval),
		db:    db.NewDB(),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
