package middleware

import (
	"context"
	"fmt"
	"goweb/utils"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(1, 5) // max 1 request per second

func Middleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		if !limiter.Allow() {
			utils.ResponseBody(writer, http.StatusTooManyRequests, "Too Many Requests", "")
			return
		}

		protectedUrl := []string{
			"/api/todo/",
		}

		for _, urls := range protectedUrl {
			if strings.HasPrefix(request.URL.Path, urls) {
				/// Get token from request header
				authHeader := request.Header.Get("Authorization")
				if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
					utils.ResponseBody(writer, http.StatusUnauthorized, "Unauthorized", "Access denied")
					return
				}

				tokenString := strings.TrimPrefix(authHeader, "Bearer ")
				tokenString = strings.TrimSpace(tokenString)

				token, err := jwt.Parse(tokenString,
					func(token *jwt.Token) (any, error) {

						if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
							return nil, fmt.Errorf("Error signing method")
						}

						return []byte(utils.JwtSecret), nil
					})
				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					// Get user id from jwt claims and put to request context for further use
					// like insert that need user id or any
					uidRaw, exists := claims["user_id"]
					if !exists {
						utils.ResponseBody(writer, http.StatusUnauthorized, "Unauthorized", "Access denied")
						return
					}

					userID, ok := uidRaw.(string)
					if !ok || userID == "" {
						utils.ResponseBody(writer, http.StatusUnauthorized, "Unauthorized", "Access denied")
						return
					}

					ctx := context.WithValue(request.Context(), "user_id", userID)
					next.ServeHTTP(writer, request.WithContext(ctx))
					return
				}
				if err != nil || !token.Valid {
					utils.ResponseBody(writer, http.StatusUnauthorized, "Unauthorized", "Access denied")
					return
				}
			}
			next.ServeHTTP(writer, request)
			return
		}
	})
}
