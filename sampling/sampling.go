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

type LTDSampler struct{}

func NewLTDSampler() Sampler {
	return &LTDSampler{}
}

func (c *LTDSampler) Sample(points []Point, sampleNum uint32) (samples []Point, err error) {
	return downsampling.LTD(points, int(sampleNum)), nil
}

type LTTBSampler struct{}

func NewLTTBSampler() Sampler {
	return &LTTBSampler{}
}

func (c *LTTBSampler) Sample(points []Point, sampleNum uint32) (samples []Point, err error) {
	return downsampling.LTTB(points, int(sampleNum)), nil
}
