package tvdb

//	AuthRequest represents the login parameters
type AuthRequest struct {
	ApiKey   string `json:"apikey"`
	Username string `json:"username"`
	UserKey  string `json:"userkey"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

//	TVDBService represents all operations in the TVDB v2.0 API
type TVDBService interface {

	//	Login to the system and get a
	Login(request AuthRequest) (AuthResponse, error)

	//	Get a specific config string
	RefreshToken(responseToRefresh AuthResponse) (AuthResponse, error)
}
