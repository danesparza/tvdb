package tvdb

//	TVDBService represents all operations in the TVDB v2.0 API
type TVDBService interface {

	//	Login to the system and get a
	Login(request AuthRequest) (AuthResponse, error)

	//	Get a specific config string
	RefreshToken(responseToRefresh AuthResponse) (AuthResponse, error)

	//	Search for a given series
	SeriesSearch(request SearchRequest) (SearchResponses, error)

	//	Get all episodes for a given seriesId
	EpisodesForSeries(request EpisodeRequest) ([]EpisodeResponse, error)
}
