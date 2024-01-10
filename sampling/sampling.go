package sampling

import (
	downsampling "github.com/haoel/downsampling/core"
)

type Point = downsampling.Point

// Sampler is the interface that wraps the basic Sample method.
// if sampleNum == 0, return original points
type Sampler interface {
	Sample(points []Point, sampleNum uint32) ([]Point, error)
}

type LTDSampleCollector struct{}

func NewLTDSampleCollector() Sampler {
	return &LTDSampleCollector{}
}

func (c *LTDSampleCollector) Sample(points []Point, sampleNum uint32) (samples []Point, err error) {
	return downsampling.LTD(points, int(sampleNum)), nil
}

type LTTBSampleCollector struct{}

func NewLTTBSampleCollector() Sampler {
	return &LTTBSampleCollector{}
}

func (c *LTTBSampleCollector) Sample(points []Point, sampleNum uint32) (samples []Point, err error) {
	return downsampling.LTTB(points, int(sampleNum)), nil
}
