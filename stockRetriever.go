package stockclient

import (
	"fmt"
	"log"
	"time"

	"github.com/jellydator/ttlcache/v3"
)

// Retriever wraps the stock client with a cache
type Retriever struct {
	cache ttlcache.Cache[string, []Stock]
}

type retrieverFunction func() ([]Stock, error)

// CreateRetriever creates the client and adds a cache with a ttl attached to it
func CreateRetriever(fun retrieverFunction, ttl time.Duration) *Retriever {
	loader := ttlcache.LoaderFunc[string, []Stock](
		func(c *ttlcache.Cache[string, []Stock], key string) *ttlcache.Item[string, []Stock] {
			log.Println("Cache stale - retrieving data")
			retrieved, err := fun()
			if err != nil {
				fmt.Println("Error", err)
			}
			item := c.Set(key, retrieved, ttlcache.DefaultTTL)
			return item
		},
	)
	cache := ttlcache.New(
		ttlcache.WithTTL[string, []Stock](ttl),
		ttlcache.WithDisableTouchOnHit[string, []Stock](),
		ttlcache.WithLoader[string, []Stock](loader),
	)

	return &Retriever{
		cache: *cache,
	}
}

// RetrieveStocks retrieve all listings from 1) cache, or 2) from source if stale
func (r *Retriever) RetrieveStocks() []Stock {
	val := r.cache.Get("stocks").Value()
	return val
}
