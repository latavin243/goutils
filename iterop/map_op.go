package iterop

type SortFn[T any] func([]T)

// MapValues takes the values of the map and sort the result (if sortFn is passed)
func MapValues[K comparable, V any](m map[K]V, sortFn SortFn[V]) []V {
	if len(m) == 0 {
		return make([]V, 0)
	}
	outputs := make([]V, 0, len(m))
	i := 0
	for _, v := range m {
		outputs = append(outputs, v)
		i++
	}
	if sortFn != nil {
		sortFn(outputs)
	}
	return outputs
}

// MapKeys takes the keys of the map and sort the result (if sortFn is passed)
func MapKeys[K comparable, V any](m map[K]V, sortFn SortFn[K]) []K {
	if len(m) == 0 {
		return make([]K, 0)
	}
	outputs := make([]K, 0, len(m))
	i := 0
	for k := range m {
		outputs = append(outputs, k)
		i++
	}
	if sortFn != nil {
		sortFn(outputs)
	}
	return outputs
}
