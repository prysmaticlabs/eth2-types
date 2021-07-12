package types

import (
	"encoding/hex"
	"testing"
)

func HexDecodeOrDie(t *testing.T, str string) []byte {
	decoded, err := hex.DecodeString(str)
	if err != nil {
		t.Errorf("hex.DecodeString(%s) unexpected error = %v", str, err)
	}
	return decoded
}
