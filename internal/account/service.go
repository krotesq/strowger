package account

import (
	"context"
	"os"
	"encoding/base64"

	"github.com/golang-jwt/jwt/v5"
	pw "github.com/krotesq/strowger/internal/password"
)

// define what functions should be implemented on the account repo
type accountRepository interface {
	findByUsername(ctx context.Context, username string) (*account, error)
}

type service struct {
	accountRepo accountRepository
}

func newService(accountRepo accountRepository) *service {
	return &service{accountRepo: accountRepo}
}

func (service *service) findByUsername(ctx context.Context, username string) (*account, error) {
	return service.accountRepo.findByUsername(ctx, username)
}

func (service *service) login(ctx context.Context, username, password string) (string, error) {
	// load account
	account, err := service.accountRepo.findByUsername(ctx, username)
	if err != nil {
		return "", err
	}
	
	// verify password
	if err := pw.Compare(password, account.PasswordHash); err != nil {
		return "", err
	}
	
	// generate jwt
	key := os.Getenv("JWT_SECRET")
	keyBytes, err := base64.StdEncoding.DecodeString(key)
	
	if err != nil {
		return "", err
	}
	
	t := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{ 
	    "iss": "strowger", 
	    "sub": account.Username,
  		},
	)
	s, err := t.SignedString(keyBytes)
	return s, nil
}