package util

import (
	"github.com/dgrijalva/jwt-go/v4"
)

const SecretKey = "secret"

func GenerateJwt(issuer string) (string, error) {
	// expirationTime := time.Now().Add(time.Hour * 24).Unix() // Token expires after 24 hours
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: issuer,
		// ExpiresAt: jwt.NewNumericDate(expirationTime), // Use jwt.NewNumericDate() with Unix timestamp
	})
	return claims.SignedString([]byte(SecretKey))
}

func ParseJwt(cookie string) (string, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil || !token.Valid {
		return "", err
	}
	claims := token.Claims.(*jwt.StandardClaims)
	return claims.Issuer, nil
}
