package aes

import (
	"encoding/base64"
	"testing"

	base64util "github.com/wind-coco/go-witls/encoding/base64"
	"github.com/wind-coco/go-witls/randutil"
)

func TestAes(t *testing.T) {

	cases := []struct {
		name   string
		key    string
		iv     string
		source string
	}{
		{
			name:   "abcdef",
			key:    "yNFhBxenWEvLZLkjCPpsPSzg",
			iv:     "ZbhEfAhlbzsLsBjS",
			source: "abcdef",
		},
		{
			name:   "hijklmn",
			key:    randutil.Strings(24),
			iv:     randutil.Strings(16),
			source: "hijklmn",
		},
	}

	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {

			cipher, err := New(cc.key, cc.iv)
			if err != nil {
				t.Errorf("%s error occured on new (%v)", cc.name, err)
			}
			encrpted := cipher.Encrypt([]byte(cc.source))
			//t.Logf("%s encrpted:%s", cc.name, encrpted)

			result, err := cipher.Decrypt(encrpted)
			if err != nil {
				t.Errorf("%s error occured on decrypt (%v)", cc.name, err)
			}
			if string(result) != cc.source {
				t.Errorf("%s want:%s;but got %s", cc.name, cc.source, result)
			}

		})
	}

}

func TestEncrypt(t *testing.T) {
	cases := []struct {
		name   string
		key    string
		iv     string
		source string
		want   string
	}{
		{
			name:   "de-single",
			key:    "VXHbxjAVuylxWqvrBCSCaopO",
			iv:     "uvjBuknnuvjBuknn",
			source: "de-single",
			want:   "x41VIwCNh+oOksbFyYdjdw==",
		},
		{
			name:   "de-single",
			key:    "woniucsdnvip8888",
			iv:     "sIVxRsEWgAHNNLYo",
			source: "嗨，您好！",
			want:   "515S8VG52TqbhUwB1T9DiA==",
		},
	}
	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {

			cipher, err := New(cc.key, cc.iv)
			cipher.Encoding = base64util.New(base64.StdEncoding)
			if err != nil {
				t.Errorf("%s error occured on new (%v)", cc.name, err)
			}
			encrpted := cipher.Encrypt([]byte(cc.source))
			if cc.want != string(encrpted) {
				t.Errorf("%s want:%s,but got:%s", cc.name, cc.want, encrpted)
			}

			resultSource, err := cipher.Decrypt(encrpted)
			if err != nil {
				t.Errorf("%s error occured on Decrypt (%v)", cc.name, err)
			}
			if cc.source != string(resultSource) {
				t.Errorf("%s want:%s,but got:%s", cc.name, cc.want, resultSource)
			}

		})
	}
}
