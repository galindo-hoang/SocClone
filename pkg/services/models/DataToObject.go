package models

type LoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type SignUpRequest struct {
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Birthday  int64  `json:"birthday"`
	Password  string `json:"password"`
}

type EditUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Birthday  string `json:"birthday"`
	Password  string `json:"password"`
}

type ValidateUserRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	OTP      string `json:"otp"`
}

type UploadImageRequest struct {
	Image       string `json:"image"`
	Name        string `json:"name"`
	ContentType string `json:"content_type"`
}

type LogInResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Image    string `json:"image"`
	UserName string `json:"user_name"`
	Token    Token  `json:"token"`
}

type RegisterResponse struct {
	Id       string `json:"id"`
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
