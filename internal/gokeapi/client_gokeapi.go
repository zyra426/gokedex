package gokeapi

import (
	"net/http"
	"time"

	"github.com/zyra426/gokedex/internal/gokecache"
)

type Client struct {
	cache      gokecache.Cache
	httpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: gokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
