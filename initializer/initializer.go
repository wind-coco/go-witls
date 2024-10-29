package initializer

import "github.com/wind-coco/go-witls/ctx"

type Initializer interface {
	Initialize() (any, error)
	Name() string
}

var initializers []Initializer

func Register(initializer ...Initializer) {
	initializers = append(initializers, initializer...)
}

func Initialize() error {

	for _, initializer := range initializers {
		result, err := initializer.Initialize()
		if err != nil {
			return err
		}
		ctx.Add(initializer.Name(), result)
	}

	return nil
}
