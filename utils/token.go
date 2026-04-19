package utils

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/google/uuid"
)

func GenerateToken() (string, string) {
	token := uuid.NewString()
	shaHash := sha256.Sum256([]byte(token))
	hash := hex.EncodeToString(shaHash[:])
	return token, hash
}
