package base64

import "encoding/base64"

type Encoding struct {
	enc *base64.Encoding
}

func New(enc *base64.Encoding) *Encoding {
	return &Encoding{
		enc: enc,
	}
}

func (this *Encoding) Encode(src []byte) []byte {
	dst := make([]byte, this.enc.EncodedLen(len(src)))
	this.enc.Encode(dst, src)
	return dst
}
func (this *Encoding) EncodeToString(src []byte) string {
	return this.enc.EncodeToString(src)
}

func (this *Encoding) Decode(encoded string) ([]byte, error) {
	bs, err := this.enc.DecodeString(encoded)
	if err != nil {
		return nil, err
	}
	return bs, nil
}
func (this *Encoding) DecodeToString(encoded string) (string, error) {
	bs, err := this.Decode(encoded)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}
