package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/arifinhermawan/blib/log"
	internalContext "github.com/arifinhermawan/probi/internal/lib/context"
	"github.com/golang-jwt/jwt/v5"
)

func (a *Auth) AuthMiddleware(endpointHandler func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx := internalContext.DefaultContext()
		authHeader := request.Header.Get("Authorization")
		writer.Header().Set("Content-Type", "application/json")
		response := struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    http.StatusUnauthorized,
			Message: "unauthorized",
		}

		if authHeader == "" {
			log.Error(ctx, nil, nil, "[AuthMiddleware] Authorization header is empty")
			writer.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(writer).Encode(response)
			return
		}

		sliced := strings.Split(authHeader, " ")
		if len(sliced) != 2 || sliced[0] != "Bearer" {
			log.Error(ctx, nil, nil, "[AuthMiddleware] Invalid Authorization format")
			writer.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(writer).Encode(response)
			return
		}

		token, err := jwt.ParseWithClaims(sliced[1], &customClaim{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(a.cfg.GetConfig().Hash.JWT), nil
		})
		if err != nil {
			log.Error(ctx, nil, err, "[AuthMiddleware] Failed to parse JWT")
			writer.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(writer).Encode(response)
			return
		}

		if claims, ok := token.Claims.(*customClaim); ok && token.Valid {
			ctx := context.WithValue(request.Context(), ContextKeyUserID, claims.UserID)
			endpointHandler(writer, request.WithContext(ctx))
		} else {
			log.Error(ctx, nil, nil, "[AuthMiddleware] Invalid claims or token")
			writer.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(writer).Encode(response)
			return
		}
	}
}
