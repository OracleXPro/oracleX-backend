package telegramx

import (
	"encoding/hex"
	"errors"
)

func Verify(botToken, hash2, initData string) (bool, error) {
	secretKey := makeHMAC([]byte("WebAppData"), []byte(botToken))
	hash := makeHMAC(secretKey, []byte(initData))

	if hex.EncodeToString(hash) == hash2 {
		return true, nil
	}
	return false, errors.New("fail to verify")
}
