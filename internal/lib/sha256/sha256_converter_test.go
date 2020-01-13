package sha256

import (
	"fmt"
	"testing"
)

func TestGetHashString(t *testing.T) {
	fmt.Print(GetHashString("abc123"))
}
