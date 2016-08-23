package tvdb

//	AuthRequest represents the login parameters
type AuthRequest struct {
	ApiKey   string `json:"apikey,omitempty"`
	Username string `json:"username,omitempty"`
	UserKey  string `json:"userkey,omitempty"`
}

type AuthResponse struct {
	Token string `json:"token"`
}
