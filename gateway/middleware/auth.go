package middleware

import (
	"context"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
)

type contextKey string

var (
	contextKeyUserID = contextKey("userID")
)

// GetUserID returns the userID extracted from the jwt token in Context
func GetUserID(ctx context.Context) string {
	return ctx.Value(contextKeyUserID).(string)
}

//JWTMiddleware returns a configured go-jwt-middleware instance
func JWTMiddleware(secret string) *jwtmiddleware.JWTMiddleware {
	return jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
}

//ExtractUser adds user to the request context
func ExtractUser(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	user := r.Context().Value("user")
	userID := r.Context().Value("userId")
	if userID == nil {
		ctx := context.WithValue(r.Context(), contextKeyUserID, user.(*jwt.Token).Claims.(jwt.MapClaims)["sub"])
		next(rw, r.WithContext(ctx))
		return
	}
	next(rw, r)
}
