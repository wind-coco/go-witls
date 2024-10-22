package hex

import (
	"testing"
)

func TestHex(t *testing.T) {

	src := "abc"
	h := New()
	dst := h.Encode([]byte(src))
	decrypted, _ := h.DecodeToString(string(dst))
	if src != decrypted {
		t.Errorf("want %s;but got %s", src, decrypted)
	}

}
