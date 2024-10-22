package randutil

import (
	"testing"
)

func TestRandString(t *testing.T) {

	cases := []struct {
		name string
		len  int
	}{
		{
			name: "",
			len:  6,
		},
		{
			name: "",
			len:  8,
		},
		{
			name: "",
			len:  16,
		},
		{
			name: "",
			len:  24,
		},
	}
	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			randStr := Strings(cc.len)
			t.Logf("%d:%s", cc.len, randStr)
			if len(randStr) != cc.len {
				t.Errorf("%s want len:%d;but got:%d", cc.name, cc.len, len(randStr))
			}
		})
	}

}
