package tvdb

//	AuthRequest represents the login parameters
type SearchRequest struct {
	Name     string
	IMDBId   string
	Zap2ItId string
}

type SeriesInfo struct {
	Aliases    []string `json:"aliases"`
	Banner     string   `json:"banner"`
	FirstAired string   `json:"firstAired"`
	Id         int      `json:"id"`
	Network    string   `json:"network"`
	Overview   string   `json:"overview"`
	SeriesName string   `json:"seriesName"`
	Status     string   `json:"status"`
}

type SearchResponses struct {
	Data []SeriesInfo `json:"data"`
}
