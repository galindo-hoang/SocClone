package models

type LogInResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Image    string `json:"image"`
	UserName string `json:"user_name"`
}

type RegisterResponse struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}

type Token struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}
