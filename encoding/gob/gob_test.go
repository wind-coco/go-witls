package gob

import "testing"

func TestGob(t *testing.T) {

	src := "abc"
	g := New()

	dst := g.EncodeToString([]byte(src))

	decrypted, err := g.DecodeToString(dst)
	if err != nil {
		t.Fatalf("error occured in [Decode] and err is :%s", err.Error())
	}
	if src != decrypted {
		t.Errorf("want %s;but got %s", src, decrypted)
	}

}
