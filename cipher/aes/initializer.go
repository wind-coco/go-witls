package aes

import (
	"github.com/wind-coco/go-witls/encoding"
	"github.com/wind-coco/go-witls/initializer"
)

type Initializer struct {
	key      string
	iv       string
	name     string
	encoding encoding.EncoderDecoder
}

func NewInitializer(name, key, iv string, encoding encoding.EncoderDecoder) initializer.Initializer {
	return &Initializer{
		key:      key,
		iv:       iv,
		name:     name,
		encoding: encoding,
	}
}

func (this *Initializer) Name() string {
	return this.name
}

func (this *Initializer) Initialize() (any, error) {
	aesCipher, err := New(this.key, this.iv)
	aesCipher.Encoding = this.encoding
	return aesCipher, err
}
