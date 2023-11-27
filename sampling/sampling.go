package sampling

import (
	downsampling "github.com/haoel/downsampling/core"
)

type Point = downsampling.Point

// SampleCollector is the interface that wraps the basic Sample method.
// if sampleNum == 0, return original points
type SampleCollector interface {
	Sample(points []Point, sampleNum uint32) ([]Point, error)
}

type LTDSampleCollector struct{}

func NewLTDSampleCollector() SampleCollector {
	return &LTDSampleCollector{}
}

func (c *LTDSampleCollector) Sample(points []Point, sampleNum uint32) (samples []Point, err error) {
	return downsampling.LTD(points, int(sampleNum)), nil
}

type LTTBSampleCollector struct{}

func NewLTTBSampleCollector() SampleCollector {
	return &LTTBSampleCollector{}
}

func (c *LTTBSampleCollector) Sample(points []Point, sampleNum uint32) (samples []Point, err error) {
	return downsampling.LTTB(points, int(sampleNum)), nil
}
