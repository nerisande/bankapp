package utils

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/google/uuid"
)

// Функция генерации случайной строки вида "e240d825d255af751f5f55af8d9671beabdf2236c0a3b4e2639b3e182d994c88"

func GenerateRandomAddress() string {
	randUUID := uuid.New().String()
	hashBytes := sha256.Sum256([]byte(randUUID))
	str := hex.EncodeToString(hashBytes[:])
	return str
}
