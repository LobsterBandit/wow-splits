package main

import (
	"encoding/json"
	"fmt"

	"github.com/lobsterbandit/wow-splits/aggregator"
)

func main() {
	aggregator.HelloAggregator()

	filePaths := aggregator.FindAllSpeedrunSplits("/World of Warcraft")

	for i, file := range filePaths {
		data, err := aggregator.ParseCharacter(file)
		if err != nil {
			fmt.Printf("\nerror reading %q: %v", file, err)
		}

		pretty, _ := json.MarshalIndent(data, "", "\t")

		fmt.Printf("\n%v: %s", i, pretty)
	}
}
