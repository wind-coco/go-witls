package noop

type Encoding struct {
}

func New() *Encoding {
	return &Encoding{}
}

func (this *Encoding) Encode(src []byte) []byte {
	return src
}
func (this *Encoding) EncodeToString(src []byte) string {
	return string(src)
}

func (this *Encoding) Decode(encoded string) ([]byte, error) {
	return []byte(encoded), nil
}
func (this *Encoding) DecodeToString(encoded string) (string, error) {
	return encoded, nil
}
