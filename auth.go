package tvdb

// AuthRequest represents the login parameters
type AuthRequest struct {
	APIKey   string `json:"apikey,omitempty"`
	Username string `json:"username,omitempty"`
	UserKey  string `json:"userkey,omitempty"`
}

// AuthResponse represents the authorization response
type AuthResponse struct {
	Token string `json:"token"`
}
