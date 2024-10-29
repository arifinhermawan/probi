package authentication

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func (svc *Service) GeneratePassword(password string) string {
	hash := hmac.New(sha256.New, []byte(svc.lib.GetConfig().Hash.Password))
	hash.Write([]byte(password))
	hashedBytes := hash.Sum(nil)

	return hex.EncodeToString(hashedBytes)
}
