package stopwatch

import (
	"time"
)

type Stopwatch struct {
	start time.Time
}

func New() *Stopwatch {
	return &Stopwatch{start: time.Now()}
}

func (s *Stopwatch) Reset() {
	s.start = time.Now()
}

func (s *Stopwatch) Record() time.Duration {
	return time.Since(s.start)
}
