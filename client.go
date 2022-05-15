package stockclient

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func RetrieveStocksDummy() *RetrievedListings {
	stocks := []Stock{Stock{ID: "DUMMY"}}
	return &RetrievedListings{
		Listings:  stocks,
		TimeStamp: time.Now().GoString(),
	}
}

func RetrieveStocksFromApi() *RetrievedListings {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://www.nasdaqomxnordic.com/webproxy/DataFeedProxy.aspx?SubSystem=Prices&Action=Search&List=M:INET:XSTO:SEEQ-SHR&List=M:INET:XSTO:SEEQ-SHR-CCP&List=M:INET:XCSE:DKEQ-SHR&List=M:INET:XCSE:DKEQ-SHR-CCP&List=M:INET:XHEL:FIEQ-SHR&List=M:INET:XHEL:FIEQ-SHR-CCP&List=M:INET:XHEL:FIEQ-SHR-IC&List=M:INET:XICE:ISEQ-SHR&List=M:INET:XSTO:SEEQ-SHR-NOK&List=M:INET:XSTO:SEEQ-SHR-AO&List=M:INET:FNSE:SEMM-NM&List=M:INET:FNDK:FNDK-CPH&List=M:INET:FNFI:SEMM-FN-HEL&List=M:INET:FNIS:ISEC-SHR&List=M:INET:FNSE:SEMM-FN-NOK&List=M:INET:FNEE:EEMM-SHR&List=M:INET:FNLV:LVMM-SHR&List=M:INET:FNLT:LTMM-SHR&List=M:INET:FNFI:SEMM-FN-HE-ERW&List=M:INET:XTAL:EEEQ-SHR&List=M:INET:XRIS:LVEQ-SHR&List=M:INET:XLIT:LTEQ-SHR&List=M:GITS:RI:RSEBA&List=M:GITS:TA:TSEBA&List=M:GITS:VI:VSEBA&json=1&callback=jQuery11110636858574942031_1622134931936", nil)
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36")
	resp, _ := client.Do(req)
	jsonStr, _ := io.ReadAll(resp.Body)
	data := serialize(jsonStr)
	defer resp.Body.Close()
	return data
}

func serialize(jsonStr []byte) *RetrievedListings {
	inputProperJson := jsonStr[42 : len(jsonStr)-2] //Remove the Jquery String prefix- and suffix
	var data RetrievedListings
	json.Unmarshal([]byte(inputProperJson), &data)
	return &data
}
