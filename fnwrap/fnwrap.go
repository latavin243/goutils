package fnwrap

import (
	"fmt"
	"time"

	"github.com/avast/retry-go"
)

type WrappedFn func() error

func New(f func() error) WrappedFn {
	return f
}

func (f WrappedFn) Do() error {
	return f()
}

func (f WrappedFn) WithCallback(callback func(error)) WrappedFn {
	return func() error {
		err := f()
		callback(err)
		return err
	}
}

func (f WrappedFn) WithTimerCallback(callback func(time.Duration, error)) WrappedFn {
	return func() error {
		start := time.Now()
		err := f()
		callback(time.Since(start), err)
		return err
	}
}

func (f WrappedFn) WithPanicRecovery() WrappedFn {
	return func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("recovered panic with error, err=%s", r)
			}
		}()
		err = f()
		return err
	}
}

// retryOpts example: retry.Attemts(3), retry.Delay(time.Second), retry.OnRetry(...)
func (f WrappedFn) WithRetry(retryOpts ...retry.Option) WrappedFn {
	retryFn := func() error {
		return f()
	}
	return func() error {
		return retry.Do(retryFn, retryOpts...)
	}
}
