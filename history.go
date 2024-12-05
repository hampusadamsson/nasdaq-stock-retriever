package stockclient

type History struct {
	Data     Data   `json:"data"`
	Messages any    `json:"messages"`
	Status   Status `json:"status"`
}
type ChartData struct {
	OrderbookID      string `json:"orderbookId"`
	AssetClass       string `json:"assetClass"`
	Isin             string `json:"isin"`
	Symbol           string `json:"symbol"`
	Company          string `json:"company"`
	TimeAsOf         string `json:"timeAsOf"`
	LastSalePrice    string `json:"lastSalePrice"`
	NetChange        string `json:"netChange"`
	PercentageChange string `json:"percentageChange"`
	DeltaIndicator   string `json:"deltaIndicator"`
	PreviousClose    string `json:"previousClose"`
}
type Z struct {
	DateTime string `json:"dateTime"`
	Value    string `json:"value"`
	High     string `json:"high"`
	Low      string `json:"low"`
	Open     string `json:"open"`
	Close    string `json:"close"`
	Volume   string `json:"volume"`
}
type Cp struct {
	Z Z       `json:"z"`
	X int     `json:"x"`
	Y float64 `json:"y"`
}
type Data struct {
	ChartData ChartData `json:"chartData"`
	Cp        []Cp      `json:"CP"`
}
