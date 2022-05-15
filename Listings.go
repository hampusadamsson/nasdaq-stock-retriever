package stockclient

type RetrievedListings struct {
	Status    string  `json:"@status"`
	TimeStamp string  `json:"@ts"`
	Sid       string  `json:"@s"`
	Listings  []Stock `json:"inst"`
}
