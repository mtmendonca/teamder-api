package middleware

import (
	"context"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mtmendonca/teamder-api/common/types"
)

const (
	contextKeyUserID = types.ContextKey("userID")
)

// GetUserID returns the userID extracted from the jwt token in Context
func GetUserID(ctx context.Context) string {
	return ctx.Value(contextKeyUserID).(string)
}

// JWTMiddleware returns a configured go-jwt-middleware instance
func JWTMiddleware(secret string) *jwtmiddleware.JWTMiddleware {
	return jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
		SigningMethod: jwt.SigningMethodHS512,
	})
}

// ExtractUser adds user to the request context
func ExtractUser(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	//Added by JWTMiddleware
	user := r.Context().Value("user")
	// Added below, no need to add twice
	userID := r.Context().Value("userID")
	if userID == nil {
		ctx := context.WithValue(r.Context(), contextKeyUserID, user.(*jwt.Token).Claims.(jwt.MapClaims)["sub"])
		next(rw, r.WithContext(ctx))
		return
	}
	next(rw, r)
}
