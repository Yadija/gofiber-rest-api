package helper

import (
	"time"

	jtoken "github.com/golang-jwt/jwt/v4"
)

func GenerateTokens(username string) (string, string, error) {
	// create access token
	at := jtoken.NewWithClaims(jtoken.SigningMethodHS256, jtoken.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Minute * 15).Unix(),
	})
	accessToken, err := at.SignedString([]byte("ACCESS_TOKEN_KEY"))
	if err != nil {
		return "", "", err
	}

	// create refresh token
	rt := jtoken.NewWithClaims(jtoken.SigningMethodHS256, jtoken.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	refreshToken, err := rt.SignedString([]byte("REFRESH_TOKEN_KEY"))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
