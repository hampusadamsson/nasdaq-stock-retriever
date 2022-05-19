package stockclient

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSerialise(t *testing.T) {
	const jsonStr = "jQuery11110636858574942031_1622134931936( {\r\n  \"@status\": \"1\",\r\n  \"@ts\": \"1652210182823\",\r\n  \"@s\": \"se20ssplink02\",\r\n  \"inst\":   [\r\n        {\r\n      \"@id\": \"SSE175638\",\r\n      \"@nm\": \"8TRA\",\r\n      \"@fnm\": \"TRATON\",\r\n      \"@isin\": \"DE000TRAT0N7\",\r\n      \"@cr\": \"SEK\",\r\n      \"@t\": \"18:00:00\",\r\n      \"@ts\": \"Closed\",\r\n      \"@mnvt\": \"\",\r\n      \"@dip\": \"2\",\r\n      \"@tr\": \"\",\r\n      \"@bp\": \"161.30\",\r\n      \"@ap\": \"161.70\",\r\n      \"@bv\": \"500\",\r\n      \"@av\": \"600\",\r\n      \"@tv\": \"95112\",\r\n      \"@hp\": \"167.10\",\r\n      \"@lp\": \"161.50\",\r\n      \"@lsp\": \"161.90\",\r\n      \"@op\": \"164.10\",\r\n      \"@cp\": \"161.90\",\r\n      \"@vf\": \"\",\r\n      \"@vt\": \"\",\r\n      \"@updt\": \"2022-05-10\"\r\n    }\r\n]\r\n} );"
	inputFmt := jsonStr[42 : len(jsonStr)-2]
	var data RetrievedListings
	json.Unmarshal([]byte(inputFmt), &data)
	assert.True(t, data.Listings[0].Status == "Closed", "Can't serialize json")
}

func TestRetrieve(t *testing.T) {
	data := RetrieveStocksDummy()
	assert.True(t, len(data.Listings[0].ID) > 2, "Invalid length of ID in return")
	for _, listing := range data.Listings {
		assert.True(t, len(listing.ID) > 2, "Invalid length of ID in return")
	}
}

func TestRetrieveWithTTLcache(t *testing.T) {
	r := CreateRetriever(RetrieveStocksDummy, 2*time.Second)
	ts := r.RetrieveStocks().TimeStamp
	time.Sleep(1 * time.Second)
	assert.Equal(t, ts, r.RetrieveStocks().TimeStamp)
	time.Sleep(2 * time.Second)
	assert.NotEqual(t, ts, r.RetrieveStocks().TimeStamp)
}

func TestSerialiseHistory(t *testing.T) {
	jsonString := "{  \"@status\": \"1\",  \"@ts\": \"1652819614811\",  \"data\": [    {      \"instData\": {        \"@id\": \"SSE992\",        \"@nm\": \"HM B\",        \"@fnm\": \"Hennes & Mauritz B\",        \"@isin\": \"SE0000106270\",        \"@tp\": \"S\",        \"@chp\": \"1.9001\",        \"@ycp\": \"130.52\"      },      \"chartData\": {        \"cp\": [          [            536544000000,            2.325          ],          [            536803200000,            2.3          ],          [            536976000000,            2.225          ],          [            1652659200000,            130.52          ]        ]      }    }  ]}"
	var d History
	json.Unmarshal([]byte(jsonString), &d)
	assert.Equal(t, "SSE992", d.Data[0].InstData.ID)
}

// func TestRetrieveSpecific(t *testing.T) {
// 	ret := RetrieveStock("SSE992")
// 	fmt.Println(ret)
// }
