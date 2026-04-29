package auth

import (
	"encoding/base64"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(sub, iss, secretBase64 string) (string, error) {
	keyBytes, err := base64.StdEncoding.DecodeString(secretBase64)
	if err != nil {
		return "", err
	}

	exp_min_str := os.Getenv("JWT_EXP_MINUTES")
	exp_min, err := strconv.Atoi(exp_min_str)
	if err != nil {
		return "", err
	}

	t := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss": iss,
			"sub": sub,
			"exp": time.Now().Add(time.Minute * time.Duration(exp_min)).Unix(),
		},
	)
	s, err := t.SignedString(keyBytes)
	return s, err
}

func ValidateJWT(tokenString, secretBase64 string) (string, error) {
	keyBytes, err := base64.StdEncoding.DecodeString(secretBase64)
	if err != nil {
		return "", err
	}

	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (any, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, errors.New("unexpected signing method")
			}
			return keyBytes, nil
		},
	)

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid claims")
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		return "", errors.New("missing subject")
	}

	return sub, nil
}
