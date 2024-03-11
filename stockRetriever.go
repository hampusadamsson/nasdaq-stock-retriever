package stockclient

import (
	"fmt"
	"log"
	"time"

	"github.com/jellydator/ttlcache/v3"
)

// Retriever wraps the stock client with a cache
type Retriever struct {
	cache ttlcache.Cache[string, RetrievedListings]
}

type retrieverFunction func() (*RetrievedListings, error)

// CreateRetriever creates the client and adds a cache with a ttl attached to it
func CreateRetriever(fun retrieverFunction, ttl time.Duration) *Retriever {
	loader := ttlcache.LoaderFunc[string, RetrievedListings](
		func(c *ttlcache.Cache[string, RetrievedListings], key string) *ttlcache.Item[string, RetrievedListings] {
			log.Println("Cache stale - retrieving data")
			retrieved, err := fun()
			if err != nil {
				fmt.Println("Error", err)
			}
			item := c.Set(key, *retrieved, ttlcache.DefaultTTL)
			return item
		},
	)
	cache := ttlcache.New(
		ttlcache.WithTTL[string, RetrievedListings](ttl),
		ttlcache.WithDisableTouchOnHit[string, RetrievedListings](),
		ttlcache.WithLoader[string, RetrievedListings](loader),
	)

	return &Retriever{
		cache: *cache,
	}
}

// RetrieveStocks retrieve all listings from 1) cache, or 2) from source if stale
func (r *Retriever) RetrieveStocks() *RetrievedListings {
	val := r.cache.Get("stocks").Value()
	return &val
}
