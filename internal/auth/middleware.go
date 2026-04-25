package auth

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/krotesq/strowger/internal/util"
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
		cookie, err := r.Cookie("token")
		if errors.Is(err, http.ErrNoCookie) {
			util.NewResponse(401, "Authentification failed", nil).Send(w)
			return
		}
		if err != nil {
			log.Println(err.Error())
			util.NewResponse(500, "Internal server error", nil).Send(w)
			return
		}

		// verify jwt
		secretBase64 := os.Getenv("JWT_SECRET")
		sub, err := ValidateToken(cookie.Value, secretBase64)
		if err != nil {
			util.NewResponse(401, "Authentification failed", nil).Send(w)
			return
		}

		log.Printf("User %s authorized with jwt\n", sub)

		ctx := WithAccountID(r.Context(), sub)

		// run next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
