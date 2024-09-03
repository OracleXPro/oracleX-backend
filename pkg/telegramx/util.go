package telegramx

import (
	"crypto/hmac"
	"crypto/sha256"
)

func makeHMAC(secret, data []byte) []byte {
	hasher := hmac.New(sha256.New, secret)
	hasher.Write(data)
	hash := hasher.Sum(nil)
	return hash
}
