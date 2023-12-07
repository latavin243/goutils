package stopwatch_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/latavin243/goutils/stopwatch"
)

func TestStopwatch(t *testing.T) {
	sw := stopwatch.New()
	time.Sleep(2 * time.Second)
	fmt.Printf("1st record, %f s\n", sw.Record().Seconds())
	time.Sleep(1 * time.Second)
	fmt.Printf("2nd record, %f s\n", sw.Record().Seconds())

	fmt.Println("[reset]")
	sw.Reset()
	time.Sleep(1 * time.Second)
	fmt.Printf("1st record after reset, %f s\n", sw.Record().Seconds())
	time.Sleep(2 * time.Second)
	fmt.Printf("2nd record after reset, %f s\n", sw.Record().Seconds())
}
