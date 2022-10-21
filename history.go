package stockclient

// History data struct
type History struct {
	Status string `json:"@status"`
	Ts     string `json:"@ts"`
	Data   []Data `json:"data"`
}

// Data data struct
type Data struct {
	InstData  InstData  `json:"InstData"`
	ChartData ChartData `json:"ChartData"`
}

// InstData data struct
type InstData struct {
	ID   string `json:"@id"`
	Nm   string `json:"@nm"`
	Fnm  string `json:"@fnm"`
	Isin string `json:"@isin"`
	Tp   string `json:"@tp"`
	Chp  string `json:"@chp"`
	Ycp  string `json:"@ycp"`
}

// ChartData data struct
type ChartData struct {
	Cp [][]float64 `json:"cp"`
}
