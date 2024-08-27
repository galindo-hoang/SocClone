package utils

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/AuthService/pkg/repositories/models"
	modelhttp "github.com/AuthService/pkg/services/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func GetValue(key string) string {
	err := godotenv.Load()
	if err != nil {
		var opts = os.Getenv(key)
		if len(opts) == 0 {
			fmt.Println(err.Error())
			return ""
		} else {
			return opts
		}
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

func BuildingToken(user models.Users) (modelhttp.Token, error) {
	claimsAccessToken := models.TokenClaims{
		UserName: user.UserName,
		Image:    user.Image,
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ID:        strconv.Itoa(user.ID),
		},
	}

	claimsRefreshToken := models.TokenClaims{
		UserName: user.UserName,
		Image:    "",
		Email:    "",
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ID:        strconv.Itoa(user.ID),
		},
	}

	refreshToken, err := generateToken(claimsRefreshToken, 24*7)
	if err != nil {
		return modelhttp.Token{}, err
	}
	accessToken, err := generateToken(claimsAccessToken, 24)
	fmt.Printf("token:::: %s\n", accessToken)
	if err != nil {
		return modelhttp.Token{}, err
	}

	return modelhttp.Token{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}, nil
}

func generateToken(claims models.TokenClaims, unix int) (string, error) {
	claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Duration(unix) * time.Hour))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(GetValue("JWT_SECRET_KEY")))

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

func JSON2Byte[T any](val T) ([]byte, error) {
	parsedBody, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	return parsedBody, nil
}

func FailOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func ValidateTokenWithId(token string, id string) error {
	parsedToken, err := jwt.ParseWithClaims(token, &models.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(GetValue("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return err
	}
	if claims, ok := parsedToken.Claims.(*models.TokenClaims); ok && parsedToken.Valid {
		if claims.ID != id {
			return errors.New("invalid token")
		} else {
			return nil
		}
	} else {
		return errors.New("invalid token")
	}
}

func IsValidToken(token string) error {
	parsedToken, err := jwt.ParseWithClaims(token, &models.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(GetValue("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return err
	}
	if claims, ok := parsedToken.Claims.(*models.TokenClaims); ok && parsedToken.Valid {
		if claims.ExpiresAt.Unix() < time.Now().Unix() {
			return errors.New("invalid token")
		} else {
			return nil
		}
	} else {
		return errors.New("invalid token")
	}
}
