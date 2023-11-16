package routinegroup

import (
	"golang.org/x/sync/errgroup"
)

type ErrGroup struct {
	group *errgroup.Group
	errs  []error
}

func NewErrGroup() Group {
	return &ErrGroup{
		group: new(errgroup.Group),
		errs:  make([]error, 0),
	}
}

func (g *ErrGroup) Submit(task func() error) {
	g.group.Go(task)
}

func (g *ErrGroup) Wait() error {
	return g.group.Wait()
}

func (g *ErrGroup) SetLimit(n int) {
	g.group.SetLimit(n)
}

func (g *ErrGroup) Close() error {
	return nil
}
