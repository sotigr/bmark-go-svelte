package system

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(str string) string {
	h := sha256.New()
	bs := h.Sum([]byte(str))
	return hex.EncodeToString(bs)
}

func ShortHash(str string) string {
	h := sha256.New()
	bs := h.Sum([]byte(str))
	hash := hex.EncodeToString(bs)
	return hash[:10]
}
