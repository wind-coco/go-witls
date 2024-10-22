package cipher

import (
	"bytes"
	"errors"
)

type Cipher interface {
	Encrypt(src []byte) []byte
	Decrypt(dst []byte) ([]byte, error)
}

// PKCS5Padding 使用PKCS5进行填充
func PKCS5Padding(src []byte, blockSize int) []byte {
	paddingLen := blockSize - len(src)%blockSize
	padding := bytes.Repeat([]byte{byte(paddingLen)}, paddingLen)
	return append(src, padding...)
}

// PKCS5Unpadding 去掉PKCS5填充
func PKCS5Unpadding(dst []byte) ([]byte, error) {
	length := len(dst)
	unPadding := int(dst[length-1])
	if (length-unPadding) > length || (length-unPadding) < 0 {
		return nil, errors.New("index out of range")
	}
	return dst[:(length - unPadding)], nil
}
