package authentication

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/arifinhermawan/blib/log"
	"github.com/golang-jwt/jwt/v5"
)

func (svc *Service) Authenticate(ctx context.Context, userID int64) (string, error) {
	metadata := map[string]interface{}{
		"user_id": userID,
	}

	existingJWT, err := svc.getJWTFromRedis(ctx, userID)
	if err != nil {
		log.Error(ctx, metadata, err, "[GenerateJWT] svc.getJWTFromRedis() got error")
		return "", err
	}

	if existingJWT != "" {
		return existingJWT, nil
	}

	cfg := svc.lib.GetConfig()
	ttl := cfg.TTL.JWT
	timeNow := svc.lib.GetTimeGMT7()
	expiredAt := timeNow.Add(time.Second * time.Duration(ttl))

	customClaim := customClaim{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredAt),
			IssuedAt:  jwt.NewNumericDate(timeNow),
			Issuer:    "probi",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaim)
	tokenString, err := token.SignedString([]byte(cfg.Hash.JWT))
	if err != nil {
		log.Error(ctx, metadata, err, "[GenerateJWT] token.SignedString() got error")
		return "", err
	}

	go func() {
		err = svc.setJWTToRedis(ctx, userID, tokenString)
		if err != nil {
			log.Warn(ctx, metadata, err, "[GenerateJWT] svc.setJWTToRedis() got error")
		}
	}()

	return tokenString, nil
}

func (svc *Service) GeneratePassword(password string) string {
	hash := hmac.New(sha256.New, []byte(svc.lib.GetConfig().Hash.Password))
	hash.Write([]byte(password))
	hashedBytes := hash.Sum(nil)

	return hex.EncodeToString(hashedBytes)
}

func (svc *Service) InvalidateJWT(ctx context.Context, userID int64) error {
	return svc.deleteJWTFromRedis(ctx, userID)
}

func (svc *Service) IsPasswordMatch(password string, encPass string) bool {
	return svc.GeneratePassword(password) == encPass
}
