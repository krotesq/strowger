package auth

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/krotesq/strowger/internal/response"
)

type contextKey string

const accountIDKey contextKey = "accountID"

func WithAccountID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, accountIDKey, id)
}

func AccountIDFromContext(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(accountIDKey).(string)
	return id, ok
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get jwt cookie from request
		cookie, err := r.Cookie("access_token")
		if errors.Is(err, http.ErrNoCookie) {
			response.Send(w, http.StatusUnauthorized, "Auth failed", nil)
			return
		}
		if err != nil {
			log.Println(err.Error())
			response.Send(w, http.StatusInternalServerError, err.Error(), nil)
			return
		}

		// verify jwt
		secretBase64 := os.Getenv("JWT_SECRET")
		sub, err := ValidateToken(cookie.Value, secretBase64)
		if err != nil {
			response.Send(w, http.StatusUnauthorized, "Auth failed", nil)
			return
		}

		log.Printf("User %s authorized with jwt\n", sub)

		ctx := WithAccountID(r.Context(), sub)

		// run next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
