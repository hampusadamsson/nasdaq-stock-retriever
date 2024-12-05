package stockclient

type RetrievedListings struct {
	Data     Data2  `json:"data"`
	Messages any    `json:"messages"`
	Status   Status `json:"status"`
}

//	type Headers struct {
//		FullName               string `json:"fullName"`
//		Currency               string `json:"currency"`
//		LastSalePrice          string `json:"lastSalePrice"`
//		NetChange              string `json:"netChange"`
//		PercentageChange       string `json:"percentageChange"`
//		BidPrice               string `json:"bidPrice"`
//		AskPrice               string `json:"askPrice"`
//		Volume                 string `json:"volume"`
//		Turnover               string `json:"turnover"`
//		GreenEquityDesignation string `json:"greenEquityDesignation"`
//		Sector                 string `json:"sector"`
//	}
type Stock struct {
	FullName               string `json:"fullName"`
	Currency               string `json:"currency"`
	PriceNotation          string `json:"priceNotation"`
	MicCode                string `json:"micCode"`
	NetChange              string `json:"netChange"`
	PercentageChange       string `json:"percentageChange"`
	BidPrice               string `json:"bidPrice"`
	AskPrice               string `json:"askPrice"`
	LastSalePrice          string `json:"lastSalePrice"`
	High                   string `json:"high"`
	Low                    string `json:"low"`
	Volume                 string `json:"volume"`
	ReportedVolume         string `json:"reportedVolume"`
	Turnover               string `json:"turnover"`
	TradesCount            string `json:"tradesCount"`
	LastTraded             string `json:"lastTraded"`
	Time                   string `json:"time"`
	NetAssetValue          string `json:"netAssetValue"`
	NavUpdatedTime         string `json:"navUpdatedTime"`
	TotalExpenseRatio      string `json:"totalExpenseRatio"`
	OrderbookID            string `json:"orderbookId"`
	AssetClass             string `json:"assetClass"`
	Symbol                 string `json:"symbol"`
	Sector                 string `json:"sector"`
	Isin                   string `json:"isin"`
	GreenEquityDesignation string `json:"greenEquityDesignation"`
	DeltaIndicator         string `json:"deltaIndicator"`
	ContractSize           string `json:"contractSize"`
	OpenInterest           string `json:"openInterest"`
	ExpirationDate         string `json:"expirationDate"`
	StrikePrice            string `json:"strikePrice"`
	NoteCode               string `json:"noteCode"`
	IssuerFullName         string `json:"issuerFullName"`
	AllTradesAvgPrice      string `json:"allTradesAvgPrice"`
	CouponRate             string `json:"couponRate"`
	ExchangeSymbol         string `json:"exchangeSymbol"`
	NoteDescription        string `json:"noteDescription"`
	Type                   string `json:"type"`
	SettlementPrice        string `json:"settlementPrice"`
	TotalVolume            string `json:"totalVolume"`
}
type InstrumentListing struct {
	// Headers Headers `json:"headers"`
	Rows []Stock `json:"rows"`
}

//	type Pagination struct {
//		PaginationID string `json:"paginationId"`
//		Total        int    `json:"total"`
//		Size         int    `json:"size"`
//		Page         int    `json:"page"`
//		TotalPages   int    `json:"totalPages"`
//	}
type Data2 struct {
	InstrumentListing InstrumentListing `json:"instrumentListing"`
	// Pagination        Pagination        `json:"pagination"`
	// Filters           any               `json:"filters"`
}
type Status struct {
	Timestamp string `json:"timestamp"`
	// RCode            int    `json:"rCode"`
	// BCodeMessage     any    `json:"bCodeMessage"`
	// DeveloperMessage string `json:"developerMessage"`
}
