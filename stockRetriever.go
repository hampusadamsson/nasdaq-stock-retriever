package stockclient

import (
	"log"
	"time"

	"github.com/jellydator/ttlcache/v3"
)

type Retriever struct {
	cache ttlcache.Cache[string, RetrievedListings]
}

type retrieverFunction func() *RetrievedListings

func CreateRetriever(fun retrieverFunction, ttl time.Duration) *Retriever {
	loader := ttlcache.LoaderFunc[string, RetrievedListings](
		func(c *ttlcache.Cache[string, RetrievedListings], key string) *ttlcache.Item[string, RetrievedListings] {
			retrieved := fun()
			item := c.Set(key, *retrieved, ttlcache.DefaultTTL)
			log.Println("Cache stale - retrieving data")
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

func (r *Retriever) RetrieveStocks() *RetrievedListings {
	val := r.cache.Get("stocks").Value()
	return &val
}
