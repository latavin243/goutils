package stopwatch_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/latavin243/goutils/stopwatch"
)

func TestStopwatch(t *testing.T) {
	var elapsed float64
	sw := stopwatch.New()

	time.Sleep(2 * time.Second)
	elapsed = sw.Record().Seconds()
	fmt.Printf("1st record, %f s\n", elapsed)
	assert.Equal(t, 2, int(elapsed))

	time.Sleep(1 * time.Second)
	elapsed = sw.Record().Seconds()
	fmt.Printf("2nd record, %f s\n", elapsed)
	assert.Equal(t, 3, int(elapsed))

	fmt.Println("[reset]")
	sw.Reset()

	time.Sleep(1 * time.Second)
	elapsed = sw.Record().Seconds()
	fmt.Printf("1st record after reset, %f s\n", elapsed)
	assert.Equal(t, 1, int(elapsed))

	time.Sleep(2 * time.Second)
	elapsed = sw.Record().Seconds()
	fmt.Printf("2nd record after reset, %f s\n", elapsed)
	assert.Equal(t, 3, int(elapsed))
}
