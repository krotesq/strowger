package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

func GenerateRefreshToken() (string, []byte, error) {
	// generate random bytes
	tokenBytes := make([]byte, 32)
	rand.Read(tokenBytes)

	// encode bytes to base64
	token := base64.StdEncoding.EncodeToString(tokenBytes)

	// hash base64 for db
	tokenHash := sha256.Sum256([]byte(token))

	return token, tokenHash[:], nil
}

func HashRefreshToken(token string) ([]byte) {
	tokenHash := sha256.Sum256([]byte(token))
	return tokenHash[:]
}
