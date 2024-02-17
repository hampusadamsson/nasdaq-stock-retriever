package stockclient

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// RetrieveStocksDummy used for testing the proper RetrieveStocksFromAPIgo
func RetrieveStocksDummy() *RetrievedListings {
	stocks := []Stock{{ID: "DUMMY"}}
	return &RetrievedListings{
		Listings:  stocks,
		TimeStamp: time.Now().GoString(),
	}
}

// FetchAllNordicAssets retrieves data for a wide list of stocks
func FetchAllNordicAssets() *RetrievedListings {
	return RetrieveStocks(RegionAll)
}

// FetchStockholmAssets retrieves data from the Stockholm stock exchange
func FetchStockholmAssets() *RetrievedListings {
	return RetrieveStocks(RegionStockholm)
}

// FetchCopoenhagenAssets retrieves data from the Copenhagen stock exchange
func FetchCopoenhagenAssets() *RetrievedListings {
	return RetrieveStocks(RegionCopenhagen)
}

// FetchBalticAssets retrieves data from the Baltic stock exchange
func FetchBalticAssets() *RetrievedListings {
	return RetrieveStocks(RegionBaltic)
}

// FetchIcelandAssets retrieves data from the Iceland stock exchange
func FetchIcelandAssets() *RetrievedListings {
	return RetrieveStocks(RegionIceland)
}

// RetrieveStocks from stock exchanges in the nordics -- possible values. RegionAll, RegionStockholm, RegionCopenhagen, RegionHelsinki, RegionIceland, RegionBaltic
func RetrieveStocks(lists string) *RetrievedListings {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://www.nasdaqomxnordic.com/webproxy/DataFeedProxy.aspx?SubSystem=Prices&Action=Search&"+lists+"&json=1", nil)
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36")
	resp, _ := client.Do(req)
	jsonStr, _ := io.ReadAll(resp.Body)
	data := serialize[RetrievedListings](jsonStr)
	defer resp.Body.Close()
	return data
}

// RetrieveStock using symbols - ex. SSE992
func RetrieveStock(symbol string) *History {
	client := &http.Client{}
	today := time.Now().String()[:10]
	url := "http://www.nasdaqomxnordic.com/webproxy/DataFeedProxy.aspx?SubSystem=History&Action=GetChartData&inst.an=id%2Cnm%2Cfnm%2Cisin%2Ctp%2Cchp%2Cycp&FromDate=1986-01-01&ToDate=" + today + "&json=true&showAdjusted=true&app=%2Faktier%2Fmicrosite-MicrositeChart-history&timezone=CET&DefaultDecimals=false&Instrument=" + symbol
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	resp, _ := client.Do(req)
	jsonStr, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	data := serialize[History](jsonStr)
	return data
}

func serialize[T any](jsonStr []byte) *T {
	var data T
	json.Unmarshal([]byte(jsonStr), &data)
	return &data
}
