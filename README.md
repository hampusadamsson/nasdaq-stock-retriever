# nasdaq-stock-retriever

![example workflow](https://github.com/hampusadamsson/nasdaq-stock-retriever/actions/workflows/go.yml/badge.svg)
![codecov](https://codecov.io/gh/hampusadamsson/nasdaq-stock-retriever/branch/main/graph/badge.svg)

Retrieves stocks information from Nasdaq nordic.
Data is fetched using a REST API client.
A ttl cache is used to limit requests.

## Example

To fetch all assets

```go
r := CreateRetriever(FetchAllNordicAssets, 2*time.Second)
s := r.RetrieveStocks()
fmt.Println(s)
```

To fetch from specific markets

```go
// Possible values are:
// FetchStockholmAssets, FetchCopoenhagenAssets, FetchBalticAssets, FetchIcelandAssets
r := CreateRetriever(FetchStockholmAssets, 2*time.Second)
s := r.RetrieveStocks()
fmt.Println(s)
```

## License

[MIT](https://choosealicense.com/licenses/mit/)
