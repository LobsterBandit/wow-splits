package main

import (
	"fmt"

	"github.com/lobsterbandit/wow-splits/aggregator"
)

func main() {
	aggregator.HelloAggregator()

	filePaths := aggregator.FindAllSpeedrunSplits("/World of Warcraft")

	for i, file := range filePaths {
		data, err := aggregator.ReadSpeedrunSplits(file)
		if err != nil {
			fmt.Printf("\nerror reading %q: %v", file, err)
		}

		fmt.Printf("\n%v: %s", i, data)
	}
}
