package ctx

var services = make(map[string]any)

func Add[T any](name string, service T) {
	services[name] = service
}
func Get[T any](name string) (T, bool) {
	service, ok := services[name]
	if !ok {
		var empty T
		return empty, false
	}
	s, ok := service.(T)
	return s, ok
}

func MustGet[T any](name string) T {
	service := services[name]
	return service.(T)
}
