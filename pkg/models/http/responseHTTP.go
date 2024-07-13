package model_http

type LogInResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Image    string `json:"image"`
	UserName string `json:"user_name"`
	Token    Token  `json:"token"`
}

type RegisterResponse struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}

type Token struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

type Response[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}
