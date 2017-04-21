package tvdb

// SeriesSearchRequest represents the request
type SeriesSearchRequest struct {
	Name     string
	IMDBID   string
	Zap2ItID string
}

// SeriesSearchData represents TV series information in the search response
type SeriesSearchData struct {
	Aliases    []string `json:"aliases"`
	Banner     string   `json:"banner"`
	FirstAired string   `json:"firstAired"`
	ID         int      `json:"id"`
	Network    string   `json:"network"`
	Overview   string   `json:"overview"`
	SeriesName string   `json:"seriesName"`
	Status     string   `json:"status"`
}

// SeriesSearchResponses represents a search response from the service
type SeriesSearchResponses struct {
	Data []SeriesSearchData `json:"data"`
}
