package exchange

type LoginCredentials struct {
	User string `json:"user"`
	PasswordHash string `json:"password_hash"`
}

type JWTResponse struct {
	AccessToken string `json:"access_token"`
}
