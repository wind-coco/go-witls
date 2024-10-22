package encoding

type Encoder interface {
	Encode(src []byte) []byte
	EncodeToString(src []byte) string
}

type Decoder interface {
	Decode(encoded string) ([]byte, error)
	DecodeToString(encoded string) (string, error)
}
type EncoderDecoder interface {
	Encoder
	Decoder
}
