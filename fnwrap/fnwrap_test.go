package fnwrap_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/latavin243/goutils/fnwrap"
)

func TestWithPanicRecovery(t *testing.T) {
	testName := "TestWithPanicRecovery"
	fmt.Printf("=== [%s] start ===\n", testName)
	defer fmt.Printf("=== [%s] finished ===\n", testName)

	f := func() error {
		panic("panicError")
	}
	err := fnwrap.New(f).WithPanicRecovery().Do()
	fmt.Printf("err=%s\n", err)
}

func TestWithTimeConsumptionCallback(t *testing.T) {
	testName := "TestWithTimeConsumptionCallback"
	fmt.Printf("=== [%s] start ===\n", testName)
	defer fmt.Printf("=== [%s] finished ===\n", testName)

	f := func() error {
		time.Sleep(time.Second)
		return nil
	}
	fnwrap.New(f).
		WithTimerCallback(func(timeConsumption time.Duration, _ error) {
			fmt.Printf("function time consumption=%fs\n", timeConsumption.Seconds())
		}).Do()
}

func TestWithRetry(t *testing.T) {
	testName := "TestWithRetry"
	fmt.Printf("=== [%s] start ===\n", testName)
	defer fmt.Printf("=== [%s] finished ===\n", testName)

	retrySettings := &fnwrap.RetrySettings{
		Attempts: 3,
		Delay:    time.Second,
		OnRetryCallback: func(n uint, err error) {
			fmt.Printf("retry failed, attempt=%d, err=%s\n", n, err)
		},
	}
	f := func() error {
		return fmt.Errorf("expected error")
	}
	fnwrap.New(f).WithRetry(retrySettings).Do()
}

func TestChainWrapper(t *testing.T) {
	testName := "TestChainWrapper"
	fmt.Printf("=== [%s] start ===\n", testName)
	defer fmt.Printf("=== [%s] finished ===\n", testName)

	f := func() error {
		time.Sleep(500 * time.Millisecond)
		panic("panicError")
	}
	retrySettings := &fnwrap.RetrySettings{
		Attempts: 3,
		Delay:    time.Second,
		OnRetryCallback: func(n uint, err error) {
			fmt.Printf("retry failed, attempt=%d, err=%s\n", n, err)
		},
	}
	timeConsumptionCallback := func(timeConsumption time.Duration, _ error) {
		fmt.Printf("function time consumption=%fs\n", timeConsumption.Seconds())
	}

	err := fnwrap.New(f).
		WithPanicRecovery().
		WithTimerCallback(timeConsumptionCallback).
		WithRetry(retrySettings).
		Do()
	fmt.Printf("wrapper chain with error, err=%s\n", err)
}