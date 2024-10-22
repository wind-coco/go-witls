package hex

import (
	"encoding/hex"
)

type h struct {
}

func New() *h {
	return &h{}
}

func (*h) Encode(src []byte) []byte {
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	return dst
}
func (*h) EncodeToString(src []byte) string {
	return hex.EncodeToString(src)
}

func (*h) Decode(encoded string) ([]byte, error) {
	bs, err := hex.DecodeString(encoded)
	if err != nil {
		return nil, err
	}
	return bs, nil
}
func (this *h) DecodeToString(encoded string) (string, error) {
	bs, err := this.Decode(encoded)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}
