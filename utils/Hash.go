package utils

import (
	"os"

	"golang.org/x/crypto/argon2"
)

var secretKey = os.Getenv("HASH_KEY")

func HashPassword(password string) []byte {
	return argon2.IDKey([]byte(password), []byte(secretKey), 1, 64*1024, 4, 32)
}
