package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/needon1997/theshop-api/internal/common/config"
	"time"
)

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

type JWTUserInfoClaim struct {
	Id       int64  `json:"id"`
	Nickname string `json:"nickname"`
	Role     uint8  `json:"role"`
	jwt.StandardClaims
}

func GetStandardClaim() jwt.StandardClaims {
	return jwt.StandardClaims{ExpiresAt: time.Now().Unix() + config.ServerConfig.JWTConfig.ExpireAt, Issuer: config.ServerConfig.JWTConfig.Issuer}
}
func NewJwtToken(claim JWTUserInfoClaim) string {
	return NewJwtTokenWithSecret(claim, config.ServerConfig.JWTConfig.Secret)
}
func NewJwtTokenWithSecret(claim JWTUserInfoClaim, secret string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedString, _ := token.SignedString([]byte(secret))
	return signedString
}
func ValidateTokenAndRetrieveInfo(tokenString string) (*JWTUserInfoClaim, error) {
	return ValidateTokenAndRetrieveInfoWithSecret(tokenString, config.ServerConfig.JWTConfig.Secret)
}
func ValidateTokenAndRetrieveInfoWithSecret(tokenString string, secret string) (*JWTUserInfoClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTUserInfoClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claim, ok := token.Claims.(*JWTUserInfoClaim); ok && token.Valid {
			return claim, nil
		} else {
			return nil, TokenInvalid
		}
	} else {
		return nil, TokenInvalid
	}
}

func RefreshToken(tokenString string) (string, error) {
	return RefreshTokenWithSecret(tokenString, config.ServerConfig.JWTConfig.Secret)
}
func RefreshTokenWithSecret(tokenString, secret string) (string, error) {
	claim, err := ValidateTokenAndRetrieveInfoWithSecret(tokenString, secret)
	if err != nil {
		return "", err
	}
	claim.ExpiresAt = time.Now().Unix() + config.ServerConfig.JWTConfig.ExpireAt
	token := NewJwtTokenWithSecret(*claim, secret)
	return token, nil
}
