package logic

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRandomKey() (string, error) {
	buffer := make([]byte, 32)

	if _, err := rand.Read(buffer); err != nil {
		return "", err
	}

	key := hex.EncodeToString(buffer)

	return key, nil
}
