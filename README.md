# nasdaq-stock-retriever

![example workflow](https://github.com/hampusadamsson/nasdaq-stock-retriever/actions/workflows/go.yml/badge.svg)
![codecov](https://codecov.io/gh/hampusadamsson/nasdaq-stock-retriever/branch/main/graph/badge.svg)

Retrieves stocks information from Nasdaq nordic.
Data is fetched using a REST API client.
A ttl cache is used to limit requests.

Live API tests are skipped in `-short` mode (used by CI).

## Example

To fetch all main-market and First North listings:

```go
r := CreateRetriever(RetrieveStocks, 2*time.Second)
s := r.RetrieveStocks()
fmt.Println(s)
```

To fetch only main markets or First North:

```go
r := CreateRetriever(RetrieveStocksMainMarkets, 2*time.Second)
r2 := CreateRetriever(RetrieveStocksFirstNorth, 2*time.Second)
```

To fetch OHLCV history for a symbol (orderbook id):

```go
h, err := RetrieveStock("SSE992")
```

Region list constants for Nasdaq Nordic markets: `RegionStockholm`, `RegionCopenhagen`,
`RegionHelsinki`, `RegionIceland`, `RegionBaltic`, `RegionAll`.

## License

[Apache-2.0](https://choosealicense.com/licenses/apache-2.0/)
