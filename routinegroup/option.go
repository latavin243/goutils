package routinegroup

type GroupType uint

type groupOptions struct {
	poolSize int
}

type Option func(*groupOptions)

var emptyOption = func(*groupOptions) {}

func withPoolSize(poolSize int) Option {
	if poolSize <= 0 {
		return emptyOption
	}
	return func(o *groupOptions) {
		o.poolSize = poolSize
	}
}
