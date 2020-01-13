package sha256

import (
	"crypto/sha256"
	"fmt"
)

func GetHashString(str string) (rStr string) {
	h := sha256.New()
	h.Write([]byte(str))
	rStr = fmt.Sprintf("%x", h.Sum(nil))
	return
}
