package iterop

const minBatchSize int = 1

// SliceFilter filters the inputs slice by the filterFn
func SliceFilter[T any](inputs []T, filterFn func(T) bool) []T {
	if len(inputs) == 0 {
		return make([]T, 0)
	}
	outputs := make([]T, 0)
	for _, v := range inputs {
		if filterFn(v) {
			outputs = append(outputs, v)
		}
	}
	return outputs
}

// SliceMap maps the inputs slice by the mapFn
func SliceMap[T1, T2 any](inputs []T1, mapFn func(T1) T2) []T2 {
	if len(inputs) == 0 {
		return make([]T2, 0)
	}
	outputs := make([]T2, len(inputs))
	for i, v := range inputs {
		outputs[i] = mapFn(v)
	}
	return outputs
}

// SliceFlatMap flatmaps the inputs slice by the flatmapFn
func SliceFlatMap[T1, T2 any](inputs []T1, flatmapFn func(input T1, collector func(T2))) []T2 {
	if len(inputs) == 0 {
		return make([]T2, 0)
	}
	outputs := make([]T2, 0)
	collector := func(v T2) {
		outputs = append(outputs, v)
	}
	for _, v := range inputs {
		flatmapFn(v, collector)
	}
	return outputs
}

// SliceReduce reduces the inputs slice by the reduceFn
func SliceReduce[T, R any](inputs []T, init R, reduceFn func(R, T) R) R {
	if len(inputs) == 0 {
		return init
	}
	output := init
	for _, v := range inputs {
		output = reduceFn(output, v)
	}
	return output
}

// SliceChunk splits the inputs into batches, each batch has batchSize elements
// e.g. SliceChunk([1,2,3,4,5,6,7,8], 3) => [[1,2,3], [4,5,6], [7,8]]
func SliceChunk[T any](inputs []T, batchSize int) (batches [][]T) {
	batchSize = max(batchSize, minBatchSize)

	rawLen := len(inputs)
	if rawLen <= batchSize {
		return [][]T{inputs}
	}
	curBatch := make([]T, 0, batchSize)
	for _, item := range inputs {
		curBatch = append(curBatch, item)
		if len(curBatch) >= batchSize {
			batches = append(batches, curBatch)
			curBatch = []T{}
		}
	}
	if len(curBatch) > 0 {
		batches = append(batches, curBatch)
	}
	return batches
}

// SliceExist checks if the elements exist in the slice
func SliceExist[T comparable](inputs []T, elements ...T) bool {
	if len(inputs) == 0 {
		return false
	}
	set := make(map[T]struct{}, len(inputs))
	for _, v := range inputs {
		set[v] = struct{}{}
	}
	for _, v := range elements {
		if _, ok := set[v]; !ok {
			return false
		}
	}
	return true
}

// SliceUnique removes the duplicated elements in the slice and keep the order
func SliceUnique[T comparable](inputs []T) []T {
	set := make(map[T]struct{}, len(inputs))
	j := 0
	for _, v := range inputs {
		_, ok := set[v]
		if ok {
			continue
		}
		set[v] = struct{}{}
		inputs[j] = v
		j++
	}
	return inputs[:j]
}

// SliceSub returns the elements in inputs but not in subs
func SliceSub[T comparable](inputs, subs []T) []T {
	if len(inputs) == 0 {
		return make([]T, 0)
	}
	if len(subs) == 0 {
		return inputs
	}
	set := make(map[T]struct{}, len(subs))
	for _, v := range subs {
		set[v] = struct{}{}
	}
	outputs := make([]T, 0)
	for _, v := range inputs {
		if _, ok := set[v]; ok {
			continue
		}
		outputs = append(outputs, v)
	}
	return outputs
}
