package stockclient

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSerialise(t *testing.T) {
	const jsonStr = `{"data":{"instrumentListing":{"headers":{"fullName":"Name","currency":"CCY","lastSalePrice":"Last","netChange":"Net Change","percentageChange":"% Change","bidPrice":"Bid","askPrice":"Ask","volume":"Volume","turnover":"Turnover","greenEquityDesignation":"Green Equity Designation","sector":"Sector"},"rows":[{"fullName":"A.P. Møller - Mærsk A","currency":"DKK","priceNotation":"","micCode":"","netChange":"+210.00","percentageChange":"+1.79%","bidPrice":"11,930.00","askPrice":"11,950.00","lastSalePrice":"11,920.00","high":"","low":"","volume":"1,107","reportedVolume":"","turnover":"13,192,420","tradesCount":"","lastTraded":"","time":"","netAssetValue":"","navUpdatedTime":"","totalExpenseRatio":"","orderbookId":"CSE3200","assetClass":"SHARES","symbol":"MAERSK A","sector":"Industrials","isin":"DK0010244425","greenEquityDesignation":"","deltaIndicator":"up","contractSize":"","openInterest":"","expirationDate":"","strikePrice":"","noteCode":"","issuerFullName":"","allTradesAvgPrice":"","couponRate":"","exchangeSymbol":"","noteDescription":"","type":"","settlementPrice":"","totalVolume":""},{"fullName":"A.P. Møller - Mærsk B","currency":"DKK","priceNotation":"","micCode":"","netChange":"+225.00","percentageChange":"+1.85%","bidPrice":"12,380.00","askPrice":"12,390.00","lastSalePrice":"12,385.00","high":"","low":"","volume":"6,625","reportedVolume":"","turnover":"81,966,660","tradesCount":"","lastTraded":"","time":"","netAssetValue":"","navUpdatedTime":"","totalExpenseRatio":"","orderbookId":"CSE3201","assetClass":"SHARES","symbol":"MAERSK B","sector":"Industrials","isin":"DK0010244508","greenEquityDesignation":"","deltaIndicator":"up","contractSize":"","openInterest":"","expirationDate":"","strikePrice":"","noteCode":"","issuerFullName":"","allTradesAvgPrice":"","couponRate":"","exchangeSymbol":"","noteDescription":"","type":"","settlementPrice":"","totalVolume":""},{"fullName":"AaB","currency":"DKK","priceNotation":"","micCode":"","netChange":"-0.60","percentageChange":"-1.75%","bidPrice":"33.60","askPrice":"34.80","lastSalePrice":"33.60","high":"","low":"","volume":"11","reportedVolume":"","turnover":"372","tradesCount":"","lastTraded":"","time":"","netAssetValue":"","navUpdatedTime":"","totalExpenseRatio":"","orderbookId":"CSE3307","assetClass":"SHARES","symbol":"AAB","sector":"Consumer Discretionary","isin":"DK0060868966","greenEquityDesignation":"","deltaIndicator":"down","contractSize":"","openInterest":"","expirationDate":"","strikePrice":"","noteCode":"","issuerFullName":"","allTradesAvgPrice":"","couponRate":"","exchangeSymbol":"","noteDescription":"","type":"","settlementPrice":"","totalVolume":""},{"fullName":"AAK","currency":"SEK","priceNotation":"","micCode":"","netChange":"-3.60","percentageChange":"-1.17%","bidPrice":"303.00","askPrice":"303.40","lastSalePrice":"303.20","high":"","low":"","volume":"67,632","reportedVolume":"","turnover":"20,595,445","tradesCount":"","lastTraded":"","time":"","netAssetValue":"","navUpdatedTime":"","totalExpenseRatio":"","orderbookId":"SSE36273","assetClass":"SHARES","symbol":"AAK","sector":"Consumer Staples","isin":"SE0011337708","greenEquityDesignation":"","deltaIndicator":"down","contractSize":"","openInterest":"","expirationDate":"","strikePrice":"","noteCode":"","issuerFullName":"","allTradesAvgPrice":"","couponRate":"","exchangeSymbol":"","noteDescription":"","type":"","settlementPrice":"","totalVolume":""},{"fullName":"ABB Ltd","currency":"SEK","priceNotation":"","micCode":"","netChange":"-1.60","percentageChange":"-0.24%","bidPrice":"642.80","askPrice":"643.00","lastSalePrice":"642.60","high":"","low":"","volume":"171,258","reportedVolume":"","turnover":"109,875,894","tradesCount":"","lastTraded":"","time":"","netAssetValue":"","navUpdatedTime":"","totalExpenseRatio":"","orderbookId":"SSE3966","assetClass":"SHARES","symbol":"ABB","sector":"Industrials","isin":"CH0012221716","greenEquityDesignation":"","deltaIndicator":"down","contractSize":"","openInterest":"","expirationDate":"","strikePrice":"","noteCode":"","issuerFullName":"","allTradesAvgPrice":"","couponRate":"","exchangeSymbol":"","noteDescription":"","type":"","settlementPrice":"","totalVolume":""},{"fullName":"Abliva","currency":"SEK","priceNotation":"","micCode":"","netChange":"-0.0032","percentageChange":"-2.28%","bidPrice":"0.1368","askPrice":"0.1382","lastSalePrice":"0.1368","high":"","low":"","volume":"922,839","reportedVolume":"","turnover":"128,189","tradesCount":"","lastTraded":"","time":"","netAssetValue":"","navUpdatedTime":"","totalExpenseRatio":"","orderbookId":"SSE58531","assetClass":"SHARES","symbol":"ABLI","sector":"Health Care","isin":"SE0002575340","greenEquityDesignation":"","deltaIndicator":"down","contractSize":"","openInterest":"","expirationDate":"","strikePrice":"","noteCode":"","issuerFullName":"","allTradesAvgPrice":"","couponRate":"","exchangeSymbol":"","noteDescription":"","type":"","settlementPrice":"","totalVolume":""},{"fullName":"AcadeMedia","currency":"SEK","priceNotation":"","micCode":"","netChange":"-0.90","percentageChange":"-1.45%","bidPrice":"60.70","askPrice":"60.80","lastSalePrice":"60.80","high":"","low":"","volume":"99,706","reportedVolume":"","turnover":"6,174,430","tradesCount":"","lastTraded":"","time":"","netAssetValue":"","navUpdatedTime":"","totalExpenseRatio":"","orderbookId":"SSE123363","assetClass":"SHARES","symbol":"ACAD","sector":"Consumer Discretionary","isin":"SE0007897079","greenEquityDesignation":"","deltaIndicator":"down","contractSize":"","openInterest":"","expirationDate":"","strikePrice":"","noteCode":"","issuerFullName":"","allTradesAvgPrice":"","couponRate":"","exchangeSymbol":"","noteDescription":"","type":"","settlementPrice":"","totalVolume":""},{"fullName":"Acrinova A","currency":"SEK","priceNotation":"","micCode":"","netChange":"0.00","percentageChange":"0.00%","bidPrice":"","askPrice":"","lastSalePrice":"9.40","high":"","low":"","volume":"319","reportedVolume":"","turnover":"2,998","tradesCount":"","lastTraded":"","time":"","netAssetValue":"","navUpdatedTime":"","totalExpenseRatio":"","orderbookId":"SSE186414","assetClass":"SHARES","symbol":"ACRI A","sector":"Real Estate","isin":"SE0015660014","greenEquityDesignation":"","deltaIndicator":"unch","contractSize":"","openInterest":"","expirationDate":"","strikePrice":"","noteCode":"","issuerFullName":"","allTradesAvgPrice":"","couponRate":"","exchangeSymbol":"","noteDescription":"","type":"","settlementPrice":"","totalVolume":""},{"fullName":"Acrinova B","currency":"SEK","priceNotation":"","micCode":"","netChange":"0.00","percentageChange":"0.00%","bidPrice":"9.12","askPrice":"9.18","lastSalePrice":"9.18","high":"","low":"","volume":"2,322","reportedVolume":"","turnover":"21,316","tradesCount":"","lastTraded":"","time":"","netAssetValue":"","navUpdatedTime":"","totalExpenseRatio":"","orderbookId":"SSE219981","assetClass":"SHARES","symbol":"ACRI B","sector":"Real Estate","isin":"SE0015660030","greenEquityDesignation":"","deltaIndicator":"unch","contractSize":"","openInterest":"","expirationDate":"","strikePrice":"","noteCode":"","issuerFullName":"","allTradesAvgPrice":"","couponRate":"","exchangeSymbol":"","noteDescription":"","type":"","settlementPrice":"","totalVolume":""},{"fullName":"Actic Group","currency":"SEK","priceNotation":"","micCode":"","netChange":"-0.18","percentageChange":"-3.13%","bidPrice":"5.58","askPrice":"5.72","lastSalePrice":"5.56","high":"","low":"","volume":"13,549","reportedVolume":"","turnover":"75,603","tradesCount":"","lastTraded":"","time":"","netAssetValue":"","navUpdatedTime":"","totalExpenseRatio":"","orderbookId":"SSE135947","assetClass":"SHARES","symbol":"ATIC","sector":"Consumer Discretionary","isin":"SE0009269467","greenEquityDesignation":"","deltaIndicator":"down","contractSize":"","openInterest":"","expirationDate":"","strikePrice":"","noteCode":"","issuerFullName":"","allTradesAvgPrice":"","couponRate":"","exchangeSymbol":"","noteDescription":"","type":"","settlementPrice":"","totalVolume":""}]},"pagination":{"paginationId":"","total":701,"size":10,"page":1,"totalPages":71},"filters":null},"messages":null,"status":{"timestamp":"2024-12-05T13:48:46+0100","rCode":200,"bCodeMessage":null,"developerMessage":""}}`
	var data RetrievedListings
	json.Unmarshal([]byte(jsonStr), &data)
	assert.True(t, data.Data.InstrumentListing.Rows[0].Symbol == "MAERSK A")
}

