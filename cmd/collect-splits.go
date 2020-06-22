package main

import (
	"fmt"

	"github.com/lobsterbandit/wow-splits/aggregator"
)

func main() {
	aggregator.HelloAggregator()

	files := aggregator.FindAllSpeedrunSplits("/World of Warcraft")

	fmt.Printf("Files found: %v\n", files)
}
