package base64

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {

	src := "abc"
	b := New(base64.StdEncoding)
	dst := b.Encode([]byte(src))
	decrypted, _ := b.DecodeToString(string(dst))
	if src != decrypted {
		t.Errorf("want %s;but got %s", src, decrypted)
	}

}

func TestBase64Encode(t *testing.T) {
	src := "hello"
	b := New(base64.StdEncoding)
	dst := b.Encode([]byte(src))
	//decrypted, _ := b.DecodeToString(string(dst))
	fmt.Println(dst)

	dstDst := b.Encode(dst)
	fmt.Println(dstDst)

	var decrypted string
	decrypted, _ = b.DecodeToString(string(dstDst))
	fmt.Println(decrypted)
	decrypted, _ = b.DecodeToString(decrypted)
	fmt.Println(decrypted)

}