func TestRetrieve(t *testing.T) {
	data, _ := RetrieveStocksDummy()
	assert.True(t, len(data.Data.InstrumentListing.Rows[0].OrderbookID) > 2, "Invalid length of ID in return")
	for _, listing := range data.Data.InstrumentListing.Rows {
		assert.True(t, len(listing.Symbol) > 1, "Invalid length of ID in return")
	}
}

// func TestReal(t *testing.T) {
// 	data := RetrieveStocksFromAPIgo()
// 	for _, v := range data.Listings {
// 		fmt.Println(v.Name)
// 	}
// 	fmt.Println(len(data.Listings))
// 	//fmt.Println(data)
// }

func TestRetrieveWithTTLcache(t *testing.T) {
	r := CreateRetriever(RetrieveStocksDummy, 2*time.Second)
	ts := r.RetrieveStocks().Status.Timestamp
	time.Sleep(1 * time.Second)
	assert.Equal(t, ts, r.RetrieveStocks().Status.Timestamp)
	time.Sleep(2 * time.Second)
	assert.NotEqual(t, ts, r.RetrieveStocks().Status.Timestamp)
}

func TestSerialiseHistory(t *testing.T) {
	jsonString := `{"data":{"chartData":{"orderbookId":"SSE36273","assetClass":"SHARES","isin":"SE0011337708","symbol":"AAK","company":"AAK","timeAsOf":"2024-12-04","lastSalePrice":"SEK 303.00","netChange":"-3.80","percentageChange":"-1.23%","deltaIndicator":"up","previousClose":"SEK 306.80"},"CP":[{"z":{"dateTime":"2024-12-03","value":"306.00","high":"306.40","low":"299.60","open":"299.60","close":"306.00","volume":"510,966"},"x":1733184000,"y":306},{"z":{"dateTime":"2024-12-04","value":"306.80","high":"308.60","low":"304.40","open":"305.40","close":"306.80","volume":"449,406"},"x":1733270400,"y":306.8}]},"messages":null,"status":{"timestamp":"2024-12-05T14:02:01+0100","rCode":200,"bCodeMessage":null,"developerMessage":""}}`
	var d History
	json.Unmarshal([]byte(jsonString), &d)
	assert.Equal(t, "SSE36273", d.Data.ChartData.OrderbookID)
}

// func TestRetrieveSpecific(t *testing.T) {
// 	ret, err := RetrieveStock("SSE992")
// 	fmt.Println(err)
// 	fmt.Println(ret)
// }

// func TestRetrieveSpecific(t *testing.T) {
// 	ret, err := RetrieveStocks()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(ret.Data.InstrumentListing.Rows[0])
// }
