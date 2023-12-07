package iterop

import "sync"

func SliceParMap[T, R any](inputs []T, mapFn func(T) R) []R {
	outputs := make([]R, len(inputs))

	var wg sync.WaitGroup
	for i, input := range inputs {
		wg.Add(1)
		go func(i int, in T) {
			defer wg.Done()
			outputs[i] = mapFn(in)
		}(i, input)
	}
	wg.Wait()
	return outputs
}

func SliceParForEach[T any](inputs []T, forEachFn func(T)) {
	var wg sync.WaitGroup
	for _, input := range inputs {
		wg.Add(1)
		go func(in T) {
			defer wg.Done()
			forEachFn(in)
		}(input)
	}
	wg.Wait()
}
