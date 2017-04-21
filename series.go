package tvdb

// SeriesRequest is used to request additional series information
type SeriesRequest struct {
	SeriesID int `json:"id"`
}

// Series represents a response for series information from the TVDB service
type Series struct {
	Added           string   `json:"added"`
	AirsDayOfWeek   string   `json:"airsDayOfWeek"`
	AirsTime        string   `json:"airsTime"`
	Aliases         []string `json:"aliases"`
	Banner          string   `json:"banner"`
	FirstAired      string   `json:"firstAired"`
	Genre           []string `json:"genre"`
	ID              int      `json:"id"`
	ImdbID          string   `json:"imdbId"`
	LastUpdated     int      `json:"lastUpdated"`
	Network         string   `json:"network"`
	NetworkID       string   `json:"networkId"`
	Overview        string   `json:"overview"`
	Rating          string   `json:"rating"`
	Runtime         string   `json:"runtime"`
	SeriesID        string   `json:"seriesId"`
	SeriesName      string   `json:"seriesName"`
	SiteRating      float64  `json:"siteRating"`
	SiteRatingCount int      `json:"siteRatingCount"`
	Status          string   `json:"status"`
	Zap2itID        string   `json:"zap2itId"`
}

// SeriesData represents the list of responses to get series information
type SeriesData struct {
	Data   Series     `json:"data"`
	Errors JSONErrors `json:"errors,omitempty"`
}

// SeriesEpisodesRequest represents a request to get episode information for a series
type SeriesEpisodesRequest struct {
	SeriesID int

	/* Query parameters */
	AiredEpisode int `json:"airedEpisode"`
	AiredSeason  int `json:"airedSeason"`
	DVDEpisode   int `json:"dvdEpisode"`
	DVDSeason    int `json:"dvdSeason"`
	IMDBId       int `json:"imdbId"`
}

// BasicEpisode represents a response for episode information from the TVDB service
type BasicEpisode struct {
	AbsoluteNumber     int    `json:"absoluteNumber"`
	AiredEpisodeNumber int    `json:"airedEpisodeNumber"`
	AiredSeason        int    `json:"airedSeason"`
	AiredSeasonID      int    `json:"airedSeasonID"`
	DVDEpisodeNumber   int    `json:"dvdEpisodeNumber"`
	DVDSeason          int    `json:"dvdSeason"`
	EpisodeName        string `json:"episodeName"`
	FirstAired         string `json:"firstAired"`
	ID                 int    `json:"id"`
	LastUpdated        int    `json:"lastUpdated"`
	Overview           string `json:"overview"`
}

// Links represents the paging information for multiple episode responses
type Links struct {
	FirstPage    int `json:"first"`
	LastPage     int `json:"last"`
	NextPage     int `json:"next"`
	PreviousPage int `json:"prev"`
}

// SeriesEpisodes represents the list of responses to get episode information
type SeriesEpisodes struct {
	Links  Links          `json:"links"`
	Data   []BasicEpisode `json:"data"`
	Errors JSONErrors     `json:"errors,omitempty"`
}

// SeriesEpisodesSummary Returns a summary of the episodes and seasons available for the series.
type SeriesEpisodesSummary struct {
	// Number of all aired episodes for this series
	AiredEpisodes string   `json:"airedEpisodes,omitempty"`
	AiredSeasons  []string `json:"airedSeasons,omitempty"`
	// Number of all dvd episodes for this series
	DvdEpisodes string   `json:"dvdEpisodes,omitempty"`
	DvdSeasons  []string `json:"dvdSeasons,omitempty"`
}

// SeriesActorsData contains information about a single actor
type SeriesActorsData struct {
	ID          int    `json:"id"`
	Image       string `json:"image"`
	ImageAdded  string `json:"imageAdded"`
	ImageAuthor int    `json:"imageAuthor"`
	LastUpdated string `json:"lastUpdated"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	SeriesID    int    `json:"seriesId"`
	SortOrder   int    `json:"sortOrder"`
}

// SeriesActors is the response of the api when asking for authors
type SeriesActors struct {
	Data []SeriesActorsData `json:"data"`
}

// SeriesImageQueryRequest used to query images for a given series and type
type SeriesImageQueryRequest struct {
	SeriesID   int    `json:"id"`
	KeyType    string `json:"keyType"`
	Resulution string `json:"resulution"`
	SubKey     string `json:"subKey"`
}

// SeriesImagesCount one single image response
type SeriesImagesCount struct {
	FileName    string             `json:"fileName"`
	ID          int                `json:"id"`
	KeyType     string             `json:"keyType"`
	LanguageID  int                `json:"languageId"`
	RatingsInfo map[string]float64 `json:"ratingsInfo"`
	Resolution  string             `json:"resulution"`
	SubKey      string             `json:"subKey"`
	Thumbnail   string             `json:"thumbnail"`
}

// SeriesImagesCounts is the response of the api when asking for images
type SeriesImagesCounts struct {
	Data []SeriesImagesCount `json:"data"`
}
