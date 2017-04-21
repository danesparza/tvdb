package tvdb

// JSONErrors contains errors that might be send back by the API
type JSONErrors struct {
	// Invalid filters passed to route
	InvalidFilters []string `json:"invalidFilters,omitempty"`
	// Invalid language or translation missing
	InvalidLanguage string `json:"invalidLanguage,omitempty"`
	// Invalid query params passed to route
	InvalidQueryParams []string `json:"invalidQueryParams,omitempty"`
}
