package account

import (
	"context"
	"errors"
	"net/netip"
	"os"
	"strconv"
	"time"

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

func (service *service) login(ctx context.Context, username, password, userAgent, ip string) (*account, string, string, error) {
	// load account
	account, err := service.repository.findByUsername(ctx, username)
	if err != nil {
		return nil, "", "", err
	}

	if !account.Active {
		return nil, "", "", errors.New("Account is deactivated. Please contact your administrator.")
	}

	// check if account has == 3 failed logins
	if account.FailedLoginAttempts >= 3 {
		return nil, "", "", errors.New("Too many failed login attempts. Please contact your administrator.")
	}

	// verify password
	if err := auth.ComparePassword(password, account.PasswordHash); err != nil {
		// increment failed login attempts
		account, _ = service.repository.incrementFailedLoginAttemptsByID(ctx, account.ID)
		return nil, "", "", err
	}

	// reset failed login attempts if > 0
	if account.FailedLoginAttempts > 0 {
		account, _ = service.repository.resetFailedLoginAttemptsByID(ctx, account.ID)
	}

	// generate jwt
	secretBase64 := os.Getenv("JWT_SECRET")
	accessToken, err := auth.GenerateJWT(account.ID, "strowger", secretBase64)

	// generate refresh token
	refreshToken, refreshTokenHash, err := auth.GenerateRefreshToken()

	// parse ip
	ipAddr, err := netip.ParseAddr(ip)
	if err != nil {
		return nil, "", "", err
	}

	// save refreshToken to db
	_, err = service.repository.createRefreshToken(ctx, account.ID, refreshTokenHash, time.Now().AddDate(0, 0, 30), userAgent, ipAddr)
	if err != nil {
		return nil, "", "", err
	}

	return account, accessToken, refreshToken, nil
}

func (service *service) refresh(ctx context.Context, token string) (string, string, error) {

	// hash token
	tokenHash := auth.HashRefreshToken(token)

	// find token in db
	refreshToken, err := service.repository.findRefreshTokenByHash(ctx, tokenHash)
	if err != nil {
		return "", "", nil
	}

	// check if token is not expired and not revoked
	if refreshToken.RevokedAt != nil {
		return "", "", errors.New("Refresh token was revoked")
	}

	if refreshToken.ExpiresAt.Before(time.Now()) {
		return "", "", errors.New("Refresh token is expired")
	}

	// revoke old refresh token
	service.repository.revokeRefreshTokenByHash(ctx, tokenHash)

	// generate new refresh token
	newToken, newTokenHash, err := auth.GenerateRefreshToken()
	if err != nil {
		return "", "", err
	}

	// parse ip
	ip, err := netip.ParseAddr("1.1.1.1")
	if err != nil {
		return "", "", err
	}

	// save signed token in db
	exp_days_str := os.Getenv("REFRESH_EXP_DAYS")
	exp_days, err := strconv.Atoi(exp_days_str)
	if err != nil {
		return "", "", err
	}
	refreshToken, err = service.repository.createRefreshToken(ctx, refreshToken.AccountID, newTokenHash, time.Now().AddDate(0, 0, exp_days), "idk", ip)
	if err != nil {
		return "", "", err
	}

	// generate new jwt
	secret := os.Getenv("JWT_SECRET")
	jwt, err := auth.GenerateJWT(refreshToken.AccountID, "strowger", secret)
	if err != nil {
		return "", "", err
	}

	return newToken, jwt, nil
}

func (service *service) me(ctx context.Context) (*account, error) {
	id, ok := auth.AccountIDFromContext(ctx)
	if !ok {
		return nil, errors.New("Unable to load id from context")
	}
	return service.repository.findByID(ctx, id)
}
