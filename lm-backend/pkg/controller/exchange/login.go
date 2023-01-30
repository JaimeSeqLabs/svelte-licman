package exchange

type LoginCredentials struct {
	User string `json:"user"`
	Password string `json:"password"`
}

type JWTResponse struct {
	AccessToken string `json:"access_token"`
}
