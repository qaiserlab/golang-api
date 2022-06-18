package helpers

import (
	"crypto/sha1"
	"fmt"
	"time"
)

func GenSalt() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func GenHash(text string, salt string) string {
	sha := sha1.New()
	sha.Write([]byte(salt + text))
	encryptedText := fmt.Sprintf("%x", sha.Sum(nil))

	return encryptedText
}
