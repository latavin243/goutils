package routinegroup

import (
	"errors"
	"io"
)

const DefaultPoolSize = 1000

var (
	ErrUnknownGroupType   = errors.New("unknown group type")
	ErrPoolNotInitialized = errors.New("pool not initialized")
)

type Group interface {
	io.Closer
	Submit(task func() error)
	Wait() error
	SetLimit(n int)
}
