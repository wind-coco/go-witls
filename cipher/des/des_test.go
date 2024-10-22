package des

import (
	"testing"

	"github.com/wind-coco/go-witls/randutil"
)

func TestDES(t *testing.T) {

	cases := []struct {
		name   string
		key    string
		iv     string
		source string
	}{
		{
			name:   "abcdef",
			key:    "bABdYXNiYjGemjEJegYpWfMv",
			iv:     "pJehULuv",
			source: "abcdef",
		},
		{
			name:   "hijklmn",
			key:    randutil.Strings(24),
			iv:     randutil.Strings(8),
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
