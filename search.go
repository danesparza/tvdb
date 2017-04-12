package tvdb

// SearchRequest represents the request
type SearchRequest struct {
	Name     string
	IMDBID   string
	Zap2ItID string
}

// SeriesInfo represents TV series information in the search response
type SeriesInfo struct {
	Aliases    []string `json:"aliases"`
	Banner     string   `json:"banner"`
	FirstAired string   `json:"firstAired"`
	ID         int      `json:"id"`
	Network    string   `json:"network"`
	Overview   string   `json:"overview"`
	SeriesName string   `json:"seriesName"`
	Status     string   `json:"status"`
}

// SearchResponses represents a search response from the service
type SearchResponses struct {
	Data []SeriesInfo `json:"data"`
}
