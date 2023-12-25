package common

import (
	"crypto/sha512"
	"encoding/hex"
	"serviceMetrica/internal/config"
)

// HashEncryption used (sha512)
func HashEncryption(password string) string {
	var sha = sha512.New() // Create hasher

	sha.Write([]byte(password + config.New().Auth.Salt))
	hashedPassword := sha.Sum(nil) // Get hashed password

	return hex.EncodeToString(hashedPassword)
}
