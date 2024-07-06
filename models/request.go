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
	Birthday  string `json:"birthday"`
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
