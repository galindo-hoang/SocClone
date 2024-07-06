package utils

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/AuthService/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func GetValue(key string) string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return os.Getenv(key)
}

func HashAndSalt(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePassword(hashedPwd string, plainPwd []byte) error {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return err
	}
	return nil
}

func BuildingToken(user models.Users) (models.Token, error) {
	claimsAccessToken := models.TokenClaims{
		UserName: user.UserName,
		Image:    user.Image,
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ID:        string(user.ID),
		},
	}

	claimsRefreshToken := models.TokenClaims{
		UserName: user.UserName,
		Image:    "",
		Email:    "",
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ID:        string(user.ID),
		},
	}

	refreshToken, err := generateToken(claimsRefreshToken, 24*7)
	if err != nil {
		return models.Token{}, err
	}
	accessToken, err := generateToken(claimsAccessToken, 24)
	if err != nil {
		return models.Token{}, err
	}

	return models.Token{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}, nil
}

func generateToken(claims models.TokenClaims, unix int) (string, error) {
	claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Duration(unix) * time.Hour))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(GetValue("JWT_SECRET_KEY"))

	if err != nil {
		return "", err
	}
	return ss, nil
}

const otpChars = "1234567890QWERTYUIOPASDFGHJKLZXCVBNM"

func GenerateOTP(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < length; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}

func ToJsonFromByte[T any](val []byte) (T, error) {
	var data = new(T)
	err := json.Unmarshal(val, &data)
	if err != nil {
		return *data, err
	}
	return *data, nil
}
