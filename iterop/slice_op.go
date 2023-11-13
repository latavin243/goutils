package iterop

const minBatchSize int = 1

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

func SliceChunk[T any](rawSlice []T, batchSize int) (batches [][]T) {
	batchSize = max(batchSize, minBatchSize)

	rawLen := len(rawSlice)
	if rawLen <= batchSize {
		return [][]T{rawSlice}
	}
	curBatch := make([]T, 0, batchSize)
	for _, item := range rawSlice {
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

func SliceRemoveDup[T comparable](arr []T) []T {
	set := make(map[T]struct{}, len(arr))
	j := 0
	for _, v := range arr {
		_, ok := set[v]
		if ok {
			continue
		}
		set[v] = struct{}{}
		arr[j] = v
		j++
	}
	return arr[:j]
}
