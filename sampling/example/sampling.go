package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/haoel/downsampling/demo/common"

	"github.com/latavin243/goutils/sampling"
)

func main() {
	const sampledCount = 100
	sampler := sampling.NewLTDSampler()

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	rawdata := common.LoadPointsFromCSV(filepath.Join(dir, "sampling/example/raw_points.csv"))
	samplePoints, err := sampler.Sample(rawdata, sampledCount)
	if err != nil {
		panic(err)
	}
	fmt.Printf("samplePoints: %+v", samplePoints)
}
