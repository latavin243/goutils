package fnwrap_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/avast/retry-go"
	"github.com/stretchr/testify/assert"

	"github.com/latavin243/goutils/fnwrap"
)

func TestWithCallback(t *testing.T) {
	a := 1
	f := func() error {
		time.Sleep(time.Second)
		a++
		return nil
	}
	var timeConsumption time.Duration
	err := fnwrap.New(f).
		WithTimerCallback(func(duration time.Duration, _ error) {
			timeConsumption = duration
			fmt.Printf("function time consumption=%fs\n", duration.Seconds())
		}).Do()
	assert.NoError(t, err)
	assert.Equal(t, 1, int(timeConsumption.Seconds()))
}

func TestWithTimerCallback(t *testing.T) {
	f := func() error {
		time.Sleep(time.Second)
		return nil
	}
	var timeConsumption time.Duration
	err := fnwrap.New(f).
		WithTimerCallback(func(duration time.Duration, _ error) {
			timeConsumption = duration
			fmt.Printf("function time consumption=%fs\n", duration.Seconds())
		}).Do()
	assert.NoError(t, err)
	assert.Equal(t, 1, int(timeConsumption.Seconds()))
}

func TestWithPanicRecovery(t *testing.T) {
	f := func() error {
		panic("panicError")
	}
	err := fnwrap.New(f).WithPanicRecovery().Do()
	assert.Error(t, err)
}

func TestWithRetry(t *testing.T) {
	retryOpts := []retry.Option{
		retry.Attempts(3),
		retry.Delay(time.Second),
		retry.OnRetry(func(n uint, err error) { fmt.Printf("retry failed, attempt=%d, err=%s\n", n, err) }),
	}
	f := func() error {
		return fmt.Errorf("expected error")
	}
	err := fnwrap.New(f).WithRetry(retryOpts...).Do()
	assert.Error(t, err)
}

func TestChainWrapper(t *testing.T) {
	testName := "TestChainWrapper"
	fmt.Printf("=== [%s] start ===\n", testName)
	defer fmt.Printf("=== [%s] finished ===\n", testName)

	f := func() error {
		time.Sleep(500 * time.Millisecond)
		panic("panicError")
	}
	retryOpts := []retry.Option{
		retry.Attempts(3),
		retry.Delay(time.Second),
		retry.OnRetry(func(n uint, err error) { fmt.Printf("retry failed, attempt=%d, err=%s\n", n, err) }),
	}
	timeConsumptionCallback := func(timeConsumption time.Duration, _ error) {
		fmt.Printf("function time consumption=%fs\n", timeConsumption.Seconds())
	}

	err := fnwrap.New(f).
		WithPanicRecovery().
		WithTimerCallback(timeConsumptionCallback).
		WithRetry(retryOpts...).
		Do()
	assert.Error(t, err)
}
