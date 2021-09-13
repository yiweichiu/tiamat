package models

type QuoteResults struct {
	LongName           string  `json:"longName"`
	ShortName          string  `json:"shortName"`
	DisplayName        string  `json:"displayName"`
	RegularMarketPrice float64 `json:"regularMarketPrice"`
}

type QuoteResponses struct {
	Err     string         `json:"error"`
	Results []QuoteResults `json:"result"`
}

type YahooQuotes struct {
	QuoteResponse QuoteResponses `json:"quoteResponse"`
}
