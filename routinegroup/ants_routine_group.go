package routinegroup

import (
	"sync"

	"github.com/panjf2000/ants/v2"
)

type AntsGroup struct {
	pool *ants.Pool
	wg   sync.WaitGroup
	errs []error
}

func NewAntsGroup(pool *ants.Pool) Group {
	return &AntsGroup{
		pool: pool,
		errs: make([]error, 0),
	}
}

func (g *AntsGroup) Submit(task func() error) {
	if g.pool == nil {
		g.errs = append(g.errs, ErrPoolNotInitialized)
		return
	}
	g.wg.Add(1)
	err := g.pool.Submit(func() {
		defer g.wg.Done()
		err := task()
		if err != nil {
			g.errs = append(g.errs, err)
		}
	})
	if err != nil {
		g.wg.Done()
		g.errs = append(g.errs, err)
	}
}

func (g *AntsGroup) Wait() error {
	if g.pool == nil {
		return ErrPoolNotInitialized
	}

	g.wg.Wait()
	if len(g.errs) > 0 {
		return g.errs[0]
	}
	return nil
}

func (g *AntsGroup) SetLimit(n int) {
	if n <= 0 {
		return
	}
	g.pool.Tune(n)
}

func (g *AntsGroup) Close() error {
	if g.pool != nil {
		g.pool.Release()
	}
	return nil
}
