package stockclient

// RetrievedListings is a meta data struct type. You are probably interested in Listings of type []Stock
type RetrievedListings struct {
	Status    string  `json:"@status"`
	TimeStamp string  `json:"@ts"`
	Sid       string  `json:"@s"`
	Listings  []Stock `json:"inst"`
}
