package ctx

import (
	"testing"
)

type A struct {
	stringKey string
}
type B struct {
}

func TestCtx(t *testing.T) {
	cases := []struct {
		name           string
		key            string
		serviceBuilder func() any
	}{
		{
			name: "",
			key:  "a",
			serviceBuilder: func() any {
				return &A{
					stringKey: "stringValue",
				}
			},
		},
	}

	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			service := cc.serviceBuilder()
			Add(cc.key, service)
			result, ok := Get[*A](cc.key)
			if !ok {
				t.Fatalf("%s have you add key %s ?", cc.name, cc.key)
			}
			if result != service {
				t.Errorf("%s result!=service", cc.name)
			}
		})
	}
}
