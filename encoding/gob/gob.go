package gob

import (
	"bytes"
	"encoding/gob"
)

type carrier struct {
	Body string
}
type g struct {
}

func New() *g {
	return &g{}
}

func (*g) Encode(src []byte) []byte {
	var w bytes.Buffer
	err := gob.NewEncoder(&w).Encode(&carrier{
		Body: string(src),
	})
	if err != nil {
		panic(err)
	}
	return w.Bytes()
}
func (g *g) EncodeToString(src []byte) string {
	return string(g.Encode(src))
}

func (*g) Decode(encoded string) ([]byte, error) {
	var carrier carrier
	err := gob.NewDecoder(bytes.NewBufferString(encoded)).Decode(&carrier)
	if err != nil {
		return nil, err
	}
	return []byte(carrier.Body), nil
}
func (g *g) DecodeToString(encoded string) (string, error) {
	bs, err := g.Decode(encoded)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}
