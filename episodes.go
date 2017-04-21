package tvdb

// Episode contains all informations about a single Episode
type Episode struct {
	AbsoluteNumber     int      `json:"absoluteNumber,omitempty"`
	AiredEpisodeNumber int      `json:"airedEpisodeNumber,omitempty"`
	AiredSeason        int      `json:"airedSeason,omitempty"`
	AirsAfterSeason    int      `json:"airsAfterSeason,omitempty"`
	AirsBeforeEpisode  int      `json:"airsBeforeEpisode,omitempty"`
	AirsBeforeSeason   int      `json:"airsBeforeSeason,omitempty"`
	Director           string   `json:"director,omitempty"`
	Directors          []string `json:"directors,omitempty"`
	DvdChapter         float32  `json:"dvdChapter,omitempty"`
	DvdDiscID          string   `json:"dvdDiscid,omitempty"`
	DvdEpisodeNumber   float32  `json:"dvdEpisodeNumber,omitempty"`
	DvdSeason          int      `json:"dvdSeason,omitempty"`
	EpisodeName        string   `json:"episodeName,omitempty"`
	Filename           string   `json:"filename,omitempty"`
	FirstAired         string   `json:"firstAired,omitempty"`
	GuestStars         []string `json:"guestStars,omitempty"`
	ID                 int      `json:"id,omitempty"`
	ImdbID             string   `json:"imdbId,omitempty"`
	LastUpdated        int      `json:"lastUpdated,omitempty"`
	LastUpdatedBy      int      `json:"lastUpdatedBy,omitempty"`
	Overview           string   `json:"overview,omitempty"`
	ProductionCode     string   `json:"productionCode,omitempty"`
	SeriesID           int      `json:"seriesId,omitempty"`
	ShowURL            string   `json:"showUrl,omitempty"`
	SiteRating         float32  `json:"siteRating,omitempty"`
	SiteRatingCount    int      `json:"siteRatingCount,omitempty"`
	ThumbAdded         string   `json:"thumbAdded,omitempty"`
	ThumbAuthor        int      `json:"thumbAuthor,omitempty"`
	ThumbHeight        string   `json:"thumbHeight,omitempty"`
	ThumbWidth         string   `json:"thumbWidth,omitempty"`
	Writers            []string `json:"writers,omitempty"`
}

// EpisodeRecordData contains the Server Response when requesting a single Episode
type EpisodeRecordData struct {
	Data   Episode    `json:"data,omitempty"`
	Errors JSONErrors `json:"errors,omitempty"`
}
