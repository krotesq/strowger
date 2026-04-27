package account

import (
	"context"
	"errors"
	"os"

	"github.com/krotesq/strowger/internal/auth"
)

type service struct {
	repository *repository
}

func newService(repository *repository) *service {
	return &service{repository: repository}
}

func (service *service) findByID(ctx context.Context, id string) (*account, error) {
	return service.repository.findByID(ctx, id)
}

func (service *service) deleteByID(ctx context.Context, id string) (*account, error) {
	return service.repository.deleteByID(ctx, id)
}

func (service *service) deactivateByID(ctx context.Context, id string) (*account, error) {
	return service.repository.deactivateByID(ctx, id)
}

func (service *service) activateByID(ctx context.Context, id string) (*account, error) {
	return service.repository.activateByID(ctx, id)
}

func (service *service) create(ctx context.Context, username, password string) (*account, error) {
	passwordHash, err := auth.HashPassword(password, 10)
	if err != nil {
		return nil, err
	}
	return service.repository.create(ctx, username, passwordHash)
}

func (service *service) login(ctx context.Context, username, password string) (*account, string, error) {
	// load account
	account, err := service.repository.findByUsername(ctx, username)
	if err != nil {
		return nil, "", err
	}

	if !account.Active {
		return nil, "", errors.New("Account is deactivated. Please contact your administrator.")
	}

	// check if account has == 3 failed logins
	if account.FailedLoginAttempts >= 3 {
		return nil, "", errors.New("Too many failed login attempts. Please contact your administrator.")
	}

	// verify password
	if err := auth.ComparePassword(password, account.PasswordHash); err != nil {
		// increment failed login attempts
		account, _ = service.repository.incrementFailedLoginAttemptsByID(ctx, account.ID)
		return nil, "", err
	}

	// reset failed login attempts if > 0
	if account.FailedLoginAttempts > 0 {
		account, _ = service.repository.resetFailedLoginAttemptsByID(ctx, account.ID)
	}

	// generate jwt
	secretBase64 := os.Getenv("JWT_SECRET")
	token, err := auth.GenerateToken(account.ID, "strowger", secretBase64)

	return account, token, nil
}
