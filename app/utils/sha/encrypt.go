package shautil

import (
	"crypto/sha256"
	"encoding/hex"
)

func EncryptString(input string) string {
	hashed := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hashed[:])
}
