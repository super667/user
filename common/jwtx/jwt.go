package jwtx

import "github.com/golang-jwt/jwt"

func GetToken(secretKey string, iat, seconds int64, name string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["name"] = name
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
