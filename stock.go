package stockclient

// Stock is the primary datatype for value of a stock
type Stock struct {
	ID       string `json:"@id"`
	Ticker   string `json:"@nm"`
	Name     string `json:"@fnm"`
	Isin     string `json:"@isin"`
	Currency string `json:"@cr"`
	Time     string `json:"@t"`
	Status   string `json:"@ts"`
	Mnvt     string `json:"@mnvt"`
	Dip      string `json:"@dip"`
	Tr       string `json:"@tr"`
	Bp       string `json:"@bp"`
	Ap       string `json:"@ap"`
	Bv       string `json:"@bv"`
	Av       string `json:"@av"`
	Tv       string `json:"@tv"`
	Hp       string `json:"@hp"`
	Lp       string `json:"@lp"`
	Lsp      string `json:"@lsp"`
	Op       string `json:"@op"`
	Cp       string `json:"@cp"`
	Vf       string `json:"@vf"`
	Vt       string `json:"@vt"`
	Updt     string `json:"@updt"`
}
